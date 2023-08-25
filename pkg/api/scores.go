package api

// Scores ORES scores representation, has nothing on https://schema.org/, it's a custom dataset.
// For more info https://ores.wikimedia.org/.
type Scores struct {
	Damaging  *ProbabilityScore `json:"damaging,omitempty"`
	GoodFaith *ProbabilityScore `json:"goodfaith,omitempty"`
}

// Probability numeric probability values form ORES models.
type Probability struct {
	False float64 `json:"false,omitempty"`
	True  float64 `json:"true,omitempty"`
}
