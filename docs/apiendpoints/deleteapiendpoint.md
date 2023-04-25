# DeleteAPIEndpoint
Available in: `APIEndpoints`

Delete an ApiEndpoint. This will also delete all associated Request Logs (if using a Postgres datastore).

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
    req := operations.DeleteAPIEndpointRequest{
        APIEndpointID: "delectus",
        APIID: "tempora",
        VersionID: "suscipit",
    }

    res, err := s.APIEndpoints.DeleteAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```