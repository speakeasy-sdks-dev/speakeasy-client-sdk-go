# Reports
(*Reports*)

## Overview

REST APIs for managing reports

### Available Operations

* [GetLintingReportSignedURL](#getlintingreportsignedurl) - Get the signed access url for the linting reports for a particular document.
* [UploadReport](#uploadreport) - Upload a report.

## GetLintingReportSignedURL

Get the signed access url for the linting reports for a particular document.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Reports.GetLintingReportSignedURL(ctx, operations.GetLintingReportSignedURLRequest{
        DocumentChecksum: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SignedAccess != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetLintingReportSignedURLRequest](../../pkg/models/operations/getlintingreportsignedurlrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetLintingReportSignedURLResponse](../../pkg/models/operations/getlintingreportsignedurlresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UploadReport

Upload a report.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Reports.UploadReport(ctx, operations.UploadReportRequestBody{
        Data: shared.Report{},
        File: operations.File{
            Content: []byte("0xA329C0ad85"),
            FileName: "sausages.m3a",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadedReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UploadReportRequestBody](../../pkg/models/operations/uploadreportrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UploadReportResponse](../../pkg/models/operations/uploadreportresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
