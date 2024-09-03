// Package api holds an API client for Wikimedia Enterprise API(s).
package api

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/klauspost/pgzip"
)

// DateFormat is the date format used for the API.
const DateFormat = "2006-01-02"

// Filter represents a filter to be applied to a dataset.
type Filter struct {
	// Field specifies the field in the dataset that the filter should be applied to.
	Field string `json:"field"`

	// Value specifies the value that the field should be compared to.
	Value interface{} `json:"value"`
}

// ReadCallback is a function that will be called with each Article object that is read from a batch or snapshot.
// You can return a custom error to stop the reading.
type ReadCallback func(art *Article) error

// Request contains properties that are used to apply filters to the API.
type Request struct {
	// Since is a parameter used only for streaming endpoints.
	// Will pick up the reading of stream from this timestamp.
	// For the articles endpoint it will be restricted to 48h.
	Since *time.Time `json:"since,omitempty"`

	// Fields represents a list of fields to retrieve from the API.
	// This is an optional argument.
	Fields []string `json:"fields,omitempty"`

	// Filters represents a list of filters to apply to the response.
	// This is an optional argument.
	Filters []*Filter `json:"filters,omitempty"`

	// Limits the amount of results from the API (for now works only with Articles API).
	// This is an optional argument.
	Limit int `json:"limit,omitempty"`

	// Provides a way to open parallel connections to realtime streaming API.
	// Allows to target subsets of partitions in each of the parallel connections.
	// The max allowed number of parallel connections to realtime API is 10, i.e., the allowed range for parts is 0 through 9.
	// Each part value lets one connect to 1/10 th of the total partitions.
	// e.g., [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 ]
	Parts []int `json:"parts,omitempty"`

	// Used for reconnection to realtime streaming API by passing this parameter.
	// This is map of partition:latest offset consumed.
	Offsets map[int]int64 `json:"offsets,omitempty"`

	// Used for reconnection to realtime streaming API by passing this parameter.
	// This is map of partition:latest event.date_published consumed.
	SincePerPartition map[int]time.Time ` json:"since_per_partition,omitempty"`
}

// CodesGetter is an interface that retrieves codes from the API.
type CodesGetter interface {
	GetCodes(ctx context.Context, req *Request) ([]*Code, error)
}

// CodeGetter is an interface that retrieves a code by ID from the API.
type CodeGetter interface {
	GetCode(ctx context.Context, idr string, req *Request) (*Code, error)
}

// LanguagesGetter is an interface that retrieves languages from the API.
type LanguagesGetter interface {
	GetLanguages(ctx context.Context, req *Request) ([]*Language, error)
}

// LanguageGetter is an interface that retrieves a language by ID from the API.
type LanguageGetter interface {
	GetLanguage(ctx context.Context, idr string, req *Request) (*Language, error)
}

// ProjectsGetter is an interface that retrieves projects from the API.
type ProjectsGetter interface {
	GetProjects(ctx context.Context, req *Request) ([]*Project, error)
}

// ProjectGetter is an interface that retrieves a project by ID from the API.
type ProjectGetter interface {
	GetProject(ctx context.Context, idr string, req *Request) (*Project, error)
}

// NamespacesGetter is an interface that retrieves namespaces from the API.
type NamespacesGetter interface {
	GetNamespaces(ctx context.Context, req *Request) ([]*Namespace, error)
}

// NamespaceGetter is an interface that retrieves a namespace by ID from the API.
type NamespaceGetter interface {
	GetNamespace(ctx context.Context, idr int, req *Request) (*Namespace, error)
}

// BatchesGetter is an interface that retrieves batches from the API.
type BatchesGetter interface {
	GetBatches(ctx context.Context, dte *time.Time, req *Request) ([]*Batch, error)
}

// BatchGetter is an interface that retrieves a realtime batch by ID from the API.
type BatchGetter interface {
	GetBatch(ctx context.Context, dte *time.Time, idr string, req *Request) (*Batch, error)
}

// BatchHeader is an interface that retrieves the header of a realtime batch by ID from the API.
type BatchHeader interface {
	HeadBatch(ctx context.Context, dte *time.Time, idr string) (*Headers, error)
}

