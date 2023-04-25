# GetAllForVersionAPIEndpoints
Available in: `APIEndpoints`

Get all ApiEndpoints for a particular apiID and versionID.

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
    req := operations.GetAllForVersionAPIEndpointsRequest{
        APIID: "quis",
        VersionID: "veritatis",
    }

    res, err := s.APIEndpoints.GetAllForVersionAPIEndpoints(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoints != nil {
        // handle response
    }
}
```