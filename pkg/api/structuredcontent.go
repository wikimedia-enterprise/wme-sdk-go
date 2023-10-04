package api

import "time"

// StructuredContent schema for Machine Readable structured contents endpoint.
type StructuredContent struct {
	// Name is the name of the structured content (article).
	Name string `json:"name,omitempty"`

	// Identifier is a unique identifier for the structured content (article, in scope of a single project).
	Identifier int `json:"identifier,omitempty"`

	// Abstract is a summary of the structured content (article).
	Abstract string `json:"abstract,omitempty"`

	// Description is a description of the structured content (article).
	Description string `json:"description,omitempty"`

	// Version is the metadata about the version of the structured content (article).
	Version *Version `json:"version,omitempty"`

	// URL is the URL of the structured content (article).
	URL string `json:"url,omitempty"`

	// DateCreated is the date and time the structured content (article) was created.
	DateCreated *time.Time `json:"date_created,omitempty"`

	// DateModified is the date and time the structured content (article) was last modified.
	DateModified *time.Time `json:"date_modified,omitempty"`

	// MainEntity is the main (Wikidata) entity of the structured content (article).
	MainEntity *Entity `json:"main_entity,omitempty"`

	// AdditionalEntities are the additional (Wikidata) entities used in the structured content (article).
	AdditionalEntities []*Entity `json:"additional_entities,omitempty"`

	// IsPartOf is the project that the structured content (article) belongs to.
	IsPartOf *Project `json:"is_part_of,omitempty"`

	// InLanguage is the language of the structured content (article).
	InLanguage *Language `json:"in_language,omitempty"`

	// InfoBox are the parts included inside the structured content (article).
	InfoBox []*Part `json:"infobox,omitempty"`

	// Image specifies the image related to the structured content (article).
	Image *Image `json:"image,omitempty"`
}

// Part represents a part of a structured content (section, field etc.).
type Part struct {
	// Name is the name of the part.
	Name string `json:"name,omitempty"`

	// Type is the type of the part, for example 'field' or 'section'.
	Type string `json:"type,omitempty"`

	// Value is the value of the part.
	Value string `json:"value,omitempty"`

	// Values are the values of the part (if there are are more than single value).
	Values []string `json:"values,omitempty"`

	// Images are the images included inside the part.
	Images []*Image `json:"images,omitempty"`

	// Links are the links included inside the part.
	Links []*Link `json:"links,omitempty"`

	// HasParts are the parts included inside the part (recursively parts can contain parts).
	HasParts []*Part `json:"has_parts,omitempty"`
}

// Link represents a link that can be found on a Wikipedia page.
type Link struct {
	// URL is the URL of the link.
	URL string `json:"url,omitempty"`

	// Text is the text of the link.
	Text string `json:"text,omitempty"`

	// Images are the images included inside  the link.
	Images []*Image `json:"images,omitempty"`
}
