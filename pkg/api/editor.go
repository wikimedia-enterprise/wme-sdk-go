package api

import (
	"time"
)

// Editor for the article version.
// Combines Person and CreativeWork with custom properties, link https://schema.org/editor.
type Editor struct {
	// Identifier is the unique identifier for this editor.
	Identifier int `json:"identifier,omitempty"`

	// Name is the human-readable name of this editor.
	Name string `json:"name,omitempty"`

	// EditCount is the number of edits made by this editor.
	EditCount int `json:"edit_count,omitempty"`

	// Groups is the list of groups this editor belongs to.
	Groups []string `json:"groups,omitempty"`

	// IsBot is true if this editor is a bot.
	IsBot bool `json:"is_bot,omitempty"`

	// IsAnonymous is true if this editor is anonymous.
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// IsAdmin is true if this editor is an admin.
	IsAdmin bool `json:"is_admin,omitempty"`

	// IsPatroller is true if this editor is a patroller.
	IsPatroller bool `json:"is_patroller,omitempty"`

	// HasAdvancedRights is true if this editor has advanced rights.
	HasAdvancedRights bool `json:"has_advanced_rights,omitempty"`

	// DateStarted is the date and time this editor started editing.
	DateStarted *time.Time `json:"date_started,omitempty"`
}
