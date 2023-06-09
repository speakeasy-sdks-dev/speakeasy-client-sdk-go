# QueryEventLogResponse


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `BoundedRequests`                                                | [][shared.BoundedRequest](../../models/shared/boundedrequest.md) | :heavy_minus_sign:                                               | OK                                                               |
| `ContentType`                                                    | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `Error`                                                          | [*shared.Error](../../models/shared/error.md)                    | :heavy_minus_sign:                                               | Default error response                                           |
| `StatusCode`                                                     | *int*                                                            | :heavy_check_mark:                                               | N/A                                                              |
| `RawResponse`                                                    | [*http.Response](https://pkg.go.dev/net/http#Response)           | :heavy_minus_sign:                                               | N/A                                                              |