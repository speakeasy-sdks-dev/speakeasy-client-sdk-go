package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetValidEmbedAccessTokensResponses struct {
	EmbedTokens []shared.EmbedToken
	Error       *shared.Error
}

type GetValidEmbedAccessTokensResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetValidEmbedAccessTokensResponses
	StatusCode  int64
}
