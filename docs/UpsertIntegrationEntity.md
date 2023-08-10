# UpsertIntegrationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Unique ID of the entity in the context of the integration&#39;s installation | 
**Title** | **string** | Title of the entity | 
**Description** | **string** | Longer text description of the entity | 
**Target** | **string** | URL to open the entity | 
**Metadata** | [**map[string]UpsertIntegrationEntityMetadataValue**](UpsertIntegrationEntityMetadataValue.md) | Metadata stored by the integration on the entity | 

## Methods

### NewUpsertIntegrationEntity

`func NewUpsertIntegrationEntity(entityId string, title string, description string, target string, metadata map[string]UpsertIntegrationEntityMetadataValue, ) *UpsertIntegrationEntity`

NewUpsertIntegrationEntity instantiates a new UpsertIntegrationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertIntegrationEntityWithDefaults

`func NewUpsertIntegrationEntityWithDefaults() *UpsertIntegrationEntity`

NewUpsertIntegrationEntityWithDefaults instantiates a new UpsertIntegrationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *UpsertIntegrationEntity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *UpsertIntegrationEntity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *UpsertIntegrationEntity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetTitle

`func (o *UpsertIntegrationEntity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpsertIntegrationEntity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpsertIntegrationEntity) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *UpsertIntegrationEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsertIntegrationEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsertIntegrationEntity) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTarget

`func (o *UpsertIntegrationEntity) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UpsertIntegrationEntity) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UpsertIntegrationEntity) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetMetadata

`func (o *UpsertIntegrationEntity) GetMetadata() map[string]UpsertIntegrationEntityMetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpsertIntegrationEntity) GetMetadataOk() (*map[string]UpsertIntegrationEntityMetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpsertIntegrationEntity) SetMetadata(v map[string]UpsertIntegrationEntityMetadataValue)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


