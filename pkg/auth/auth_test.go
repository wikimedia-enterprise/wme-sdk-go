package auth_test

import (
	"context"
	"embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

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
