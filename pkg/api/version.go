package api

// PreviousVersion is the representation for an article's previous version.
type PreviousVersion struct {
	Identifier         int `json:"identifier,omitempty"`
	NumberOfCharacters int `json:"number_of_characters,omitempty"`
}

// Version representation for the article.
// Mainly modeled after https://schema.org/Thing.
type Version struct {
	Identifier          int      `json:"identifier,omitempty"`
	Comment             string   `json:"comment,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	IsMinorEdit         bool     `json:"is_minor_edit,omitempty"`
	IsFlaggedStable     bool     `json:"is_flagged_stable,omitempty"`
	IsBreakingNews      bool     `json:"is_breaking_news,omitempty"`
	HasTagNeedsCitation bool     `json:"has_tag_needs_citation,omitempty"`
	Scores              *Scores  `json:"scores,omitempty"`
	Editor              *Editor  `json:"editor,omitempty"`
	NumberOfCharacters  int      `json:"number_of_characters,omitempty"`
	Size                *Size    `json:"size,omitempty"`
	Event               *Event   `json:"event,omitempty"`
}
