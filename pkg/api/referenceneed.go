package api

// ReferenceNeedData represents the structured response data for the reference-need Liftwing API.
type ReferenceNeedData struct {
	// ReferenceNeedScore estimates the likelihood that a statement requires a citation.
	ReferenceNeedScore float64 `json:"reference_need_score"`
}
