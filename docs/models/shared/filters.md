# Filters

Filters are used to query requests.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Filters`                                   | [][Filter](../../models/shared/filter.md)   | :heavy_check_mark:                          | A list of filters to apply to the query.    |
| `Limit`                                     | *int64*                                     | :heavy_check_mark:                          | The maximum number of results to return.    |
| `Offset`                                    | *int64*                                     | :heavy_check_mark:                          | The offset to start the query from.         |
| `Operator`                                  | *string*                                    | :heavy_check_mark:                          | The operator to use when combining filters. |