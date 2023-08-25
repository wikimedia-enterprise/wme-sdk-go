package api

// Namespace representation of mediawiki namespace.
// There's nothing related to this in https://schema.org/, we used  https://schema.org/Thing.
type Namespace struct {
	Name          string `json:"name,omitempty"`
	AlternateName string `json:"alternate_name,omitempty"`
	Identifier    int    `json:"identifier"`
	Event         *Event `json:"event,omitempty"`
}
