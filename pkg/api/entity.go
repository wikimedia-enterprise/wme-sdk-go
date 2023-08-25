package api

// Entity schema for wikidata article.
// Right now will just be a copy of initial wikidata schema.
// Partially uses https://schema.org/Thing.
type Entity struct {
	Identifier string   `json:"identifier,omitempty"`
	URL        string   `json:"url,omitempty"`
	Aspects    []string `json:"aspects,omitempty"`
}
