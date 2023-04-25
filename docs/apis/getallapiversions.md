# GetAllAPIVersions
Available in: `Apis`

Get all Api versions for a particular ApiEndpoint.
Supports filtering the versions based on metadata attributes.

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
    req := operations.GetAllAPIVersionsRequest{
        APIID: "nam",
        Metadata: map[string][]string{
            "occaecati": []string{
                "deleniti",
            },
            "hic": []string{
                "totam",
                "beatae",
                "commodi",
                "molestiae",
            },
            "modi": []string{
                "impedit",
            },
        },
        Op: &operations.GetAllAPIVersionsOp{
            And: false,
        },
    }

    res, err := s.Apis.GetAllAPIVersions(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
}
```