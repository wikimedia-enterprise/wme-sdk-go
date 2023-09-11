package api

// ProbabilityScore probability score representation for ORES models.
// Read more about ORES here: https://www.mediawiki.org/wiki/ORES.
type ProbabilityScore struct {
	// Prediction is the prediction of the model.
	Prediction bool `json:"prediction,omitempty"`

	// Probability is the probability of the prediction.
	Probability *Probability `json:"probability,omitempty"`
}
