package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type RegisterSchemaV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type RegisterSchemaV1RequestBodyFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type RegisterSchemaV1RequestBody struct {
	File RegisterSchemaV1RequestBodyFile `multipartForm:"file"`
}

type RegisterSchemaV1Request struct {
	PathParams RegisterSchemaV1PathParams
	Request    RegisterSchemaV1RequestBody `request:"mediaType=multipart/form-data"`
}

type RegisterSchemaV1Responses struct {
	Error       *shared.Error
	RawResponse []byte
}

type RegisterSchemaV1Response struct {
	ContentType string
	Responses   map[int64]map[string]RegisterSchemaV1Responses
	StatusCode  int64
}
