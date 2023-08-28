package api

// Size representation according to https://schema.org/QuantitativeValue.
type Size struct {
	// Value is the size value.
	Value float64 `json:"value,omitempty"`

	// UnitText is the unit of the size value.
	// For example: 'MB'.
	UnitText string `json:"unit_text,omitempty"`
}
