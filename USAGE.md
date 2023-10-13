<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
		Metadata: map[string][]string{
			"South": []string{
				"Southwest",
			},
		},
		Op: &operations.GetApisOp{
			And: false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Apis != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->