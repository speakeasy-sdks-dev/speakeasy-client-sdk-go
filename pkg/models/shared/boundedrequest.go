package shared

import (
	"time"
)

type BoundedRequest struct {
	APIEndpointID     *string             `json:"api_endpoint_id"`
	APIID             *string             `json:"api_id"`
	CreatedAt         *time.Time          `json:"created_at"`
	CustomerID        *string             `json:"customer_id"`
	Latency           *int64              `json:"latency"`
	Metadata          map[string][]string `json:"metadata"`
	Method            *string             `json:"method"`
	Path              *string             `json:"path"`
	RequestFinishTime *time.Time          `json:"request_finish_time"`
	RequestID         *string             `json:"request_id"`
	RequestStartTime  *time.Time          `json:"request_start_time"`
	Status            *int64              `json:"status"`
	VersionID         *string             `json:"version_id"`
	WorkspaceID       *string             `json:"workspace_id"`
}
