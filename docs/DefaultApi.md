# \DefaultApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AskQuery**](DefaultApi.md#AskQuery) | **Post** /search/ask | Ask a question to an AI across spaces that is accessible by the currently authenticated target.
[**AskQueryInSpace**](DefaultApi.md#AskQueryInSpace) | **Post** /spaces/{spaceId}/search/ask | Ask a question to an AI within the context of the space.
[**AskQueryInSpaceWithGet**](DefaultApi.md#AskQueryInSpaceWithGet) | **Get** /spaces/{spaceId}/search/ask | Ask a question to an AI within the context of the space.
[**AskQueryWithGet**](DefaultApi.md#AskQueryWithGet) | **Get** /search/ask | Ask a question to an AI across spaces that is accessible by the currently authenticated target.
[**GetCurrentRevision**](DefaultApi.md#GetCurrentRevision) | **Get** /spaces/{spaceId}/content | Get the current primary content revision for a space
[**GetPageById**](DefaultApi.md#GetPageById) | **Get** /spaces/{spaceId}/content/page/{pageId} | Get a page by its ID in the primary content.
[**GetPageByPath**](DefaultApi.md#GetPageByPath) | **Get** /spaces/{spaceId}/content/path/{pagePath} | Get a page by its path in the primary content.
[**GetPageInChangeRequestById**](DefaultApi.md#GetPageInChangeRequestById) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content/page/{pageId} | Get a page by its ID in a change request.
[**GetPageInChangeRequestByPath**](DefaultApi.md#GetPageInChangeRequestByPath) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content/path/{pagePath} | Get a page by its path in a change request.
[**GetPageInRevisionById**](DefaultApi.md#GetPageInRevisionById) | **Get** /spaces/{spaceId}/revisions/{revisionId}/page/{pageId} | Get a page by its ID in a revision.
[**GetPageInRevisionByPath**](DefaultApi.md#GetPageInRevisionByPath) | **Get** /spaces/{spaceId}/revisions/{revisionId}/path/{pagePath} | Get a page by its path in a revision.
[**GetRecommendedQuestions**](DefaultApi.md#GetRecommendedQuestions) | **Post** /search/questions | Get a list of questions recommended by AI for a list of content.
[**GetRecommendedQuestionsInSpace**](DefaultApi.md#GetRecommendedQuestionsInSpace) | **Get** /spaces/{spaceId}/search/questions | Get a list of questions that can be asked in a space.
[**GetRevisionById**](DefaultApi.md#GetRevisionById) | **Get** /spaces/{spaceId}/revisions/{revisionId} | Get a specific revision in a space
[**GetRevisionOfChangeRequestById**](DefaultApi.md#GetRevisionOfChangeRequestById) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content | Get the latest content revision for a change request.
[**ImportContent**](DefaultApi.md#ImportContent) | **Post** /spaces/{spaceId}/content/import | Import content in a space.
[**ImportContentInChangeRequest**](DefaultApi.md#ImportContentInChangeRequest) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/content/import | Import content in a change request.
[**ImportContentInChangeRequestPageById**](DefaultApi.md#ImportContentInChangeRequestPageById) | **Post** /spaces/{spaceId}/change-requests/{changeRequestId}/content/page/{pageId}/import | Import external content into a page of a change-request by its ID.
[**ImportContentInPageById**](DefaultApi.md#ImportContentInPageById) | **Post** /spaces/{spaceId}/content/page/{pageId}/import | Import external content into a page by its ID.
[**InstallIntegrationOnSpace**](DefaultApi.md#InstallIntegrationOnSpace) | **Post** /integrations/{integrationName}/installations/{installationId}/spaces | Install integration on a space using an existing installation
[**ListFiles**](DefaultApi.md#ListFiles) | **Get** /spaces/{spaceId}/content/files | List all files for the latest primary revision content for a space
[**ListFilesInChangeRequestById**](DefaultApi.md#ListFilesInChangeRequestById) | **Get** /spaces/{spaceId}/change-requests/{changeRequestId}/content/files | List all files in the latest content of the change-request
[**ListFilesInRevisionById**](DefaultApi.md#ListFilesInRevisionById) | **Get** /spaces/{spaceId}/revisions/{revisionId}/files | List all files in a revision
[**RenderIntegrationUIWithGet**](DefaultApi.md#RenderIntegrationUIWithGet) | **Get** /integrations/{integrationName}/render | Render an integration UI in the context of an installation.
[**RenderIntegrationUIWithPost**](DefaultApi.md#RenderIntegrationUIWithPost) | **Post** /integrations/{integrationName}/render | Render an integration UI in the context of an installation.
[**SearchSpaceContent**](DefaultApi.md#SearchSpaceContent) | **Get** /spaces/{spaceId}/search | Search content in a space
[**UninstallIntegrationFromSpace**](DefaultApi.md#UninstallIntegrationFromSpace) | **Delete** /integrations/{integrationName}/installations/{installationId}/spaces/{spaceId} | Uninstall the integration from a space
[**UpdateIntegrationSpaceInstallation**](DefaultApi.md#UpdateIntegrationSpaceInstallation) | **Patch** /integrations/{integrationName}/installations/{installationId}/spaces/{spaceId} | Update external IDs and configurations of an integration&#39;s installation for a space



## AskQuery

> AskQueryWithGet200Response AskQuery(ctx).SearchAIQuery(searchAIQuery).Execute()

Ask a question to an AI across spaces that is accessible by the currently authenticated target.

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
    searchAIQuery := *openapiclient.NewSearchAIQuery("Query_example") // SearchAIQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AskQuery(context.Background()).SearchAIQuery(searchAIQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AskQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AskQuery`: AskQueryWithGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AskQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAskQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchAIQuery** | [**SearchAIQuery**](SearchAIQuery.md) |  | 

### Return type

[**AskQueryWithGet200Response**](AskQueryWithGet200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AskQueryInSpace

> AskQueryWithGet200Response AskQueryInSpace(ctx, spaceId).SearchAIQuery(searchAIQuery).Execute()

Ask a question to an AI within the context of the space.

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
    searchAIQuery := *openapiclient.NewSearchAIQuery("Query_example") // SearchAIQuery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AskQueryInSpace(context.Background(), spaceId).SearchAIQuery(searchAIQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AskQueryInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AskQueryInSpace`: AskQueryWithGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AskQueryInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiAskQueryInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchAIQuery** | [**SearchAIQuery**](SearchAIQuery.md) |  | 

### Return type

[**AskQueryWithGet200Response**](AskQueryWithGet200Response.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AskQueryInSpaceWithGet

> AskQueryWithGet200Response AskQueryInSpaceWithGet(ctx, spaceId).Query(query).Execute()

Ask a question to an AI within the context of the space.

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
    spaceId := "spaceId_example" // string | The unique id of the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AskQueryInSpaceWithGet(context.Background(), spaceId).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AskQueryInSpaceWithGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AskQueryInSpaceWithGet`: AskQueryWithGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AskQueryInSpaceWithGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiAskQueryInSpaceWithGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 


### Return type

[**AskQueryWithGet200Response**](AskQueryWithGet200Response.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AskQueryWithGet

> AskQueryWithGet200Response AskQueryWithGet(ctx).Query(query).Execute()

Ask a question to an AI across spaces that is accessible by the currently authenticated target.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AskQueryWithGet(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AskQueryWithGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AskQueryWithGet`: AskQueryWithGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AskQueryWithGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAskQueryWithGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 

### Return type

[**AskQueryWithGet200Response**](AskQueryWithGet200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentRevision

> Revision GetCurrentRevision(ctx, spaceId).Execute()

Get the current primary content revision for a space

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
    resp, r, err := apiClient.DefaultApi.GetCurrentRevision(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentRevision`: Revision
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Revision**](Revision.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageById

> RevisionPage GetPageById(ctx, spaceId, pageId).Format(format).Execute()

Get a page by its ID in the primary content.

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
    pageId := "pageId_example" // string | The unique id of the page
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPageById(context.Background(), spaceId, pageId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageById`: RevisionPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**pageId** | **string** | The unique id of the page | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Output format for the content. | 

### Return type

[**RevisionPage**](RevisionPage.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageByPath

> GetPageByPath200Response GetPageByPath(ctx, spaceId, pagePath).Format(format).Execute()

Get a page by its path in the primary content.

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
    pagePath := "pagePath_example" // string | The path of the page in the revision.
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPageByPath(context.Background(), spaceId, pagePath).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPageByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageByPath`: GetPageByPath200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPageByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**pagePath** | **string** | The path of the page in the revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Output format for the content. | 

### Return type

[**GetPageByPath200Response**](GetPageByPath200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageInChangeRequestById

> RevisionPage GetPageInChangeRequestById(ctx, spaceId, changeRequestId, pageId).Format(format).Execute()

Get a page by its ID in a change request.

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
    pageId := "pageId_example" // string | The unique id of the page
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPageInChangeRequestById(context.Background(), spaceId, changeRequestId, pageId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPageInChangeRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageInChangeRequestById`: RevisionPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPageInChangeRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**pageId** | **string** | The unique id of the page | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageInChangeRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **string** | Output format for the content. | 

### Return type

[**RevisionPage**](RevisionPage.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageInChangeRequestByPath

> GetPageByPath200Response GetPageInChangeRequestByPath(ctx, spaceId, changeRequestId, pagePath).Format(format).Execute()

Get a page by its path in a change request.

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
    pagePath := "pagePath_example" // string | The path of the page in the revision.
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPageInChangeRequestByPath(context.Background(), spaceId, changeRequestId, pagePath).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPageInChangeRequestByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageInChangeRequestByPath`: GetPageByPath200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPageInChangeRequestByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**pagePath** | **string** | The path of the page in the revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageInChangeRequestByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **string** | Output format for the content. | 

### Return type

[**GetPageByPath200Response**](GetPageByPath200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageInRevisionById

> RevisionPage GetPageInRevisionById(ctx, spaceId, revisionId, pageId).Format(format).Execute()

Get a page by its ID in a revision.

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
    revisionId := "revisionId_example" // string | The unique id of the revision
    pageId := "pageId_example" // string | The unique id of the page
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPageInRevisionById(context.Background(), spaceId, revisionId, pageId).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPageInRevisionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageInRevisionById`: RevisionPage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPageInRevisionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**revisionId** | **string** | The unique id of the revision | 
**pageId** | **string** | The unique id of the page | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageInRevisionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **string** | Output format for the content. | 

### Return type

[**RevisionPage**](RevisionPage.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPageInRevisionByPath

> GetPageByPath200Response GetPageInRevisionByPath(ctx, spaceId, revisionId, pagePath).Format(format).Execute()

Get a page by its path in a revision.

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
    revisionId := "revisionId_example" // string | The unique id of the revision
    pagePath := "pagePath_example" // string | The path of the page in the revision.
    format := "format_example" // string | Output format for the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetPageInRevisionByPath(context.Background(), spaceId, revisionId, pagePath).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPageInRevisionByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageInRevisionByPath`: GetPageByPath200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPageInRevisionByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**revisionId** | **string** | The unique id of the revision | 
**pagePath** | **string** | The path of the page in the revision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageInRevisionByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **format** | **string** | Output format for the content. | 

### Return type

[**GetPageByPath200Response**](GetPageByPath200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedQuestions

> SearchAIRecommendedQuestions GetRecommendedQuestions(ctx).GetRecommendedQuestionsRequest(getRecommendedQuestionsRequest).Execute()

Get a list of questions recommended by AI for a list of content.

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
    getRecommendedQuestionsRequest := *openapiclient.NewGetRecommendedQuestionsRequest([]string{"Documents_example"}) // GetRecommendedQuestionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRecommendedQuestions(context.Background()).GetRecommendedQuestionsRequest(getRecommendedQuestionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRecommendedQuestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendedQuestions`: SearchAIRecommendedQuestions
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRecommendedQuestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendedQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getRecommendedQuestionsRequest** | [**GetRecommendedQuestionsRequest**](GetRecommendedQuestionsRequest.md) |  | 

### Return type

[**SearchAIRecommendedQuestions**](SearchAIRecommendedQuestions.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedQuestionsInSpace

> SearchAIRecommendedQuestions GetRecommendedQuestionsInSpace(ctx, spaceId).Execute()

Get a list of questions that can be asked in a space.

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
    resp, r, err := apiClient.DefaultApi.GetRecommendedQuestionsInSpace(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRecommendedQuestionsInSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendedQuestionsInSpace`: SearchAIRecommendedQuestions
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRecommendedQuestionsInSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendedQuestionsInSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchAIRecommendedQuestions**](SearchAIRecommendedQuestions.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionById

> Revision GetRevisionById(ctx, spaceId, revisionId).Execute()

Get a specific revision in a space

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
    revisionId := "revisionId_example" // string | The unique id of the revision

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRevisionById(context.Background(), spaceId, revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRevisionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionById`: Revision
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRevisionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**revisionId** | **string** | The unique id of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Revision**](Revision.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionOfChangeRequestById

> Revision GetRevisionOfChangeRequestById(ctx, spaceId, changeRequestId).Execute()

Get the latest content revision for a change request.

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
    resp, r, err := apiClient.DefaultApi.GetRevisionOfChangeRequestById(context.Background(), spaceId, changeRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRevisionOfChangeRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionOfChangeRequestById`: Revision
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRevisionOfChangeRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionOfChangeRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Revision**](Revision.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportContent

> ImportContentResult ImportContent(ctx, spaceId).RequestImportContent(requestImportContent).Execute()

Import content in a space.

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
    requestImportContent := *openapiclient.NewRequestImportContent("Url_example", openapiclient.ImportContentSource("website")) // RequestImportContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ImportContent(context.Background(), spaceId).RequestImportContent(requestImportContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportContent`: ImportContentResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestImportContent** | [**RequestImportContent**](RequestImportContent.md) |  | 

### Return type

[**ImportContentResult**](ImportContentResult.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportContentInChangeRequest

> ImportContentResult ImportContentInChangeRequest(ctx, spaceId, changeRequestId).RequestImportContent(requestImportContent).Execute()

Import content in a change request.

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
    requestImportContent := *openapiclient.NewRequestImportContent("Url_example", openapiclient.ImportContentSource("website")) // RequestImportContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ImportContentInChangeRequest(context.Background(), spaceId, changeRequestId).RequestImportContent(requestImportContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportContentInChangeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportContentInChangeRequest`: ImportContentResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportContentInChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportContentInChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestImportContent** | [**RequestImportContent**](RequestImportContent.md) |  | 

### Return type

[**ImportContentResult**](ImportContentResult.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportContentInChangeRequestPageById

> ImportContentResult ImportContentInChangeRequestPageById(ctx, spaceId, changeRequestId, pageId).RequestImportContent(requestImportContent).Execute()

Import external content into a page of a change-request by its ID.

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
    pageId := "pageId_example" // string | The unique id of the page
    requestImportContent := *openapiclient.NewRequestImportContent("Url_example", openapiclient.ImportContentSource("website")) // RequestImportContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ImportContentInChangeRequestPageById(context.Background(), spaceId, changeRequestId, pageId).RequestImportContent(requestImportContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportContentInChangeRequestPageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportContentInChangeRequestPageById`: ImportContentResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportContentInChangeRequestPageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 
**pageId** | **string** | The unique id of the page | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportContentInChangeRequestPageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestImportContent** | [**RequestImportContent**](RequestImportContent.md) |  | 

### Return type

[**ImportContentResult**](ImportContentResult.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportContentInPageById

> ImportContentResult ImportContentInPageById(ctx, spaceId, pageId).RequestImportContent(requestImportContent).Execute()

Import external content into a page by its ID.

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
    pageId := "pageId_example" // string | The unique id of the page
    requestImportContent := *openapiclient.NewRequestImportContent("Url_example", openapiclient.ImportContentSource("website")) // RequestImportContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ImportContentInPageById(context.Background(), spaceId, pageId).RequestImportContent(requestImportContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportContentInPageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportContentInPageById`: ImportContentResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportContentInPageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**pageId** | **string** | The unique id of the page | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportContentInPageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestImportContent** | [**RequestImportContent**](RequestImportContent.md) |  | 

### Return type

[**ImportContentResult**](ImportContentResult.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallIntegrationOnSpace

> IntegrationSpaceInstallation InstallIntegrationOnSpace(ctx, integrationName, installationId).InstallIntegrationOnSpaceRequest(installIntegrationOnSpaceRequest).Execute()

Install integration on a space using an existing installation

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
    integrationName := "integrationName_example" // string | Name of the integration
    installationId := "installationId_example" // string | Identifier of the installation
    installIntegrationOnSpaceRequest := *openapiclient.NewInstallIntegrationOnSpaceRequest("Space_example") // InstallIntegrationOnSpaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.InstallIntegrationOnSpace(context.Background(), integrationName, installationId).InstallIntegrationOnSpaceRequest(installIntegrationOnSpaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InstallIntegrationOnSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallIntegrationOnSpace`: IntegrationSpaceInstallation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InstallIntegrationOnSpace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**installationId** | **string** | Identifier of the installation | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallIntegrationOnSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installIntegrationOnSpaceRequest** | [**InstallIntegrationOnSpaceRequest**](InstallIntegrationOnSpaceRequest.md) |  | 

### Return type

[**IntegrationSpaceInstallation**](IntegrationSpaceInstallation.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFiles

> ListFiles200Response ListFiles(ctx, spaceId).Page(page).Limit(limit).Execute()

List all files for the latest primary revision content for a space

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
    resp, r, err := apiClient.DefaultApi.ListFiles(context.Background(), spaceId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFiles`: ListFiles200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListFiles200Response**](ListFiles200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilesInChangeRequestById

> ListFiles200Response ListFilesInChangeRequestById(ctx, spaceId, changeRequestId).Page(page).Limit(limit).Execute()

List all files in the latest content of the change-request

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
    resp, r, err := apiClient.DefaultApi.ListFilesInChangeRequestById(context.Background(), spaceId, changeRequestId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFilesInChangeRequestById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFilesInChangeRequestById`: ListFiles200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFilesInChangeRequestById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**changeRequestId** | [**GetChangeRequestByIdChangeRequestIdParameter**](.md) | The unique ID of the change request or its number identifier in the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesInChangeRequestByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListFiles200Response**](ListFiles200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilesInRevisionById

> ListFiles200Response ListFilesInRevisionById(ctx, spaceId, revisionId).Page(page).Limit(limit).Execute()

List all files in a revision

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
    revisionId := "revisionId_example" // string | The unique id of the revision
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListFilesInRevisionById(context.Background(), spaceId, revisionId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFilesInRevisionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFilesInRevisionById`: ListFiles200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFilesInRevisionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 
**revisionId** | **string** | The unique id of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesInRevisionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListFiles200Response**](ListFiles200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderIntegrationUIWithGet

> ContentKitRenderOutput RenderIntegrationUIWithGet(ctx, integrationName).Request(request).Execute()

Render an integration UI in the context of an installation.

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
    integrationName := "integrationName_example" // string | Name of the integration
    request := "request_example" // string | LZ-string compressed JSON request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RenderIntegrationUIWithGet(context.Background(), integrationName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RenderIntegrationUIWithGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenderIntegrationUIWithGet`: ContentKitRenderOutput
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RenderIntegrationUIWithGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenderIntegrationUIWithGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | **string** | LZ-string compressed JSON request | 

### Return type

[**ContentKitRenderOutput**](ContentKitRenderOutput.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderIntegrationUIWithPost

> ContentKitRenderOutput RenderIntegrationUIWithPost(ctx, integrationName).RequestRenderIntegrationUI(requestRenderIntegrationUI).Execute()

Render an integration UI in the context of an installation.

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
    integrationName := "integrationName_example" // string | Name of the integration
    requestRenderIntegrationUI := *openapiclient.NewRequestRenderIntegrationUI("ComponentId_example", map[string]interface{}(123), openapiclient.ContentKitContext{ContentKitContextConfigurationAccount: openapiclient.NewContentKitContextConfigurationAccount("Theme_example", "Type_example", "OrganizationId_example")}) // RequestRenderIntegrationUI | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RenderIntegrationUIWithPost(context.Background(), integrationName).RequestRenderIntegrationUI(requestRenderIntegrationUI).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RenderIntegrationUIWithPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenderIntegrationUIWithPost`: ContentKitRenderOutput
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RenderIntegrationUIWithPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenderIntegrationUIWithPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestRenderIntegrationUI** | [**RequestRenderIntegrationUI**](RequestRenderIntegrationUI.md) |  | 

### Return type

[**ContentKitRenderOutput**](ContentKitRenderOutput.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSpaceContent

> SearchSpaceContent200Response SearchSpaceContent(ctx, spaceId).Query(query).Page(page).Limit(limit).Execute()

Search content in a space

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
    spaceId := "spaceId_example" // string | The unique id of the space
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchSpaceContent(context.Background(), spaceId).Query(query).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchSpaceContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSpaceContent`: SearchSpaceContent200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchSpaceContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSpaceContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**SearchSpaceContent200Response**](SearchSpaceContent200Response.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallIntegrationFromSpace

> UninstallIntegrationFromSpace(ctx, integrationName, installationId, spaceId).Execute()

Uninstall the integration from a space

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
    integrationName := "integrationName_example" // string | Name of the integration
    installationId := "installationId_example" // string | Identifier of the installation
    spaceId := "spaceId_example" // string | The unique id of the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UninstallIntegrationFromSpace(context.Background(), integrationName, installationId, spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UninstallIntegrationFromSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**installationId** | **string** | Identifier of the installation | 
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallIntegrationFromSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIntegrationSpaceInstallation

> IntegrationSpaceInstallation UpdateIntegrationSpaceInstallation(ctx, integrationName, installationId, spaceId).RequestUpdateIntegrationInstallation(requestUpdateIntegrationInstallation).Execute()

Update external IDs and configurations of an integration's installation for a space

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
    integrationName := "integrationName_example" // string | Name of the integration
    installationId := "installationId_example" // string | Identifier of the installation
    spaceId := "spaceId_example" // string | The unique id of the space
    requestUpdateIntegrationInstallation := *openapiclient.NewRequestUpdateIntegrationInstallation() // RequestUpdateIntegrationInstallation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateIntegrationSpaceInstallation(context.Background(), integrationName, installationId, spaceId).RequestUpdateIntegrationInstallation(requestUpdateIntegrationInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIntegrationSpaceInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIntegrationSpaceInstallation`: IntegrationSpaceInstallation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIntegrationSpaceInstallation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**installationId** | **string** | Identifier of the installation | 
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationSpaceInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestUpdateIntegrationInstallation** | [**RequestUpdateIntegrationInstallation**](RequestUpdateIntegrationInstallation.md) |  | 

### Return type

[**IntegrationSpaceInstallation**](IntegrationSpaceInstallation.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

