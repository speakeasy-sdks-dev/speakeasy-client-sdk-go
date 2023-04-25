# GetRequestFromEventLog
Available in: `Requests`

Get information about a particular request.

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
    req := operations.GetRequestFromEventLogRequest{
        RequestID: "sequi",
    }

    res, err := s.Requests.GetRequestFromEventLog(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UnboundedRequest != nil {
        // handle response
    }
}
```