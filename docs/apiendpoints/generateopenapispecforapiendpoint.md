# GenerateOpenAPISpecForAPIEndpoint
Available in: `APIEndpoints`

This endpoint will generate a new operation in any registered OpenAPI document if the operation does not already exist in the document.
Returns the original document and the newly generated document allowing a diff to be performed to see what has changed.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasy.New(
        speakeasy.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GenerateOpenAPISpecForAPIEndpointRequest{
        APIEndpointID: "voluptatum",
        APIID: "iusto",
        VersionID: "excepturi",
    }

    res, err := s.APIEndpoints.GenerateOpenAPISpecForAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateOpenAPISpecDiff != nil {
        // handle response
    }
}
```