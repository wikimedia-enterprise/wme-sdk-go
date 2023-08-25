package api

import (
	"time"
)

// Editor for the article version.
// Combines Person and CreativeWork with custom properties, link https://schema.org/editor.
type Editor struct {
	Identifier        int        `json:"identifier,omitempty"`
	Name              string     `json:"name,omitempty"`
	EditCount         int        `json:"edit_count,omitempty"`
	Groups            []string   `json:"groups,omitempty"`
	IsBot             bool       `json:"is_bot,omitempty"`
	IsAnonymous       bool       `json:"is_anonymous,omitempty"`
	IsAdmin           bool       `json:"is_admin,omitempty"`
	IsPatroller       bool       `json:"is_patroller,omitempty"`
	HasAdvancedRights bool       `json:"has_advanced_rights,omitempty"`
	DateStarted       *time.Time `json:"date_started,omitempty"`
}
