# FindAPIEndpoint
Available in: `APIEndpoints`

Find an ApiEndpoint via its displayName (set by operationId from a registered OpenAPI schema).
This is useful for finding the ID of an ApiEndpoint to use in the /v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID} endpoints.

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
    req := operations.FindAPIEndpointRequest{
        APIID: "molestiae",
        DisplayName: "minus",
        VersionID: "placeat",
    }

    res, err := s.APIEndpoints.FindAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```