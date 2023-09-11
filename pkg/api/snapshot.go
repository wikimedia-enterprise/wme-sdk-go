package api

import "time"

// Batch represents metadata for the daily snapshot in WME API.
type Snapshot struct {
	// Identifier is the unique identifier for the snapshot.
	Identifier string `json:"identifier,omitempty"`

	// Version of the snapshot as a md5 checksum.
	Version string `json:"version,omitempty"`

	// DateModified date and time the snapshot was last modified.
	DateModified *time.Time `json:"date_modified,omitempty"`

	// IsPartOf the project that this snapshot belongs to.
	IsPartOf *Project `json:"is_part_of,omitempty"`

	// InLanguage the language of the contents of the snapshot.
	InLanguage *Language `json:"in_language,omitempty"`

	// Namespace of the snapshot.
	Namespace *Namespace `json:"namespace,omitempty"`

	// Size of the snapshot.
	Size *Size `json:"size,omitempty"`
}
