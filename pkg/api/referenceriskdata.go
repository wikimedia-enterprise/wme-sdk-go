package api

// ReferenceRiskData represents the structured response data for the reference-need Liftwing API.
type ReferenceRiskData struct {
	// ReferenceRiskScore quantifies Proportion of deprecated/blacklisted domains
	ReferenceRiskScore float64 `json:"reference_risk_score"`
}
