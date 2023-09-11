package api

// Protection level for the article, does not comply with https://schema.org/ custom data.
// Read more about protection here: https://www.mediawiki.org/wiki/Manual:Protection.
type Protection struct {
	// Type is the type of the protection.
	Type string `json:"type,omitempty"`

	// Level is the level of the protection.
	Level string `json:"level,omitempty"`

	// Expiry is the expiry of the protection.
	Expiry string `json:"expiry,omitempty"`
}
