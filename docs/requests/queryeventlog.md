# QueryEventLog
Available in: `Requests`

Supports retrieving a list of request captured by the SDK for this workspace.
Allows the filtering of requests on a number of criteria such as ApiID, VersionID, Path, Method, etc.

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
    req := operations.QueryEventLogRequest{
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "ipsam",
                    Operator: "id",
                    Value: "possimus",
                },
                shared.Filter{
                    Key: "aut",
                    Operator: "quasi",
                    Value: "error",
                },
                shared.Filter{
                    Key: "temporibus",
                    Operator: "laborum",
                    Value: "quasi",
                },
                shared.Filter{
                    Key: "reiciendis",
                    Operator: "voluptatibus",
                    Value: "vero",
                },
            },
            Limit: 468651,
            Offset: 509624,
            Operator: "voluptatibus",
        },
    }

    res, err := s.Requests.QueryEventLog(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BoundedRequests != nil {
        // handle response
    }
}
```