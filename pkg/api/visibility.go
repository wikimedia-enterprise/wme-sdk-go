package api

// Visibility representing visibility changes for parts of the article.
// Custom dataset, not modeletd after https://schema.org/.
type Visibility struct {
	Text    bool `json:"text,omitempty"`
	Editor  bool `json:"editor,omitempty"`
	Comment bool `json:"comment,omitempty"`
}
