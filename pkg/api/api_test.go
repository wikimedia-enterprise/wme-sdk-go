package api_test

import (
	"context"
	"embed"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/api"
)

//go:embed testdata/*
var testData embed.FS

type readAllTestSuite struct {
	suite.Suite
	ctx context.Context
	rdr io.ReadCloser
	clt api.API
	pth string
}

func (s *readAllTestSuite) SetupTest() {
	var err error
	s.rdr, err = testData.Open(s.pth)
	s.NoError(err)

	s.clt = api.NewClient()
}

func (s *readAllTestSuite) TestReadAll() {
	nmc := 0
	err := s.clt.ReadAll(s.ctx, s.rdr, func(art *api.Article) error {
		s.NotEmpty(art.Name)
		s.NotEmpty(art.Identifier)
		nmc++
		return nil
	})

	s.NoError(err)
	s.NotZero(nmc)
}

func TestReadAll(t *testing.T) {
	for _, testCase := range []*readAllTestSuite{
		{
			pth: "testdata/simplewiki_namespace_0.tar.gz",
		},
	} {
		suite.Run(t, testCase)
	}
}

type setAccessTokenTestSuite struct {
	suite.Suite
	clt *api.Client
	tkn string
}

func (s *setAccessTokenTestSuite) SetupTest() {
	s.tkn = "foo"
	s.clt = new(api.Client)
}

func (s *setAccessTokenTestSuite) TestSetAccessToken() {
	s.clt.SetAccessToken(s.tkn)

	s.Equal(s.tkn, s.clt.AccessToken)
}

func TestSetAccessToken(t *testing.T) {
	suite.Run(t, new(setAccessTokenTestSuite))
}
