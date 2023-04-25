# UpsertPlugin
Available in: `Plugins`

Upsert a plugin

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/types"
)

func main() {
    s := speakeasy.New(
        speakeasy.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.Plugin{
        Code: "velit",
        CreatedAt: types.MustTimeFromString("2022-09-06T22:51:09.401Z"),
        EvalHash: speakeasy.String("quis"),
        PluginID: "vitae",
        Title: "Miss",
        UpdatedAt: types.MustTimeFromString("2022-05-14T10:37:30.872Z"),
        WorkspaceID: "odit",
    }

    res, err := s.Plugins.UpsertPlugin(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plugin != nil {
        // handle response
    }
}
```