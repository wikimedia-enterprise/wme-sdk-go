package api

import "time"

// Batch represents metadata for the realtime batch in WME API.
type Batch struct {
	Name         string     `json:"name,omitempty"`          // Name of the batch.
	Identifier   string     `json:"identifier,omitempty"`    // Unique identifier for the batch.
	Version      string     `json:"version,omitempty"`       // Version of the batch.
	DateModified *time.Time `json:"date_modified,omitempty"` // Time the batch was last modified.
	IsPartOf     *Project   `json:"is_part_of,omitempty"`    // The project that this batch belongs to.
	InLanguage   *Language  `json:"in_language,omitempty"`   // The language of the contents of the batch.
	Namespace    *Namespace `json:"namespace,omitempty"`     // The namespace of the batch.
	Size         *Size      `json:"size,omitempty"`          // The size of the batch.
}
