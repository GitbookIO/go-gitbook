# \IntegrationsApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIntegrationInstallationToken**](IntegrationsApi.md#CreateIntegrationInstallationToken) | **Post** /integrations/{integrationName}/installations/{installationId}/tokens | Create an integration installation API token
[**GetIntegrationByName**](IntegrationsApi.md#GetIntegrationByName) | **Get** /integrations/{integrationName} | Get a specific integration by its name
[**GetIntegrationEvent**](IntegrationsApi.md#GetIntegrationEvent) | **Get** /integrations/{integrationName}/events/{eventId} | Get a specific integration event by its id
[**GetIntegrationInstallationById**](IntegrationsApi.md#GetIntegrationInstallationById) | **Get** /integrations/{integrationName}/installations/{installationId} | Get a specific integration&#39;s installation by its ID
[**GetIntegrationSpaceInstallation**](IntegrationsApi.md#GetIntegrationSpaceInstallation) | **Get** /integrations/{integrationName}/installations/{installationId}/spaces/{spaceId} | Get a specific integration&#39;s space installation
[**InstallIntegration**](IntegrationsApi.md#InstallIntegration) | **Post** /integrations/{integrationName}/installations | Install integration on a target organization
[**ListIntegrationEvents**](IntegrationsApi.md#ListIntegrationEvents) | **Get** /integrations/{integrationName}/events | List all integration events
[**ListIntegrationInstallations**](IntegrationsApi.md#ListIntegrationInstallations) | **Get** /integrations/{integrationName}/installations | Fetch a list of installations of an integration
[**ListIntegrationSpaceInstallations**](IntegrationsApi.md#ListIntegrationSpaceInstallations) | **Get** /integrations/{integrationName}/spaces | Fetch a list of space installations of an integration
[**ListIntegrations**](IntegrationsApi.md#ListIntegrations) | **Get** /integrations | List all public integrations
[**ListSpaceIntegrationsBlocks**](IntegrationsApi.md#ListSpaceIntegrationsBlocks) | **Get** /spaces/{spaceId}/integration-blocks | List integrations blocks for a space
[**PublishIntegration**](IntegrationsApi.md#PublishIntegration) | **Post** /integrations/{integrationName} | Publish an integration
[**RemoveIntegrationDevSpace**](IntegrationsApi.md#RemoveIntegrationDevSpace) | **Delete** /integrations/{integrationName}/spaces/{spaceId}/dev | Remove the development space for an integration
[**UninstallIntegration**](IntegrationsApi.md#UninstallIntegration) | **Delete** /integrations/{integrationName}/installations/{installationId} | Uninstall the integration from a target organization
[**UnpublishIntegration**](IntegrationsApi.md#UnpublishIntegration) | **Delete** /integrations/{integrationName} | Unpublish an integration
[**UpdateIntegrationDevSpace**](IntegrationsApi.md#UpdateIntegrationDevSpace) | **Put** /integrations/{integrationName}/spaces/{spaceId}/dev | Update the development space for an integration
[**UpdateIntegrationInstallation**](IntegrationsApi.md#UpdateIntegrationInstallation) | **Patch** /integrations/{integrationName}/installations/{installationId} | Update external IDs and configurations of an integration&#39;s installation



## CreateIntegrationInstallationToken

> APITemporaryToken CreateIntegrationInstallationToken(ctx, integrationName, installationId).Execute()

Create an integration installation API token



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.CreateIntegrationInstallationToken(context.Background(), integrationName, installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.CreateIntegrationInstallationToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIntegrationInstallationToken`: APITemporaryToken
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.CreateIntegrationInstallationToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**installationId** | **string** | Identifier of the installation | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIntegrationInstallationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**APITemporaryToken**](APITemporaryToken.md)

### Authorization

[integration](../README.md#integration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationByName

> Integration GetIntegrationByName(ctx, integrationName).Execute()

Get a specific integration by its name

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GetIntegrationByName(context.Background(), integrationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationByName`: Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationEvent

> GetIntegrationEvent200Response GetIntegrationEvent(ctx, integrationName, eventId).Execute()

Get a specific integration event by its id

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
    eventId := "eventId_example" // string | ID of the integration event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GetIntegrationEvent(context.Background(), integrationName, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationEvent`: GetIntegrationEvent200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**eventId** | **string** | ID of the integration event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetIntegrationEvent200Response**](GetIntegrationEvent200Response.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationInstallationById

> IntegrationInstallation GetIntegrationInstallationById(ctx, integrationName, installationId).Execute()

Get a specific integration's installation by its ID

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.GetIntegrationInstallationById(context.Background(), integrationName, installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationInstallationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationInstallationById`: IntegrationInstallation
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationInstallationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**installationId** | **string** | Identifier of the installation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationInstallationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IntegrationInstallation**](IntegrationInstallation.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationSpaceInstallation

> IntegrationSpaceInstallation GetIntegrationSpaceInstallation(ctx, integrationName, installationId, spaceId).Execute()

Get a specific integration's space installation

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
    resp, r, err := apiClient.IntegrationsApi.GetIntegrationSpaceInstallation(context.Background(), integrationName, installationId, spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationSpaceInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationSpaceInstallation`: IntegrationSpaceInstallation
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationSpaceInstallation`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetIntegrationSpaceInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IntegrationSpaceInstallation**](IntegrationSpaceInstallation.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallIntegration

> IntegrationInstallation InstallIntegration(ctx, integrationName).OrganizationTarget(organizationTarget).Execute()

Install integration on a target organization

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
    organizationTarget := *openapiclient.NewOrganizationTarget("Organization_example") // OrganizationTarget | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.InstallIntegration(context.Background(), integrationName).OrganizationTarget(organizationTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.InstallIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallIntegration`: IntegrationInstallation
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.InstallIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationTarget** | [**OrganizationTarget**](OrganizationTarget.md) |  | 

### Return type

[**IntegrationInstallation**](IntegrationInstallation.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationEvents

> ListIntegrationEvents200Response ListIntegrationEvents(ctx, integrationName).Page(page).Limit(limit).Execute()

List all integration events

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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.ListIntegrationEvents(context.Background(), integrationName).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationEvents`: ListIntegrationEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListIntegrationEvents200Response**](ListIntegrationEvents200Response.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationInstallations

> ListIntegrationInstallations200Response ListIntegrationInstallations(ctx, integrationName).Page(page).Limit(limit).ExternalId(externalId).Execute()

Fetch a list of installations of an integration

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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    externalId := "externalId_example" // string | External Id to filter by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.ListIntegrationInstallations(context.Background(), integrationName).Page(page).Limit(limit).ExternalId(externalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrationInstallations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationInstallations`: ListIntegrationInstallations200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrationInstallations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationInstallationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **externalId** | **string** | External Id to filter by | 

### Return type

[**ListIntegrationInstallations200Response**](ListIntegrationInstallations200Response.md)

### Authorization

[integration](../README.md#integration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationSpaceInstallations

> ListIntegrationSpaceInstallations200Response ListIntegrationSpaceInstallations(ctx, integrationName).Page(page).Limit(limit).ExternalId(externalId).Execute()

Fetch a list of space installations of an integration

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
    page := "page_example" // string | Identifier of the page results to fetch. (optional)
    limit := float32(8.14) // float32 | The number of results per page (optional)
    externalId := "externalId_example" // string | External Id to filter by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.ListIntegrationSpaceInstallations(context.Background(), integrationName).Page(page).Limit(limit).ExternalId(externalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrationSpaceInstallations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationSpaceInstallations`: ListIntegrationSpaceInstallations200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrationSpaceInstallations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationSpaceInstallationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 
 **externalId** | **string** | External Id to filter by | 

### Return type

[**ListIntegrationSpaceInstallations200Response**](ListIntegrationSpaceInstallations200Response.md)

### Authorization

[integration](../README.md#integration)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrations

> ListIntegrations200Response ListIntegrations(ctx).Page(page).Limit(limit).Execute()

List all public integrations

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
    resp, r, err := apiClient.IntegrationsApi.ListIntegrations(context.Background()).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrations`: ListIntegrations200Response
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Identifier of the page results to fetch. | 
 **limit** | **float32** | The number of results per page | 

### Return type

[**ListIntegrations200Response**](ListIntegrations200Response.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpaceIntegrationsBlocks

> []SpaceIntegrationBlocksInner ListSpaceIntegrationsBlocks(ctx, spaceId).Execute()

List integrations blocks for a space

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
    resp, r, err := apiClient.IntegrationsApi.ListSpaceIntegrationsBlocks(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.ListSpaceIntegrationsBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpaceIntegrationsBlocks`: []SpaceIntegrationBlocksInner
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.ListSpaceIntegrationsBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSpaceIntegrationsBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SpaceIntegrationBlocksInner**](SpaceIntegrationBlocksInner.md)

### Authorization

[integration](../README.md#integration), [integration-installation](../README.md#integration-installation), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishIntegration

> Integration PublishIntegration(ctx, integrationName).RequestPublishIntegration(requestPublishIntegration).Execute()

Publish an integration

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
    requestPublishIntegration := *openapiclient.NewRequestPublishIntegration("Title_example", "Description_example", []openapiclient.IntegrationScope{openapiclient.IntegrationScope("entities:write")}, "Script_example") // RequestPublishIntegration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.PublishIntegration(context.Background(), integrationName).RequestPublishIntegration(requestPublishIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.PublishIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishIntegration`: Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.PublishIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestPublishIntegration** | [**RequestPublishIntegration**](RequestPublishIntegration.md) |  | 

### Return type

[**Integration**](Integration.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveIntegrationDevSpace

> RemoveIntegrationDevSpace(ctx, integrationName, spaceId).Execute()

Remove the development space for an integration

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
    spaceId := "spaceId_example" // string | The unique id of the space

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationsApi.RemoveIntegrationDevSpace(context.Background(), integrationName, spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.RemoveIntegrationDevSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIntegrationDevSpaceRequest struct via the builder pattern


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


## UninstallIntegration

> UninstallIntegration(ctx, integrationName, installationId).Execute()

Uninstall the integration from a target organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationsApi.UninstallIntegration(context.Background(), integrationName, installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.UninstallIntegration``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallIntegrationRequest struct via the builder pattern


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


## UnpublishIntegration

> UnpublishIntegration(ctx, integrationName).Execute()

Unpublish an integration

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationsApi.UnpublishIntegration(context.Background(), integrationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.UnpublishIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishIntegrationRequest struct via the builder pattern


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


## UpdateIntegrationDevSpace

> UpdateIntegrationDevSpace(ctx, integrationName, spaceId).UpdateIntegrationDevSpaceRequest(updateIntegrationDevSpaceRequest).Execute()

Update the development space for an integration

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
    spaceId := "spaceId_example" // string | The unique id of the space
    updateIntegrationDevSpaceRequest := *openapiclient.NewUpdateIntegrationDevSpaceRequest("TunnelUrl_example") // UpdateIntegrationDevSpaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationsApi.UpdateIntegrationDevSpace(context.Background(), integrationName, spaceId).UpdateIntegrationDevSpaceRequest(updateIntegrationDevSpaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.UpdateIntegrationDevSpace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**spaceId** | **string** | The unique id of the space | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationDevSpaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIntegrationDevSpaceRequest** | [**UpdateIntegrationDevSpaceRequest**](UpdateIntegrationDevSpaceRequest.md) |  | 

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


## UpdateIntegrationInstallation

> IntegrationInstallation UpdateIntegrationInstallation(ctx, integrationName, installationId).RequestUpdateIntegrationInstallation(requestUpdateIntegrationInstallation).Execute()

Update external IDs and configurations of an integration's installation

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
    requestUpdateIntegrationInstallation := *openapiclient.NewRequestUpdateIntegrationInstallation() // RequestUpdateIntegrationInstallation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsApi.UpdateIntegrationInstallation(context.Background(), integrationName, installationId).RequestUpdateIntegrationInstallation(requestUpdateIntegrationInstallation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.UpdateIntegrationInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIntegrationInstallation`: IntegrationInstallation
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.UpdateIntegrationInstallation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationName** | **string** | Name of the integration | 
**installationId** | **string** | Identifier of the installation | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestUpdateIntegrationInstallation** | [**RequestUpdateIntegrationInstallation**](RequestUpdateIntegrationInstallation.md) |  | 

### Return type

[**IntegrationInstallation**](IntegrationInstallation.md)

### Authorization

[user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

