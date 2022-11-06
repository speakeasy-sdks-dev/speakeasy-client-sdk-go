package shared

// Filter
// A filter is a key-value pair that can be used to filter a list of requests.
type Filter struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
