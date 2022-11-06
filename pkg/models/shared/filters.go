package shared

// Filters
// Filters are used to query requests.
type Filters struct {
	Filters  []Filter `json:"filters"`
	Limit    int64    `json:"limit"`
	Offset   int64    `json:"offset"`
	Operator string   `json:"operator"`
}
