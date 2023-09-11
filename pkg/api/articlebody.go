package api

// ArticleBody schema for article content.
// Not fully compliant with https://schema.org/articleBody, we need multiple article bodies.
type ArticleBody struct {
	// HTML of the article, as a Parsoid HTML.
	// See the specification at https://www.mediawiki.org/wiki/Specs/HTML/2.8.0.
	HTML string `json:"html,omitempty"`

	// WikiText of the article.
	// For more information visit https://www.mediawiki.org/wiki/Wikitext.
	// Specification: https://www.mediawiki.org/wiki/Specs/wikitext/1.0.0.
	WikiText string `json:"wikitext,omitempty"`
}
