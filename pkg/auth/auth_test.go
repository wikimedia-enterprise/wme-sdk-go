package auth_test

import (
	"context"
	"embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/auth"
)

//go:embed testdata/*
var testData embed.FS

type endpointTestSuite struct {
	suite.Suite
	sts int
	pth string
	fph string
	clt auth.API
	ctx context.Context
	srv *httptest.Server
	hrq any
}

func (s *endpointTestSuite) SetupSuite() {
	dta := []byte{}

	if len(s.fph) > 0 {
		fle, err := testData.Open(s.fph)

		if err != nil {
			log.Fatal(err)
		}

		defer fle.Close()
		dta, err = io.ReadAll(fle)

		if err != nil {
			log.Fatal(err)
		}
	}

	rtr := http.NewServeMux()
	rtr.HandleFunc(s.pth, func(w http.ResponseWriter, r *http.Request) {
		if s.hrq != nil {
			s.NoError(json.NewDecoder(r.Body).Decode(s.hrq))
		}

		w.WriteHeader(s.sts)
		_, _ = w.Write(dta)
	})

	s.srv = httptest.NewServer(rtr)
	s.ctx = context.Background()

	clt := auth.NewClient().(*auth.Client)
	clt.BaseURL = s.srv.URL
	s.clt = clt
}

func (s *endpointTestSuite) TearDownSuite() {
	s.srv.Close()
}

type loginTestSuite struct {
	endpointTestSuite
	req *auth.LoginRequest
}

func (s *loginTestSuite) SetupSuite() {
	s.pth = "/login"
	s.hrq = new(auth.LoginRequest)
	s.endpointTestSuite.SetupSuite()
}

func (s *loginTestSuite) TestLogin() {
	res, err := s.clt.Login(s.ctx, s.req)

	s.Assert().Equal(s.req, s.hrq)

	if s.sts != http.StatusOK {
		s.Empty(res)
		s.Error(err)
	} else {
		s.NotEmpty(res.IDToken)
		s.NotEmpty(res.AccessToken)
		s.NotEmpty(res.RefreshToken)
		s.NotEmpty(res.ExpiresIn)
		s.NoError(err)
	}
}

func TestLogin(t *testing.T) {
	req := &auth.LoginRequest{
		Username: "foo",
		Password: "bar",
	}

	for _, tcs := range []*loginTestSuite{
		{
			req: req,
			endpointTestSuite: endpointTestSuite{
				sts: http.StatusOK,
				fph: "testdata/login.json",
			},
		},
		{
			req: req,
			endpointTestSuite: endpointTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type refreshTokenTestSuite struct {
	endpointTestSuite
	req *auth.RefreshTokenRequest
}

func (s *refreshTokenTestSuite) SetupSuite() {
	s.pth = "/token-refresh"
	s.hrq = new(auth.RefreshTokenRequest)
	s.endpointTestSuite.SetupSuite()
}

func (s *refreshTokenTestSuite) TestRefreshToken() {
	res, err := s.clt.RefreshToken(s.ctx, s.req)

	s.Assert().Equal(s.req, s.hrq)

	if s.sts != http.StatusOK {
		s.Empty(res)
		s.Error(err)
	} else {
		s.NotEmpty(res.IDToken)
		s.NotEmpty(res.AccessToken)
		s.NotEmpty(res.ExpiresIn)
		s.NoError(err)
	}
}

func TestRefreshToken(t *testing.T) {
	req := &auth.RefreshTokenRequest{
		Username:     "foo",
		RefreshToken: "bar",
	}

	for _, tsc := range []*refreshTokenTestSuite{
		{
			req: req,
			endpointTestSuite: endpointTestSuite{
				sts: http.StatusOK,
				fph: "testdata/token-refresh.json",
			},
		},
		{
			req: req,
			endpointTestSuite: endpointTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tsc)
	}
}

type revokeTokenTestSuite struct {
	endpointTestSuite
	req *auth.RevokeTokenRequest
}

func (s *revokeTokenTestSuite) SetupSuite() {
	s.pth = "/token-revoke"
	s.hrq = new(auth.RevokeTokenRequest)
	s.endpointTestSuite.SetupSuite()
}

func (s *revokeTokenTestSuite) TestRevokeToken() {
	err := s.clt.RevokeToken(s.ctx, s.req)

	s.Assert().Equal(s.req, s.hrq)

	if s.sts != http.StatusOK {
		s.Error(err)
	} else {
		s.NoError(err)
	}
}

func TestRevokeToken(t *testing.T) {
	req := &auth.RevokeTokenRequest{
		RefreshToken: "foo",
	}

	for _, tsc := range []*revokeTokenTestSuite{
		{
			req: req,
			endpointTestSuite: endpointTestSuite{
				sts: http.StatusOK,
			},
		},
		{
			req: req,
			endpointTestSuite: endpointTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tsc)
	}
}

// MockAPI is a mock implementation of the API interface
type MockAPI struct {
	mock.Mock
}

func (m *MockAPI) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auth.LoginResponse), args.Error(1)
}

func (m *MockAPI) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auth.RefreshTokenResponse), args.Error(1)
}

