# ListSpaceRelations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | [**[]SpaceRelation**](SpaceRelation.md) |  | 

## Methods

### NewListSpaceRelations200Response

`func NewListSpaceRelations200Response(items []SpaceRelation, ) *ListSpaceRelations200Response`

NewListSpaceRelations200Response instantiates a new ListSpaceRelations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpaceRelations200ResponseWithDefaults

`func NewListSpaceRelations200ResponseWithDefaults() *ListSpaceRelations200Response`

NewListSpaceRelations200ResponseWithDefaults instantiates a new ListSpaceRelations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ListSpaceRelations200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListSpaceRelations200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListSpaceRelations200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListSpaceRelations200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *ListSpaceRelations200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListSpaceRelations200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListSpaceRelations200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListSpaceRelations200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListSpaceRelations200Response) GetItems() []SpaceRelation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListSpaceRelations200Response) GetItemsOk() (*[]SpaceRelation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListSpaceRelations200Response) SetItems(v []SpaceRelation)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


