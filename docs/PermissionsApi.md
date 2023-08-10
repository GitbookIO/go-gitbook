# \PermissionsApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPermissionsAggregateInCollection**](PermissionsApi.md#ListPermissionsAggregateInCollection) | **Get** /collections/{collectionId}/permissions/aggregate | List permissions for all users in a collection.
[**ListPermissionsAggregateInSpace**](PermissionsApi.md#ListPermissionsAggregateInSpace) | **Get** /spaces/{spaceId}/permissions/aggregate | List permissions for all users in a space.
[**ListSpacesForOrganizationMember**](PermissionsApi.md#ListSpacesForOrganizationMember) | **Get** /orgs/{organizationId}/members/{userId}/spaces | List permissions accross all spaces for a member of an organization



## ListPermissionsAggregateInCollection

> ListPermissionsAggregateInSpace200Response ListPermissionsAggregateInCollection(ctx, collectionId).Page(page).Limit(limit).Role(role).Execute()

List permissions for all users in a collection.

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
    role := openapiclient.MemberRole("admin") // MemberRole | If defined, only members with this role will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.ListPermissionsAggregateInCollection(context.Background(), collectionId).Page(page).Limit(limit).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.ListPermissionsAggregateInCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissionsAggregateInCollection`: ListPermissionsAggregateInSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.ListPermissionsAggregateInCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The unique id of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsAggregateInCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **role** | [**MemberRole**](MemberRole.md) | If defined, only members with this role will be returned. | 

### Return type

[**ListPermissionsAggregateInSpace200Response**](ListPermissionsAggregateInSpace200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissionsAggregateInSpace

> ListPermissionsAggregateInSpace200Response ListPermissionsAggregateInSpace(ctx, spaceId).Page(page).Limit(limit).Role(role).Execute()

List permissions for all users in a space.

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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    role := openapiclient.MemberRole("admin") // MemberRole | If defined, only members with this role will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.ListPermissionsAggregateInSpace(context.Background(), spaceId).Page(page).Limit(limit).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.ListPermissionsAggregateInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissionsAggregateInSpace`: ListPermissionsAggregateInSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.ListPermissionsAggregateInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsAggregateInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **role** | [**MemberRole**](MemberRole.md) | If defined, only members with this role will be returned. | 

### Return type

[**ListPermissionsAggregateInSpace200Response**](ListPermissionsAggregateInSpace200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpacesForOrganizationMember

> ListSpacesForOrganizationMember200Response ListSpacesForOrganizationMember(ctx, organizationId, userId).Page(page).Limit(limit).Execute()

List permissions accross all spaces for a member of an organization

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
    organizationId := "organizationId_example" // string | The unique id of the organization
    userId := "userId_example" // string | The unique ID of the User
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionsApi.ListSpacesForOrganizationMember(context.Background(), organizationId, userId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsApi.ListSpacesForOrganizationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpacesForOrganizationMember`: ListSpacesForOrganizationMember200Response
    fmt.Fprintf(os.Stdout, "Response from `PermissionsApi.ListSpacesForOrganizationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpacesForOrganizationMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListSpacesForOrganizationMember200Response**](ListSpacesForOrganizationMember200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

