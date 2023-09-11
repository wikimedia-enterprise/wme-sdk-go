package api

import (
	"time"
)

// Type of events supported by the system.
const (
	EventTypeUpdate           = "update"
	EventTypeDelete           = "delete"
	EventTypeVisibilityChange = "visibility-change"
)

// Event meta data for every event that happens in the system.
type Event struct {
	// Identifier is the unique identifier for this event (UUID).
	Identifier string `json:"identifier,omitempty"`

	// Type of the event, see EventType* constants.
	Type string `json:"type,omitempty"`

	// DateCreated is the date and time this event was created.
	DateCreated *time.Time `json:"date_created,omitempty"`

	// DatePublished is the date and time this event was published to the stream after processing.
	DatePublished *time.Time `json:"date_published,omitempty"`

	// Partition is the partition this event belongs to.
	Partition *int `json:"partition,omitempty"`

	// Offset what is the offset of the event inside the specified partition.
	Offset *int64 `json:"offset,omitempty"`
}
