# SpaceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | [**IntegrationEntity**](IntegrationEntity.md) |  | 
**Environments** | **[]string** | Environments the entity is linked to | 
**Urls** | [**SpaceEntityUrls**](SpaceEntityUrls.md) |  | 

## Methods

### NewSpaceEntity

`func NewSpaceEntity(entity IntegrationEntity, environments []string, urls SpaceEntityUrls, ) *SpaceEntity`

NewSpaceEntity instantiates a new SpaceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceEntityWithDefaults

`func NewSpaceEntityWithDefaults() *SpaceEntity`

NewSpaceEntityWithDefaults instantiates a new SpaceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *SpaceEntity) GetEntity() IntegrationEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SpaceEntity) GetEntityOk() (*IntegrationEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SpaceEntity) SetEntity(v IntegrationEntity)`

SetEntity sets Entity field to given value.


### GetEnvironments

`func (o *SpaceEntity) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *SpaceEntity) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *SpaceEntity) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.


### GetUrls

`func (o *SpaceEntity) GetUrls() SpaceEntityUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SpaceEntity) GetUrlsOk() (*SpaceEntityUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SpaceEntity) SetUrls(v SpaceEntityUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


