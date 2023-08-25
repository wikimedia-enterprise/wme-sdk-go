package api

// ArticleBody schema for article content.
// Not fully compliant with https://schema.org/articleBody, we need multiple article bodies.
type ArticleBody struct {
	HTML     string `json:"html,omitempty"`
	WikiText string `json:"wikitext,omitempty"`
}
