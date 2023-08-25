package api

// Size representation according to https://schema.org/QuantitativeValue.
type Size struct {
	Value    float64 `json:"value,omitempty"`
	UnitText string  `json:"unit_text,omitempty"`
}
