# \UrlsApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContentByUrl**](UrlsApi.md#GetContentByUrl) | **Get** /urls/content | Resolve a URL to a content (space, collection, page)



## GetContentByUrl

> GetContentByUrl200Response GetContentByUrl(ctx).Url(url).Execute()

Resolve a URL to a content (space, collection, page)

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
    url := "url_example" // string | URL to resolve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UrlsApi.GetContentByUrl(context.Background()).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UrlsApi.GetContentByUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentByUrl`: GetContentByUrl200Response
    fmt.Fprintf(os.Stdout, "Response from `UrlsApi.GetContentByUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** | URL to resolve | 

### Return type

[**GetContentByUrl200Response**](GetContentByUrl200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

