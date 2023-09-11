package api

import (
	"time"
)

// Project representation of mediawiki project according to https://schema.org/Project.
type Project struct {
	// Name is the human-readable name of this project.
	Name string `json:"name,omitempty"`

	// Identifier is the unique identifier for this project.
	// For example for English wikipedia it would be 'enwiki'.
	Identifier string `json:"identifier,omitempty"`

	// URL is the URL of the project.
	URL string `json:"url,omitempty"`

	// Code is the code of the project.
	// For example 'wiki' means that it's a Wikipedia project.
	Code string `json:"code,omitempty"`

	// Version is the version of the project in form of md5 checksum.
	Version string `json:"version,omitempty"`

	// DateModified is the date and time the project was last modified.
	DateModified *time.Time `json:"date_modified,omitempty"`

	// InLanguage is the language of the project.
	InLanguage *Language `json:"in_language,omitempty"`

	// Size is the size of the project.
	Size *Size `json:"size,omitempty"`
}
