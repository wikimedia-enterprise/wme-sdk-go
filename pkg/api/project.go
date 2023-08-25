package api

import (
	"time"
)

// Project representation of mediawiki project according to https://schema.org/Project.
type Project struct {
	Name           string     `json:"name,omitempty"`
	Identifier     string     `json:"identifier,omitempty"`
	URL            string     `json:"url,omitempty"`
	Version        string     `json:"version,omitempty"`
	AdditionalType string     `json:"additional_type,omitempty"`
	DateModified   *time.Time `json:"date_modified,omitempty"`
	InLanguage     *Language  `json:"in_language,omitempty"`
	Size           *Size      `json:"size,omitempty"` // note that there's intentional `sizes` instead of `size` because size is ksqldb keyword
	Event          *Event     `json:"event,omitempty"`
}
