package api_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/wikimedia-enterprise/wme-sdk-go/pkg/api"
)

func createHandler(sts int, dta string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(sts)
		_, _ = w.Write([]byte(dta))
	}
}

type apiTestSuite struct {
	suite.Suite
	clt api.API
	ctx context.Context
	sra *httptest.Server
	srr *httptest.Server
	req *api.Request
	cid string
	cds string
	cde string
	lid string
	lgs string
	lgg string
	pid string
	pgs string
	pjt string
	nid int
	nss string
	nsp string
	bdt string
	bid string
	bts string
	bth string
	sid string
	sps string
	spt string
	anm string
	ats string
	sts int
	act string
	err error
	snm string
	sns string
}

func (s *apiTestSuite) createAPIServer() http.Handler {
	rtr := http.NewServeMux()

	rtr.HandleFunc("/v2/codes", createHandler(s.sts, s.cds))
	rtr.HandleFunc(fmt.Sprintf("/v2/codes/%s", s.cid), createHandler(s.sts, s.cde))

	rtr.HandleFunc("/v2/languages", createHandler(s.sts, s.lgs))
	rtr.HandleFunc(fmt.Sprintf("/v2/languages/%s", s.lid), createHandler(s.sts, s.lgg))

	rtr.HandleFunc("/v2/projects", createHandler(s.sts, s.pgs))
	rtr.HandleFunc(fmt.Sprintf("/v2/projects/%s", s.pid), createHandler(s.sts, s.pjt))

	rtr.HandleFunc("/v2/namespaces", createHandler(s.sts, s.nss))
	rtr.HandleFunc(fmt.Sprintf("/v2/namespaces/%d", s.nid), createHandler(s.sts, s.nsp))

	rtr.HandleFunc(fmt.Sprintf("/v2/batches/%s", s.bdt), createHandler(s.sts, s.bts))
	rtr.HandleFunc(fmt.Sprintf("/v2/batches/%s/%s", s.bdt, s.bid), createHandler(s.sts, s.bth))

	rtr.HandleFunc("/v2/snapshots", createHandler(s.sts, s.sps))
	rtr.HandleFunc(fmt.Sprintf("/v2/snapshots/%s", s.sid), createHandler(s.sts, s.spt))

	rtr.HandleFunc(fmt.Sprintf("/v2/articles/%s", s.anm), createHandler(s.sts, s.ats))
	rtr.HandleFunc(fmt.Sprintf("/v2/structured-contents/%s", s.snm), createHandler(s.sts, s.sns))

	return rtr
}

func (s *apiTestSuite) createRealtimeServer() http.Handler {
	rtr := http.NewServeMux()
	return rtr
}

func (s *apiTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.sra = httptest.NewServer(s.createAPIServer())
	s.srr = httptest.NewServer(s.createRealtimeServer())
	s.clt = api.NewClient(func(clt *api.Client) {
		clt.BaseUrl = s.sra.URL
		clt.BaseUrl = fmt.Sprintf("%s/", s.sra.URL)
		clt.RealtimeURL = fmt.Sprintf("%s/", s.srr.URL)
	})
}

func (s *apiTestSuite) TearDownSuite() {
	s.sra.Close()
	s.srr.Close()
}

func (s *apiTestSuite) TestSetAccessToken() {
	s.clt.SetAccessToken(s.act)

	s.Assert().Equal(s.act, s.clt.(*api.Client).AccessToken)
}

func (s *apiTestSuite) TestGetCodes() {
	cds, err := s.clt.GetCodes(s.ctx, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(cds)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(cds)
	}
}

func (s *apiTestSuite) TestGetCode() {
	cde, err := s.clt.GetCode(s.ctx, s.cid, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(cde)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(cde)
	}
}

func (s *apiTestSuite) TestGetLanguages() {
	lgs, err := s.clt.GetLanguages(s.ctx, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(lgs)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(lgs)
	}
}

func (s *apiTestSuite) TestGetLanguage() {
	lgn, err := s.clt.GetLanguage(s.ctx, s.lid, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(lgn)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(lgn)
	}
}

func (s *apiTestSuite) TestGetProjects() {
	pgs, err := s.clt.GetProjects(s.ctx, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(pgs)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(pgs)
	}
}

func (s *apiTestSuite) TestGetProject() {
	pjt, err := s.clt.GetProject(s.ctx, s.pid, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(pjt)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(pjt)
	}
}

func (s *apiTestSuite) TestGetNamespaces() {
	nss, err := s.clt.GetNamespaces(s.ctx, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(nss)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(nss)
	}
}

func (s *apiTestSuite) TestGetNamespace() {
	nps, err := s.clt.GetNamespace(s.ctx, s.nid, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(nps)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(nps)
	}
}

func (s *apiTestSuite) TestGetBatches() {
	dte, _ := time.Parse(api.DateFormat, s.bdt)
	bts, err := s.clt.GetBatches(s.ctx, &dte, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(bts)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(bts)
	}
}

func (s *apiTestSuite) TestGetBatch() {
	dte, _ := time.Parse(api.DateFormat, s.bdt)
	bth, err := s.clt.GetBatch(s.ctx, &dte, s.bid, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(bth)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(bth)
	}
}

func (s *apiTestSuite) TestGetSnapshots() {
	sps, err := s.clt.GetSnapshots(s.ctx, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(sps)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(sps)
	}
}

func (s *apiTestSuite) TestGetSnapshot() {
	spt, err := s.clt.GetSnapshot(s.ctx, s.sid, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(spt)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(spt)
	}
}

func (s *apiTestSuite) TestGetArticles() {
	ats, err := s.clt.GetArticles(s.ctx, s.anm, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(ats)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(ats)
	}
}

