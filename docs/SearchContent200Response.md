# SearchContent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | [**[]SearchSpaceResult**](SearchSpaceResult.md) |  | 

## Methods

### NewSearchContent200Response

`func NewSearchContent200Response(items []SearchSpaceResult, ) *SearchContent200Response`

NewSearchContent200Response instantiates a new SearchContent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchContent200ResponseWithDefaults

`func NewSearchContent200ResponseWithDefaults() *SearchContent200Response`

NewSearchContent200ResponseWithDefaults instantiates a new SearchContent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *SearchContent200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SearchContent200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SearchContent200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *SearchContent200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *SearchContent200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchContent200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchContent200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SearchContent200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *SearchContent200Response) GetItems() []SearchSpaceResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchContent200Response) GetItemsOk() (*[]SearchSpaceResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchContent200Response) SetItems(v []SearchSpaceResult)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


