# ListEnvironments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | Pointer to [**[]Environment**](Environment.md) |  | [optional] 

## Methods

### NewListEnvironments200Response

`func NewListEnvironments200Response() *ListEnvironments200Response`

NewListEnvironments200Response instantiates a new ListEnvironments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironments200ResponseWithDefaults

`func NewListEnvironments200ResponseWithDefaults() *ListEnvironments200Response`

NewListEnvironments200ResponseWithDefaults instantiates a new ListEnvironments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ListEnvironments200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListEnvironments200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListEnvironments200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListEnvironments200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *ListEnvironments200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListEnvironments200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListEnvironments200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListEnvironments200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListEnvironments200Response) GetItems() []Environment`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListEnvironments200Response) GetItemsOk() (*[]Environment, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListEnvironments200Response) SetItems(v []Environment)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListEnvironments200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


