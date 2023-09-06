# \SpacesApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChangeRequest**](SpacesApi.md#CreateChangeRequest) | **Post** /spaces/{spaceId}/change-requests | Create a new change request for a space.
[**CreateSpace**](SpacesApi.md#CreateSpace) | **Post** /orgs/{organizationId}/spaces | Create an organization space
[**CreateSpaceRelation**](SpacesApi.md#CreateSpaceRelation) | **Post** /spaces/{spaceId}/relations | Create a new relation between a source space and a target space
[**DeleteCommentInChangeRequest**](SpacesApi.md#DeleteCommentInChangeRequest) | **Delete** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId} | Delete a comment in a change request.
[**DeleteCommentInSpace**](SpacesApi.md#DeleteCommentInSpace) | **Delete** /spaces/{spaceId}/comments/{commentId} | Delete a comment in a space.
[**DeleteCommentReplyInChangeRequest**](SpacesApi.md#DeleteCommentReplyInChangeRequest) | **Delete** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies/{commentReplyId} | Delete a comment reply in a change request.
[**DeleteCommentReplyInSpace**](SpacesApi.md#DeleteCommentReplyInSpace) | **Delete** /spaces/{spaceId}/comments/{commentId}/replies/{commentReplyId} | Delete a comment reply in a space.
[**DeleteSpaceRelation**](SpacesApi.md#DeleteSpaceRelation) | **Delete** /spaces/{spaceId}/relations/{targetSpaceId} | Delete a relation between spaces
[**DuplicateSpace**](SpacesApi.md#DuplicateSpace) | **Post** /spaces/{spaceId}/duplicate | Create a duplicate of the space.
[**ExportToGitRepository**](SpacesApi.md#ExportToGitRepository) | **Post** /spaces/{spaceId}/git/export | Export the space content to a Git repository.
[**GetChangeRequestById**](SpacesApi.md#GetChangeRequestById) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId} | Get the change request with the given id.
[**GetCommentInChangeRequest**](SpacesApi.md#GetCommentInChangeRequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId} | Get a comment in a change request.
[**GetCommentInSpace**](SpacesApi.md#GetCommentInSpace) | **Get** /spaces/{spaceId}/comments/{commentId} | Get a comment in a space.
[**GetCommentReplyInChangeRequest**](SpacesApi.md#GetCommentReplyInChangeRequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies/{commentReplyId} | Get a comment reply in a change request.
[**GetCommentReplyInSpace**](SpacesApi.md#GetCommentReplyInSpace) | **Get** /spaces/{spaceId}/comments/{commentId}/replies/{commentReplyId} | Get a comment reply in a space.
[**GetContentAnalyticsForSpaceById**](SpacesApi.md#GetContentAnalyticsForSpaceById) | **Get** /spaces/{spaceId}/insights/content | Get content analytics for a given space.
[**GetContributorsByChangeRequestId**](SpacesApi.md#GetContributorsByChangeRequestId) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/contributors | Get all contributors for the change request with the given id.
[**GetRequestedReviewersByChangeRequestId**](SpacesApi.md#GetRequestedReviewersByChangeRequestId) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/requested-reviewers | Get all requested reviewers for a change request.
[**GetReviewsByChangeRequestId**](SpacesApi.md#GetReviewsByChangeRequestId) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/reviews | Get all reviews for a change request.
[**GetSearchAnalyticsForSpaceById**](SpacesApi.md#GetSearchAnalyticsForSpaceById) | **Get** /spaces/{spaceId}/insights/search | Get an overview of the top search queries in a space.
[**GetSpaceById**](SpacesApi.md#GetSpaceById) | **Get** /spaces/{spaceId} | Get the details about a space.
[**GetSpaceCustomFields**](SpacesApi.md#GetSpaceCustomFields) | **Get** /spaces/{spaceId}/custom-fields | Get the values of the custom fields set on a space
[**GetSpaceGitInfo**](SpacesApi.md#GetSpaceGitInfo) | **Get** /spaces/{spaceId}/git/info | Get metadata about the Git Sync provider currently set up on the space
[**GetSpacePublishingAuthById**](SpacesApi.md#GetSpacePublishingAuthById) | **Get** /spaces/{spaceId}/publishing/auth | Get the publishing authentication configuration for a space.
[**GetSpaceRelation**](SpacesApi.md#GetSpaceRelation) | **Get** /spaces/{spaceId}/relations/{targetSpaceId} | Get the relation between 2 spaces.
[**GetTrafficAnalyticsForSpaceById**](SpacesApi.md#GetTrafficAnalyticsForSpaceById) | **Get** /spaces/{spaceId}/insights/traffic | Get traffic page views for a given space
[**ImportGitRepository**](SpacesApi.md#ImportGitRepository) | **Post** /spaces/{spaceId}/git/import | Import a Git repository
[**ListChangeRequestsForSpace**](SpacesApi.md#ListChangeRequestsForSpace) | **Get** /spaces/{spaceId}/change-requests | List change requests for a space.
[**ListCommentRepliesInChangeRequest**](SpacesApi.md#ListCommentRepliesInChangeRequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies | List all the replies to a comment in a change request.
[**ListCommentRepliesInSpace**](SpacesApi.md#ListCommentRepliesInSpace) | **Get** /spaces/{spaceId}/comments/{commentId}/replies | List all the replies to a comment in a space.
[**ListCommentsInChangeRequest**](SpacesApi.md#ListCommentsInChangeRequest) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/comments | List all the comments in a change request.
[**ListCommentsInSpace**](SpacesApi.md#ListCommentsInSpace) | **Get** /spaces/{spaceId}/comments | List all the comments in a space.
[**ListPermissionsAggregateInCollection**](SpacesApi.md#ListPermissionsAggregateInCollection) | **Get** /collections/{collectionId}/permissions/aggregate | List permissions for all users in a collection.
[**ListPermissionsAggregateInSpace**](SpacesApi.md#ListPermissionsAggregateInSpace) | **Get** /spaces/{spaceId}/permissions/aggregate | List permissions for all users in a space.
[**ListSpaceRelations**](SpacesApi.md#ListSpaceRelations) | **Get** /spaces/{spaceId}/relations | List all relations for a space
[**ListSpaceRelationsInOrganization**](SpacesApi.md#ListSpaceRelationsInOrganization) | **Get** /orgs/{organizationId}/space-relations | List all relations between spaces in an organization
[**ListSpacesForOrganizationMember**](SpacesApi.md#ListSpacesForOrganizationMember) | **Get** /orgs/{organizationId}/members/{userId}/spaces | List permissions accross all spaces for a member of an organization
[**MergeChangeRequest**](SpacesApi.md#MergeChangeRequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/merge | Merge a change request in the primary content of a space.
[**PostCommentInChangeRequest**](SpacesApi.md#PostCommentInChangeRequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/comments | Post a new comment in a change request.
[**PostCommentInSpace**](SpacesApi.md#PostCommentInSpace) | **Post** /spaces/{spaceId}/comments | Post a new comment in a space.
[**PostCommentReplyInChangeRequest**](SpacesApi.md#PostCommentReplyInChangeRequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies | Post a new reply to a comment in a change request.
[**PostCommentReplyInSpace**](SpacesApi.md#PostCommentReplyInSpace) | **Post** /spaces/{spaceId}/comments/{commentId}/replies | Post a new reply to a comment in a space.
[**RequestReviewersForChangeRequest**](SpacesApi.md#RequestReviewersForChangeRequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/requested-reviewers | Request reviewers on a change request. Note that requesting a review from teams is not yet supported.
[**SubmitChangeRequestReview**](SpacesApi.md#SubmitChangeRequestReview) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/reviews | Submit a review for a change request.
[**TrackViewInSpaceById**](SpacesApi.md#TrackViewInSpaceById) | **Post** /spaces/{spaceId}/insights/track_view | 
[**UpdateChangeRequest**](SpacesApi.md#UpdateChangeRequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/update | Update a change-request with changes from primary content.
[**UpdateCommentInChangeRequest**](SpacesApi.md#UpdateCommentInChangeRequest) | **Put** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId} | Update a comment in a change request.
[**UpdateCommentInSpace**](SpacesApi.md#UpdateCommentInSpace) | **Put** /spaces/{spaceId}/comments/{commentId} | Update a comment in a space.
[**UpdateCommentReplyInChangeRequest**](SpacesApi.md#UpdateCommentReplyInChangeRequest) | **Put** /spaces/{spaceId}/change-requests/{changeRequestId}/comments/{commentId}/replies/{commentReplyId} | Update a comment reply in a change request.
[**UpdateCommentReplyInSpace**](SpacesApi.md#UpdateCommentReplyInSpace) | **Put** /spaces/{spaceId}/comments/{commentId}/replies/{commentReplyId} | Update a comment reply in a space.
[**UpdateSpaceById**](SpacesApi.md#UpdateSpaceById) | **Patch** /spaces/{spaceId} | Update the details of a space
[**UpdateSpaceCustomFields**](SpacesApi.md#UpdateSpaceCustomFields) | **Patch** /spaces/{spaceId}/custom-fields | Update the custom fields in a space
[**UpdateSpacePublishingAuthById**](SpacesApi.md#UpdateSpacePublishingAuthById) | **Post** /spaces/{spaceId}/publishing/auth | Update the publishing authentication configuration for a space.



## CreateChangeRequest

> CreateChangeRequest201Response CreateChangeRequest(ctx, spaceId).RequestCreateChangeRequest(requestCreateChangeRequest).Execute()

Create a new change request for a space.

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
    requestCreateChangeRequest := *openapiclient.NewRequestCreateChangeRequest() // RequestCreateChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.CreateChangeRequest(context.Background(), spaceId).RequestCreateChangeRequest(requestCreateChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.CreateChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChangeRequest`: CreateChangeRequest201Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.CreateChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestCreateChangeRequest** | [**RequestCreateChangeRequest**](RequestCreateChangeRequest.md) |  | 

### Return type

[**CreateChangeRequest201Response**](CreateChangeRequest201Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpace

> Space CreateSpace(ctx, organizationId).RequestCreateSpace(requestCreateSpace).Execute()

Create an organization space

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
    requestCreateSpace := *openapiclient.NewRequestCreateSpace() // RequestCreateSpace |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.CreateSpace(context.Background(), organizationId).RequestCreateSpace(requestCreateSpace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.CreateSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpace`: Space
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.CreateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestCreateSpace** | [**RequestCreateSpace**](RequestCreateSpace.md) |  | 

### Return type

[**Space**](Space.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpaceRelation

> SpaceRelation CreateSpaceRelation(ctx, spaceId).CreateSpaceRelationRequest(createSpaceRelationRequest).Execute()

Create a new relation between a source space and a target space

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
    createSpaceRelationRequest := *openapiclient.NewCreateSpaceRelationRequest("Target_example", openapiclient.SpaceRelationType("dependsOn")) // CreateSpaceRelationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.CreateSpaceRelation(context.Background(), spaceId).CreateSpaceRelationRequest(createSpaceRelationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.CreateSpaceRelation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpaceRelation`: SpaceRelation
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.CreateSpaceRelation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpaceRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSpaceRelationRequest** | [**CreateSpaceRelationRequest**](CreateSpaceRelationRequest.md) |  | 

### Return type

[**SpaceRelation**](SpaceRelation.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommentInChangeRequest

> DeleteCommentInChangeRequest(ctx, spaceId, changeRequestId, commentId).Execute()

Delete a comment in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.DeleteCommentInChangeRequest(context.Background(), spaceId, changeRequestId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.DeleteCommentInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentInChangeRequestRequest struct via the builder pattern


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


## DeleteCommentInSpace

> DeleteCommentInSpace(ctx, spaceId, commentId).Execute()

Delete a comment in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.DeleteCommentInSpace(context.Background(), spaceId, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.DeleteCommentInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentInSpaceRequest struct via the builder pattern


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


## DeleteCommentReplyInChangeRequest

> DeleteCommentReplyInChangeRequest(ctx, spaceId, changeRequestId, commentId, commentReplyId).Execute()

Delete a comment reply in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment
    commentReplyId := "commentReplyId_example" // string | The unique id of the comment reply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.DeleteCommentReplyInChangeRequest(context.Background(), spaceId, changeRequestId, commentId, commentReplyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.DeleteCommentReplyInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 
**commentReplyId** | **string** | The unique id of the comment reply | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentReplyInChangeRequestRequest struct via the builder pattern


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


## DeleteCommentReplyInSpace

> DeleteCommentReplyInSpace(ctx, spaceId, commentId, commentReplyId).Execute()

Delete a comment reply in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment
    commentReplyId := "commentReplyId_example" // string | The unique id of the comment reply

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.DeleteCommentReplyInSpace(context.Background(), spaceId, commentId, commentReplyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.DeleteCommentReplyInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 
**commentReplyId** | **string** | The unique id of the comment reply | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentReplyInSpaceRequest struct via the builder pattern


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


## DeleteSpaceRelation

> DeleteSpaceRelation(ctx, spaceId, targetSpaceId).Execute()

Delete a relation between spaces

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
    targetSpaceId := "targetSpaceId_example" // string | The ID of the other space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.DeleteSpaceRelation(context.Background(), spaceId, targetSpaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.DeleteSpaceRelation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**targetSpaceId** | **string** | The ID of the other space | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpaceRelationRequest struct via the builder pattern


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


## DuplicateSpace

> Space DuplicateSpace(ctx, spaceId).Execute()

Create a duplicate of the space.

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
    resp, r, err := apiClient.SpacesApi.DuplicateSpace(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.DuplicateSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DuplicateSpace`: Space
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.DuplicateSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Space**](Space.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportToGitRepository

> ExportToGitRepository(ctx, spaceId).RequestExportToGitRepository(requestExportToGitRepository).Execute()

Export the space content to a Git repository.

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
    requestExportToGitRepository := *openapiclient.NewRequestExportToGitRepository("Url_example", "Ref_example", "CommitMessage_example") // RequestExportToGitRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.ExportToGitRepository(context.Background(), spaceId).RequestExportToGitRepository(requestExportToGitRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ExportToGitRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportToGitRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestExportToGitRepository** | [**RequestExportToGitRepository**](RequestExportToGitRepository.md) |  | 

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


## GetChangeRequestById

> ChangeRequest GetChangeRequestById(ctx, spaceId, changeRequestId).Execute()

Get the change request with the given id.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetChangeRequestById(context.Background(), spaceId, changeRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetChangeRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChangeRequestById`: ChangeRequest
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetChangeRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChangeRequest**](ChangeRequest.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentInChangeRequest

> Comment GetCommentInChangeRequest(ctx, spaceId, changeRequestId, commentId).Format(format).Execute()

Get a comment in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetCommentInChangeRequest(context.Background(), spaceId, changeRequestId, commentId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetCommentInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommentInChangeRequest`: Comment
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetCommentInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **string** | Output format for the content. | 

### Return type

[**Comment**](Comment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentInSpace

> Comment GetCommentInSpace(ctx, spaceId, commentId).Format(format).Execute()

Get a comment in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetCommentInSpace(context.Background(), spaceId, commentId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetCommentInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommentInSpace`: Comment
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetCommentInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Output format for the content. | 

### Return type

[**Comment**](Comment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentReplyInChangeRequest

> CommentReply GetCommentReplyInChangeRequest(ctx, spaceId, changeRequestId, commentId, commentReplyId).Format(format).Execute()

Get a comment reply in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment
    commentReplyId := "commentReplyId_example" // string | The unique id of the comment reply
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetCommentReplyInChangeRequest(context.Background(), spaceId, changeRequestId, commentId, commentReplyId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetCommentReplyInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommentReplyInChangeRequest`: CommentReply
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetCommentReplyInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 
**commentReplyId** | **string** | The unique id of the comment reply | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentReplyInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | **string** | Output format for the content. | 

### Return type

[**CommentReply**](CommentReply.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentReplyInSpace

> CommentReply GetCommentReplyInSpace(ctx, spaceId, commentId, commentReplyId).Format(format).Execute()

Get a comment reply in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment
    commentReplyId := "commentReplyId_example" // string | The unique id of the comment reply
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetCommentReplyInSpace(context.Background(), spaceId, commentId, commentReplyId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetCommentReplyInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommentReplyInSpace`: CommentReply
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetCommentReplyInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 
**commentReplyId** | **string** | The unique id of the comment reply | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentReplyInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **string** | Output format for the content. | 

### Return type

[**CommentReply**](CommentReply.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentAnalyticsForSpaceById

> AnalyticsContentPages GetContentAnalyticsForSpaceById(ctx, spaceId).Execute()

Get content analytics for a given space.

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
    resp, r, err := apiClient.SpacesApi.GetContentAnalyticsForSpaceById(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetContentAnalyticsForSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentAnalyticsForSpaceById`: AnalyticsContentPages
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetContentAnalyticsForSpaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentAnalyticsForSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsContentPages**](AnalyticsContentPages.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContributorsByChangeRequestId

> GetContributorsByChangeRequestId200Response GetContributorsByChangeRequestId(ctx, spaceId, changeRequestId).Execute()

Get all contributors for the change request with the given id.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetContributorsByChangeRequestId(context.Background(), spaceId, changeRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetContributorsByChangeRequestId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContributorsByChangeRequestId`: GetContributorsByChangeRequestId200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetContributorsByChangeRequestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContributorsByChangeRequestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContributorsByChangeRequestId200Response**](GetContributorsByChangeRequestId200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestedReviewersByChangeRequestId

> GetRequestedReviewersByChangeRequestId200Response GetRequestedReviewersByChangeRequestId(ctx, spaceId, changeRequestId).Page(page).Limit(limit).Execute()

Get all requested reviewers for a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetRequestedReviewersByChangeRequestId(context.Background(), spaceId, changeRequestId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetRequestedReviewersByChangeRequestId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestedReviewersByChangeRequestId`: GetRequestedReviewersByChangeRequestId200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetRequestedReviewersByChangeRequestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestedReviewersByChangeRequestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**GetRequestedReviewersByChangeRequestId200Response**](GetRequestedReviewersByChangeRequestId200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReviewsByChangeRequestId

> GetReviewsByChangeRequestId200Response GetReviewsByChangeRequestId(ctx, spaceId, changeRequestId).Format(format).Page(page).Limit(limit).Execute()

Get all reviews for a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    format := "format_example" // string | Output format for the content. (optional)
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetReviewsByChangeRequestId(context.Background(), spaceId, changeRequestId).Format(format).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetReviewsByChangeRequestId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReviewsByChangeRequestId`: GetReviewsByChangeRequestId200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetReviewsByChangeRequestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewsByChangeRequestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Output format for the content. | 
 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**GetReviewsByChangeRequestId200Response**](GetReviewsByChangeRequestId200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchAnalyticsForSpaceById

> AnalyticsTopSearches GetSearchAnalyticsForSpaceById(ctx, spaceId).Period(period).Execute()

Get an overview of the top search queries in a space.

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
    period := openapiclient.AnalyticsSearchPeriod("last_month") // AnalyticsSearchPeriod |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetSearchAnalyticsForSpaceById(context.Background(), spaceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetSearchAnalyticsForSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchAnalyticsForSpaceById`: AnalyticsTopSearches
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetSearchAnalyticsForSpaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchAnalyticsForSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | [**AnalyticsSearchPeriod**](AnalyticsSearchPeriod.md) |  | 

### Return type

[**AnalyticsTopSearches**](AnalyticsTopSearches.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceById

> Space GetSpaceById(ctx, spaceId).Execute()

Get the details about a space.

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
    resp, r, err := apiClient.SpacesApi.GetSpaceById(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpaceById`: Space
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetSpaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Space**](Space.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceCustomFields

> []CustomFieldValuesInner GetSpaceCustomFields(ctx, spaceId).Execute()

Get the values of the custom fields set on a space

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
    resp, r, err := apiClient.SpacesApi.GetSpaceCustomFields(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetSpaceCustomFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpaceCustomFields`: []CustomFieldValuesInner
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetSpaceCustomFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomFieldValuesInner**](CustomFieldValuesInner.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceGitInfo

> GitSyncState GetSpaceGitInfo(ctx, spaceId).Execute()

Get metadata about the Git Sync provider currently set up on the space

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
    resp, r, err := apiClient.SpacesApi.GetSpaceGitInfo(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetSpaceGitInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpaceGitInfo`: GitSyncState
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetSpaceGitInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceGitInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitSyncState**](GitSyncState.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpacePublishingAuthById

> ContentPublishingAuth GetSpacePublishingAuthById(ctx, spaceId).Execute()

Get the publishing authentication configuration for a space.

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
    resp, r, err := apiClient.SpacesApi.GetSpacePublishingAuthById(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetSpacePublishingAuthById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpacePublishingAuthById`: ContentPublishingAuth
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetSpacePublishingAuthById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpacePublishingAuthByIdRequest struct via the builder pattern


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


## GetSpaceRelation

> SpaceRelation GetSpaceRelation(ctx, spaceId, targetSpaceId).Execute()

Get the relation between 2 spaces.

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
    targetSpaceId := "targetSpaceId_example" // string | The ID of the other space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetSpaceRelation(context.Background(), spaceId, targetSpaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetSpaceRelation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpaceRelation`: SpaceRelation
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetSpaceRelation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**targetSpaceId** | **string** | The ID of the other space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceRelationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SpaceRelation**](SpaceRelation.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrafficAnalyticsForSpaceById

> AnalyticsTrafficPageViews GetTrafficAnalyticsForSpaceById(ctx, spaceId).Interval(interval).Execute()

Get traffic page views for a given space



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
    interval := openapiclient.AnalyticsTrafficInterval("daily") // AnalyticsTrafficInterval |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.GetTrafficAnalyticsForSpaceById(context.Background(), spaceId).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.GetTrafficAnalyticsForSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrafficAnalyticsForSpaceById`: AnalyticsTrafficPageViews
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.GetTrafficAnalyticsForSpaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrafficAnalyticsForSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interval** | [**AnalyticsTrafficInterval**](AnalyticsTrafficInterval.md) |  | 

### Return type

[**AnalyticsTrafficPageViews**](AnalyticsTrafficPageViews.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportGitRepository

> ImportGitRepository(ctx, spaceId).RequestImportGitRepository(requestImportGitRepository).Execute()

Import a Git repository

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
    requestImportGitRepository := *openapiclient.NewRequestImportGitRepository("Url_example", "Ref_example") // RequestImportGitRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.ImportGitRepository(context.Background(), spaceId).RequestImportGitRepository(requestImportGitRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ImportGitRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportGitRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestImportGitRepository** | [**RequestImportGitRepository**](RequestImportGitRepository.md) |  | 

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


## ListChangeRequestsForSpace

> ListChangeRequestsForSpace200Response ListChangeRequestsForSpace(ctx, spaceId).Page(page).Limit(limit).Status(status).Execute()

List change requests for a space.

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
    status := openapiclient.ChangeRequestStatus("draft") // ChangeRequestStatus | If defined, only change requests matching this status will be returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.ListChangeRequestsForSpace(context.Background(), spaceId).Page(page).Limit(limit).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListChangeRequestsForSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChangeRequestsForSpace`: ListChangeRequestsForSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListChangeRequestsForSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChangeRequestsForSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **status** | [**ChangeRequestStatus**](ChangeRequestStatus.md) | If defined, only change requests matching this status will be returned | 

### Return type

[**ListChangeRequestsForSpace200Response**](ListChangeRequestsForSpace200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommentRepliesInChangeRequest

> ListCommentRepliesInChangeRequest200Response ListCommentRepliesInChangeRequest(ctx, spaceId, changeRequestId, commentId).Page(page).Limit(limit).Format(format).Execute()

List all the replies to a comment in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.ListCommentRepliesInChangeRequest(context.Background(), spaceId, changeRequestId, commentId).Page(page).Limit(limit).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListCommentRepliesInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommentRepliesInChangeRequest`: ListCommentRepliesInChangeRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListCommentRepliesInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommentRepliesInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **format** | **string** | Output format for the content. | 

### Return type

[**ListCommentRepliesInChangeRequest200Response**](ListCommentRepliesInChangeRequest200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommentRepliesInSpace

> ListCommentRepliesInChangeRequest200Response ListCommentRepliesInSpace(ctx, spaceId, commentId).Page(page).Limit(limit).Format(format).Execute()

List all the replies to a comment in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.ListCommentRepliesInSpace(context.Background(), spaceId, commentId).Page(page).Limit(limit).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListCommentRepliesInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommentRepliesInSpace`: ListCommentRepliesInChangeRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListCommentRepliesInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommentRepliesInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **format** | **string** | Output format for the content. | 

### Return type

[**ListCommentRepliesInChangeRequest200Response**](ListCommentRepliesInChangeRequest200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommentsInChangeRequest

> ListCommentsInChangeRequest200Response ListCommentsInChangeRequest(ctx, spaceId, changeRequestId).Page(page).Limit(limit).Format(format).Status(status).Execute()

List all the comments in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    format := "format_example" // string | Output format for the content. (optional)
    status := "status_example" // string | When provided, only comments with the given status are returned. Only \"all\" is supported for now. (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.ListCommentsInChangeRequest(context.Background(), spaceId, changeRequestId).Page(page).Limit(limit).Format(format).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListCommentsInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommentsInChangeRequest`: ListCommentsInChangeRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListCommentsInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommentsInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **format** | **string** | Output format for the content. | 
 **status** | **string** | When provided, only comments with the given status are returned. Only \&quot;all\&quot; is supported for now. | [default to &quot;all&quot;]

### Return type

[**ListCommentsInChangeRequest200Response**](ListCommentsInChangeRequest200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommentsInSpace

> ListCommentsInChangeRequest200Response ListCommentsInSpace(ctx, spaceId).Page(page).Limit(limit).Status(status).Format(format).Execute()

List all the comments in a space.

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
    status := "status_example" // string | When provided, only comments with the given status are returned. Only \"all\" is supported for now. (optional) (default to "all")
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.ListCommentsInSpace(context.Background(), spaceId).Page(page).Limit(limit).Status(status).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListCommentsInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommentsInSpace`: ListCommentsInChangeRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListCommentsInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommentsInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **status** | **string** | When provided, only comments with the given status are returned. Only \&quot;all\&quot; is supported for now. | [default to &quot;all&quot;]
 **format** | **string** | Output format for the content. | 

### Return type

[**ListCommentsInChangeRequest200Response**](ListCommentsInChangeRequest200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.SpacesApi.ListPermissionsAggregateInCollection(context.Background(), collectionId).Page(page).Limit(limit).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListPermissionsAggregateInCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissionsAggregateInCollection`: ListPermissionsAggregateInSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListPermissionsAggregateInCollection`: %v\n", resp)
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
    resp, r, err := apiClient.SpacesApi.ListPermissionsAggregateInSpace(context.Background(), spaceId).Page(page).Limit(limit).Role(role).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListPermissionsAggregateInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissionsAggregateInSpace`: ListPermissionsAggregateInSpace200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListPermissionsAggregateInSpace`: %v\n", resp)
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


## ListSpaceRelations

> ListSpaceRelations200Response ListSpaceRelations(ctx, spaceId).Page(page).Limit(limit).Execute()

List all relations for a space

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.ListSpaceRelations(context.Background(), spaceId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListSpaceRelations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpaceRelations`: ListSpaceRelations200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListSpaceRelations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpaceRelationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListSpaceRelations200Response**](ListSpaceRelations200Response.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpaceRelationsInOrganization

> ListSpaceRelationsInOrganization200Response ListSpaceRelationsInOrganization(ctx, organizationId).Page(page).Limit(limit).Execute()

List all relations between spaces in an organization

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
    resp, r, err := apiClient.SpacesApi.ListSpaceRelationsInOrganization(context.Background(), organizationId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListSpaceRelationsInOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpaceRelationsInOrganization`: ListSpaceRelationsInOrganization200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListSpaceRelationsInOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The unique id of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpaceRelationsInOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListSpaceRelationsInOrganization200Response**](ListSpaceRelationsInOrganization200Response.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

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
    resp, r, err := apiClient.SpacesApi.ListSpacesForOrganizationMember(context.Background(), organizationId, userId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.ListSpacesForOrganizationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpacesForOrganizationMember`: ListSpacesForOrganizationMember200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.ListSpacesForOrganizationMember`: %v\n", resp)
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


## MergeChangeRequest

> MergeChangeRequest200Response MergeChangeRequest(ctx, spaceId, changeRequestId).Execute()

Merge a change request in the primary content of a space.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.MergeChangeRequest(context.Background(), spaceId, changeRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.MergeChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MergeChangeRequest`: MergeChangeRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.MergeChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MergeChangeRequest200Response**](MergeChangeRequest200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentInChangeRequest

> Comment PostCommentInChangeRequest(ctx, spaceId, changeRequestId).PostCommentSchema(postCommentSchema).Execute()

Post a new comment in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    postCommentSchema := *openapiclient.NewPostCommentSchema(openapiclient.Document{JSONDocument: openapiclient.NewJSONDocument(*openapiclient.NewJSONDocumentDocument([]map[string]interface{}{map[string]interface{}(123)}))}) // PostCommentSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.PostCommentInChangeRequest(context.Background(), spaceId, changeRequestId).PostCommentSchema(postCommentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.PostCommentInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommentInChangeRequest`: Comment
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.PostCommentInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommentInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postCommentSchema** | [**PostCommentSchema**](PostCommentSchema.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentInSpace

> Comment PostCommentInSpace(ctx, spaceId).PostCommentSchema(postCommentSchema).Execute()

Post a new comment in a space.

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
    postCommentSchema := *openapiclient.NewPostCommentSchema(openapiclient.Document{JSONDocument: openapiclient.NewJSONDocument(*openapiclient.NewJSONDocumentDocument([]map[string]interface{}{map[string]interface{}(123)}))}) // PostCommentSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.PostCommentInSpace(context.Background(), spaceId).PostCommentSchema(postCommentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.PostCommentInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommentInSpace`: Comment
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.PostCommentInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommentInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCommentSchema** | [**PostCommentSchema**](PostCommentSchema.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentReplyInChangeRequest

> CommentReply PostCommentReplyInChangeRequest(ctx, spaceId, changeRequestId, commentId).PostCommentReplySchema(postCommentReplySchema).Format(format).Execute()

Post a new reply to a comment in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment
    postCommentReplySchema := *openapiclient.NewPostCommentReplySchema(openapiclient.Document{JSONDocument: openapiclient.NewJSONDocument(*openapiclient.NewJSONDocumentDocument([]map[string]interface{}{map[string]interface{}(123)}))}) // PostCommentReplySchema | 
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.PostCommentReplyInChangeRequest(context.Background(), spaceId, changeRequestId, commentId).PostCommentReplySchema(postCommentReplySchema).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.PostCommentReplyInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommentReplyInChangeRequest`: CommentReply
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.PostCommentReplyInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommentReplyInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postCommentReplySchema** | [**PostCommentReplySchema**](PostCommentReplySchema.md) |  | 
 **format** | **string** | Output format for the content. | 

### Return type

[**CommentReply**](CommentReply.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentReplyInSpace

> CommentReply PostCommentReplyInSpace(ctx, spaceId, commentId).PostCommentReplySchema(postCommentReplySchema).Execute()

Post a new reply to a comment in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment
    postCommentReplySchema := *openapiclient.NewPostCommentReplySchema(openapiclient.Document{JSONDocument: openapiclient.NewJSONDocument(*openapiclient.NewJSONDocumentDocument([]map[string]interface{}{map[string]interface{}(123)}))}) // PostCommentReplySchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.PostCommentReplyInSpace(context.Background(), spaceId, commentId).PostCommentReplySchema(postCommentReplySchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.PostCommentReplyInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommentReplyInSpace`: CommentReply
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.PostCommentReplyInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommentReplyInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postCommentReplySchema** | [**PostCommentReplySchema**](PostCommentReplySchema.md) |  | 

### Return type

[**CommentReply**](CommentReply.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestReviewersForChangeRequest

> RequestReviewersForChangeRequest200Response RequestReviewersForChangeRequest(ctx, spaceId, changeRequestId).RequestReviewersForChangeRequestRequest(requestReviewersForChangeRequestRequest).Execute()

Request reviewers on a change request. Note that requesting a review from teams is not yet supported.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    requestReviewersForChangeRequestRequest := *openapiclient.NewRequestReviewersForChangeRequestRequest([]string{"Users_example"}) // RequestReviewersForChangeRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.RequestReviewersForChangeRequest(context.Background(), spaceId, changeRequestId).RequestReviewersForChangeRequestRequest(requestReviewersForChangeRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.RequestReviewersForChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestReviewersForChangeRequest`: RequestReviewersForChangeRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.RequestReviewersForChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestReviewersForChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestReviewersForChangeRequestRequest** | [**RequestReviewersForChangeRequestRequest**](RequestReviewersForChangeRequestRequest.md) |  | 

### Return type

[**RequestReviewersForChangeRequest200Response**](RequestReviewersForChangeRequest200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitChangeRequestReview

> ChangeRequestReview SubmitChangeRequestReview(ctx, spaceId, changeRequestId).SubmitChangeRequestReviewRequest(submitChangeRequestReviewRequest).Execute()

Submit a review for a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    submitChangeRequestReviewRequest := *openapiclient.NewSubmitChangeRequestReviewRequest(openapiclient.ChangeRequestReviewStatus("changes-requested")) // SubmitChangeRequestReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.SubmitChangeRequestReview(context.Background(), spaceId, changeRequestId).SubmitChangeRequestReviewRequest(submitChangeRequestReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.SubmitChangeRequestReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitChangeRequestReview`: ChangeRequestReview
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.SubmitChangeRequestReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitChangeRequestReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **submitChangeRequestReviewRequest** | [**SubmitChangeRequestReviewRequest**](SubmitChangeRequestReviewRequest.md) |  | 

### Return type

[**ChangeRequestReview**](ChangeRequestReview.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackViewInSpaceById

> TrackViewInSpaceById(ctx, spaceId).RequestSpaceTrackPageView(requestSpaceTrackPageView).Execute()





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
    requestSpaceTrackPageView := *openapiclient.NewRequestSpaceTrackPageView("PageId_example", *openapiclient.NewRequestSpaceTrackPageViewVisitor("AnonymousId_example", map[string]string{"key": "Inner_example"}, "UserAgent_example"), "Url_example", "Referrer_example") // RequestSpaceTrackPageView | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.TrackViewInSpaceById(context.Background(), spaceId).RequestSpaceTrackPageView(requestSpaceTrackPageView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.TrackViewInSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackViewInSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestSpaceTrackPageView** | [**RequestSpaceTrackPageView**](RequestSpaceTrackPageView.md) |  | 

### Return type

 (empty response body)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChangeRequest

> UpdateChangeRequest200Response UpdateChangeRequest(ctx, spaceId, changeRequestId).Execute()

Update a change-request with changes from primary content.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.UpdateChangeRequest(context.Background(), spaceId, changeRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChangeRequest`: UpdateChangeRequest200Response
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.UpdateChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UpdateChangeRequest200Response**](UpdateChangeRequest200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommentInChangeRequest

> Comment UpdateCommentInChangeRequest(ctx, spaceId, changeRequestId, commentId).UpdateCommentSchema(updateCommentSchema).Execute()

Update a comment in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment
    updateCommentSchema := *openapiclient.NewUpdateCommentSchema() // UpdateCommentSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.UpdateCommentInChangeRequest(context.Background(), spaceId, changeRequestId, commentId).UpdateCommentSchema(updateCommentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateCommentInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommentInChangeRequest`: Comment
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.UpdateCommentInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateCommentSchema** | [**UpdateCommentSchema**](UpdateCommentSchema.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommentInSpace

> Comment UpdateCommentInSpace(ctx, spaceId, commentId).UpdateCommentSchema(updateCommentSchema).Execute()

Update a comment in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment
    updateCommentSchema := *openapiclient.NewUpdateCommentSchema() // UpdateCommentSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.UpdateCommentInSpace(context.Background(), spaceId, commentId).UpdateCommentSchema(updateCommentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateCommentInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommentInSpace`: Comment
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.UpdateCommentInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCommentSchema** | [**UpdateCommentSchema**](UpdateCommentSchema.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommentReplyInChangeRequest

> CommentReply UpdateCommentReplyInChangeRequest(ctx, spaceId, changeRequestId, commentId, commentReplyId).UpdateCommentSchema(updateCommentSchema).Execute()

Update a comment reply in a change request.

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
    changeRequestId := openapiclient.getChangeRequestById_changeRequestId_parameter{Int32: new(int32)} // GetChangeRequestByIdChangeRequestIdParameter | The unique ID of the change request or its number identifier in the space
    commentId := "commentId_example" // string | The unique id of the comment
    commentReplyId := "commentReplyId_example" // string | The unique id of the comment reply
    updateCommentSchema := *openapiclient.NewUpdateCommentSchema() // UpdateCommentSchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.UpdateCommentReplyInChangeRequest(context.Background(), spaceId, changeRequestId, commentId, commentReplyId).UpdateCommentSchema(updateCommentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateCommentReplyInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommentReplyInChangeRequest`: CommentReply
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.UpdateCommentReplyInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**commentId** | **string** | The unique id of the comment | 
**commentReplyId** | **string** | The unique id of the comment reply | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentReplyInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateCommentSchema** | [**UpdateCommentSchema**](UpdateCommentSchema.md) |  | 

### Return type

[**CommentReply**](CommentReply.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCommentReplyInSpace

> CommentReply UpdateCommentReplyInSpace(ctx, spaceId, commentId, commentReplyId).UpdateCommentReplySchema(updateCommentReplySchema).Execute()

Update a comment reply in a space.

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
    commentId := "commentId_example" // string | The unique id of the comment
    commentReplyId := "commentReplyId_example" // string | The unique id of the comment reply
    updateCommentReplySchema := *openapiclient.NewUpdateCommentReplySchema() // UpdateCommentReplySchema | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.UpdateCommentReplyInSpace(context.Background(), spaceId, commentId, commentReplyId).UpdateCommentReplySchema(updateCommentReplySchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateCommentReplyInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCommentReplyInSpace`: CommentReply
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.UpdateCommentReplyInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**commentId** | **string** | The unique id of the comment | 
**commentReplyId** | **string** | The unique id of the comment reply | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentReplyInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateCommentReplySchema** | [**UpdateCommentReplySchema**](UpdateCommentReplySchema.md) |  | 

### Return type

[**CommentReply**](CommentReply.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpaceById

> Space UpdateSpaceById(ctx, spaceId).UpdateSpaceByIdRequest(updateSpaceByIdRequest).Execute()

Update the details of a space

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
    updateSpaceByIdRequest := *openapiclient.NewUpdateSpaceByIdRequest() // UpdateSpaceByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.UpdateSpaceById(context.Background(), spaceId).UpdateSpaceByIdRequest(updateSpaceByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpaceById`: Space
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.UpdateSpaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSpaceByIdRequest** | [**UpdateSpaceByIdRequest**](UpdateSpaceByIdRequest.md) |  | 

### Return type

[**Space**](Space.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpaceCustomFields

> UpdateSpaceCustomFields(ctx, spaceId).UpdateCustomFieldValues(updateCustomFieldValues).Execute()

Update the custom fields in a space

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
    updateCustomFieldValues := *openapiclient.NewUpdateCustomFieldValues(map[string]UpdateCustomFieldValuesValuesValue{"key": *openapiclient.NewUpdateCustomFieldValuesValuesValue(openapiclient.CustomFieldValue{ArrayOfString: new([]string)})}) // UpdateCustomFieldValues | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpacesApi.UpdateSpaceCustomFields(context.Background(), spaceId).UpdateCustomFieldValues(updateCustomFieldValues).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateSpaceCustomFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpaceCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomFieldValues** | [**UpdateCustomFieldValues**](UpdateCustomFieldValues.md) |  | 

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


## UpdateSpacePublishingAuthById

> ContentPublishingAuth UpdateSpacePublishingAuthById(ctx, spaceId).RequestUpdateContentPublishingAuth(requestUpdateContentPublishingAuth).Execute()

Update the publishing authentication configuration for a space.

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
    requestUpdateContentPublishingAuth := *openapiclient.NewRequestUpdateContentPublishingAuth() // RequestUpdateContentPublishingAuth | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpacesApi.UpdateSpacePublishingAuthById(context.Background(), spaceId).RequestUpdateContentPublishingAuth(requestUpdateContentPublishingAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpacesApi.UpdateSpacePublishingAuthById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpacePublishingAuthById`: ContentPublishingAuth
    fmt.Fprintf(os.Stdout, "Response from `SpacesApi.UpdateSpacePublishingAuthById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpacePublishingAuthByIdRequest struct via the builder pattern


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

