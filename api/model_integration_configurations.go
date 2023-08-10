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

// checks if the IntegrationConfigurations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationConfigurations{}

// IntegrationConfigurations struct for IntegrationConfigurations
type IntegrationConfigurations struct {
	Account *IntegrationConfiguration `json:"account,omitempty"`
	Space   *IntegrationConfiguration `json:"space,omitempty"`
}

// NewIntegrationConfigurations instantiates a new IntegrationConfigurations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfigurations() *IntegrationConfigurations {
	this := IntegrationConfigurations{}
	return &this
}

// NewIntegrationConfigurationsWithDefaults instantiates a new IntegrationConfigurations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigurationsWithDefaults() *IntegrationConfigurations {
	this := IntegrationConfigurations{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IntegrationConfigurations) GetAccount() IntegrationConfiguration {
	if o == nil || IsNil(o.Account) {
		var ret IntegrationConfiguration
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurations) GetAccountOk() (*IntegrationConfiguration, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IntegrationConfigurations) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IntegrationConfiguration and assigns it to the Account field.
func (o *IntegrationConfigurations) SetAccount(v IntegrationConfiguration) {
	o.Account = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *IntegrationConfigurations) GetSpace() IntegrationConfiguration {
	if o == nil || IsNil(o.Space) {
		var ret IntegrationConfiguration
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurations) GetSpaceOk() (*IntegrationConfiguration, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *IntegrationConfigurations) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given IntegrationConfiguration and assigns it to the Space field.
func (o *IntegrationConfigurations) SetSpace(v IntegrationConfiguration) {
	o.Space = &v
}

func (o IntegrationConfigurations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationConfigurations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	return toSerialize, nil
}

type NullableIntegrationConfigurations struct {
	value *IntegrationConfigurations
	isSet bool
}

func (v NullableIntegrationConfigurations) Get() *IntegrationConfigurations {
	return v.value
}

func (v *NullableIntegrationConfigurations) Set(val *IntegrationConfigurations) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfigurations) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfigurations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfigurations(val *IntegrationConfigurations) *NullableIntegrationConfigurations {
	return &NullableIntegrationConfigurations{value: val, isSet: true}
}

func (v NullableIntegrationConfigurations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfigurations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
