package api

import "time"

// Headers is representation of headers is in the response of HEAD request.
type Headers struct {
	// ContentLength is the size of the file in bytes.
	// Read more about Content-Length header here: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length.
	ContentLength int `json:"content_length,omitempty"`

	// ETag is the entity tag of the file.
	// Read more about ETag header here: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag.
	ETag string `json:"etag,omitempty"`

	// LastModified is the date and time the file was last modified.
	// Read more about Last-Modified header here: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified.
	LastModified *time.Time `json:"last_modified,omitempty"`

	// ContentType is the content type of the file.
	// Read more about Content-Type header here: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type.
	ContentType string `json:"content_type,omitempty"`

	// AcceptRanges is the accept ranges of the file.
	// Read more about Accept-Ranges header here: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Ranges.
	AcceptRanges string `json:"accept_ranges,omitempty"`
}
