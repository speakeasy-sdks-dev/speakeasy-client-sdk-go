package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GeneratePostmanCollectionForAPIEndpointRequest struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GeneratePostmanCollectionForAPIEndpointResponse struct {
	ContentType       string
	Error             *shared.Error
	PostmanCollection []byte
	StatusCode        int
	RawResponse       *http.Response
}
