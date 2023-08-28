package api

import "time"

// Batch represents metadata for the realtime batch in WME API.
type Batch struct {
	// Unique identifier for the batch.
	Identifier string `json:"identifier,omitempty"`

	// Version of the batch as a md5 checksum.
	Version string `json:"version,omitempty"`

	// DateModified date and time the batch was last modified.
	DateModified *time.Time `json:"date_modified,omitempty"`

	// IsPartOf the project that this batch belongs to.
	IsPartOf *Project `json:"is_part_of,omitempty"`

	// InLanguage the language of the contents of the batch.
	InLanguage *Language `json:"in_language,omitempty"`

	// Namespace of the batch.
	Namespace *Namespace `json:"namespace,omitempty"`

	// Size of the batch.
	Size *Size `json:"size,omitempty"`
}
