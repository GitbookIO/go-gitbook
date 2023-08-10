# AnalyticsSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Searches** | **float32** | Number of searches done by users. | 
**Hits** | **float32** | Number of objects matching this search. | 
**PageHits** | **float32** | Number of pages matching this search. | 
**SectionHits** | **float32** | Number of sections matching this search. | 

## Methods

### NewAnalyticsSearchQuery

`func NewAnalyticsSearchQuery(query string, searches float32, hits float32, pageHits float32, sectionHits float32, ) *AnalyticsSearchQuery`

NewAnalyticsSearchQuery instantiates a new AnalyticsSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsSearchQueryWithDefaults

`func NewAnalyticsSearchQueryWithDefaults() *AnalyticsSearchQuery`

NewAnalyticsSearchQueryWithDefaults instantiates a new AnalyticsSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *AnalyticsSearchQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AnalyticsSearchQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AnalyticsSearchQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetSearches

`func (o *AnalyticsSearchQuery) GetSearches() float32`

GetSearches returns the Searches field if non-nil, zero value otherwise.

### GetSearchesOk

`func (o *AnalyticsSearchQuery) GetSearchesOk() (*float32, bool)`

GetSearchesOk returns a tuple with the Searches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearches

`func (o *AnalyticsSearchQuery) SetSearches(v float32)`

SetSearches sets Searches field to given value.


### GetHits

`func (o *AnalyticsSearchQuery) GetHits() float32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *AnalyticsSearchQuery) GetHitsOk() (*float32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *AnalyticsSearchQuery) SetHits(v float32)`

SetHits sets Hits field to given value.


### GetPageHits

`func (o *AnalyticsSearchQuery) GetPageHits() float32`

GetPageHits returns the PageHits field if non-nil, zero value otherwise.

### GetPageHitsOk

`func (o *AnalyticsSearchQuery) GetPageHitsOk() (*float32, bool)`

GetPageHitsOk returns a tuple with the PageHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageHits

`func (o *AnalyticsSearchQuery) SetPageHits(v float32)`

SetPageHits sets PageHits field to given value.


### GetSectionHits

`func (o *AnalyticsSearchQuery) GetSectionHits() float32`

GetSectionHits returns the SectionHits field if non-nil, zero value otherwise.

### GetSectionHitsOk

`func (o *AnalyticsSearchQuery) GetSectionHitsOk() (*float32, bool)`

GetSectionHitsOk returns a tuple with the SectionHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionHits

`func (o *AnalyticsSearchQuery) SetSectionHits(v float32)`

SetSectionHits sets SectionHits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


