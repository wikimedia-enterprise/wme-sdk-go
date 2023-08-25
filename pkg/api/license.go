package api

// License representation according to https://schema.org/license.
type License struct {
	Name       string `json:"name,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty"`
}
