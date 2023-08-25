package api

import "time"

// Headers is representation of headers is in the response of HEAD request.
type Headers struct {
	ContentLength int        `json:"content_length,omitempty"`
	ETag          string     `json:"etag,omitempty"`
	LastModified  *time.Time `json:"last_modified,omitempty"`
	ContentType   string     `json:"content_type,omitempty"`
	AcceptRanges  string     `json:"accept_ranges,omitempty"`
}
