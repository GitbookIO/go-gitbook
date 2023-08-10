# AddSpaceEntitiesRequestEntitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | **string** | ID of the integration owning the entity | 
**EntityId** | **string** | ID of the entity in the integration | 
**Environments** | **[]string** | Environments the entity is linked to | 

## Methods

### NewAddSpaceEntitiesRequestEntitiesInner

`func NewAddSpaceEntitiesRequestEntitiesInner(integration string, entityId string, environments []string, ) *AddSpaceEntitiesRequestEntitiesInner`

NewAddSpaceEntitiesRequestEntitiesInner instantiates a new AddSpaceEntitiesRequestEntitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSpaceEntitiesRequestEntitiesInnerWithDefaults

`func NewAddSpaceEntitiesRequestEntitiesInnerWithDefaults() *AddSpaceEntitiesRequestEntitiesInner`

NewAddSpaceEntitiesRequestEntitiesInnerWithDefaults instantiates a new AddSpaceEntitiesRequestEntitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *AddSpaceEntitiesRequestEntitiesInner) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AddSpaceEntitiesRequestEntitiesInner) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AddSpaceEntitiesRequestEntitiesInner) SetIntegration(v string)`

SetIntegration sets Integration field to given value.


### GetEntityId

`func (o *AddSpaceEntitiesRequestEntitiesInner) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AddSpaceEntitiesRequestEntitiesInner) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AddSpaceEntitiesRequestEntitiesInner) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEnvironments

`func (o *AddSpaceEntitiesRequestEntitiesInner) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *AddSpaceEntitiesRequestEntitiesInner) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *AddSpaceEntitiesRequestEntitiesInner) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


