# \OrganizationsApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemberToOrganizationTeamById**](OrganizationsApi.md#AddMemberToOrganizationTeamById) | **Put** /orgs/{organizationId}/teams/{teamId}/members/{userId} | Add or update a team membership
[**CreateEnvironment**](OrganizationsApi.md#CreateEnvironment) | **Post** /orgs/{organizationId}/environments | Create a new environment within an organization
[**CreateOrganization**](OrganizationsApi.md#CreateOrganization) | **Post** /orgs | Create an organization
[**CreateOrganizationCustomField**](OrganizationsApi.md#CreateOrganizationCustomField) | **Post** /orgs/{organizationId}/custom-fields | Create a new custom field in an orgamization
[**CreateOrganizationTeam**](OrganizationsApi.md#CreateOrganizationTeam) | **Put** /orgs/{organizationId}/teams | Create organization team
[**DeleteEnvironment**](OrganizationsApi.md#DeleteEnvironment) | **Delete** /orgs/{organizationId}/environments/{environmentName} | Delete an environment in an organization
[**DeleteMemberFromOrganizationTeamById**](OrganizationsApi.md#DeleteMemberFromOrganizationTeamById) | **Delete** /orgs/{organizationId}/teams/{teamId}/members/{userId} | Delete members from a team
[**DeleteOrganizationCustomField**](OrganizationsApi.md#DeleteOrganizationCustomField) | **Delete** /orgs/{organizationId}/custom-fields/{fieldName} | Delete a custom field in an organization
[**GetEnvironmentByName**](OrganizationsApi.md#GetEnvironmentByName) | **Get** /orgs/{organizationId}/environments/{environmentName} | Get an environment by its name
[**GetMemberInOrganizationById**](OrganizationsApi.md#GetMemberInOrganizationById) | **Get** /orgs/{organizationId}/members/{userId} | Get specified organization member
[**GetOrganizationBillingPortal**](OrganizationsApi.md#GetOrganizationBillingPortal) | **Get** /orgs/{organizationId}/billing | Get the billing portal for an organization
[**GetOrganizationById**](OrganizationsApi.md#GetOrganizationById) | **Get** /orgs/{organizationId} | Get an organization by its ID
[**GetOrganizationCustomFieldByName**](OrganizationsApi.md#GetOrganizationCustomFieldByName) | **Get** /orgs/{organizationId}/custom-fields/{fieldName} | Get a custom field by its name
[**GetTeamInOrganizationById**](OrganizationsApi.md#GetTeamInOrganizationById) | **Get** /orgs/{organizationId}/teams/{teamId} | Get specified organization team
[**InviteUsersToOrganization**](OrganizationsApi.md#InviteUsersToOrganization) | **Post** /orgs/{organizationId}/invites | Invite users to a given organization based on a list of emails
[**JoinOrganizationWithInvite**](OrganizationsApi.md#JoinOrganizationWithInvite) | **Post** /orgs/{organizationId}/invites/{inviteId} | Use an invite to join an organization.
[**ListCollectionsInOrganizationById**](OrganizationsApi.md#ListCollectionsInOrganizationById) | **Get** /orgs/{organizationId}/collections | List organization collections
[**ListDirectorySyncGroups**](OrganizationsApi.md#ListDirectorySyncGroups) | **Get** /orgs/{organizationId}/dsync/groups | Lists the groups exposed to the synced Directory on an organization.
[**ListEnvironments**](OrganizationsApi.md#ListEnvironments) | **Get** /orgs/{organizationId}/environments | Get the environments in an organization
[**ListMembersInOrganizationById**](OrganizationsApi.md#ListMembersInOrganizationById) | **Get** /orgs/{organizationId}/members | List organization members
[**ListOrganizationCustomFields**](OrganizationsApi.md#ListOrganizationCustomFields) | **Get** /orgs/{organizationId}/custom-fields | Get the custom fields for spaces in an organization
[**ListOrganizationsForAuthenticatedUser**](OrganizationsApi.md#ListOrganizationsForAuthenticatedUser) | **Get** /orgs | Get the list of organizations for the currently authenticated user
[**ListSpacesInOrganizationById**](OrganizationsApi.md#ListSpacesInOrganizationById) | **Get** /orgs/{organizationId}/spaces | List organization spaces
[**ListSpacesWithGitSyncInOrganizationById**](OrganizationsApi.md#ListSpacesWithGitSyncInOrganizationById) | **Get** /orgs/{organizationId}/spaces/gitsync | List organization spaces including Git sync metadata
[**ListTeamMembersInOrganizationById**](OrganizationsApi.md#ListTeamMembersInOrganizationById) | **Get** /orgs/{organizationId}/teams/{teamId}/members | List team members
[**ListTeamsInOrganizationById**](OrganizationsApi.md#ListTeamsInOrganizationById) | **Get** /orgs/{organizationId}/teams | List organization teams
[**RemoveMemberFromOrganizationById**](OrganizationsApi.md#RemoveMemberFromOrganizationById) | **Delete** /orgs/{organizationId}/members/{userId} | Delete a member from an organization
[**RemoveTeamFromOrganizationById**](OrganizationsApi.md#RemoveTeamFromOrganizationById) | **Delete** /orgs/{organizationId}/teams/{teamId} | Delete a team in an organization
[**RequestOrganizationUpgrade**](OrganizationsApi.md#RequestOrganizationUpgrade) | **Post** /orgs/{organizationId}/request_upgrade | Send a request to ask the organization&#39;s admin to upgrade it.
[**SearchOrganizationContent**](OrganizationsApi.md#SearchOrganizationContent) | **Get** /orgs/{organizationId}/search | Search content in an organization
[**SetUserAsSSOMemberForOrganization**](OrganizationsApi.md#SetUserAsSSOMemberForOrganization) | **Post** /orgs/{organizationId}/members/{userId}/sso | Set a user as an SSO member of an organization
[**SetupDirectorySync**](OrganizationsApi.md#SetupDirectorySync) | **Post** /orgs/{organizationId}/dsync | Set up Directory Sync in an organization.
[**SyncDirectorySyncGroupsToTeams**](OrganizationsApi.md#SyncDirectorySyncGroupsToTeams) | **Post** /orgs/{organizationId}/dsync/teams | Syncs a list of group/team unique identifiers pairs together.
[**TransferOrganization**](OrganizationsApi.md#TransferOrganization) | **Post** /orgs/{organizationId}/transfer | Transfer one organization (source) into another organization (target).
[**UpdateEnvironment**](OrganizationsApi.md#UpdateEnvironment) | **Patch** /orgs/{organizationId}/environments/{environmentName} | Update an existing environment within an organization
[**UpdateMemberInOrganizationById**](OrganizationsApi.md#UpdateMemberInOrganizationById) | **Patch** /orgs/{organizationId}/members/{userId} | Update specified organization member
[**UpdateMembersInOrganizationTeam**](OrganizationsApi.md#UpdateMembersInOrganizationTeam) | **Put** /orgs/{organizationId}/teams/{teamId}/members | Updates members of an organization team
[**UpdateOrganizationCustomField**](OrganizationsApi.md#UpdateOrganizationCustomField) | **Patch** /orgs/{organizationId}/custom-fields/{fieldName} | Update a custom field in an organization
[**UpdateOrganizationMemberLastSeenAt**](OrganizationsApi.md#UpdateOrganizationMemberLastSeenAt) | **Post** /orgs/{organizationId}/ping | Update organization member&#39;s \&quot;last seen at\&quot; timestamp.
[**UpdateTeamInOrganizationById**](OrganizationsApi.md#UpdateTeamInOrganizationById) | **Patch** /orgs/{organizationId}/teams/{teamId} | Update specified organization team
[**UpgradeOrganizationPlan**](OrganizationsApi.md#UpgradeOrganizationPlan) | **Post** /orgs/{organizationId}/billing | Upgrade an organization&#39;s billing plan



## AddMemberToOrganizationTeamById

> AddMemberToOrganizationTeamById(ctx, organizationId, teamId, userId).AddMemberToOrganizationTeamByIdRequest(addMemberToOrganizationTeamByIdRequest).Execute()

Add or update a team membership



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
    teamId := "teamId_example" // string | The unique ID of the Team
    userId := "userId_example" // string | The unique ID of the User
    addMemberToOrganizationTeamByIdRequest := *openapiclient.NewAddMemberToOrganizationTeamByIdRequest() // AddMemberToOrganizationTeamByIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsApi.AddMemberToOrganizationTeamById(context.Background(), organizationId, teamId, userId).AddMemberToOrganizationTeamByIdRequest(addMemberToOrganizationTeamByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.AddMemberToOrganizationTeamById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**teamId** | **string** | The unique ID of the Team | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberToOrganizationTeamByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addMemberToOrganizationTeamByIdRequest** | [**AddMemberToOrganizationTeamByIdRequest**](AddMemberToOrganizationTeamByIdRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironment

> Environment CreateEnvironment(ctx, organizationId).CreateEnvironment(createEnvironment).Execute()

Create a new environment within an organization

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
    createEnvironment := *openapiclient.NewCreateEnvironment("Name_example", "Title_example") // CreateEnvironment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateEnvironment(context.Background(), organizationId).CreateEnvironment(createEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEnvironment** | [**CreateEnvironment**](CreateEnvironment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> Organization CreateOrganization(ctx).RequestCreateOrganization(requestCreateOrganization).Execute()

Create an organization

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
    requestCreateOrganization := *openapiclient.NewRequestCreateOrganization("Title_example") // RequestCreateOrganization |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganization(context.Background()).RequestCreateOrganization(requestCreateOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreateOrganization** | [**RequestCreateOrganization**](RequestCreateOrganization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationCustomField

> CustomField CreateOrganizationCustomField(ctx, organizationId).CreateOrganizationCustomFieldRequest(createOrganizationCustomFieldRequest).Execute()

Create a new custom field in an orgamization

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
    createOrganizationCustomFieldRequest := *openapiclient.NewCreateOrganizationCustomFieldRequest("Name_example", openapiclient.CustomFieldType("text")) // CreateOrganizationCustomFieldRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationCustomField(context.Background(), organizationId).CreateOrganizationCustomFieldRequest(createOrganizationCustomFieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationCustomField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCustomField`: CustomField
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCustomFieldRequest** | [**CreateOrganizationCustomFieldRequest**](CreateOrganizationCustomFieldRequest.md) |  | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationTeam

> OrganizationTeam CreateOrganizationTeam(ctx, organizationId).CreateOrganizationTeamRequest(createOrganizationTeamRequest).Execute()

Create organization team



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
    createOrganizationTeamRequest := *openapiclient.NewCreateOrganizationTeamRequest("Title_example") // CreateOrganizationTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.CreateOrganizationTeam(context.Background(), organizationId).CreateOrganizationTeamRequest(createOrganizationTeamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.CreateOrganizationTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationTeam`: OrganizationTeam
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.CreateOrganizationTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationTeamRequest** | [**CreateOrganizationTeamRequest**](CreateOrganizationTeamRequest.md) |  | 

### Return type

[**OrganizationTeam**](OrganizationTeam.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, organizationId, environmentName).Execute()

Delete an environment in an organization

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
    environmentName := "environmentName_example" // string | The unique name of the environment within the organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsApi.DeleteEnvironment(context.Background(), organizationId, environmentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**environmentName** | **string** | The unique name of the environment within the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemberFromOrganizationTeamById

> DeleteMemberFromOrganizationTeamById(ctx, organizationId, teamId, userId).Execute()

Delete members from a team



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
    teamId := "teamId_example" // string | The unique ID of the Team
    userId := "userId_example" // string | The unique ID of the User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsApi.DeleteMemberFromOrganizationTeamById(context.Background(), organizationId, teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteMemberFromOrganizationTeamById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**teamId** | **string** | The unique ID of the Team | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberFromOrganizationTeamByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationCustomField

> DeleteOrganizationCustomField(ctx, organizationId, fieldName).Execute()

Delete a custom field in an organization

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
    fieldName := "fieldName_example" // string | The name of the custom field

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsApi.DeleteOrganizationCustomField(context.Background(), organizationId, fieldName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeleteOrganizationCustomField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**fieldName** | **string** | The name of the custom field | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentByName

> Environment GetEnvironmentByName(ctx, organizationId, environmentName).Execute()

Get an environment by its name

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
    environmentName := "environmentName_example" // string | The unique name of the environment within the organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetEnvironmentByName(context.Background(), organizationId, environmentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetEnvironmentByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentByName`: Environment
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetEnvironmentByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**environmentName** | **string** | The unique name of the environment within the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Environment**](Environment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemberInOrganizationById

> OrganizationMember GetMemberInOrganizationById(ctx, organizationId, userId).Execute()

Get specified organization member



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetMemberInOrganizationById(context.Background(), organizationId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetMemberInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMemberInOrganizationById`: OrganizationMember
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetMemberInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBillingPortal

> BillingPortal GetOrganizationBillingPortal(ctx, organizationId).Execute()

Get the billing portal for an organization

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
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationBillingPortal(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationBillingPortal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBillingPortal`: BillingPortal
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationBillingPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBillingPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingPortal**](BillingPortal.md)

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationById

> Organization GetOrganizationById(ctx, organizationId).Execute()

Get an organization by its ID

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
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationById(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationById`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCustomFieldByName

> CustomField GetOrganizationCustomFieldByName(ctx, organizationId, fieldName).Execute()

Get a custom field by its name

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
    fieldName := "fieldName_example" // string | The name of the custom field

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizationCustomFieldByName(context.Background(), organizationId, fieldName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationCustomFieldByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCustomFieldByName`: CustomField
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationCustomFieldByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**fieldName** | **string** | The name of the custom field | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCustomFieldByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomField**](CustomField.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamInOrganizationById

> OrganizationTeam GetTeamInOrganizationById(ctx, organizationId, teamId).Execute()

Get specified organization team



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
    teamId := "teamId_example" // string | The unique ID of the Team

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetTeamInOrganizationById(context.Background(), organizationId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetTeamInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamInOrganizationById`: OrganizationTeam
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetTeamInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**teamId** | **string** | The unique ID of the Team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationTeam**](OrganizationTeam.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUsersToOrganization

> InviteUsersToOrganization200Response InviteUsersToOrganization(ctx, organizationId).RequestInviteUsersToOrganization(requestInviteUsersToOrganization).Execute()

Invite users to a given organization based on a list of emails

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
    requestInviteUsersToOrganization := *openapiclient.NewRequestInviteUsersToOrganization([]openapiclient.RequestInviteUsersToOrganizationEmailsInner{openapiclient.RequestInviteUsersToOrganization_emails_inner{RequestInviteUsersToOrganizationEmailsInnerOneOf: openapiclient.NewRequestInviteUsersToOrganizationEmailsInnerOneOf("Email_example", openapiclient.MemberRoleOrGuest{MemberRole: penapiclient.MemberRole("admin")})}}) // RequestInviteUsersToOrganization | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.InviteUsersToOrganization(context.Background(), organizationId).RequestInviteUsersToOrganization(requestInviteUsersToOrganization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.InviteUsersToOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteUsersToOrganization`: InviteUsersToOrganization200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.InviteUsersToOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteUsersToOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInviteUsersToOrganization** | [**RequestInviteUsersToOrganization**](RequestInviteUsersToOrganization.md) |  | 

### Return type

[**InviteUsersToOrganization200Response**](InviteUsersToOrganization200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JoinOrganizationWithInvite

> map[string]interface{} JoinOrganizationWithInvite(ctx, organizationId, inviteId).Execute()

Use an invite to join an organization.

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
    inviteId := "inviteId_example" // string | The unique id of the invite

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.JoinOrganizationWithInvite(context.Background(), organizationId, inviteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.JoinOrganizationWithInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoinOrganizationWithInvite`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.JoinOrganizationWithInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**inviteId** | **string** | The unique id of the invite | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoinOrganizationWithInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionsInOrganizationById

> ListCollectionsInOrganizationById200Response ListCollectionsInOrganizationById(ctx, organizationId).Page(page).Limit(limit).Nested(nested).Execute()

List organization collections



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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    nested := true // bool | If true, all nested collections will be listed (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListCollectionsInOrganizationById(context.Background(), organizationId).Page(page).Limit(limit).Nested(nested).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListCollectionsInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCollectionsInOrganizationById`: ListCollectionsInOrganizationById200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListCollectionsInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **nested** | **bool** | If true, all nested collections will be listed | [default to true]

### Return type

[**ListCollectionsInOrganizationById200Response**](ListCollectionsInOrganizationById200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.OrganizationsApi.ListDirectorySyncGroups(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListDirectorySyncGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDirectorySyncGroups`: ListDirectorySyncGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListDirectorySyncGroups`: %v\n", resp)
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


## ListEnvironments

> ListEnvironments200Response ListEnvironments(ctx, organizationId).Page(page).Limit(limit).Execute()

Get the environments in an organization

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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListEnvironments(context.Background(), organizationId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironments`: ListEnvironments200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListEnvironments200Response**](ListEnvironments200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMembersInOrganizationById

> ListMembersInOrganizationById200Response ListMembersInOrganizationById(ctx, organizationId).Page(page).Limit(limit).Role(role).Search(search).Execute()

List organization members



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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    role := openapiclient.listMembersInOrganizationById_role_parameter{MemberRole: penapiclient.MemberRole("admin")} // ListMembersInOrganizationByIdRoleParameter | The Role to filter the member list by (optional)
    search := "search_example" // string | A query to filter the member list (displayName and email) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListMembersInOrganizationById(context.Background(), organizationId).Page(page).Limit(limit).Role(role).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListMembersInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembersInOrganizationById`: ListMembersInOrganizationById200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListMembersInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **role** | [**ListMembersInOrganizationByIdRoleParameter**](ListMembersInOrganizationByIdRoleParameter.md) | The Role to filter the member list by | 
 **search** | **string** | A query to filter the member list (displayName and email) | 

### Return type

[**ListMembersInOrganizationById200Response**](ListMembersInOrganizationById200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationCustomFields

> ListOrganizationCustomFields200Response ListOrganizationCustomFields(ctx, organizationId).Page(page).Limit(limit).Execute()

Get the custom fields for spaces in an organization

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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListOrganizationCustomFields(context.Background(), organizationId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListOrganizationCustomFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationCustomFields`: ListOrganizationCustomFields200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListOrganizationCustomFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListOrganizationCustomFields200Response**](ListOrganizationCustomFields200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizationsForAuthenticatedUser

> ListOrganizationsForAuthenticatedUser200Response ListOrganizationsForAuthenticatedUser(ctx).Page(page).Limit(limit).Execute()

Get the list of organizations for the currently authenticated user

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
    resp, r, err := apiClient.OrganizationsApi.ListOrganizationsForAuthenticatedUser(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListOrganizationsForAuthenticatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizationsForAuthenticatedUser`: ListOrganizationsForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListOrganizationsForAuthenticatedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsForAuthenticatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListOrganizationsForAuthenticatedUser200Response**](ListOrganizationsForAuthenticatedUser200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpacesInOrganizationById

> ListSpacesForAuthenticatedUser200Response ListSpacesInOrganizationById(ctx, organizationId).Page(page).Limit(limit).Visibility(visibility).Execute()

List organization spaces



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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    visibility := openapiclient.ContentVisibility("public") // ContentVisibility | If defined, only content with this visibility will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListSpacesInOrganizationById(context.Background(), organizationId).Page(page).Limit(limit).Visibility(visibility).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListSpacesInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpacesInOrganizationById`: ListSpacesForAuthenticatedUser200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListSpacesInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpacesInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **visibility** | [**ContentVisibility**](ContentVisibility.md) | If defined, only content with this visibility will be returned. | 

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


## ListSpacesWithGitSyncInOrganizationById

> ListSpacesWithGitSyncInOrganizationById200Response ListSpacesWithGitSyncInOrganizationById(ctx, organizationId).Page(page).Limit(limit).Status(status).Execute()

List organization spaces including Git sync metadata



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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    status := openapiclient.GitSyncOperationState("running") // GitSyncOperationState | If defined, only spaces with matching Git sync status are returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListSpacesWithGitSyncInOrganizationById(context.Background(), organizationId).Page(page).Limit(limit).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListSpacesWithGitSyncInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpacesWithGitSyncInOrganizationById`: ListSpacesWithGitSyncInOrganizationById200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListSpacesWithGitSyncInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpacesWithGitSyncInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **status** | [**GitSyncOperationState**](GitSyncOperationState.md) | If defined, only spaces with matching Git sync status are returned | 

### Return type

[**ListSpacesWithGitSyncInOrganizationById200Response**](ListSpacesWithGitSyncInOrganizationById200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamMembersInOrganizationById

> ListMembersInOrganizationById200Response ListTeamMembersInOrganizationById(ctx, organizationId, teamId).Page(page).Limit(limit).Execute()

List team members



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
    teamId := "teamId_example" // string | The unique ID of the Team
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListTeamMembersInOrganizationById(context.Background(), organizationId, teamId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListTeamMembersInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamMembersInOrganizationById`: ListMembersInOrganizationById200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListTeamMembersInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**teamId** | **string** | The unique ID of the Team | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamMembersInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListMembersInOrganizationById200Response**](ListMembersInOrganizationById200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamsInOrganizationById

> ListTeamsInOrganizationById200Response ListTeamsInOrganizationById(ctx, organizationId).Page(page).Limit(limit).Title(title).Execute()

List organization teams



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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    title := "title_example" // string | If provided, only teams whose name contains the given parameter will be returned. Case insensitive. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListTeamsInOrganizationById(context.Background(), organizationId).Page(page).Limit(limit).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListTeamsInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamsInOrganizationById`: ListTeamsInOrganizationById200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListTeamsInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **title** | **string** | If provided, only teams whose name contains the given parameter will be returned. Case insensitive. | 

### Return type

[**ListTeamsInOrganizationById200Response**](ListTeamsInOrganizationById200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMemberFromOrganizationById

> RemoveMemberFromOrganizationById(ctx, organizationId, userId).Execute()

Delete a member from an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsApi.RemoveMemberFromOrganizationById(context.Background(), organizationId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.RemoveMemberFromOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMemberFromOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTeamFromOrganizationById

> RemoveTeamFromOrganizationById(ctx, organizationId, teamId).Execute()

Delete a team in an organization



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
    teamId := "teamId_example" // string | The unique ID of the Team

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsApi.RemoveTeamFromOrganizationById(context.Background(), organizationId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.RemoveTeamFromOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**teamId** | **string** | The unique ID of the Team | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTeamFromOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestOrganizationUpgrade

> map[string]interface{} RequestOrganizationUpgrade(ctx, organizationId).Execute()

Send a request to ask the organization's admin to upgrade it.

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
    resp, r, err := apiClient.OrganizationsApi.RequestOrganizationUpgrade(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.RequestOrganizationUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestOrganizationUpgrade`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.RequestOrganizationUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestOrganizationUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrganizationContent

> SearchContent200Response SearchOrganizationContent(ctx, organizationId).Query(query).Page(page).Limit(limit).Execute()

Search content in an organization

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
    organizationId := "organizationId_example" // string | The unique id of the organization
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.SearchOrganizationContent(context.Background(), organizationId).Query(query).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.SearchOrganizationContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchOrganizationContent`: SearchContent200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.SearchOrganizationContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrganizationContentRequest struct via the builder pattern


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


## SetUserAsSSOMemberForOrganization

> OrganizationMember SetUserAsSSOMemberForOrganization(ctx, organizationId, userId).Execute()

Set a user as an SSO member of an organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.SetUserAsSSOMemberForOrganization(context.Background(), organizationId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.SetUserAsSSOMemberForOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUserAsSSOMemberForOrganization`: OrganizationMember
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.SetUserAsSSOMemberForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserAsSSOMemberForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[user](../README.md#user)

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
    resp, r, err := apiClient.OrganizationsApi.SetupDirectorySync(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.SetupDirectorySync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetupDirectorySync`: SetupDirectorySync200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.SetupDirectorySync`: %v\n", resp)
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
    resp, r, err := apiClient.OrganizationsApi.SyncDirectorySyncGroupsToTeams(context.Background(), organizationId).SyncDirectorySyncGroupsToTeamsRequest(syncDirectorySyncGroupsToTeamsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.SyncDirectorySyncGroupsToTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncDirectorySyncGroupsToTeams`: SyncDirectorySyncGroupsToTeams200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.SyncDirectorySyncGroupsToTeams`: %v\n", resp)
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


## TransferOrganization

> TransferOrganization200Response TransferOrganization(ctx, organizationId).TransferOrganizationRequest(transferOrganizationRequest).Execute()

Transfer one organization (source) into another organization (target).

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
    transferOrganizationRequest := *openapiclient.NewTransferOrganizationRequest("Source_example") // TransferOrganizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.TransferOrganization(context.Background(), organizationId).TransferOrganizationRequest(transferOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.TransferOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferOrganization`: TransferOrganization200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.TransferOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferOrganizationRequest** | [**TransferOrganizationRequest**](TransferOrganizationRequest.md) |  | 

### Return type

[**TransferOrganization200Response**](TransferOrganization200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironment

> Environment UpdateEnvironment(ctx, organizationId, environmentName).UpdateEnvironment(updateEnvironment).Execute()

Update an existing environment within an organization

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
    environmentName := "environmentName_example" // string | The unique name of the environment within the organization
    updateEnvironment := *openapiclient.NewUpdateEnvironment() // UpdateEnvironment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateEnvironment(context.Background(), organizationId, environmentName).UpdateEnvironment(updateEnvironment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**environmentName** | **string** | The unique name of the environment within the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEnvironment** | [**UpdateEnvironment**](UpdateEnvironment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemberInOrganizationById

> OrganizationMember UpdateMemberInOrganizationById(ctx, organizationId, userId).UpdateMemberInOrganizationByIdRequest(updateMemberInOrganizationByIdRequest).Execute()

Update specified organization member



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
    updateMemberInOrganizationByIdRequest := *openapiclient.NewUpdateMemberInOrganizationByIdRequest() // UpdateMemberInOrganizationByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateMemberInOrganizationById(context.Background(), organizationId, userId).UpdateMemberInOrganizationByIdRequest(updateMemberInOrganizationByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateMemberInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMemberInOrganizationById`: OrganizationMember
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateMemberInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**userId** | **string** | The unique ID of the User | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMemberInOrganizationByIdRequest** | [**UpdateMemberInOrganizationByIdRequest**](UpdateMemberInOrganizationByIdRequest.md) |  | 

### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMembersInOrganizationTeam

> UpdateMembersInOrganizationTeam(ctx, organizationId, teamId).UpdateMembersInOrganizationTeamRequest(updateMembersInOrganizationTeamRequest).Execute()

Updates members of an organization team



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
    teamId := "teamId_example" // string | The unique ID of the Team
    updateMembersInOrganizationTeamRequest := *openapiclient.NewUpdateMembersInOrganizationTeamRequest() // UpdateMembersInOrganizationTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsApi.UpdateMembersInOrganizationTeam(context.Background(), organizationId, teamId).UpdateMembersInOrganizationTeamRequest(updateMembersInOrganizationTeamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateMembersInOrganizationTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**teamId** | **string** | The unique ID of the Team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMembersInOrganizationTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMembersInOrganizationTeamRequest** | [**UpdateMembersInOrganizationTeamRequest**](UpdateMembersInOrganizationTeamRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCustomField

> CustomField UpdateOrganizationCustomField(ctx, organizationId, fieldName).UpdateOrganizationCustomFieldRequest(updateOrganizationCustomFieldRequest).Execute()

Update a custom field in an organization

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
    fieldName := "fieldName_example" // string | The name of the custom field
    updateOrganizationCustomFieldRequest := *openapiclient.NewUpdateOrganizationCustomFieldRequest() // UpdateOrganizationCustomFieldRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateOrganizationCustomField(context.Background(), organizationId, fieldName).UpdateOrganizationCustomFieldRequest(updateOrganizationCustomFieldRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationCustomField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCustomField`: CustomField
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateOrganizationCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**fieldName** | **string** | The name of the custom field | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationCustomFieldRequest** | [**UpdateOrganizationCustomFieldRequest**](UpdateOrganizationCustomFieldRequest.md) |  | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationMemberLastSeenAt

> UpdateOrganizationMemberLastSeenAt(ctx, organizationId).Execute()

Update organization member's \"last seen at\" timestamp.



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
    r, err := apiClient.OrganizationsApi.UpdateOrganizationMemberLastSeenAt(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateOrganizationMemberLastSeenAt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationMemberLastSeenAtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamInOrganizationById

> OrganizationTeam UpdateTeamInOrganizationById(ctx, organizationId, teamId).UpdateTeamInOrganizationByIdRequest(updateTeamInOrganizationByIdRequest).Execute()

Update specified organization team



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
    teamId := "teamId_example" // string | The unique ID of the Team
    updateTeamInOrganizationByIdRequest := *openapiclient.NewUpdateTeamInOrganizationByIdRequest("Title_example") // UpdateTeamInOrganizationByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpdateTeamInOrganizationById(context.Background(), organizationId, teamId).UpdateTeamInOrganizationByIdRequest(updateTeamInOrganizationByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpdateTeamInOrganizationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamInOrganizationById`: OrganizationTeam
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpdateTeamInOrganizationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 
**teamId** | **string** | The unique ID of the Team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamInOrganizationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateTeamInOrganizationByIdRequest** | [**UpdateTeamInOrganizationByIdRequest**](UpdateTeamInOrganizationByIdRequest.md) |  | 

### Return type

[**OrganizationTeam**](OrganizationTeam.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeOrganizationPlan

> BillingUpgrade UpgradeOrganizationPlan(ctx, organizationId).RequestUpgradeOrganizationBilling(requestUpgradeOrganizationBilling).Execute()

Upgrade an organization's billing plan

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
    requestUpgradeOrganizationBilling := *openapiclient.NewRequestUpgradeOrganizationBilling(openapiclient.BillingProduct("free"), openapiclient.BillingInterval("monthly")) // RequestUpgradeOrganizationBilling | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.UpgradeOrganizationPlan(context.Background(), organizationId).RequestUpgradeOrganizationBilling(requestUpgradeOrganizationBilling).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UpgradeOrganizationPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeOrganizationPlan`: BillingUpgrade
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.UpgradeOrganizationPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeOrganizationPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpgradeOrganizationBilling** | [**RequestUpgradeOrganizationBilling**](RequestUpgradeOrganizationBilling.md) |  | 

### Return type

[**BillingUpgrade**](BillingUpgrade.md)

### Authorization

[user-internal](../README.md#user-internal)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

