<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v2"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v2/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
		Metadata: map[string][]string{
			"key": []string{
				"string",
			},
		},
		Op: &operations.QueryParamOp{
			And: false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->