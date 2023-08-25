package api

// Delta represents the change description between two versions for certain dataset.
type Delta struct {
	Increase             int `json:"increase,omitempty"`
	Decrease             int `json:"decrease,omitempty"`
	Sum                  int `json:"sum,omitempty"`
	ProportionalIncrease int `json:"proportional_increase,omitempty"`
	ProportionalDecrease int `json:"proportional_decrease,omitempty"`
}
