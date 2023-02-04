package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetPluginsResponse struct {
	ContentType string
	Error       *shared.Error
	Plugins     []shared.Plugin
	StatusCode  int64
}
