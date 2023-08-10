# \CollectionsApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCollectionById**](CollectionsApi.md#GetCollectionById) | **Get** /collections/{collectionId} | Get the details about a collection using its ID
[**GetCollectionPublishingAuthById**](CollectionsApi.md#GetCollectionPublishingAuthById) | **Get** /collections/{collectionId}/publishing/auth | Get the publishing authentication configuration for a collection.
[**ListSpacesInCollectionById**](CollectionsApi.md#ListSpacesInCollectionById) | **Get** /collections/{collectionId}/spaces | List all the spaces in a collection by its ID.
[**UpdateCollectionById**](CollectionsApi.md#UpdateCollectionById) | **Patch** /collections/{collectionId} | Update the details of a collection
[**UpdateCollectionPublishingAuthById**](CollectionsApi.md#UpdateCollectionPublishingAuthById) | **Post** /collections/{collectionId}/publishing/auth | Update the publishing authentication configuration for a collection.



## GetCollectionById

> Collection GetCollectionById(ctx, collectionId).Execute()

Get the details about a collection using its ID

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
    collectionId := "collectionId_example" // string | The unique id of the collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.GetCollectionById(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.GetCollectionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionById`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.GetCollectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The unique id of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionPublishingAuthById

> ContentPublishingAuth GetCollectionPublishingAuthById(ctx, collectionId).Execute()

Get the publishing authentication configuration for a collection.

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
    collectionId := "collectionId_example" // string | The unique id of the collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.GetCollectionPublishingAuthById(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.GetCollectionPublishingAuthById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionPublishingAuthById`: ContentPublishingAuth
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.GetCollectionPublishingAuthById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The unique id of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionPublishingAuthByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentPublishingAuth**](ContentPublishingAuth.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpacesInCollectionById

> ListSpacesForAuthenticatedUser200Response ListSpacesInCollectionById(ctx, collectionId).Page(page).Limit(limit).Execute()

List all the spaces in a collection by its ID.

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
    collectionId := "collectionId_example" // string | The unique id of the collection
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.ListSpacesInCollectionById(context.Background(), collectionId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.ListSpacesInCollectionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpacesInCollectionById`: ListSpacesForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.ListSpacesInCollectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The unique id of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpacesInCollectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListSpacesForAuthenticatedUser200Response**](ListSpacesForAuthenticatedUser200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollectionById

> Collection UpdateCollectionById(ctx, collectionId).UpdateCollectionByIdRequest(updateCollectionByIdRequest).Execute()

Update the details of a collection

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
    collectionId := "collectionId_example" // string | The unique id of the collection
    updateCollectionByIdRequest := *openapiclient.NewUpdateCollectionByIdRequest() // UpdateCollectionByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.UpdateCollectionById(context.Background(), collectionId).UpdateCollectionByIdRequest(updateCollectionByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.UpdateCollectionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCollectionById`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.UpdateCollectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The unique id of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCollectionByIdRequest** | [**UpdateCollectionByIdRequest**](UpdateCollectionByIdRequest.md) |  | 

### Return type

[**Collection**](Collection.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollectionPublishingAuthById

> ContentPublishingAuth UpdateCollectionPublishingAuthById(ctx, collectionId).RequestUpdateContentPublishingAuth(requestUpdateContentPublishingAuth).Execute()

Update the publishing authentication configuration for a collection.

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
    collectionId := "collectionId_example" // string | The unique id of the collection
    requestUpdateContentPublishingAuth := *openapiclient.NewRequestUpdateContentPublishingAuth() // RequestUpdateContentPublishingAuth | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.UpdateCollectionPublishingAuthById(context.Background(), collectionId).RequestUpdateContentPublishingAuth(requestUpdateContentPublishingAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.UpdateCollectionPublishingAuthById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCollectionPublishingAuthById`: ContentPublishingAuth
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.UpdateCollectionPublishingAuthById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The unique id of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionPublishingAuthByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdateContentPublishingAuth** | [**RequestUpdateContentPublishingAuth**](RequestUpdateContentPublishingAuth.md) |  | 

### Return type

[**ContentPublishingAuth**](ContentPublishingAuth.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

