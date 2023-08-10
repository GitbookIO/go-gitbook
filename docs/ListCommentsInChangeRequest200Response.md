# ListCommentsInChangeRequest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | [**[]Comment**](Comment.md) |  | 

## Methods

### NewListCommentsInChangeRequest200Response

`func NewListCommentsInChangeRequest200Response(items []Comment, ) *ListCommentsInChangeRequest200Response`

NewListCommentsInChangeRequest200Response instantiates a new ListCommentsInChangeRequest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentsInChangeRequest200ResponseWithDefaults

`func NewListCommentsInChangeRequest200ResponseWithDefaults() *ListCommentsInChangeRequest200Response`

NewListCommentsInChangeRequest200ResponseWithDefaults instantiates a new ListCommentsInChangeRequest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ListCommentsInChangeRequest200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListCommentsInChangeRequest200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListCommentsInChangeRequest200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListCommentsInChangeRequest200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *ListCommentsInChangeRequest200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCommentsInChangeRequest200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCommentsInChangeRequest200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCommentsInChangeRequest200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListCommentsInChangeRequest200Response) GetItems() []Comment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListCommentsInChangeRequest200Response) GetItemsOk() (*[]Comment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListCommentsInChangeRequest200Response) SetItems(v []Comment)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


