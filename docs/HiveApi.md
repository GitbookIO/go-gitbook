# \HiveApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateHiveAccessToken**](HiveApi.md#GenerateHiveAccessToken) | **Post** /internal/hive/token | Returns a token to authenticate with Hive.
[**GenerateSpaceHiveReadAccessToken**](HiveApi.md#GenerateSpaceHiveReadAccessToken) | **Post** /spaces/{spaceId}/hive/token | Returns a token to authenticate with Hive to read content from a given space.



## GenerateHiveAccessToken

> HiveAccessToken GenerateHiveAccessToken(ctx).GenerateHiveAccessTokenRequest(generateHiveAccessTokenRequest).Execute()

Returns a token to authenticate with Hive.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GitbookIO/go-gitbook"
)

func main() {
    generateHiveAccessTokenRequest := *openapiclient.NewGenerateHiveAccessTokenRequest() // GenerateHiveAccessTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HiveApi.GenerateHiveAccessToken(context.Background()).GenerateHiveAccessTokenRequest(generateHiveAccessTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HiveApi.GenerateHiveAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateHiveAccessToken`: HiveAccessToken
    fmt.Fprintf(os.Stdout, "Response from `HiveApi.GenerateHiveAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateHiveAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateHiveAccessTokenRequest** | [**GenerateHiveAccessTokenRequest**](GenerateHiveAccessTokenRequest.md) |  | 

### Return type

[**HiveAccessToken**](HiveAccessToken.md)

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSpaceHiveReadAccessToken

> HiveAccessToken GenerateSpaceHiveReadAccessToken(ctx, spaceId).Execute()

Returns a token to authenticate with Hive to read content from a given space.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GitbookIO/go-gitbook"
)

func main() {
    spaceId := "spaceId_example" // string | The unique id of the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HiveApi.GenerateSpaceHiveReadAccessToken(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HiveApi.GenerateSpaceHiveReadAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSpaceHiveReadAccessToken`: HiveAccessToken
    fmt.Fprintf(os.Stdout, "Response from `HiveApi.GenerateSpaceHiveReadAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSpaceHiveReadAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HiveAccessToken**](HiveAccessToken.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

