# EntityPropertySchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the property in the object | 
**Title** | **string** | Title displayed to the users | 
**Description** | Pointer to **string** | Description of the property | [optional] 
**Deprecated** | Pointer to **bool** | If true, the property is no longer required and not taken into consideration | [optional] 

## Methods

### NewEntityPropertySchemaAllOf

`func NewEntityPropertySchemaAllOf(name string, title string, ) *EntityPropertySchemaAllOf`

NewEntityPropertySchemaAllOf instantiates a new EntityPropertySchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertySchemaAllOfWithDefaults

`func NewEntityPropertySchemaAllOfWithDefaults() *EntityPropertySchemaAllOf`

NewEntityPropertySchemaAllOfWithDefaults instantiates a new EntityPropertySchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityPropertySchemaAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityPropertySchemaAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityPropertySchemaAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *EntityPropertySchemaAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EntityPropertySchemaAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EntityPropertySchemaAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *EntityPropertySchemaAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityPropertySchemaAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityPropertySchemaAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityPropertySchemaAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeprecated

`func (o *EntityPropertySchemaAllOf) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *EntityPropertySchemaAllOf) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *EntityPropertySchemaAllOf) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *EntityPropertySchemaAllOf) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


