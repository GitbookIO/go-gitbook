# EntityPropertySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the property in the object | 
**Title** | **string** | Title displayed to the users | 
**Description** | Pointer to **string** | Description of the property | [optional] 
**Deprecated** | Pointer to **bool** | If true, the property is no longer required and not taken into consideration | [optional] 
**Type** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] 
**Values** | **[]map[string]interface{}** |  | 
**Entity** | **map[string]interface{}** |  | 

## Methods

### NewEntityPropertySchema

`func NewEntityPropertySchema(name string, title string, type_ string, values []map[string]interface{}, entity map[string]interface{}, ) *EntityPropertySchema`

NewEntityPropertySchema instantiates a new EntityPropertySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertySchemaWithDefaults

`func NewEntityPropertySchemaWithDefaults() *EntityPropertySchema`

NewEntityPropertySchemaWithDefaults instantiates a new EntityPropertySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityPropertySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityPropertySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityPropertySchema) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *EntityPropertySchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EntityPropertySchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EntityPropertySchema) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *EntityPropertySchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityPropertySchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityPropertySchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityPropertySchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeprecated

`func (o *EntityPropertySchema) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *EntityPropertySchema) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *EntityPropertySchema) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *EntityPropertySchema) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetType

`func (o *EntityPropertySchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityPropertySchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityPropertySchema) SetType(v string)`

SetType sets Type field to given value.


### GetRole

`func (o *EntityPropertySchema) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EntityPropertySchema) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EntityPropertySchema) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EntityPropertySchema) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetValues

`func (o *EntityPropertySchema) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntityPropertySchema) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntityPropertySchema) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.


### GetEntity

`func (o *EntityPropertySchema) GetEntity() map[string]interface{}`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *EntityPropertySchema) GetEntityOk() (*map[string]interface{}, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *EntityPropertySchema) SetEntity(v map[string]interface{})`

SetEntity sets Entity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


