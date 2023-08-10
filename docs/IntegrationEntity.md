# IntegrationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Unique ID of the entity in the context of the integration&#39;s installation | 
**Title** | **string** | Title of the entity | 
**Description** | **string** | Longer text description of the entity | 
**Target** | **string** | URL to open the entity | 
**Metadata** | [**map[string]UpsertIntegrationEntityMetadataValue**](UpsertIntegrationEntityMetadataValue.md) | Metadata stored by the integration on the entity | 
**Id** | **string** | Unique ID for the entity in GitBook | 
**Integration** | **string** | ID of the integration. | 

## Methods

### NewIntegrationEntity

`func NewIntegrationEntity(entityId string, title string, description string, target string, metadata map[string]UpsertIntegrationEntityMetadataValue, id string, integration string, ) *IntegrationEntity`

NewIntegrationEntity instantiates a new IntegrationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEntityWithDefaults

`func NewIntegrationEntityWithDefaults() *IntegrationEntity`

NewIntegrationEntityWithDefaults instantiates a new IntegrationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *IntegrationEntity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IntegrationEntity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IntegrationEntity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetTitle

`func (o *IntegrationEntity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IntegrationEntity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IntegrationEntity) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *IntegrationEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationEntity) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTarget

`func (o *IntegrationEntity) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IntegrationEntity) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IntegrationEntity) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetMetadata

`func (o *IntegrationEntity) GetMetadata() map[string]UpsertIntegrationEntityMetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IntegrationEntity) GetMetadataOk() (*map[string]UpsertIntegrationEntityMetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IntegrationEntity) SetMetadata(v map[string]UpsertIntegrationEntityMetadataValue)`

SetMetadata sets Metadata field to given value.


### GetId

`func (o *IntegrationEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationEntity) SetId(v string)`

SetId sets Id field to given value.


### GetIntegration

`func (o *IntegrationEntity) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationEntity) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationEntity) SetIntegration(v string)`

SetIntegration sets Integration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


