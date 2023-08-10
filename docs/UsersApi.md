# \UsersApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthenticatedUser**](UsersApi.md#GetAuthenticatedUser) | **Get** /user | Get profile of authenticated user
[**GetUserById**](UsersApi.md#GetUserById) | **Get** /users/{userId} | Get a user by its ID
[**ListSpacesForAuthenticatedUser**](UsersApi.md#ListSpacesForAuthenticatedUser) | **Get** /user/spaces | List spaces for the authenticated user



## GetAuthenticatedUser

> User GetAuthenticatedUser(ctx).Execute()

Get profile of authenticated user



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
    resp, r, err := apiClient.UsersApi.GetAuthenticatedUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthenticatedUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatedUserRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> User GetUserById(ctx, userId).Execute()

Get a user by its ID



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
    userId := "userId_example" // string | The unique ID of the User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserById(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserById`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpacesForAuthenticatedUser

> ListSpacesForAuthenticatedUser200Response ListSpacesForAuthenticatedUser(ctx).Page(page).Limit(limit).Execute()

List spaces for the authenticated user



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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ListSpacesForAuthenticatedUser(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListSpacesForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpacesForAuthenticatedUser`: ListSpacesForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListSpacesForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpacesForAuthenticatedUserRequest struct via the builder pattern


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

