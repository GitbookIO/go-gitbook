# IntegrationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | [**map[string]IntegrationConfigurationSchemaPropertiesValue**](IntegrationConfigurationSchemaPropertiesValue.md) |  | 
**Required** | Pointer to **[]string** |  | [optional] 
**ComponentId** | **string** | ID of the ContentKit component defined in the integration | 

## Methods

### NewIntegrationConfiguration

`func NewIntegrationConfiguration(properties map[string]IntegrationConfigurationSchemaPropertiesValue, componentId string, ) *IntegrationConfiguration`

NewIntegrationConfiguration instantiates a new IntegrationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationConfigurationWithDefaults

`func NewIntegrationConfigurationWithDefaults() *IntegrationConfiguration`

NewIntegrationConfigurationWithDefaults instantiates a new IntegrationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *IntegrationConfiguration) GetProperties() map[string]IntegrationConfigurationSchemaPropertiesValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IntegrationConfiguration) GetPropertiesOk() (*map[string]IntegrationConfigurationSchemaPropertiesValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IntegrationConfiguration) SetProperties(v map[string]IntegrationConfigurationSchemaPropertiesValue)`

SetProperties sets Properties field to given value.


### GetRequired

`func (o *IntegrationConfiguration) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IntegrationConfiguration) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IntegrationConfiguration) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IntegrationConfiguration) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetComponentId

`func (o *IntegrationConfiguration) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *IntegrationConfiguration) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *IntegrationConfiguration) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


