package shared

// Error
// The `Status` type defines a logical error model
type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
