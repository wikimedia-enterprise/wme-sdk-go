package api

// StructuredCitation represents a collection of inline citations within an article,
// linking specific content to corresponding references.
type StructuredCitation struct {
	//A unique identifier for the article.
	Identifier string `json:"identifier,omitempty"`
	//The reference group
	Group string `json:"group,omitempty"`
	//The inline citation as it appears in the article
	Text string `json:"text,omitempty"`
}
