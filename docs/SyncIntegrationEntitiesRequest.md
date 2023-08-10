# SyncIntegrationEntitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]UpsertIntegrationEntity**](UpsertIntegrationEntity.md) | Array of entities to synchronize. Entities will be created and updated, missing entities will be deleted. | 

## Methods

### NewSyncIntegrationEntitiesRequest

`func NewSyncIntegrationEntitiesRequest(entities []UpsertIntegrationEntity, ) *SyncIntegrationEntitiesRequest`

NewSyncIntegrationEntitiesRequest instantiates a new SyncIntegrationEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncIntegrationEntitiesRequestWithDefaults

`func NewSyncIntegrationEntitiesRequestWithDefaults() *SyncIntegrationEntitiesRequest`

NewSyncIntegrationEntitiesRequestWithDefaults instantiates a new SyncIntegrationEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SyncIntegrationEntitiesRequest) GetEntities() []UpsertIntegrationEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SyncIntegrationEntitiesRequest) GetEntitiesOk() (*[]UpsertIntegrationEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SyncIntegrationEntitiesRequest) SetEntities(v []UpsertIntegrationEntity)`

SetEntities sets Entities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


