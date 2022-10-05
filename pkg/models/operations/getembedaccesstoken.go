package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetEmbedAccessTokenQueryParams struct {
	Description *string         `queryParam:"style=form,explode=true,name=description"`
	Duration    *int64          `queryParam:"style=form,explode=true,name=duration"`
	Filters     *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type GetEmbedAccessTokenRequest struct {
	QueryParams GetEmbedAccessTokenQueryParams
}

type GetEmbedAccessTokenResponses struct {
	EmbedAccessTokenResponse *shared.EmbedAccessTokenResponse
	Error                    *shared.Error
}

type GetEmbedAccessTokenResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetEmbedAccessTokenResponses
	StatusCode  int64
}
