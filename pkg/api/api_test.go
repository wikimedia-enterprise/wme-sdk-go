package api_test

import (
	"context"
	"embed"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/api"
)

//go:embed testdata/*
var testData embed.FS

type newClientTestSuite struct {
	suite.Suite
}

func (s *newClientTestSuite) TestNewClient() {
	clt := api.NewClient().(*api.Client)

	s.NotNil(clt)
	s.NotNil(clt.HTTPClient)
	s.NotZero(clt.DownloadMinChunkSize)
	s.NotZero(clt.DownloadChunkSize)
	s.NotZero(clt.DownloadConcurrency)
	s.NotZero(clt.ScannerBufferSize)
	s.NotZero(clt.BaseUrl)
	s.NotZero(clt.RealtimeURL)
}

func (s *newClientTestSuite) TestNewClientWithOpts() {
	httpClient := new(http.Client)
	downloadMinChunkSize := 100
	downloadChunkSize := 25
	downloadConcurrency := 2
	scannerBufferSize := 100
	baseUrl := "https://foo.bar"
	realtimeURL := "https://foo.bar/realtime"
	opt := func(clt *api.Client) {
		clt.HTTPClient = httpClient
		clt.DownloadMinChunkSize = downloadMinChunkSize
		clt.DownloadChunkSize = downloadChunkSize
		clt.DownloadConcurrency = downloadConcurrency
		clt.ScannerBufferSize = scannerBufferSize
		clt.BaseUrl = baseUrl
		clt.RealtimeURL = realtimeURL
	}
	clt := api.NewClient(opt).(*api.Client)

	s.NotNil(clt)
	s.Equal(httpClient, clt.HTTPClient)
	s.Equal(downloadMinChunkSize, clt.DownloadMinChunkSize)
	s.Equal(downloadChunkSize, clt.DownloadChunkSize)
	s.Equal(downloadConcurrency, clt.DownloadConcurrency)
	s.Equal(scannerBufferSize, clt.ScannerBufferSize)
	s.Equal(baseUrl, clt.BaseUrl)
	s.Equal(realtimeURL, clt.RealtimeURL)
}

func TestNewClient(t *testing.T) {
	suite.Run(t, new(newClientTestSuite))
}

type readAllTestSuite struct {
	suite.Suite
	ctx context.Context
	rdr io.ReadCloser
	clt api.API
}

func (s *readAllTestSuite) SetupTest() {
	var err error
	s.rdr, err = testData.Open("testdata/simplewiki_namespace_0.tar.gz")
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
	suite.Run(t, new(readAllTestSuite))
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
