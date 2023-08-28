package api

// Entity schema for wikidata article.
// Right now will just be a copy of initial wikidata schema.
// Partially uses https://schema.org/Thing.
type Entity struct {
	// Identifier is the unique identifier for this entity.
	// For example 'Q937' is the identifier for 'Albert Einstein' see https://www.wikidata.org/wiki/Q937.
	Identifier string `json:"identifier,omitempty"`

	// URL is the URL of the entity.
	// Example: https://www.wikidata.org/wiki/Q937.
	URL string `json:"url,omitempty"`

	// Aspects of the entity that are being used in this article.
	Aspects []string `json:"aspects,omitempty"`
}
