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

// checks if the IntegrationConfigurationPropertiesValueAllOf1OneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationConfigurationPropertiesValueAllOf1OneOf1{}

// IntegrationConfigurationPropertiesValueAllOf1OneOf1 struct for IntegrationConfigurationPropertiesValueAllOf1OneOf1
type IntegrationConfigurationPropertiesValueAllOf1OneOf1 struct {
	Type    string   `json:"type"`
	Default *float32 `json:"default,omitempty"`
}

// NewIntegrationConfigurationPropertiesValueAllOf1OneOf1 instantiates a new IntegrationConfigurationPropertiesValueAllOf1OneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfigurationPropertiesValueAllOf1OneOf1(type_ string) *IntegrationConfigurationPropertiesValueAllOf1OneOf1 {
	this := IntegrationConfigurationPropertiesValueAllOf1OneOf1{}
	this.Type = type_
	return &this
}

// NewIntegrationConfigurationPropertiesValueAllOf1OneOf1WithDefaults instantiates a new IntegrationConfigurationPropertiesValueAllOf1OneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigurationPropertiesValueAllOf1OneOf1WithDefaults() *IntegrationConfigurationPropertiesValueAllOf1OneOf1 {
	this := IntegrationConfigurationPropertiesValueAllOf1OneOf1{}
	return &this
}

// GetType returns the Type field value
func (o *IntegrationConfigurationPropertiesValueAllOf1OneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValueAllOf1OneOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationConfigurationPropertiesValueAllOf1OneOf1) SetType(v string) {
	o.Type = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *IntegrationConfigurationPropertiesValueAllOf1OneOf1) GetDefault() float32 {
	if o == nil || IsNil(o.Default) {
		var ret float32
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValueAllOf1OneOf1) GetDefaultOk() (*float32, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *IntegrationConfigurationPropertiesValueAllOf1OneOf1) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given float32 and assigns it to the Default field.
func (o *IntegrationConfigurationPropertiesValueAllOf1OneOf1) SetDefault(v float32) {
	o.Default = &v
}

func (o IntegrationConfigurationPropertiesValueAllOf1OneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationConfigurationPropertiesValueAllOf1OneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1 struct {
	value *IntegrationConfigurationPropertiesValueAllOf1OneOf1
	isSet bool
}

func (v NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1) Get() *IntegrationConfigurationPropertiesValueAllOf1OneOf1 {
	return v.value
}

func (v *NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1) Set(val *IntegrationConfigurationPropertiesValueAllOf1OneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfigurationPropertiesValueAllOf1OneOf1(val *IntegrationConfigurationPropertiesValueAllOf1OneOf1) *NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1 {
	return &NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1{value: val, isSet: true}
}

func (v NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfigurationPropertiesValueAllOf1OneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
