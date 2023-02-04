package shared

// RequestMetadata
// Key-Value pairs associated with a request
type RequestMetadata struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}
