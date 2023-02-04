package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type UpsertPluginRequest struct {
	Request shared.Plugin `request:"mediaType=application/json"`
}

type UpsertPluginResponse struct {
	ContentType string
	Error       *shared.Error
	Plugin      *shared.Plugin
	StatusCode  int64
}
