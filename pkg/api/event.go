package api

import (
	"time"
)

// Type of events supported by the system.
const (
	EventTypeUpdate           = "update"
	EventTypeCreate           = "create"
	EventTypeDelete           = "delete"
	EventTypeVisibilityChange = "visibility-change"
)

// Event meta data for every event that happens in the system.
type Event struct {
	Identifier    string     `json:"identifier,omitempty"`
	Type          string     `json:"type,omitempty"`
	DateCreated   *time.Time `json:"date_created,omitempty"`
	FailCount     int        `json:"fail_count,omitempty"`
	FailReason    string     `json:"fail_reason,omitempty"`
	DatePublished *time.Time `json:"date_published,omitempty"`
	Partition     *int       `json:"partition,omitempty"`
	Offset        *int64     `json:"offset,omitempty"`
}
