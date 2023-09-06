# EntitySchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | **float32** | Count of entities created in this schema. | 
**Integration** | Pointer to [**Integration**](Integration.md) |  | [optional] 
**Urls** | [**EntitySchemaAllOfUrls**](EntitySchemaAllOfUrls.md) |  | 

## Methods

### NewEntitySchemaAllOf

`func NewEntitySchemaAllOf(entities float32, urls EntitySchemaAllOfUrls, ) *EntitySchemaAllOf`

NewEntitySchemaAllOf instantiates a new EntitySchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaAllOfWithDefaults

`func NewEntitySchemaAllOfWithDefaults() *EntitySchemaAllOf`

NewEntitySchemaAllOfWithDefaults instantiates a new EntitySchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *EntitySchemaAllOf) GetEntities() float32`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *EntitySchemaAllOf) GetEntitiesOk() (*float32, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *EntitySchemaAllOf) SetEntities(v float32)`

SetEntities sets Entities field to given value.


### GetIntegration

`func (o *EntitySchemaAllOf) GetIntegration() Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *EntitySchemaAllOf) GetIntegrationOk() (*Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *EntitySchemaAllOf) SetIntegration(v Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *EntitySchemaAllOf) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUrls

`func (o *EntitySchemaAllOf) GetUrls() EntitySchemaAllOfUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *EntitySchemaAllOf) GetUrlsOk() (*EntitySchemaAllOfUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *EntitySchemaAllOf) SetUrls(v EntitySchemaAllOfUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


