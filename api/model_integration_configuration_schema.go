/*
GitBook API

The GitBook API

API version: 0.0.1-beta
Contact: support@gitbook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the IntegrationConfigurationSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationConfigurationSchema{}

// IntegrationConfigurationSchema Schema for a configuration
type IntegrationConfigurationSchema struct {
	Properties map[string]IntegrationConfigurationSchemaPropertiesValue `json:"properties"`
	Required   []string                                                 `json:"required,omitempty"`
}

// NewIntegrationConfigurationSchema instantiates a new IntegrationConfigurationSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfigurationSchema(properties map[string]IntegrationConfigurationSchemaPropertiesValue) *IntegrationConfigurationSchema {
	this := IntegrationConfigurationSchema{}
	this.Properties = properties
	return &this
}

// NewIntegrationConfigurationSchemaWithDefaults instantiates a new IntegrationConfigurationSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigurationSchemaWithDefaults() *IntegrationConfigurationSchema {
	this := IntegrationConfigurationSchema{}
	return &this
}

// GetProperties returns the Properties field value
func (o *IntegrationConfigurationSchema) GetProperties() map[string]IntegrationConfigurationSchemaPropertiesValue {
	if o == nil {
		var ret map[string]IntegrationConfigurationSchemaPropertiesValue
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationSchema) GetPropertiesOk() (*map[string]IntegrationConfigurationSchemaPropertiesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *IntegrationConfigurationSchema) SetProperties(v map[string]IntegrationConfigurationSchemaPropertiesValue) {
	o.Properties = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *IntegrationConfigurationSchema) GetRequired() []string {
	if o == nil || IsNil(o.Required) {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationSchema) GetRequiredOk() ([]string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *IntegrationConfigurationSchema) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *IntegrationConfigurationSchema) SetRequired(v []string) {
	o.Required = v
}

func (o IntegrationConfigurationSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationConfigurationSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableIntegrationConfigurationSchema struct {
	value *IntegrationConfigurationSchema
	isSet bool
}

func (v NullableIntegrationConfigurationSchema) Get() *IntegrationConfigurationSchema {
	return v.value
}

func (v *NullableIntegrationConfigurationSchema) Set(val *IntegrationConfigurationSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfigurationSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfigurationSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfigurationSchema(val *IntegrationConfigurationSchema) *NullableIntegrationConfigurationSchema {
	return &NullableIntegrationConfigurationSchema{value: val, isSet: true}
}

func (v NullableIntegrationConfigurationSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfigurationSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}