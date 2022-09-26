package shared

import (
	"time"
)

type UsageMetric struct {
	APIEndpointID   string    `json:"api_endpoint_id"`
	APIEndpointPath string    `json:"api_endpoint_path"`
	APIID           string    `json:"api_id"`
	CreatedAt       time.Time `json:"created_at"`
	CustomerID      string    `json:"customer_id"`
	Status          int64     `json:"status"`
	VersionID       string    `json:"version_id"`
	WorkspaceID     string    `json:"workspace_id"`
}
