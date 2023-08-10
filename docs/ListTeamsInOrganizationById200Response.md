# ListTeamsInOrganizationById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | [**[]OrganizationTeam**](OrganizationTeam.md) |  | 

## Methods

### NewListTeamsInOrganizationById200Response

`func NewListTeamsInOrganizationById200Response(items []OrganizationTeam, ) *ListTeamsInOrganizationById200Response`

NewListTeamsInOrganizationById200Response instantiates a new ListTeamsInOrganizationById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTeamsInOrganizationById200ResponseWithDefaults

`func NewListTeamsInOrganizationById200ResponseWithDefaults() *ListTeamsInOrganizationById200Response`

NewListTeamsInOrganizationById200ResponseWithDefaults instantiates a new ListTeamsInOrganizationById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *ListTeamsInOrganizationById200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListTeamsInOrganizationById200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListTeamsInOrganizationById200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListTeamsInOrganizationById200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *ListTeamsInOrganizationById200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListTeamsInOrganizationById200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListTeamsInOrganizationById200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListTeamsInOrganizationById200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *ListTeamsInOrganizationById200Response) GetItems() []OrganizationTeam`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListTeamsInOrganizationById200Response) GetItemsOk() (*[]OrganizationTeam, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListTeamsInOrganizationById200Response) SetItems(v []OrganizationTeam)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


