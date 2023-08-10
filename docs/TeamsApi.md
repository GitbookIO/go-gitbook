# \TeamsApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTeamsForOrganizationMember**](TeamsApi.md#ListTeamsForOrganizationMember) | **Get** /orgs/{organizationId}/members/{userId}/teams | List all teams an organization member is part of



## ListTeamsForOrganizationMember

> ListTeamsForOrganizationMember200Response ListTeamsForOrganizationMember(ctx, organizationId, userId).Page(page).Limit(limit).Title(title).Execute()

List all teams an organization member is part of

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
    title := "title_example" // string | If provided, only teams whose name contains the given parameter will be returned. Case insensitive. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.ListTeamsForOrganizationMember(context.Background(), organizationId, userId).Page(page).Limit(limit).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.ListTeamsForOrganizationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsForOrganizationMember`: ListTeamsForOrganizationMember200Response
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.ListTeamsForOrganizationMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsForOrganizationMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **title** | **string** | If provided, only teams whose name contains the given parameter will be returned. Case insensitive. | 

### Return type

[**ListTeamsForOrganizationMember200Response**](ListTeamsForOrganizationMember200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

