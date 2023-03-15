package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GeneratePostmanCollectionRequest struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GeneratePostmanCollectionResponse struct {
	ContentType       string
	Error             *shared.Error
	PostmanCollection []byte
	StatusCode        int
	RawResponse       *http.Response
}
