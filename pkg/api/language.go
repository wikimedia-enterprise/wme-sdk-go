package api

// Language representation according to https://schema.org/Language.
type Language struct {
	// Identifier is the unique identifier for this language.
	// For example 'en' is the identifier for 'English'.
	Identifier string `json:"identifier,omitempty"`

	// Name is the human-readable name of this language.
	Name string `json:"name,omitempty"`

	// AlternateName is an alias for the language.
	AlternateName string `json:"alternate_name,omitempty"`

	// Direction is the direction of the language.
	// For example 'ltr' is the direction for 'English'.
	Direction string `json:"direction,omitempty"`
}
