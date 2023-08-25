package api

import "time"

// Batch represents metadata for the daily snapshot in WME API.
type Snapshot struct {
	Name         string     `json:"name,omitempty"`          // Name of the snapshot.
	Identifier   string     `json:"identifier,omitempty"`    // Unique identifier for the snapshot.
	Version      string     `json:"version,omitempty"`       // Version of the snapshot.
	DateModified *time.Time `json:"date_modified,omitempty"` // Time the snapshot was last modified.
	IsPartOf     *Project   `json:"is_part_of,omitempty"`    // The project that this snapshot belongs to.
	InLanguage   *Language  `json:"in_language,omitempty"`   // The language of the contents of the snapshot.
	Namespace    *Namespace `json:"namespace,omitempty"`     // The namespace of the snapshot.
	Size         *Size      `json:"size,omitempty"`          // The size of the snapshot.
}
