package api

// Scores ORES scores representation, has nothing on https://schema.org/, it's a custom dataset.
// For more info https://ores.wikimedia.org/.
type Scores struct {
	RevertRisk *ProbabilityScore `json:"revertrisk,omitempty" avro:"revertrisk"`
}

// Probability numeric probability values form ORES models.
type Probability struct {
	// False is the probability of the prediction being false.
	False float64 `json:"false,omitempty"`

	// True is the probability of the prediction being true.
	True float64 `json:"true,omitempty"`
}
