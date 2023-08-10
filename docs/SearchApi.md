# \SearchApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchContent**](SearchApi.md#SearchContent) | **Get** /search | Search content across spaces that is accessible by the currently authenticated target



## SearchContent

> SearchContent200Response SearchContent(ctx).Query(query).Page(page).Limit(limit).Execute()

Search content across spaces that is accessible by the currently authenticated target

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
    query := "query_example" // string | 
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchContent(context.Background()).Query(query).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContent`: SearchContent200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**SearchContent200Response**](SearchContent200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

