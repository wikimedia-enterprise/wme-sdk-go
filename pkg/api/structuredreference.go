package api

// Reference represents a structured source linked to citations,
// containing metadata and formatted text for verification.
type StructuredReference struct {
	//A unique identifier for the article.
	Identifier string `json:"identifier,omitempty"`
	//The reference group
	Group string `json:"group,omitempty"`
	//The type of reference
	Type string `json:"type,omitempty"`
	//Additional structured details about the reference, such as author name, publisher, year, ISBN, etc.
	Metadata map[string]string `json:"metadata,omitempty"`
	//The reference as it appears in the article.
	Text *StructuredReferenceText `json:"text,omitempty"`
	//If the reference is from a bibliography, this stores its original source details.
	Source *StructuredReferenceText `json:"source,omitempty"`
}
