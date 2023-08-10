# IntegrationConfigurationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | [**map[string]IntegrationConfigurationSchemaPropertiesValue**](IntegrationConfigurationSchemaPropertiesValue.md) |  | 
**Required** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIntegrationConfigurationSchema

`func NewIntegrationConfigurationSchema(properties map[string]IntegrationConfigurationSchemaPropertiesValue, ) *IntegrationConfigurationSchema`

NewIntegrationConfigurationSchema instantiates a new IntegrationConfigurationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationConfigurationSchemaWithDefaults

`func NewIntegrationConfigurationSchemaWithDefaults() *IntegrationConfigurationSchema`

NewIntegrationConfigurationSchemaWithDefaults instantiates a new IntegrationConfigurationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *IntegrationConfigurationSchema) GetProperties() map[string]IntegrationConfigurationSchemaPropertiesValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IntegrationConfigurationSchema) GetPropertiesOk() (*map[string]IntegrationConfigurationSchemaPropertiesValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IntegrationConfigurationSchema) SetProperties(v map[string]IntegrationConfigurationSchemaPropertiesValue)`

SetProperties sets Properties field to given value.


### GetRequired

`func (o *IntegrationConfigurationSchema) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IntegrationConfigurationSchema) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IntegrationConfigurationSchema) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IntegrationConfigurationSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


