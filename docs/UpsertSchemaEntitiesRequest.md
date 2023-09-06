# UpsertSchemaEntitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | Pointer to [**UpsertSchemaEntitiesRequestDelete**](UpsertSchemaEntitiesRequestDelete.md) |  | [optional] 
**Entities** | [**[]UpsertEntity**](UpsertEntity.md) | Array of entities to create/update. | 

## Methods

### NewUpsertSchemaEntitiesRequest

`func NewUpsertSchemaEntitiesRequest(entities []UpsertEntity, ) *UpsertSchemaEntitiesRequest`

NewUpsertSchemaEntitiesRequest instantiates a new UpsertSchemaEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertSchemaEntitiesRequestWithDefaults

`func NewUpsertSchemaEntitiesRequestWithDefaults() *UpsertSchemaEntitiesRequest`

NewUpsertSchemaEntitiesRequestWithDefaults instantiates a new UpsertSchemaEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *UpsertSchemaEntitiesRequest) GetDelete() UpsertSchemaEntitiesRequestDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpsertSchemaEntitiesRequest) GetDeleteOk() (*UpsertSchemaEntitiesRequestDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpsertSchemaEntitiesRequest) SetDelete(v UpsertSchemaEntitiesRequestDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpsertSchemaEntitiesRequest) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetEntities

`func (o *UpsertSchemaEntitiesRequest) GetEntities() []UpsertEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *UpsertSchemaEntitiesRequest) GetEntitiesOk() (*[]UpsertEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *UpsertSchemaEntitiesRequest) SetEntities(v []UpsertEntity)`

SetEntities sets Entities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


