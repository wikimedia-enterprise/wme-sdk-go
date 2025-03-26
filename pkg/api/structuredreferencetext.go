package api

type StructuredReferenceText struct {
	//A formatted citation string from the bibliography.
	Value string `json:"value,omitempty"`
	//Links associated with the source
	Links []*Link `json:"links,omitempty"`
}
