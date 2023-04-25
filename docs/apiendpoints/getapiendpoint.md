# GetAPIEndpoint
Available in: `APIEndpoints`

Get an ApiEndpoint.

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
    req := operations.GetAPIEndpointRequest{
        APIEndpointID: "deserunt",
        APIID: "perferendis",
        VersionID: "ipsam",
    }

    res, err := s.APIEndpoints.GetAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```