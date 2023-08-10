# GetIntegrationEntities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**ListNext**](ListNext.md) |  | [optional] 
**Count** | Pointer to **float32** | Total count of objects in the list | [optional] 
**Items** | [**[]IntegrationEntity**](IntegrationEntity.md) |  | 

## Methods

### NewGetIntegrationEntities200Response

`func NewGetIntegrationEntities200Response(items []IntegrationEntity, ) *GetIntegrationEntities200Response`

NewGetIntegrationEntities200Response instantiates a new GetIntegrationEntities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrationEntities200ResponseWithDefaults

`func NewGetIntegrationEntities200ResponseWithDefaults() *GetIntegrationEntities200Response`

NewGetIntegrationEntities200ResponseWithDefaults instantiates a new GetIntegrationEntities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *GetIntegrationEntities200Response) GetNext() ListNext`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetIntegrationEntities200Response) GetNextOk() (*ListNext, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetIntegrationEntities200Response) SetNext(v ListNext)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetIntegrationEntities200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCount

`func (o *GetIntegrationEntities200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetIntegrationEntities200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetIntegrationEntities200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetIntegrationEntities200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *GetIntegrationEntities200Response) GetItems() []IntegrationEntity`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetIntegrationEntities200Response) GetItemsOk() (*[]IntegrationEntity, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetIntegrationEntities200Response) SetItems(v []IntegrationEntity)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


