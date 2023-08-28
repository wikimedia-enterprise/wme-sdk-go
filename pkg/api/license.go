package api

// License representation according to https://schema.org/license.
type License struct {
	// Name is the human-readable name of this license.
	Name string `json:"name,omitempty"`

	// Identifier is the unique identifier for this license.
	// Example: 'CC-BY-SA-3.0'.
	Identifier string `json:"identifier,omitempty"`

	// URL is the URL of the license.
	// Example: https://creativecommons.org/licenses/by-sa/3.0/.
	URL string `json:"url,omitempty"`
}
