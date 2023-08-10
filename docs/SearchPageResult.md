# SearchPageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Path** | **string** |  | 
**Sections** | Pointer to [**[]SearchSectionResult**](SearchSectionResult.md) |  | [optional] 
**Urls** | [**SearchPageResultUrls**](SearchPageResultUrls.md) |  | 

## Methods

### NewSearchPageResult

`func NewSearchPageResult(id string, title string, path string, urls SearchPageResultUrls, ) *SearchPageResult`

NewSearchPageResult instantiates a new SearchPageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPageResultWithDefaults

`func NewSearchPageResultWithDefaults() *SearchPageResult`

NewSearchPageResultWithDefaults instantiates a new SearchPageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchPageResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchPageResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchPageResult) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *SearchPageResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchPageResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchPageResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPath

`func (o *SearchPageResult) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SearchPageResult) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SearchPageResult) SetPath(v string)`

SetPath sets Path field to given value.


### GetSections

`func (o *SearchPageResult) GetSections() []SearchSectionResult`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SearchPageResult) GetSectionsOk() (*[]SearchSectionResult, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SearchPageResult) SetSections(v []SearchSectionResult)`

SetSections sets Sections field to given value.

### HasSections

`func (o *SearchPageResult) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetUrls

`func (o *SearchPageResult) GetUrls() SearchPageResultUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SearchPageResult) GetUrlsOk() (*SearchPageResultUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SearchPageResult) SetUrls(v SearchPageResultUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


