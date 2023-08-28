package api

// Visibility representing visibility changes for parts of the article.
// Custom dataset, not modeled after https://schema.org/.
type Visibility struct {
	// Text is the visibility of the text ('true' for visible, 'false` for hidden`).
	Text bool `json:"text,omitempty"`

	// Edit is the visibility of the editor ('true' for visible, 'false` for hidden`).
	Editor bool `json:"editor,omitempty"`

	// Comment is the visibility of the comment ('true' for visible, 'false` for hidden`).
	Comment bool `json:"comment,omitempty"`
}
