# RegisterSchema
Available in: `Schemas`

Allows uploading a schema for a particular API version.
This will be used to populate ApiEndpoints and used as a base for any schema generation if present.

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
    req := operations.RegisterSchemaRequest{
        RequestBody: operations.RegisterSchemaRequestBody{
            File: operations.RegisterSchemaRequestBodyFile{
                Content: []byte("quae"),
                File: "ipsum",
            },
        },
        APIID: "quidem",
        VersionID: "molestias",
    }

    res, err := s.Schemas.RegisterSchema(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```