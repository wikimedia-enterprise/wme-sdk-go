package api

// ProbabilityScore probability score representation for ORES models.
type ProbabilityScore struct {
	Prediction  bool         `json:"prediction,omitempty"`
	Probability *Probability `json:"probability,omitempty"`
}