func (s *apiTestSuite) TestGetStructuredContents() {
	sts, err := s.clt.GetStructuredContents(s.ctx, s.snm, s.req)

	if s.err != nil {
		s.Assert().Error(err)
		s.Assert().Empty(sts)
	} else {
		s.Assert().NoError(err)
		s.Assert().NotEmpty(sts)
	}
}

func TestAPI(t *testing.T) {
	for _, testCase := range []*apiTestSuite{
		{
			act: "test",
			cid: "wiki",
			cds: `[
				{
					"identifier": "wiki",
					"name": "Wikipedia",
					"description": "..."
				},
				{
					"identifier": "wikibooks",
					"name": "Wikibooks",
					"description": "..."
				}
			]`,
			cde: `{
				"identifier": "wiki",
				"name": "Wikipedia",
				"description": "..."
			}`,
			lid: "en",
			lgs: `[
				{
					"identifier": "en",
					"name": "...",
					"alternate_name": "...",
					"direction": "ltr"
				},
				{
					"identifier": "ff",
					"name": "...",
					"alternate_name": "...",
					"direction": "ltr"
				}
			]`,
			lgg: `{
				"identifier": "en",
				"name": "...",
				"alternate_name": "...",
				"direction": "ltr"
			}`,
			pid: "enwiki",
			pgs: `[
				{
					"name": "...",
					"identifier": "fiwikivoyage",
					"url": "https://fi.wikivoyage.org",
					"code": "wikivoyage",
					"in_language": {
						"identifier": "fi"
					}
				},
				{
					"name": "...",
					"identifier": "enwiki",
					"url": "https://en.wikipedia.org",
					"code": "wiki",
					"in_language": {
						"identifier": "en"
					}
				}
			]`,
			pjt: `{
				"name": "...",
				"identifier": "enwiki",
				"url": "https://en.wikipedia.org",
				"code": "wiki",
				"in_language": {
					"identifier": "en"
				}
			}`,
			nid: 6,
			nss: `[
				{
					"name": "...",
					"identifier": 6,
					"description": "..."
				},
				{
					"name": "...",
					"identifier": 0,
					"description": "..."
				}
			]`,
			nsp: `{
				"name": "...",
				"identifier": 6,
				"description": "..."
			}`,
			bdt: "2023-02-28",
			bid: "enwiki_namespace_0",
			bts: `[
				{
					"identifier": "enwiki_namespace_0",
					"version": "f9bc4266b42ac15a3a7f881c6e38aed7",
					"date_modified": "2023-02-28T19:15:24.182031535Z",
					"is_part_of": {
						"identifier": "abwiki"
					},
					"in_language": {
						"identifier": "ab"
					},
					"namespace": {
						"identifier": 0
					},
					"size": {
						"value": 0.002,
						"unit_text": "MB"
					}
				},
				{
					"identifier": "abwiki_namespace_10",
					"version": "05f21bfaa910a50fad318a8fc2c0ccb5",
					"date_modified": "2023-02-28T19:15:26.611157307Z",
					"is_part_of": {
						"identifier": "abwiki"
					},
					"in_language": {
						"identifier": "ab"
					},
					"namespace": {
						"identifier": 10
					},
					"size": {
						"value": 0.005,
						"unit_text": "MB"
					}
				}
			]`,
			bth: `{
				"identifier": "enwiki_namespace_0",
				"version": "f9bc4266b42ac15a3a7f881c6e38aed7",
				"date_modified": "2023-02-28T19:15:24.182031535Z",
				"is_part_of": {
					"identifier": "abwiki"
				},
				"in_language": {
					"identifier": "ab"
				},
				"namespace": {
					"identifier": 0
				},
				"size": {
					"value": 0.002,
					"unit_text": "MB"
				}
			}`,
			sid: "abwiki_namespace_10",
			sps: `[
				{
					"identifier": "abwiki_namespace_0",
					"version": "93ec1fbd34b79e7c0302d5f2691445ab",
					"date_modified": "2023-02-28T02:24:03.458822229Z",
					"is_part_of": {
						"identifier": "abwiki"
					},
					"in_language": {
						"identifier": "ab"
					},
					"namespace": {
						"identifier": 0
					},
					"size": {
						"value": 14.561,
						"unit_text": "MB"
					}
				},
				{
					"identifier": "abwiki_namespace_10",
					"version": "cae69fac8c2f3c980ce0fb3623191245",
					"date_modified": "2023-02-28T02:24:08.184474026Z",
					"is_part_of": {
						"identifier": "abwiki"
					},
					"in_language": {
						"identifier": "ab"
					},
					"namespace": {
						"identifier": 10
					},
					"size": {
						"value": 11.482,
						"unit_text": "MB"
					}
				}
			]`,
			spt: `{
				"identifier": "abwiki_namespace_10",
				"version": "cae69fac8c2f3c980ce0fb3623191245",
				"date_modified": "2023-02-28T02:24:08.184474026Z",
				"is_part_of": {
					"identifier": "abwiki"
				},
				"in_language": {
					"identifier": "ab"
				},
				"namespace": {
					"identifier": 10
				},
				"size": {
					"value": 11.482,
					"unit_text": "MB"
				}
			}`,
			anm: "Earth",
			ats: `[
				{
					"name": "Earth",
					"is_part_of": {
						"identifier": "enwiki"
					}
				},
				{
					"name": "Earth",
					"is_part_of": {
						"identifier": "enwikinews"
					}
				}
			]`,
			snm: "Ninja",
			sns: `[
				{
					"name": "Ninja",
					"is_part_of": {
						"identifier": "enwiki"
					}
				},
				{
					"name": "Ninja",
					"is_part_of": {
						"identifier": "enwikinews"
					}
				}
			]`,
			sts: http.StatusOK,
		},
	} {
		suite.Run(t, testCase)
	}
}
