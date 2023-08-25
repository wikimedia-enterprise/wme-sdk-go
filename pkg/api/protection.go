package api

// Protection level for the article, does not comply with https://schema.org/ custom data.
type Protection struct {
	Type   string `json:"type,omitempty"`
	Level  string `json:"level,omitempty"`
	Expiry string `json:"expiry,omitempty"`
}
