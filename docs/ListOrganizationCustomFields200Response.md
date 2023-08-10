# ListOrganizationCustomFields200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | [**[]CustomField**](CustomField.md) |  | 

## Methods

### NewListOrganizationCustomFields200Response

`func NewListOrganizationCustomFields200Response(items []CustomField, ) *ListOrganizationCustomFields200Response`

NewListOrganizationCustomFields200Response instantiates a new ListOrganizationCustomFields200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationCustomFields200ResponseWithDefaults

`func NewListOrganizationCustomFields200ResponseWithDefaults() *ListOrganizationCustomFields200Response`

NewListOrganizationCustomFields200ResponseWithDefaults instantiates a new ListOrganizationCustomFields200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ListOrganizationCustomFields200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListOrganizationCustomFields200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListOrganizationCustomFields200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListOrganizationCustomFields200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *ListOrganizationCustomFields200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOrganizationCustomFields200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOrganizationCustomFields200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListOrganizationCustomFields200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListOrganizationCustomFields200Response) GetItems() []CustomField`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOrganizationCustomFields200Response) GetItemsOk() (*[]CustomField, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOrganizationCustomFields200Response) SetItems(v []CustomField)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


