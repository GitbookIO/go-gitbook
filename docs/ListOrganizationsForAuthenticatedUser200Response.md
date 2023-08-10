# ListOrganizationsForAuthenticatedUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | [**[]Organization**](Organization.md) |  | 

## Methods

### NewListOrganizationsForAuthenticatedUser200Response

`func NewListOrganizationsForAuthenticatedUser200Response(items []Organization, ) *ListOrganizationsForAuthenticatedUser200Response`

NewListOrganizationsForAuthenticatedUser200Response instantiates a new ListOrganizationsForAuthenticatedUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrganizationsForAuthenticatedUser200ResponseWithDefaults

`func NewListOrganizationsForAuthenticatedUser200ResponseWithDefaults() *ListOrganizationsForAuthenticatedUser200Response`

NewListOrganizationsForAuthenticatedUser200ResponseWithDefaults instantiates a new ListOrganizationsForAuthenticatedUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ListOrganizationsForAuthenticatedUser200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListOrganizationsForAuthenticatedUser200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListOrganizationsForAuthenticatedUser200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListOrganizationsForAuthenticatedUser200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *ListOrganizationsForAuthenticatedUser200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListOrganizationsForAuthenticatedUser200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListOrganizationsForAuthenticatedUser200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListOrganizationsForAuthenticatedUser200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListOrganizationsForAuthenticatedUser200Response) GetItems() []Organization`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListOrganizationsForAuthenticatedUser200Response) GetItemsOk() (*[]Organization, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListOrganizationsForAuthenticatedUser200Response) SetItems(v []Organization)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


