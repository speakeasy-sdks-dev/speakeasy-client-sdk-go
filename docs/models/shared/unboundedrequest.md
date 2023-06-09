# UnboundedRequest

An UnboundedRequest represents the HAR content capture by Speakeasy when logging a request.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CreatedAt`                                | [time.Time](https://pkg.go.dev/time#Time)  | :heavy_check_mark:                         | Creation timestamp.                        |
| `Har`                                      | *string*                                   | :heavy_check_mark:                         | The HAR content of the request.            |
| `HarSizeBytes`                             | *int64*                                    | :heavy_check_mark:                         | The size of the HAR content in bytes.      |
| `RequestID`                                | *string*                                   | :heavy_check_mark:                         | The ID of this request.                    |
| `WorkspaceID`                              | *string*                                   | :heavy_check_mark:                         | The workspace ID this request was made to. |