// Copyright 2023 GitBook, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the IntegrationConfigurationComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationConfigurationComponent{}

// IntegrationConfigurationComponent ContentKit component for configuration
type IntegrationConfigurationComponent struct {
	// ID of the ContentKit component defined in the integration
	ComponentId string `json:"componentId"`
}

// NewIntegrationConfigurationComponent instantiates a new IntegrationConfigurationComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfigurationComponent(componentId string) *IntegrationConfigurationComponent {
	this := IntegrationConfigurationComponent{}
	this.ComponentId = componentId
	return &this
}

// NewIntegrationConfigurationComponentWithDefaults instantiates a new IntegrationConfigurationComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigurationComponentWithDefaults() *IntegrationConfigurationComponent {
	this := IntegrationConfigurationComponent{}
	return &this
}

// GetComponentId returns the ComponentId field value
func (o *IntegrationConfigurationComponent) GetComponentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationComponent) GetComponentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentId, true
}

// SetComponentId sets field value
func (o *IntegrationConfigurationComponent) SetComponentId(v string) {
	o.ComponentId = v
}

func (o IntegrationConfigurationComponent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationConfigurationComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["componentId"] = o.ComponentId
	return toSerialize, nil
}

type NullableIntegrationConfigurationComponent struct {
	value *IntegrationConfigurationComponent
	isSet bool
}

func (v NullableIntegrationConfigurationComponent) Get() *IntegrationConfigurationComponent {
	return v.value
}

func (v *NullableIntegrationConfigurationComponent) Set(val *IntegrationConfigurationComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfigurationComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfigurationComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfigurationComponent(val *IntegrationConfigurationComponent) *NullableIntegrationConfigurationComponent {
	return &NullableIntegrationConfigurationComponent{value: val, isSet: true}
}

func (v NullableIntegrationConfigurationComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfigurationComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
