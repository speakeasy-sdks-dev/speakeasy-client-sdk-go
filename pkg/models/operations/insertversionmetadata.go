// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type InsertVersionMetadataRequest struct {
	// A JSON representation of the metadata to insert.
	VersionMetadataInput shared.VersionMetadataInput `request:"mediaType=application/json"`
	// The ID of the Api to insert metadata for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to insert metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type InsertVersionMetadataResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// OK
	VersionMetadata *shared.VersionMetadata
}
