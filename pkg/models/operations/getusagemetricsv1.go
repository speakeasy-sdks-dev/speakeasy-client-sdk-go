package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetUsageMetricsV1PathParams struct {
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

type GetUsageMetricsV1QueryParams struct {
	Filters *string `queryParam:"style=form,explode=true,name=filters"`
}

type GetUsageMetricsV1Request struct {
	PathParams  GetUsageMetricsV1PathParams
	QueryParams GetUsageMetricsV1QueryParams
}

type GetUsageMetricsV1Responses struct {
	Error        *shared.Error
	UsageMetrics []shared.UsageMetric
}

type GetUsageMetricsV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetUsageMetricsV1Responses
	StatusCode  int64
}
