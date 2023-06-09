# EmbedToken

A representation of an embed token granted for working with Speakeasy components.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `CreatedAt`                                 | [time.Time](https://pkg.go.dev/time#Time)   | :heavy_check_mark:                          | Creation timestamp.                         |
| `CreatedBy`                                 | *string*                                    | :heavy_check_mark:                          | The ID of the user that created this token. |
| `Description`                               | *string*                                    | :heavy_check_mark:                          | A detailed description of the EmbedToken.   |
| `ExpiresAt`                                 | [time.Time](https://pkg.go.dev/time#Time)   | :heavy_check_mark:                          | The time this token expires.                |
| `Filters`                                   | *string*                                    | :heavy_check_mark:                          | The filters applied to this token.          |
| `ID`                                        | *string*                                    | :heavy_check_mark:                          | The ID of this EmbedToken.                  |
| `LastUsed`                                  | [*time.Time](https://pkg.go.dev/time#Time)  | :heavy_minus_sign:                          | The last time this token was used.          |
| `RevokedAt`                                 | [*time.Time](https://pkg.go.dev/time#Time)  | :heavy_minus_sign:                          | The time this token was revoked.            |
| `RevokedBy`                                 | **string*                                   | :heavy_minus_sign:                          | The ID of the user that revoked this token. |
| `WorkspaceID`                               | *string*                                    | :heavy_check_mark:                          | The workspace ID this token belongs to.     |