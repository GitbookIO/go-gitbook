# \ApiApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiInformation**](ApiApi.md#GetApiInformation) | **Get** / | Get information about the state of the GitBook API



## GetApiInformation

> ApiInformation GetApiInformation(ctx).Execute()

Get information about the state of the GitBook API



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiApi.GetApiInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiApi.GetApiInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiInformation`: ApiInformation
    fmt.Fprintf(os.Stdout, "Response from `ApiApi.GetApiInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiInformationRequest struct via the builder pattern


### Return type

[**ApiInformation**](ApiInformation.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

