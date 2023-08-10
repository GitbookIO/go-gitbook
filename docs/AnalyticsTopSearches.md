# AnalyticsTopSearches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Searches** | **float32** | Number of searches done by users. | 
**Queries** | [**[]AnalyticsSearchQuery**](AnalyticsSearchQuery.md) | Top queries searched for this content. | 

## Methods

### NewAnalyticsTopSearches

`func NewAnalyticsTopSearches(searches float32, queries []AnalyticsSearchQuery, ) *AnalyticsTopSearches`

NewAnalyticsTopSearches instantiates a new AnalyticsTopSearches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsTopSearchesWithDefaults

`func NewAnalyticsTopSearchesWithDefaults() *AnalyticsTopSearches`

NewAnalyticsTopSearchesWithDefaults instantiates a new AnalyticsTopSearches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearches

`func (o *AnalyticsTopSearches) GetSearches() float32`

GetSearches returns the Searches field if non-nil, zero value otherwise.

### GetSearchesOk

`func (o *AnalyticsTopSearches) GetSearchesOk() (*float32, bool)`

GetSearchesOk returns a tuple with the Searches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearches

`func (o *AnalyticsTopSearches) SetSearches(v float32)`

SetSearches sets Searches field to given value.


### GetQueries

`func (o *AnalyticsTopSearches) GetQueries() []AnalyticsSearchQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *AnalyticsTopSearches) GetQueriesOk() (*[]AnalyticsSearchQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *AnalyticsTopSearches) SetQueries(v []AnalyticsSearchQuery)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


