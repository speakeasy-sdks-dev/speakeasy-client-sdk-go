# GetApis
Available in: `Apis`

Get a list of all Apis and their versions for a given workspace.
Supports filtering the APIs based on metadata attributes.

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
    req := operations.GetApisRequest{
        Metadata: map[string][]string{
            "esse": []string{
                "excepturi",
            },
            "aspernatur": []string{
                "ad",
            },
            "natus": []string{
                "iste",
            },
        },
        Op: &operations.GetApisOp{
            And: false,
        },
    }

    res, err := s.Apis.GetApis(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
}
```