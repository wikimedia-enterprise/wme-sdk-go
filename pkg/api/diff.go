package api

// Diff representats the difference between current and previous version.
type Diff struct {
	LongestNewRepeatedCharacter int    `json:"longest_new_repeated_character,omitempty"`
	Words                       *Delta `json:"words,omitempty"`
	NonWords                    *Delta `json:"non_words,omitempty"`
	NonSafeWords                *Delta `json:"non_safe_words,omitempty"`
	InformalWords               *Delta `json:"informal_words,omitempty"`
	UppercaseLetters            *Delta `json:"uppercase_letters,omitempty"`
	Size                        *Size  `json:"size,omitempty"`
}