// BatchReader is an interface that reads a realtime batch data by ID from the API.
type BatchReader interface {
	ReadBatch(ctx context.Context, dte *time.Time, idr string, cbk ReadCallback) error
}

// BatchDownloader is an interface that downloads a realtime batch `tar.gz` by ID file from the API.
type BatchDownloader interface {
	DownloadBatch(ctx context.Context, dte *time.Time, idr string, wsk io.WriteSeeker) error
}

// SnapshotsGetter is an interface for getting multiple snapshots.
type SnapshotsGetter interface {
	GetSnapshots(ctx context.Context, req *Request) ([]*Snapshot, error)
}

// SnapshotGetter is an interface for getting a single snapshot by ID.
type SnapshotGetter interface {
	GetSnapshot(ctx context.Context, idr string, req *Request) (*Snapshot, error)
}

// SnapshotHeader is an interface for getting the headers of a single snapshot by ID.
type SnapshotHeader interface {
	HeadSnapshot(ctx context.Context, idr string) (*Headers, error)
}

// SnapshotDownloader is an interface for downloading a single snapshot by ID to a writer.
type SnapshotDownloader interface {
	DownloadSnapshot(ctx context.Context, idr string, wsk io.WriteSeeker) error
}

// SnapshotReader is an interface for reading the contents of a single snapshot by ID with a callback function.
type SnapshotReader interface {
	ReadSnapshot(ctx context.Context, idr string, cbk ReadCallback) error
}

// ArticlesGetter is an interface for getting a lits of articles by name.
type ArticlesGetter interface {
	GetArticles(ctx context.Context, nme string, req *Request) ([]*Article, error)
}

// StructuredContentsGetter is an interface for getting a lits of structured contents by name.
type StructuredContentsGetter interface {
	GetStructuredContents(ctx context.Context, nme string, req *Request) ([]*StructuredContent, error)
}

// ArticlesStreamer is an interface for getting all the article changes in realtime.
type ArticlesStreamer interface {
	StreamArticles(ctx context.Context, req *Request, cbk ReadCallback) error
}

// AllReader is an interface for reading all the contents of a reader with a callback function.
type AllReader interface {
	ReadAll(ctx context.Context, rdr io.Reader, cbk ReadCallback) error
}

// AccessTokenSetter is an interface for setting an access token.
type AccessTokenSetter interface {
	SetAccessToken(tkn string)
}

// ChunksGetter is an interface for getting multiple snapshot chunks by snapshot ID.
type ChunksGetter interface {
	GetChunks(ctx context.Context, sid string, req *Request) ([]*Snapshot, error)
}

// ChunkGetter is an interface for getting a single chunk by ID.
type ChunkGetter interface {
	GetChunk(ctx context.Context, sid string, idr string, req *Request) (*Snapshot, error)
}

// ChunksHeader is an interface for getting the headers of a single chunk by ID.
type ChunksHeader interface {
	HeadChunk(ctx context.Context, sid string, idr string) (*Headers, error)
}

// ChunksDownloader is an interface for downloading a single chunk by ID to a writer.
type ChunksDownloader interface {
	DownloadChunk(ctx context.Context, sid string, idr string, wsk io.WriteSeeker) error
}

// ChunksReader is an interface for reading the contents of a single chunk by ID with a callback function.
type ChunksReader interface {
	ReadChunk(ctx context.Context, sid string, idr string, cbk ReadCallback) error
}

// API interface tha encapsulates the whole functionality of the client.
// Can be used with composition in unit testing.
type API interface {
	AllReader
	AccessTokenSetter
	CodesGetter
	CodeGetter
	LanguagesGetter
	LanguageGetter
	ProjectsGetter
	ProjectGetter
	NamespacesGetter
	NamespaceGetter
	BatchesGetter
	BatchGetter
	BatchHeader
	BatchReader
	BatchDownloader
	SnapshotsGetter
	SnapshotGetter
	SnapshotHeader
	SnapshotReader
	SnapshotDownloader
	ChunksGetter
	ChunkGetter
	ChunksHeader
	ChunksDownloader
	ChunksReader
	ArticlesGetter
	StructuredContentsGetter
	ArticlesStreamer
}

