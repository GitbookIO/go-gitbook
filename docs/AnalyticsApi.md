# \AnalyticsApi

All URIs are relative to *https://api.gitbook.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContentAnalyticsForSpaceById**](AnalyticsApi.md#GetContentAnalyticsForSpaceById) | **Get** /spaces/{spaceId}/insights/content | Get content analytics for a given space.
[**GetSearchAnalyticsForSpaceById**](AnalyticsApi.md#GetSearchAnalyticsForSpaceById) | **Get** /spaces/{spaceId}/insights/search | Get an overview of the top search queries in a space.
[**GetTrafficAnalyticsForSpaceById**](AnalyticsApi.md#GetTrafficAnalyticsForSpaceById) | **Get** /spaces/{spaceId}/insights/traffic | Get traffic page views for a given space
[**TrackViewInSpaceById**](AnalyticsApi.md#TrackViewInSpaceById) | **Post** /spaces/{spaceId}/insights/track_view | 



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
    resp, r, err := apiClient.AnalyticsApi.GetContentAnalyticsForSpaceById(context.Background(), spaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetContentAnalyticsForSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentAnalyticsForSpaceById`: AnalyticsContentPages
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetContentAnalyticsForSpaceById`: %v\n", resp)
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
    resp, r, err := apiClient.AnalyticsApi.GetSearchAnalyticsForSpaceById(context.Background(), spaceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetSearchAnalyticsForSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchAnalyticsForSpaceById`: AnalyticsTopSearches
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetSearchAnalyticsForSpaceById`: %v\n", resp)
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
    resp, r, err := apiClient.AnalyticsApi.GetTrafficAnalyticsForSpaceById(context.Background(), spaceId).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetTrafficAnalyticsForSpaceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrafficAnalyticsForSpaceById`: AnalyticsTrafficPageViews
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetTrafficAnalyticsForSpaceById`: %v\n", resp)
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
    r, err := apiClient.AnalyticsApi.TrackViewInSpaceById(context.Background(), spaceId).RequestSpaceTrackPageView(requestSpaceTrackPageView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.TrackViewInSpaceById``: %v\n", err)
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

[integration](../README.md#integration), [user-internal](../README.md#user-internal), [user](../README.md#user), [user-staff](../README.md#user-staff)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

