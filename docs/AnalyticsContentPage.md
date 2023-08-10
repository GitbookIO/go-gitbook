# AnalyticsContentPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **string** |  | 
**Title** | **string** |  | 
**PageViews** | **float32** |  | 
**Feedbacks** | Pointer to [**AnalyticsContentPageFeedbacks**](AnalyticsContentPageFeedbacks.md) |  | [optional] 

## Methods

### NewAnalyticsContentPage

`func NewAnalyticsContentPage(page string, title string, pageViews float32, ) *AnalyticsContentPage`

NewAnalyticsContentPage instantiates a new AnalyticsContentPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsContentPageWithDefaults

`func NewAnalyticsContentPageWithDefaults() *AnalyticsContentPage`

NewAnalyticsContentPageWithDefaults instantiates a new AnalyticsContentPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *AnalyticsContentPage) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AnalyticsContentPage) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AnalyticsContentPage) SetPage(v string)`

SetPage sets Page field to given value.


### GetTitle

`func (o *AnalyticsContentPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AnalyticsContentPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AnalyticsContentPage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPageViews

`func (o *AnalyticsContentPage) GetPageViews() float32`

GetPageViews returns the PageViews field if non-nil, zero value otherwise.

### GetPageViewsOk

`func (o *AnalyticsContentPage) GetPageViewsOk() (*float32, bool)`

GetPageViewsOk returns a tuple with the PageViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageViews

`func (o *AnalyticsContentPage) SetPageViews(v float32)`

SetPageViews sets PageViews field to given value.


### GetFeedbacks

`func (o *AnalyticsContentPage) GetFeedbacks() AnalyticsContentPageFeedbacks`

GetFeedbacks returns the Feedbacks field if non-nil, zero value otherwise.

### GetFeedbacksOk

`func (o *AnalyticsContentPage) GetFeedbacksOk() (*AnalyticsContentPageFeedbacks, bool)`

GetFeedbacksOk returns a tuple with the Feedbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbacks

`func (o *AnalyticsContentPage) SetFeedbacks(v AnalyticsContentPageFeedbacks)`

SetFeedbacks sets Feedbacks field to given value.

### HasFeedbacks

`func (o *AnalyticsContentPage) HasFeedbacks() bool`

HasFeedbacks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


