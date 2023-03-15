package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type UpsertAPIRequest struct {
	APIInput shared.APIInput `request:"mediaType=application/json"`
	APIID    string          `pathParam:"style=simple,explode=false,name=apiID"`
}

type UpsertAPIResponse struct {
	API         *shared.API
	ContentType string
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
