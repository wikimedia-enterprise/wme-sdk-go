package api_test

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
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

type baseEntityTestSuite struct {
	suite.Suite
	sts int
	fph string
	pth string
	ctx context.Context
	req *api.Request
	srv *httptest.Server
	clt api.API
}

func (s *baseEntityTestSuite) SetupSuite() {
	fle, err := testData.Open(s.fph)

	if err != nil {
		log.Fatal(err)
	}

	dta, err := io.ReadAll(fle)

	if err != nil {
		log.Fatal(err)
	}

	rtr := http.NewServeMux()
	rtr.HandleFunc(fmt.Sprintf("/v2/%s", s.pth), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(s.sts)
		w.Write(dta)
	})

	s.ctx = context.Background()
	s.req = new(api.Request)
	s.srv = httptest.NewServer(rtr)
	s.clt = api.NewClient(func(clt *api.Client) {
		clt.BaseUrl = fmt.Sprintf("%s/", s.srv.URL)
	})
}

func (s *baseEntityTestSuite) TearDownSuite() {
	s.srv.Close()
}

type getCodesTestSuite struct {
	baseEntityTestSuite
}

func (s *getCodesTestSuite) SetupSuite() {
	s.pth = "codes"
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getCodesTestSuite) TestGetCodes() {
	cds, err := s.clt.GetCodes(s.ctx, s.req)

	for _, cde := range cds {
		s.NotEmpty(cde.Identifier)
		s.NotEmpty(cde.Name)
		s.NotEmpty(cde.Description)
	}

	if s.sts != http.StatusOK {
		s.Empty(cds)
		s.Error(err)
	} else {
		s.NotEmpty(cds)
		s.NoError(err)
	}
}

func TestGetCodes(t *testing.T) {
	for _, tcs := range []*getCodesTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/codes.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type getCodeTestSuite struct {
	baseEntityTestSuite
	idr string
}

func (s *getCodeTestSuite) SetupSuite() {
	s.idr = "simplewiki_namepace_0"
	s.pth = fmt.Sprintf("codes/%s", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getCodeTestSuite) TestGetCode() {
	cde, err := s.clt.GetCode(s.ctx, s.idr, s.req)

	if s.sts != http.StatusOK {
		s.Empty(cde.Identifier)
		s.Empty(cde.Name)
		s.Empty(cde.Description)
		s.Error(err)
	} else {
		s.NotEmpty(cde.Identifier)
		s.NotEmpty(cde.Name)
		s.NotEmpty(cde.Description)
		s.NoError(err)
	}
}

func TestGetCode(t *testing.T) {
	for _, tcs := range []*getCodeTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/code.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type getLanguagesTestSuite struct {
	baseEntityTestSuite
}

func (s *getLanguagesTestSuite) SetupSuite() {
	s.pth = "languages"
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getLanguagesTestSuite) TestGetLanguages() {
	lgs, err := s.clt.GetLanguages(s.ctx, s.req)

	for _, lng := range lgs {
		s.NotEmpty(lng.Identifier)
		s.NotEmpty(lng.Name)
		s.NotEmpty(lng.AlternateName)
		s.NotEmpty(lng.Direction)
	}

	if s.sts != http.StatusOK {
		s.Empty(lgs)
		s.Error(err)
	} else {
		s.NotEmpty(lgs)
		s.NoError(err)
	}
}

func TestGetLanguages(t *testing.T) {
	for _, tcs := range []*getLanguagesTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/languages.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type getLanguageTestSuite struct {
	baseEntityTestSuite
	idr string
}

func (s *getLanguageTestSuite) SetupSuite() {
	s.idr = "en"
	s.pth = fmt.Sprintf("languages/%s", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getLanguageTestSuite) TestGetLanguage() {
	lng, err := s.clt.GetLanguage(s.ctx, s.idr, s.req)

	if s.sts != http.StatusOK {
		s.Empty(lng.Identifier)
		s.Empty(lng.Name)
		s.Empty(lng.AlternateName)
		s.Empty(lng.Direction)
		s.Error(err)
	} else {
		s.NotEmpty(lng.Identifier)
		s.NotEmpty(lng.Name)
		s.NotEmpty(lng.AlternateName)
		s.NotEmpty(lng.Direction)
		s.NoError(err)
	}
}

func TestGetLanguage(t *testing.T) {
	for _, tcs := range []*getLanguageTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/language.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type getProjectsTestSuite struct {
	baseEntityTestSuite
}

func (s *getProjectsTestSuite) SetupSuite() {
	s.pth = "projects"
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getProjectsTestSuite) TestGetProjects() {
	prs, err := s.clt.GetProjects(s.ctx, s.req)

	for _, prj := range prs {
		s.NotEmpty(prj.Name)
		s.NotEmpty(prj.Identifier)
		s.NotEmpty(prj.URL)
		s.NotEmpty(prj.Code)
		s.NotNil(prj.InLanguage)
		s.NotEmpty(prj.InLanguage.Identifier)
	}

	if s.sts != http.StatusOK {
		s.Empty(prs)
		s.Error(err)
	} else {
		s.NotEmpty(prs)
		s.NoError(err)
	}
}

func TestGetProjects(t *testing.T) {
	for _, tcs := range []*getProjectsTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/projects.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type getProjectTestSuite struct {
	baseEntityTestSuite
	idr string
}

func (s *getProjectTestSuite) SetupSuite() {
	s.idr = "enwiki"
	s.pth = fmt.Sprintf("projects/%s", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getProjectTestSuite) TestGetProject() {
	prj, err := s.clt.GetProject(s.ctx, s.idr, s.req)

	if s.sts != http.StatusOK {
		s.Empty(prj.Name)
		s.Empty(prj.Identifier)
		s.Empty(prj.URL)
		s.Empty(prj.Code)
		s.Nil(prj.InLanguage)
		s.Error(err)
	} else {
		s.NotEmpty(prj.Name)
		s.NotEmpty(prj.Identifier)
		s.NotEmpty(prj.URL)
		s.NotEmpty(prj.Code)
		s.NotNil(prj.InLanguage)
		s.NotEmpty(prj.InLanguage.Identifier)
		s.NoError(err)
	}
}

func TestGetProject(t *testing.T) {
	for _, tcs := range []*getProjectTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/project.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusUnauthorized,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}
