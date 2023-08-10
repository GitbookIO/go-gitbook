# AnalyticsTrafficPageViews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **float32** | Total number of page views over the period. | 
**Views** | [**[]AnalyticsTrafficPageViewsViewsInner**](AnalyticsTrafficPageViewsViewsInner.md) | Page views per interval (day, week, month). | 

## Methods

### NewAnalyticsTrafficPageViews

`func NewAnalyticsTrafficPageViews(count float32, views []AnalyticsTrafficPageViewsViewsInner, ) *AnalyticsTrafficPageViews`

NewAnalyticsTrafficPageViews instantiates a new AnalyticsTrafficPageViews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsTrafficPageViewsWithDefaults

`func NewAnalyticsTrafficPageViewsWithDefaults() *AnalyticsTrafficPageViews`

NewAnalyticsTrafficPageViewsWithDefaults instantiates a new AnalyticsTrafficPageViews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AnalyticsTrafficPageViews) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnalyticsTrafficPageViews) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnalyticsTrafficPageViews) SetCount(v float32)`

SetCount sets Count field to given value.


### GetViews

`func (o *AnalyticsTrafficPageViews) GetViews() []AnalyticsTrafficPageViewsViewsInner`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *AnalyticsTrafficPageViews) GetViewsOk() (*[]AnalyticsTrafficPageViewsViewsInner, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *AnalyticsTrafficPageViews) SetViews(v []AnalyticsTrafficPageViewsViewsInner)`

SetViews sets Views field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


