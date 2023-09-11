package api

// PreviousVersion is the representation for an article's previous version.
type PreviousVersion struct {
	// Identifier is the identifier of the previous version.
	Identifier int `json:"identifier,omitempty"`

	// NumberOfCharacters is the number of characters in the previous version.
	NumberOfCharacters int `json:"number_of_characters,omitempty"`
}

// Version representation for the article.
// Mainly modeled after https://schema.org/Thing.
type Version struct {
	// Identifier is the identifier of the version.
	Identifier int `json:"identifier,omitempty"`

	// Comment is the comment from editor for this particular version of the edit.
	Comment string `json:"comment,omitempty"`

	// Tags is the list of tags for this particular version of the edit.
	// Read more about tags: https://www.mediawiki.org/wiki/Manual:Tags.
	Tags []string `json:"tags,omitempty"`

	// IsMinorEdit is the flag for minor edit.
	IsMinorEdit bool `json:"is_minor_edit,omitempty"`

	// IsFlaggedStable is the version flagged stable.
	IsFlaggedStable bool `json:"is_flagged_stable,omitempty"`

	// IsBreakingNews is the version flagged as breaking news.
	IsBreakingNews bool `json:"is_breaking_news,omitempty"`

	// HasTagNeedsCitation is the version flagged as needs citation.
	HasTagNeedsCitation bool `json:"has_tag_needs_citation,omitempty"`

	// Score is ORES scores for the edit.
	Scores *Scores `json:"scores,omitempty"`

	// Editor is the information about the editor who made this edit.
	Editor *Editor `json:"editor,omitempty"`

	// NumberOfCharacters is the number of characters in the version.
	NumberOfCharacters int `json:"number_of_characters,omitempty"`

	// Size is the size of the version.
	Size *Size `json:"size,omitempty"`
}
