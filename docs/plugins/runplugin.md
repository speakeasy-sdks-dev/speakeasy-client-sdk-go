# RunPlugin
Available in: `Plugins`

Run a plugin

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

func main() {
    s := speakeasy.New(
        speakeasy.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.RunPluginRequest{
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "repellat",
                    Operator: "mollitia",
                    Value: "occaecati",
                },
            },
            Limit: 253291,
            Offset: 414369,
            Operator: "quam",
        },
        PluginID: "molestiae",
    }

    res, err := s.Plugins.RunPlugin(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BoundedRequests != nil {
        // handle response
    }
}
```