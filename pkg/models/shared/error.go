package shared

type Error struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}