func (m *MockAPI) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

type HelperTestSuite struct {
	suite.Suite
	mockAPI *MockAPI
	helper  *auth.Helper
}

func (suite *HelperTestSuite) SetupTest() {
	suite.mockAPI = new(MockAPI)
	suite.helper = &auth.Helper{
		Username: "test_user",
		Password: "test_password",
		API:      suite.mockAPI,
	}
	os.Remove("tokenstore.json")
}

func (suite *HelperTestSuite) TearDownTest() {
	os.Remove("tokenstore.json")
}

func (suite *HelperTestSuite) TestGetToken_NewToken() {
	suite.mockAPI.On("Login", mock.Anything, &auth.LoginRequest{Username: "test_user", Password: "test_password"}).Return(&auth.LoginResponse{
		AccessToken:  "new_access_token",
		RefreshToken: "new_refresh_token",
	}, nil)

	token, err := suite.helper.GetToken()

	suite.NoError(err)
	suite.Equal("new_access_token", token)
	suite.mockAPI.AssertExpectations(suite.T())
}

func (suite *HelperTestSuite) TestGetToken_ExistingValidToken() {
	tokenStore := &auth.Tokenstore{
		AccessToken:             "existing_access_token",
		AccessTokenGeneratedAt:  time.Now(),
		RefreshToken:            "existing_refresh_token",
		RefreshTokenGeneratedAt: time.Now(),
	}

	data, _ := json.Marshal(tokenStore)
	os.WriteFile("tokenstore.json", data, 0600)

	token, err := suite.helper.GetToken()

	suite.NoError(err)
	suite.Equal("existing_access_token", token)
	suite.mockAPI.AssertNotCalled(suite.T(), "Login")
	suite.mockAPI.AssertNotCalled(suite.T(), "RefreshToken")
}

func (suite *HelperTestSuite) TestGetToken_ExpiredAccessToken() {
	tokenStore := &auth.Tokenstore{
		AccessToken:             "expired_access_token",
		AccessTokenGeneratedAt:  time.Now().Add(-25 * time.Hour),
		RefreshToken:            "valid_refresh_token",
		RefreshTokenGeneratedAt: time.Now(),
	}

	data, _ := json.Marshal(tokenStore)
	os.WriteFile("tokenstore.json", data, 0600)

	suite.mockAPI.On("RefreshToken", mock.Anything, &auth.RefreshTokenRequest{
		Username:     "test_user",
		RefreshToken: "valid_refresh_token",
	}).Return(&auth.RefreshTokenResponse{
		AccessToken: "new_access_token",
	}, nil)

	token, err := suite.helper.GetToken()

	suite.NoError(err)
	suite.Equal("new_access_token", token)
	suite.mockAPI.AssertExpectations(suite.T())
}

func (suite *HelperTestSuite) TestGetToken_ExpiredRefreshToken() {
	tokenStore := &auth.Tokenstore{
		AccessToken:             "expired_access_token",
		AccessTokenGeneratedAt:  time.Now().Add(-25 * time.Hour),
		RefreshToken:            "expired_refresh_token",
		RefreshTokenGeneratedAt: time.Now().Add(-31 * 24 * time.Hour),
	}

	data, _ := json.Marshal(tokenStore)
	os.WriteFile("tokenstore.json", data, 0600)

	suite.mockAPI.On("Login", mock.Anything, &auth.LoginRequest{Username: "test_user", Password: "test_password"}).Return(&auth.LoginResponse{
		AccessToken:  "new_access_token",
		RefreshToken: "new_refresh_token",
	}, nil)

	token, err := suite.helper.GetToken()

	suite.NoError(err)
	suite.Equal("new_access_token", token)
	suite.mockAPI.AssertExpectations(suite.T())
}

func TestHelperTestSuite(t *testing.T) {
	suite.Run(t, new(HelperTestSuite))
}
