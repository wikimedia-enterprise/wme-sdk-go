package api_test

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

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
	s.NotZero(clt.BaseURL)
	s.NotZero(clt.RealtimeURL)
}

func (s *newClientTestSuite) TestNewClientWithOpts() {
	httpClient := new(http.Client)
	downloadMinChunkSize := 100
	downloadChunkSize := 25
	downloadConcurrency := 2
	scannerBufferSize := 100
	baseURL := "https://foo.bar"
	realtimeURL := "https://foo.bar/realtime"
	opt := func(clt *api.Client) {
		clt.HTTPClient = httpClient
		clt.DownloadMinChunkSize = downloadMinChunkSize
		clt.DownloadChunkSize = downloadChunkSize
		clt.DownloadConcurrency = downloadConcurrency
		clt.ScannerBufferSize = scannerBufferSize
		clt.BaseURL = baseURL
		clt.RealtimeURL = realtimeURL
	}
	clt := api.NewClient(opt).(*api.Client)

	s.NotNil(clt)
	s.Equal(httpClient, clt.HTTPClient)
	s.Equal(downloadMinChunkSize, clt.DownloadMinChunkSize)
	s.Equal(downloadChunkSize, clt.DownloadChunkSize)
	s.Equal(downloadConcurrency, clt.DownloadConcurrency)
	s.Equal(scannerBufferSize, clt.ScannerBufferSize)
	s.Equal(baseURL, clt.BaseURL)
	s.Equal(realtimeURL, clt.RealtimeURL)
}

func TestNewClient(t *testing.T) {
	suite.Run(t, new(newClientTestSuite))
}

type readAllTestSuite struct {
	suite.Suite
	ctx context.Context
	rcr io.ReadCloser
	clt api.API
}

func (s *readAllTestSuite) SetupSuite() {
	var err error
	s.rcr, err = testData.Open("testdata/simplewiki_namespace_0.tar.gz")
	s.NoError(err)

	s.clt = api.NewClient()
}

func (s *readAllTestSuite) TearDownSuite() {
	s.rcr.Close()
}

