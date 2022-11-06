package shared

import (
	"time"
)

// UnboundedRequest
// An UnboundedRequest represents the HAR content capture by Speakeasy when logging a request.
type UnboundedRequest struct {
	CreatedAt    time.Time `json:"created_at"`
	Har          string    `json:"har"`
	HarSizeBytes int64     `json:"har_size_bytes"`
	RequestID    string    `json:"request_id"`
	WorkspaceID  string    `json:"workspace_id"`
}
