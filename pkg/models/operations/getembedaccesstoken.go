package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetEmbedAccessTokenRequest struct {
	Description *string         `queryParam:"style=form,explode=true,name=description"`
	Duration    *int64          `queryParam:"style=form,explode=true,name=duration"`
	Filters     *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type GetEmbedAccessTokenResponse struct {
	ContentType              string
	EmbedAccessTokenResponse *shared.EmbedAccessTokenResponse
	Error                    *shared.Error
	StatusCode               int
	RawResponse              *http.Response
}
