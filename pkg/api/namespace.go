package api

// Namespace representation of mediawiki namespace.
// There's nothing related to this in https://schema.org/, we used  https://schema.org/Thing.
// Read more about namespaces here: https://en.wikipedia.org/wiki/Wikipedia:Namespace.
type Namespace struct {
	// Name is the human-readable name of this namespace.
	Name string `json:"name,omitempty"`

	// AlternateName is an alias for the namespace.
	AlternateName string `json:"alternate_name,omitempty"`

	// Identifier is the unique identifier for this namespace.
	Identifier int `json:"identifier"`
}
