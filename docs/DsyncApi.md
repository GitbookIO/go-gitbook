# \DsyncApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDirectorySyncGroups**](DsyncApi.md#ListDirectorySyncGroups) | **Get** /orgs/{organizationId}/dsync/groups | Lists the groups exposed to the synced Directory on an organization.
[**SetupDirectorySync**](DsyncApi.md#SetupDirectorySync) | **Post** /orgs/{organizationId}/dsync | Set up Directory Sync in an organization.
[**SyncDirectorySyncGroupsToTeams**](DsyncApi.md#SyncDirectorySyncGroupsToTeams) | **Post** /orgs/{organizationId}/dsync/teams | Syncs a list of group/team unique identifiers pairs together.



## ListDirectorySyncGroups

> ListDirectorySyncGroups200Response ListDirectorySyncGroups(ctx, organizationId).Execute()

Lists the groups exposed to the synced Directory on an organization.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DsyncApi.ListDirectorySyncGroups(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DsyncApi.ListDirectorySyncGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDirectorySyncGroups`: ListDirectorySyncGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `DsyncApi.ListDirectorySyncGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDirectorySyncGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListDirectorySyncGroups200Response**](ListDirectorySyncGroups200Response.md)

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetupDirectorySync

> SetupDirectorySync200Response SetupDirectorySync(ctx, organizationId).Execute()

Set up Directory Sync in an organization.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DsyncApi.SetupDirectorySync(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DsyncApi.SetupDirectorySync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetupDirectorySync`: SetupDirectorySync200Response
    fmt.Fprintf(os.Stdout, "Response from `DsyncApi.SetupDirectorySync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetupDirectorySyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetupDirectorySync200Response**](SetupDirectorySync200Response.md)

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncDirectorySyncGroupsToTeams

> SyncDirectorySyncGroupsToTeams200Response SyncDirectorySyncGroupsToTeams(ctx, organizationId).SyncDirectorySyncGroupsToTeamsRequest(syncDirectorySyncGroupsToTeamsRequest).Execute()

Syncs a list of group/team unique identifiers pairs together.



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
    syncDirectorySyncGroupsToTeamsRequest := *openapiclient.NewSyncDirectorySyncGroupsToTeamsRequest([]openapiclient.SyncDirectorySyncGroupsToTeamsRequestToSyncInner{*openapiclient.NewSyncDirectorySyncGroupsToTeamsRequestToSyncInner("GroupId_example")}) // SyncDirectorySyncGroupsToTeamsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DsyncApi.SyncDirectorySyncGroupsToTeams(context.Background(), organizationId).SyncDirectorySyncGroupsToTeamsRequest(syncDirectorySyncGroupsToTeamsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DsyncApi.SyncDirectorySyncGroupsToTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncDirectorySyncGroupsToTeams`: SyncDirectorySyncGroupsToTeams200Response
    fmt.Fprintf(os.Stdout, "Response from `DsyncApi.SyncDirectorySyncGroupsToTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncDirectorySyncGroupsToTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syncDirectorySyncGroupsToTeamsRequest** | [**SyncDirectorySyncGroupsToTeamsRequest**](SyncDirectorySyncGroupsToTeamsRequest.md) |  | 

### Return type

[**SyncDirectorySyncGroupsToTeams200Response**](SyncDirectorySyncGroupsToTeams200Response.md)

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

