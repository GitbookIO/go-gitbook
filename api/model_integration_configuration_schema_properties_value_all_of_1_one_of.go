/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf{}

// IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf struct for IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf
type IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf struct {
	Type    string  `json:"type"`
	Default *string `json:"default,omitempty"`
}

// NewIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf instantiates a new IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf(type_ string) *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf {
	this := IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf{}
	this.Type = type_
	return &this
}

// NewIntegrationConfigurationSchemaPropertiesValueAllOf1OneOfWithDefaults instantiates a new IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigurationSchemaPropertiesValueAllOf1OneOfWithDefaults() *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf {
	this := IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf{}
	return &this
}

// GetType returns the Type field value
func (o *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) SetType(v string) {
	o.Type = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) GetDefault() string {
	if o == nil || IsNil(o.Default) {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) GetDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) SetDefault(v string) {
	o.Default = &v
}

func (o IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf struct {
	value *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf
	isSet bool
}

func (v NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) Get() *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf {
	return v.value
}

func (v *NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) Set(val *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf(val *IntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) *NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf {
	return &NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf{value: val, isSet: true}
}

func (v NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfigurationSchemaPropertiesValueAllOf1OneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
