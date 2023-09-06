# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** | Unique ID of the entity in the context of the integration&#39;s entity type | 
**Properties** | [**map[string]UpsertEntityPropertiesValue**](UpsertEntityPropertiesValue.md) | Map of values stored as properties on the entity | 
**Id** | **string** | Unique ID for the entity in GitBook | 
**Type** | **string** | Type of an entity | 
**Urls** | [**EntityAllOfUrls**](EntityAllOfUrls.md) |  | 

## Methods

### NewEntity

`func NewEntity(entityId string, properties map[string]UpsertEntityPropertiesValue, id string, type_ string, urls EntityAllOfUrls, ) *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *Entity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Entity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Entity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetProperties

`func (o *Entity) GetProperties() map[string]UpsertEntityPropertiesValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Entity) GetPropertiesOk() (*map[string]UpsertEntityPropertiesValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Entity) SetProperties(v map[string]UpsertEntityPropertiesValue)`

SetProperties sets Properties field to given value.


### GetId

`func (o *Entity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entity) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Entity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entity) SetType(v string)`

SetType sets Type field to given value.


### GetUrls

`func (o *Entity) GetUrls() EntityAllOfUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Entity) GetUrlsOk() (*EntityAllOfUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Entity) SetUrls(v EntityAllOfUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


