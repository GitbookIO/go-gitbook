# UpsertEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Unique ID of the entity in the context of the integration&#39;s entity type | 
**Properties** | [**map[string]UpsertEntityPropertiesValue**](UpsertEntityPropertiesValue.md) | Map of values stored as properties on the entity | 

## Methods

### NewUpsertEntity

`func NewUpsertEntity(entityId string, properties map[string]UpsertEntityPropertiesValue, ) *UpsertEntity`

NewUpsertEntity instantiates a new UpsertEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertEntityWithDefaults

`func NewUpsertEntityWithDefaults() *UpsertEntity`

NewUpsertEntityWithDefaults instantiates a new UpsertEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *UpsertEntity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *UpsertEntity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *UpsertEntity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetProperties

`func (o *UpsertEntity) GetProperties() map[string]UpsertEntityPropertiesValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UpsertEntity) GetPropertiesOk() (*map[string]UpsertEntityPropertiesValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UpsertEntity) SetProperties(v map[string]UpsertEntityPropertiesValue)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


