# EntityRawSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of an entity | 
**Title** | [**EntityRawSchemaTitle**](EntityRawSchemaTitle.md) |  | 
**Properties** | [**[]EntityPropertySchema**](EntityPropertySchema.md) | Ordered list of all properties stored in entities. | 

## Methods

### NewEntityRawSchema

`func NewEntityRawSchema(type_ string, title EntityRawSchemaTitle, properties []EntityPropertySchema, ) *EntityRawSchema`

NewEntityRawSchema instantiates a new EntityRawSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRawSchemaWithDefaults

`func NewEntityRawSchemaWithDefaults() *EntityRawSchema`

NewEntityRawSchemaWithDefaults instantiates a new EntityRawSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntityRawSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityRawSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityRawSchema) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *EntityRawSchema) GetTitle() EntityRawSchemaTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EntityRawSchema) GetTitleOk() (*EntityRawSchemaTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EntityRawSchema) SetTitle(v EntityRawSchemaTitle)`

SetTitle sets Title field to given value.


### GetProperties

`func (o *EntityRawSchema) GetProperties() []EntityPropertySchema`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EntityRawSchema) GetPropertiesOk() (*[]EntityPropertySchema, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EntityRawSchema) SetProperties(v []EntityPropertySchema)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


