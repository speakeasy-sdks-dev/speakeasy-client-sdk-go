# GetEmbedAccessToken
Available in: `Embeds`

Returns an embed access token for the current workspace. This can be used to authenticate access to externally embedded content.
Filters can be applied allowing views to be filtered to things like particular customerIds.

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
    req := operations.GetEmbedAccessTokenRequest{
        Description: speakeasy.String("laborum"),
        Duration: speakeasy.Int64(170909),
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "corporis",
                    Operator: "explicabo",
                    Value: "nobis",
                },
            },
            Limit: 315428,
            Offset: 607831,
            Operator: "nemo",
        },
    }

    res, err := s.Embeds.GetEmbedAccessToken(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.EmbedAccessTokenResponse != nil {
        // handle response
    }
}
```