func (s *readAllTestSuite) TestReadAll() {
	nmc := 0
	err := s.clt.ReadAll(s.ctx, s.rcr, func(art *api.Article) error {
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
	mtd string
}

func (s *baseEntityTestSuite) SetupSuite() {
	rtr := http.NewServeMux()
	var hdr func(w http.ResponseWriter, r *http.Request)

	switch s.mtd {
	case http.MethodHead:
		hdr = func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Length", "4827640")
			w.Header().Set("Content-Type", "binary/octet-stream")
			w.Header().Set("ETag", "528262227e37be50594b5a0ac0bcb752")
			w.Header().Set("Last-Modified", "Mon, 04 Sep 2023 11:08:50 UTC")
			w.WriteHeader(s.sts)
		}
	case http.MethodGet:
	default:
		fle, err := testData.Open(s.fph)

		if err != nil {
			log.Fatal(err)
		}

		dta, err := io.ReadAll(fle)

		if err != nil {
			log.Fatal(err)
		}

		hdr = func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(s.sts)
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write(dta)
		}
	}

	rtr.HandleFunc(fmt.Sprintf("/v2/%s", s.pth), hdr)

	s.ctx = context.Background()
	s.req = new(api.Request)
	s.srv = httptest.NewServer(rtr)
	s.clt = api.NewClient(func(clt *api.Client) {
		clt.BaseURL = fmt.Sprintf("%s/", s.srv.URL)
		clt.RealtimeURL = fmt.Sprintf("%s/", s.srv.URL)
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

type getNamespacesTestSuite struct {
	baseEntityTestSuite
}

func (s *getNamespacesTestSuite) SetupSuite() {
	s.pth = "namespaces"
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getNamespacesTestSuite) TestGetNamespaces() {
	nss, err := s.clt.GetNamespaces(s.ctx, s.req)

	for _, nsp := range nss {
		if nsp.Identifier != 0 {
			s.NotEmpty(nsp.Identifier)
		}

		s.NotEmpty(nsp.Name)
		s.NotEmpty(nsp.Description)
	}

	if s.sts != http.StatusOK {
		s.Empty(nss)
		s.Error(err)
	} else {
		s.NotEmpty(nss)
		s.NoError(err)
	}
}

func TestGetNamespaces(t *testing.T) {
	for _, tcs := range []*getNamespacesTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/namespaces.json",
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

type getNamespaceTestSuite struct {
	baseEntityTestSuite
	idr int
}

func (s *getNamespaceTestSuite) SetupSuite() {
	s.idr = 14
	s.pth = fmt.Sprintf("namespaces/%d", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getNamespaceTestSuite) TestGetNamespace() {
	nsp, err := s.clt.GetNamespace(s.ctx, s.idr, s.req)

	if s.sts != http.StatusOK {
		s.Empty(nsp.Identifier)
		s.Empty(nsp.Name)
		s.Empty(nsp.Description)
		s.Error(err)
	} else {
		s.NotEmpty(nsp.Identifier)
		s.NotEmpty(nsp.Name)
		s.NotEmpty(nsp.Description)
		s.NoError(err)
	}
}

func TestGetNamespace(t *testing.T) {
	for _, tcs := range []*getNamespaceTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/namespace.json",
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

type getBatchesTestSuite struct {
	baseEntityTestSuite
	dte *time.Time
}

func (s *getBatchesTestSuite) SetupSuite() {
	dtn := time.Now()
	s.dte = &dtn
	s.pth = fmt.Sprintf("batches/%s", s.dte.Format(api.DateFormat))
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getBatchesTestSuite) TestGetBatches() {
	bts, err := s.clt.GetBatches(s.ctx, s.dte, s.req)

	for _, bth := range bts {
		s.NotEmpty(bth.Identifier)
		s.NotEmpty(bth.Version)
		s.NotEmpty(bth.DateModified)
		s.NotEmpty(bth.IsPartOf)
		s.NotEmpty(bth.InLanguage)
		s.NotNil(bth.Namespace)
		s.NotEmpty(bth.Size)
	}

	if s.sts != http.StatusOK {
		s.Empty(bts)
		s.Error(err)
	} else {
		s.NotEmpty(bts)
		s.NoError(err)
	}
}

func TestGetBatches(t *testing.T) {
	for _, tcs := range []*getBatchesTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/batches.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type getBatchTestSuite struct {
	baseEntityTestSuite
	dte *time.Time
	idr string
}

func (s *getBatchTestSuite) SetupSuite() {
	dtn := time.Now()
	s.dte = &dtn
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("batches/%s/%s", s.dte.Format(api.DateFormat), s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getBatchTestSuite) TestGetBatch() {
	bth, err := s.clt.GetBatch(s.ctx, s.dte, s.idr, s.req)

	if s.sts != http.StatusOK {
		s.Empty(bth)
		s.Error(err)
	} else {
		s.NotEmpty(bth.Identifier)
		s.NotEmpty(bth.Version)
		s.NotEmpty(bth.DateModified)
		s.NotEmpty(bth.IsPartOf)
		s.NotEmpty(bth.InLanguage)
		s.NotNil(bth.Namespace)
		s.NotEmpty(bth.Size)
		s.NoError(err)
	}
}

func TestGetBatch(t *testing.T) {
	for _, tcs := range []*getBatchTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/batch.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type headBatchTestSuite struct {
	baseEntityTestSuite
	idr string
	dte *time.Time
}

func (s *headBatchTestSuite) SetupSuite() {
	dtn := time.Now()
	s.dte = &dtn
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("batches/%s/%s/download", s.dte.Format(api.DateFormat), s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *headBatchTestSuite) TestHeadBatch() {
	bth, err := s.clt.HeadBatch(s.ctx, s.dte, s.idr)

	if s.sts != http.StatusOK {
		s.Empty(bth)
		s.Error(err)
	} else {
		s.NotEmpty(bth)
		s.NotEmpty(bth.ContentLength)
		s.NotEmpty(bth.ContentType)
		s.NotEmpty(bth.ETag)
		s.NotEmpty(bth.LastModified)
		s.NoError(err)
	}
}

func TestHeadBatch(t *testing.T) {
	for _, tcs := range []*headBatchTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				mtd: http.MethodHead,
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type downloadBatchTestSuite struct {
	baseEntityTestSuite
	idr string
	dte *time.Time
}

func (s *downloadBatchTestSuite) SetupSuite() {
	dtn := time.Now()
	s.dte = &dtn
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("batches/%s/%s/download", s.dte.Format(api.DateFormat), s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *downloadBatchTestSuite) TestDownloadBatch() {
	tmf, err := os.CreateTemp("", "bth_tmp.tar.gz")

	if err != nil {
		log.Fatal(err)
	}

	defer tmf.Close()
	err = s.clt.DownloadBatch(s.ctx, s.dte, s.idr, tmf)

	if s.sts != http.StatusOK {
		s.Error(err)
	} else {
		s.NoError(err)
	}
}

func TestDownloadBatch(t *testing.T) {
	for _, tcs := range []*downloadBatchTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/simplewiki_namespace_0.tar.gz",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type readBathTestSuite struct {
	baseEntityTestSuite
	idr string
	dte *time.Time
}

func (s *readBathTestSuite) SetupSuite() {
	dtn := time.Now()
	s.dte = &dtn
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("batches/%s/%s/download", s.dte.Format(api.DateFormat), s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *readBathTestSuite) TestReadBatch() {
	nmc := 0
	err := s.clt.ReadBatch(s.ctx, s.dte, s.idr, func(art *api.Article) error {
		s.NotEmpty(art.Name)
		s.NotEmpty(art.Identifier)
		nmc++
		return nil
	})

	if s.sts != http.StatusOK {
		s.Error(err)
		s.Zero(nmc)
	} else {
		s.NoError(err)
		s.NotZero(nmc)
	}
}

func TestReadBatch(t *testing.T) {
	for _, tcs := range []*readBathTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/simplewiki_namespace_0.tar.gz",
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

type getSnapshotsTestSuite struct {
	baseEntityTestSuite
}

func (s *getSnapshotsTestSuite) SetupSuite() {
	s.pth = "snapshots"
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getSnapshotsTestSuite) TestGetSnapshots() {
	sps, err := s.clt.GetSnapshots(s.ctx, s.req)

	for _, spt := range sps {
		s.NotEmpty(spt.Identifier)
		s.NotEmpty(spt.Version)
		s.NotEmpty(spt.DateModified)
		s.NotEmpty(spt.IsPartOf)
		s.NotEmpty(spt.InLanguage)
		s.NotNil(spt.Namespace)
		s.NotEmpty(spt.Size)
	}

	if s.sts != http.StatusOK {
		s.Empty(sps)
		s.Error(err)
	} else {
		s.NotEmpty(sps)
		s.NoError(err)
	}
}

func TestGetSnapshots(t *testing.T) {
	for _, tcs := range []*getSnapshotsTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/snapshots.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type getSnapshotTestSuite struct {
	baseEntityTestSuite
	idr string
}

func (s *getSnapshotTestSuite) SetupSuite() {
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("snapshots/%s", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getSnapshotTestSuite) TestGetSnapshot() {
	spt, err := s.clt.GetSnapshot(s.ctx, s.idr, s.req)

	if s.sts != http.StatusOK {
		s.Empty(spt.Identifier)
		s.Empty(spt.Version)
		s.Empty(spt.DateModified)
		s.Empty(spt.IsPartOf)
		s.Empty(spt.InLanguage)
		s.Nil(spt.Namespace)
		s.Empty(spt.Size)
		s.Error(err)
	} else {
		s.NotEmpty(spt.Identifier)
		s.NotEmpty(spt.Version)
		s.NotEmpty(spt.DateModified)
		s.NotEmpty(spt.IsPartOf)
		s.NotEmpty(spt.InLanguage)
		s.NotNil(spt.Namespace)
		s.NotEmpty(spt.Size)
		s.NoError(err)
	}
}

func TestGetSnapshot(t *testing.T) {
	for _, tcs := range []*getSnapshotTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/snapshot.json",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type headSnapshotTestSuite struct {
	baseEntityTestSuite
	idr string
}

func (s *headSnapshotTestSuite) SetupSuite() {
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("snapshots/%s/download", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *headSnapshotTestSuite) TestHeadSnapshot() {
	shs, err := s.clt.HeadSnapshot(s.ctx, s.idr)

	if s.sts != http.StatusOK {
		s.Empty(shs)
		s.Error(err)
	} else {
		s.NotEmpty(shs)
		s.NotEmpty(shs.ContentLength)
		s.NotEmpty(shs.ContentType)
		s.NotEmpty(shs.ETag)
		s.NotEmpty(shs.LastModified)
		s.NoError(err)
	}
}

func TestHeadSnapshot(t *testing.T) {
	for _, tcs := range []*headSnapshotTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				mtd: http.MethodHead,
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type downloadSnapshotTestSuite struct {
	baseEntityTestSuite
	idr string
}

func (s *downloadSnapshotTestSuite) SetupSuite() {
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("snapshots/%s/download", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *downloadSnapshotTestSuite) TestDownloadSnapshot() {
	tmf, err := os.CreateTemp("", "spt_tmp.tar.gz")

	if err != nil {
		log.Fatal(err)
	}

	defer tmf.Close()
	err = s.clt.DownloadSnapshot(s.ctx, s.idr, tmf)

	if s.sts != http.StatusOK {
		s.Error(err)
	} else {
		s.NoError(err)
	}
}

func TestDownloadSnapshot(t *testing.T) {
	for _, tcs := range []*downloadSnapshotTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/simplewiki_namespace_0.tar.gz",
			},
		},
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusNotFound,
				fph: "testdata/error.json",
			},
		},
	} {
		suite.Run(t, tcs)
	}
}

type readSnapshotTestSuite struct {
	baseEntityTestSuite
	idr string
}

func (s *readSnapshotTestSuite) SetupSuite() {
	s.idr = "simplewiki_namespace_0"
	s.pth = fmt.Sprintf("snapshots/%s/download", s.idr)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *readSnapshotTestSuite) TestReadSnapshot() {
	nmc := 0
	err := s.clt.ReadSnapshot(s.ctx, s.idr, func(art *api.Article) error {
		s.NotEmpty(art.Name)
		s.NotEmpty(art.Identifier)
		nmc++
		return nil
	})

	if s.sts != http.StatusOK {
		s.Error(err)
		s.Zero(nmc)
	} else {
		s.NoError(err)
		s.NotZero(nmc)
	}
}

func TestReadSnapshot(t *testing.T) {
	for _, tcs := range []*readSnapshotTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/simplewiki_namespace_0.tar.gz",
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

type getArticlesTestSuite struct {
	baseEntityTestSuite
	nme string
}

func (s *getArticlesTestSuite) SetupSuite() {
	s.nme = "Squirrel"
	s.pth = fmt.Sprintf("articles/%s", s.nme)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getArticlesTestSuite) TestGetArticles() {
	ars, err := s.clt.GetArticles(s.ctx, s.nme, s.req)

	for _, art := range ars {
		s.NotEmpty(art.Name)
		s.NotEmpty(art.Identifier)
		s.NotEmpty(art.Abstract)
		s.NotEmpty(art.URL)
	}

	if s.sts != http.StatusOK {
		s.Empty(ars)
		s.Error(err)
	} else {
		s.NotEmpty(ars)
		s.NoError(err)
	}
}

func TestGetArticles(t *testing.T) {
	for _, tcs := range []*getArticlesTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/articles.json",
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

type getStructuredContentsTestSuite struct {
	baseEntityTestSuite
	nme string
}

func (s *getStructuredContentsTestSuite) SetupSuite() {
	s.nme = "Squirrel"
	s.pth = fmt.Sprintf("structured-contents/%s", s.nme)
	s.baseEntityTestSuite.SetupSuite()
}

func (s *getStructuredContentsTestSuite) TestGetStructuredContents() {
	scs, err := s.clt.GetStructuredContents(s.ctx, s.nme, s.req)

	for _, sct := range scs {
		s.NotEmpty(sct.Name)
		s.NotEmpty(sct.Identifier)
		s.NotEmpty(sct.URL)
		s.NotEmpty(sct.Abstract)
	}

	if s.sts != http.StatusOK {
		s.Empty(scs)
		s.Error(err)
	} else {
		s.NotEmpty(scs)
		s.NoError(err)
	}
}

func TestGetStructuredContents(t *testing.T) {
	for _, tcs := range []*getStructuredContentsTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/structured-contents.json",
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

type streamArticlesTestSuite struct {
	baseEntityTestSuite
}

func (s *streamArticlesTestSuite) SetupSuite() {
	s.pth = "articles"
	s.baseEntityTestSuite.SetupSuite()
}

func (s *streamArticlesTestSuite) TestStreamArticles() {
	nmc := 0
	err := s.clt.StreamArticles(s.ctx, s.req, func(art *api.Article) error {
		s.NotEmpty(art.Name)
		s.NotEmpty(art.Identifier)
		nmc++
		return nil
	})

	if s.sts != http.StatusOK {
		s.Error(err)
		s.Zero(nmc)
	} else {
		s.NoError(err)
		s.NotZero(nmc)
	}
}

func TestStreamArticles(t *testing.T) {
	for _, tcs := range []*streamArticlesTestSuite{
		{
			baseEntityTestSuite: baseEntityTestSuite{
				sts: http.StatusOK,
				fph: "testdata/articles.ndjson",
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
