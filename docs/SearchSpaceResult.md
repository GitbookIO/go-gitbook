# SearchSpaceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Pages** | [**[]SearchPageResult**](SearchPageResult.md) |  | 

## Methods

### NewSearchSpaceResult

`func NewSearchSpaceResult(id string, title string, pages []SearchPageResult, ) *SearchSpaceResult`

NewSearchSpaceResult instantiates a new SearchSpaceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSpaceResultWithDefaults

`func NewSearchSpaceResultWithDefaults() *SearchSpaceResult`

NewSearchSpaceResultWithDefaults instantiates a new SearchSpaceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchSpaceResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchSpaceResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchSpaceResult) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *SearchSpaceResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchSpaceResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchSpaceResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPages

`func (o *SearchSpaceResult) GetPages() []SearchPageResult`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *SearchSpaceResult) GetPagesOk() (*[]SearchPageResult, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *SearchSpaceResult) SetPages(v []SearchPageResult)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


