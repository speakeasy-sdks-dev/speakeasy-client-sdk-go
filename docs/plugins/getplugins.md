# GetPlugins
Available in: `Plugins`

Get all plugins for the current workspace.

## Example Usage
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-api/speakeasy-client-sdk-go"
)

func main() {
    s := speakeasy.New(
        speakeasy.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plugins.GetPlugins(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plugins != nil {
        // handle response
    }
}
```