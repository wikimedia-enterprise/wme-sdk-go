package api

// Language representation accroding to https://schema.org/Language.
type Language struct {
	Identifier    string `json:"identifier,omitempty"`
	Name          string `json:"name,omitempty"`
	AlternateName string `json:"alternate_name,omitempty"`
	Direction     string `json:"direction,omitempty"`
}