// NewClient returns a new instance of the Client that implements the API interface.
// The function takes in optional functional options that allow the caller to configure
// the client with custom settings.
func NewClient(ops ...func(clt *Client)) API {
	clt := &Client{
		HTTPClient:           &http.Client{},
		DownloadMinChunkSize: 5242880,
		DownloadChunkSize:    5242880 * 5,
		DownloadConcurrency:  10,
		ScannerBufferSize:    20971520,
		UserAgent:            "",
		BaseURL:              "https://api.enterprise.wikimedia.com/",
		RealtimeURL:          "https://realtime.enterprise.wikimedia.com/",
	}

	for _, opt := range ops {
		opt(clt)
	}

	return clt
}

// Client is a struct that represents an HTTP client used to interact with the API.
type Client struct {
	// HTTPClient is the HTTP client used to send requests.
	HTTPClient *http.Client

	// UserAgent is the user-agent header value sent with each request.
	UserAgent string

	// BaseUrl is the base URL for all API requests.
	BaseURL string

	// RealtimeURL is the base URL for all realtime API requests.
	RealtimeURL string

	// AccessToken is the access token used to authenticate requests.
	AccessToken string

	// DownloadMinChunkSize is the minimum chunk size used for downloading resources.
	DownloadMinChunkSize int

	// DownloadChunkSize is the chunk size used for downloading resources.
	DownloadChunkSize int

	// DownloadConcurrency is the number of simultaneous downloads allowed.
	DownloadConcurrency int

	// ScannerBufferSize is the buffer size for the scanner when it reads from the API.
	ScannerBufferSize int
}

func (c *Client) newRequest(ctx context.Context, url string, mtd string, pth string, req *Request) (*http.Request, error) {
	dta := []byte{}

	if req != nil {
		bdy, err := json.Marshal(req)

		if err != nil {
			return nil, err
		}

		dta = bdy
	}

	hrq, err := http.NewRequestWithContext(ctx, mtd, fmt.Sprintf("%sv2/%s", url, pth), bytes.NewReader(dta))

	if err != nil {
		return nil, err
	}

	hrq.Header.Set("User-Agent", c.UserAgent)
	hrq.Header.Set("Content-Type", "application/json")
	hrq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	return hrq, nil
}

func (c *Client) do(hrq *http.Request) (*http.Response, error) {
	res, err := c.HTTPClient.Do(hrq)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		dta, err := io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			return nil, err
		}

		if len(string(dta)) == 0 {
			return nil, errors.New(res.Status)
		}

		return nil, errors.New(string(dta))
	}

	return res, nil
}

func (c *Client) getEntity(ctx context.Context, req *Request, pth string, val interface{}) error {
	hrq, err := c.newRequest(ctx, c.BaseURL, http.MethodPost, pth, req)

	if err != nil {
		return err
	}

	res, err := c.do(hrq)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(val)
}

