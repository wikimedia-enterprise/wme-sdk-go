package api

import (
	"time"
)

// Article schema for wikipedia article.
// Tries to compliant with https://schema.org/Article.
type Article struct {
	// Name is the name of the article.
	Name string `json:"name,omitempty"`

	// Abstract is a summary of the article.
	Abstract string `json:"abstract,omitempty"`

	// Identifier is a unique identifier for the article (in scope of a single project).
	Identifier int `json:"identifier,omitempty"`

	// DateCreated is the date and time the article was created.
	DateCreated *time.Time `json:"date_created,omitempty"`

	// DateModified is the date and time the article was last modified.
	DateModified *time.Time `json:"date_modified,omitempty"`

	// DatePreviouslyModified is the date and time the article was previously modified.
	DatePreviouslyModified *time.Time `json:"date_previously_modified,omitempty"`

	// Protection specifies the access restrictions for the article.
	Protection []*Protection `json:"protection,omitempty"`

	// Version is the metadata about the version of the article.
	Version *Version `json:"version,omitempty"`

	// PreviousVersion is the metadata about the previous version of the article.
	PreviousVersion *PreviousVersion `json:"previous_version,omitempty"`

	// URL is the URL of the article.
	URL string `json:"url,omitempty"`

	// WatchersCount is the number of watchers for the article.
	WatchersCount int `json:"watchers_count,omitempty"`

	// Namespace is the namespace of the article.
	Namespace *Namespace `json:"namespace,omitempty"`

	// InLanguage is the language of the article.
	InLanguage *Language `json:"in_language,omitempty"`

	// MainEntity is the main (Wikidata) entity of the article.
	MainEntity *Entity `json:"main_entity,omitempty"`

	// AdditionalEntities are the additional (Wikidata) entities used in the article.
	AdditionalEntities []*Entity `json:"additional_entities,omitempty"`

	// Categories are the categories of the article.
	Categories []*Category `json:"categories,omitempty"`

	// Templates are the templates used in the article.
	Templates []*Template `json:"templates,omitempty"`

	// Redirects are the redirects for the article.
	Redirects []*Redirect `json:"redirects,omitempty"`

	// IsPartOf is the project that the article belongs to.
	IsPartOf *Project `json:"is_part_of,omitempty"`

	// ArticleBody is the body of the article.
	ArticleBody *ArticleBody `json:"article_body,omitempty"`

	// License specifies the license for the article.
	License []*License `json:"license,omitempty"`

	// Visibility specifies the visibility of the article.
	Visibility *Visibility `json:"visibility,omitempty"`

	// Event specifies the event related to the article.
	Event *Event `json:"event,omitempty"`

	// Image specifies the image related to the article.
	Image *Image `json:"image,omitempty"`
}

// Image schema for article image.
// Compliant with https://schema.org/ImageObject,
type Image struct {
	// ContentUrl is the URL of the image.
	ContentURL string `json:"content_url,omitempty"`

	// Width is the width of the image.
	Width int `json:"width,omitempty" avro:"width"`

	// Height is the height of the image.
	Height int `json:"height,omitempty" avro:"height"`

	// AlternativeText is the alternative text of the image.
	AlternativeText string `json:"alternative_text,omitempty"`

	// Caption is the caption of the image.
	Caption string `json:"caption,omitempty"`
}

// Category article category representation.
type Category struct {
	// Name is the name of the category.
	Name string `json:"name,omitempty"`

	// URL is the URL of the category.
	URL string `json:"url,omitempty"`
}

// Redirect article redirect representation.
type Redirect struct {
	// Name is the name of the redirect.
	Name string `json:"name,omitempty"`

	// URL is the URL of the redirect.
	URL string `json:"url,omitempty"`
}

// Template article template representation.
type Template struct {
	// Name is the name of the template.
	Name string `json:"name,omitempty"`

	// URL is the URL of the template.
	URL string `json:"url,omitempty"`
}
