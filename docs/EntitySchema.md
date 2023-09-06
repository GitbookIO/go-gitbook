# EntitySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of an entity | 
**Title** | [**EntityRawSchemaTitle**](EntityRawSchemaTitle.md) |  | 
**Properties** | [**[]EntityPropertySchema**](EntityPropertySchema.md) | Ordered list of all properties stored in entities. | 
**Entities** | **float32** | Count of entities created in this schema. | 
**Integration** | Pointer to [**Integration**](Integration.md) |  | [optional] 
**Urls** | [**EntitySchemaAllOfUrls**](EntitySchemaAllOfUrls.md) |  | 

## Methods

### NewEntitySchema

`func NewEntitySchema(type_ string, title EntityRawSchemaTitle, properties []EntityPropertySchema, entities float32, urls EntitySchemaAllOfUrls, ) *EntitySchema`

NewEntitySchema instantiates a new EntitySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaWithDefaults

`func NewEntitySchemaWithDefaults() *EntitySchema`

NewEntitySchemaWithDefaults instantiates a new EntitySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntitySchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitySchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitySchema) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *EntitySchema) GetTitle() EntityRawSchemaTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EntitySchema) GetTitleOk() (*EntityRawSchemaTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EntitySchema) SetTitle(v EntityRawSchemaTitle)`

SetTitle sets Title field to given value.


### GetProperties

`func (o *EntitySchema) GetProperties() []EntityPropertySchema`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EntitySchema) GetPropertiesOk() (*[]EntityPropertySchema, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EntitySchema) SetProperties(v []EntityPropertySchema)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *EntitySchema) GetEntities() float32`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *EntitySchema) GetEntitiesOk() (*float32, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *EntitySchema) SetEntities(v float32)`

SetEntities sets Entities field to given value.


### GetIntegration

`func (o *EntitySchema) GetIntegration() Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *EntitySchema) GetIntegrationOk() (*Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *EntitySchema) SetIntegration(v Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *EntitySchema) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUrls

`func (o *EntitySchema) GetUrls() EntitySchemaAllOfUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *EntitySchema) GetUrlsOk() (*EntitySchemaAllOfUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *EntitySchema) SetUrls(v EntitySchemaAllOfUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