func (c *Client) readLoop(ctx context.Context, rdr io.Reader, cbk ReadCallback) error {
	scn := bufio.NewScanner(rdr)
	scn.Buffer([]byte{}, c.ScannerBufferSize)

	for scn.Scan() {
		art := new(Article)

		if err := json.Unmarshal(scn.Bytes(), art); err != nil {
			return err
		}

		if err := cbk(art); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) readEntity(ctx context.Context, pth string, cbk ReadCallback) error {
	hrq, err := c.newRequest(ctx, c.BaseURL, http.MethodGet, pth, nil)

	if err != nil {
		return err
	}

	res, err := c.do(hrq)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	return c.ReadAll(ctx, res.Body, cbk)
}

func (c *Client) headEntity(ctx context.Context, pth string) (*Headers, error) {
	hrq, err := c.newRequest(ctx, c.BaseURL, http.MethodHead, pth, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.do(hrq)

	if err != nil {
		return nil, err
	}

	hdr := &Headers{
		ETag:         strings.Trim(res.Header.Get("ETag"), "\""),
		ContentType:  res.Header.Get("Content-Type"),
		AcceptRanges: res.Header.Get("Accept-Ranges"),
	}

	if lmf := res.Header.Get("Last-Modified"); len(lmf) > 0 {
		lmd, err := time.Parse(time.RFC1123, lmf)

		if err != nil {
			return nil, err
		}

		hdr.LastModified = &lmd
	}

	if ctl := res.Header.Get("Content-Length"); len(ctl) > 0 {
		cti, err := strconv.Atoi(ctl)

		if err != nil {
			return nil, err
		}

		hdr.ContentLength = cti
	}

	return hdr, nil
}

type chunk struct {
	start int
	end   int
	data  []byte
}

func (c *Client) downloadEntity(ctx context.Context, pth string, wrr io.WriteSeeker) error {
	hds, err := c.headEntity(ctx, pth)

	if err != nil {
		return err
	}

	csz := c.DownloadChunkSize

	if hds.ContentLength < c.DownloadMinChunkSize {
		csz = c.DownloadMinChunkSize
	}

	cks := []*chunk{}

	for i := 0; true; i++ {
		cnk := &chunk{
			start: i * csz,
			end:   (i * csz) + csz,
		}

		if cnk.end > hds.ContentLength {
			cnk.end = hds.ContentLength
		}

		cks = append(cks, cnk)

		if cnk.end == hds.ContentLength {
			break
		}
	}

	ers := make(chan error, len(cks)*2)
	cds := make(chan *chunk, len(cks))

	go func() {
		for cnk := range cds {
			if _, err := wrr.Seek(int64(cnk.start), 0); err != nil {
				ers <- err
				return
			}

			if _, err := io.CopyN(wrr, bytes.NewReader(cnk.data), int64(cnk.end-cnk.start)); err != nil {
				ers <- err
				return
			}

			ers <- nil
		}
	}()

	dcs := c.DownloadConcurrency
	smr := make(chan struct{}, dcs)

	for _, cnk := range cks {
		go func(cnk *chunk) {
			smr <- struct{}{}
			defer func() {
				ers <- nil
				<-smr
			}()

			hrq, err := c.newRequest(ctx, c.BaseURL, http.MethodGet, pth, nil)
			hrq.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", cnk.start, cnk.end))

			if err != nil {
				ers <- err
				return
			}

			res, err := c.do(hrq)

			if err != nil {
				ers <- err
				return
			}

			defer res.Body.Close()
			cnk.data, err = io.ReadAll(res.Body)

			if err != nil {
				ers <- err
				return
			}

			cds <- cnk
		}(cnk)
	}

	for i := 0; i < cap(ers); i++ {
		if err := <-ers; err != nil {
			return err
		}
	}

	close(cds)

	return nil
}

func (c *Client) subscribeToEntity(ctx context.Context, pth string, req *Request, cbk ReadCallback) error {
	hrq, err := c.newRequest(ctx, c.RealtimeURL, http.MethodGet, pth, req)

	if err != nil {
		return err
	}

	hrq.Header.Set("Cache-Control", "no-cache")
	hrq.Header.Set("Accept", "application/x-ndjson")
	hrq.Header.Set("Connection", "keep-alive")
	res, err := c.do(hrq)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	return c.readLoop(ctx, res.Body, cbk)
}

// ReadAll reads the contents of the given io.Reader and calls the given ReadCallback function
// with each chunk of data read.
func (c *Client) ReadAll(ctx context.Context, rdr io.Reader, cbk ReadCallback) error {
	gzr, err := pgzip.NewReader(rdr)

	if err != nil {
		return err
	}

	trr := tar.NewReader(gzr)

	for {
		_, err := trr.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err := c.readLoop(ctx, trr, cbk); err != nil {
			return err
		}
	}

	return nil
}

// SetAccessToken sets the access token for the client.
func (c *Client) SetAccessToken(tkn string) {
	c.AccessToken = tkn
}

// GetCodes retrieves a list of codes, and returns an error if any.
func (c *Client) GetCodes(ctx context.Context, req *Request) ([]*Code, error) {
	cds := []*Code{}
	return cds, c.getEntity(ctx, req, "codes", &cds)
}

// GetCode retrieves a code by ID, and returns an error if any.
func (c *Client) GetCode(ctx context.Context, idr string, req *Request) (*Code, error) {
	cde := new(Code)
	return cde, c.getEntity(ctx, req, fmt.Sprintf("codes/%s", idr), cde)
}

// GetLanguages retrieves a list of languages, and returns an error if any.
func (c *Client) GetLanguages(ctx context.Context, req *Request) ([]*Language, error) {
	lgs := []*Language{}
	return lgs, c.getEntity(ctx, req, "languages", &lgs)
}

// GetLanguage retrieves a language by ID, and returns an error if any.
func (c *Client) GetLanguage(ctx context.Context, idr string, req *Request) (*Language, error) {
	lng := new(Language)
	return lng, c.getEntity(ctx, req, fmt.Sprintf("languages/%s", idr), lng)
}

// GetProjects retrieves a list of projects, and returns an error if any.
func (c *Client) GetProjects(ctx context.Context, req *Request) ([]*Project, error) {
	prs := []*Project{}
	return prs, c.getEntity(ctx, req, "projects", &prs)
}

// GetProject retrieves a project by ID, and returns an error if any.
func (c *Client) GetProject(ctx context.Context, idr string, req *Request) (*Project, error) {
	prj := new(Project)
	return prj, c.getEntity(ctx, req, fmt.Sprintf("projects/%s", idr), prj)
}

// GetNamespaces retrieves a list of namespaces, and returns an error if any.
func (c *Client) GetNamespaces(ctx context.Context, req *Request) ([]*Namespace, error) {
	nss := []*Namespace{}
	return nss, c.getEntity(ctx, req, "namespaces", &nss)
}

// GetNamespaces retrieves a namespaces by ID, and returns an error if any.
func (c *Client) GetNamespace(ctx context.Context, idr int, req *Request) (*Namespace, error) {
	nsp := new(Namespace)
	return nsp, c.getEntity(ctx, req, fmt.Sprintf("namespaces/%d", idr), nsp)
}

// GetBatches retrieves a list of batches for a specific date and request, and returns an error if any.
func (c *Client) GetBatches(ctx context.Context, dte *time.Time, req *Request) ([]*Batch, error) {
	bts := []*Batch{}
	return bts, c.getEntity(ctx, req, fmt.Sprintf("batches/%s", dte.Format(DateFormat)), &bts)
}

// GetBatch retrieves a single batch for a specific date and ID, and returns an error if any.
func (c *Client) GetBatch(ctx context.Context, dte *time.Time, idr string, req *Request) (*Batch, error) {
	bth := new(Batch)
	return bth, c.getEntity(ctx, req, fmt.Sprintf("batches/%s/%s", dte.Format(DateFormat), idr), bth)
}

// HeadBatch retrieves only the headers of a single batch for a specific date and ID, and returns an error if any.
func (c *Client) HeadBatch(ctx context.Context, dte *time.Time, idr string) (*Headers, error) {
	return c.headEntity(ctx, fmt.Sprintf("batches/%s/%s/download", dte.Format(DateFormat), idr))
}

// ReadBatch reads the contents of a single batch for a specific date and ID, and invokes the specified callback function for each chunk read.
func (c *Client) ReadBatch(ctx context.Context, dte *time.Time, idr string, cbk ReadCallback) error {
	return c.readEntity(ctx, fmt.Sprintf("batches/%s/%s/download", dte.Format(DateFormat), idr), cbk)
}

// DownloadBatch downloads the contents of a single batch for a specific date and ID, and writes the data to the specified WriteSeeker.
func (c *Client) DownloadBatch(ctx context.Context, dte *time.Time, idr string, wsk io.WriteSeeker) error {
	return c.downloadEntity(ctx, fmt.Sprintf("batches/%s/%s/download", dte.Format(DateFormat), idr), wsk)
}

// GetSnapshots retrieves a list of all snapshots and returns an error if any.
func (c *Client) GetSnapshots(ctx context.Context, req *Request) ([]*Snapshot, error) {
	sps := []*Snapshot{}
	return sps, c.getEntity(ctx, req, "snapshots", &sps)
}

// GetSnapshot retrieves a single snapshot for a specific ID and returns an error if any.
func (c *Client) GetSnapshot(ctx context.Context, idr string, req *Request) (*Snapshot, error) {
	snp := new(Snapshot)
	return snp, c.getEntity(ctx, req, fmt.Sprintf("snapshots/%s", idr), snp)
}

// HeadSnapshot retrieves only the headers of a single snapshot for a specific ID, and returns an error if any.
func (c *Client) HeadSnapshot(ctx context.Context, idr string) (*Headers, error) {
	return c.headEntity(ctx, fmt.Sprintf("snapshots/%s/download", idr))
}

// ReadSnapshot reads the contents of a single snapshots for a specific ID, and invokes the specified callback function for each chunk read.
func (c *Client) ReadSnapshot(ctx context.Context, idr string, cbk ReadCallback) error {
	return c.readEntity(ctx, fmt.Sprintf("snapshots/%s/download", idr), cbk)
}

// DownloadSnapshot downloads the contents of a single snapshot for a specific ID, and writes the data to the specified WriteSeeker.
func (c *Client) DownloadSnapshot(ctx context.Context, idr string, wsk io.WriteSeeker) error {
	return c.downloadEntity(ctx, fmt.Sprintf("snapshots/%s/download", idr), wsk)
}

// GetChunks retrieves a list of all snapshot chunks and returns an error if any.
func (c *Client) GetChunks(ctx context.Context, sid string, req *Request) ([]*Snapshot, error) {
	sps := []*Snapshot{}
	return sps, c.getEntity(ctx, req, fmt.Sprintf("snapshots/%s/chunks", sid), &sps)
}

// GetChunk retrieves a single snapshot chunk for a specific ID and returns an error if any.
func (c *Client) GetChunk(ctx context.Context, sid string, idr string, req *Request) (*Snapshot, error) {
	snp := new(Snapshot)
	return snp, c.getEntity(ctx, req, fmt.Sprintf("snapshots/%s/chunks/%s", sid, idr), snp)
}

// HeadChunk retrieves only the headers of a single snapshot chunk for a specific ID, and returns an error if any.
func (c *Client) HeadChunk(ctx context.Context, sid string, idr string) (*Headers, error) {
	return c.headEntity(ctx, fmt.Sprintf("snapshots/%s/chunks/%s/download", sid, idr))
}

// ReadChunk reads the contents of a single snapshot chunk for a specific ID, and invokes the specified callback function for each chunk read.
func (c *Client) ReadChunk(ctx context.Context, sid string, idr string, cbk ReadCallback) error {
	return c.readEntity(ctx, fmt.Sprintf("snapshots/%s/chunks/%s/download", sid, idr), cbk)
}

// DownloadChunk downloads the contents of a single snapshot chunk for a specific ID, and writes the data to the specified WriteSeeker.
func (c *Client) DownloadChunk(ctx context.Context, sid string, idr string, wsk io.WriteSeeker) error {
	return c.downloadEntity(ctx, fmt.Sprintf("snapshots/%s/chunks/%s/download", sid, idr), wsk)
}

// GetArticles retrieves articles from the API based on the given name and request parameters.
func (c *Client) GetArticles(ctx context.Context, nme string, req *Request) ([]*Article, error) {
	ats := []*Article{}
	return ats, c.getEntity(ctx, req, fmt.Sprintf("articles/%s", nme), &ats)
}

// GetStructuredContents retrieves structured contents from the API based on the given name and request parameters.
func (c *Client) GetStructuredContents(ctx context.Context, nme string, req *Request) ([]*StructuredContent, error) {
	ats := []*StructuredContent{}
	return ats, c.getEntity(ctx, req, fmt.Sprintf("structured-contents/%s", nme), &ats)
}

// StreamArticles streams all available articles from the server and applies a callback function to each article
// as they arrive. The callback function must implement the ReadCallback interface.
func (c *Client) StreamArticles(ctx context.Context, req *Request, cbk ReadCallback) error {
	return c.subscribeToEntity(ctx, "articles", req, cbk)
}
