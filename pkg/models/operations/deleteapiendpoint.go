package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIEndpointRequest struct {
	PathParams DeleteAPIEndpointPathParams
}

type DeleteAPIEndpointResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
