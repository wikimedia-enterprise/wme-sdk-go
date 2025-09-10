package api

// StructuredTable represents a structured table .
type StructuredTable struct {
	// Identifier is a unique string that identifies the structured table.
	Identifier string `json:"identifier,omitempty"`

	// Headers contains the header cells of the table.
	Headers [][]*StructuredTableCell `json:"headers,omitempty"`

	// Rows contains the main body cells of the table.
	Rows [][]*StructuredTableCell `json:"rows,omitempty"`

	// Footers contains the footer cells of the table.
	Footers [][]*StructuredTableCell `json:"footers,omitempty"`

	// ConfidenceScore indicates the confidence (between 0.0 and 1.0)
	// associated with this structured table extraction.
	// A higher score suggests a structural and a well formed table.
	ConfidenceScore *float64 `json:"confidence_score,omitempty"`
}

// StructuredTableCell represents a single cell in a structured table.
type StructuredTableCell struct {
	// Value is the textual content of the cell.
	Value string `json:"value,omitempty"`

	// NestedTable represents a table contained within this cell, if present.
	NestedTable *StructuredTable `json:"nested_table,omitempty"`
}
