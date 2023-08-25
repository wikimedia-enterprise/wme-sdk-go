package api

// Code is a project code representation in WME API.
type Code struct {
	// Identifier is the unique identifier for this code.
	Identifier string `json:"identifier,omitempty"`
	// Name is the human-readable name of this code.
	Name string `json:"name,omitempty"`
	// Description is a description of this code.
	Description string `json:"description,omitempty"`
}
