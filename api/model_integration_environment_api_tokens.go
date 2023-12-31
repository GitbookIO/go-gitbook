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

// checks if the IntegrationEnvironmentApiTokens type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationEnvironmentApiTokens{}

// IntegrationEnvironmentApiTokens struct for IntegrationEnvironmentApiTokens
type IntegrationEnvironmentApiTokens struct {
	// API authentication token representing the integration.
	Integration string `json:"integration"`
	// API authentication token representing the current installation.
	Installation *string `json:"installation,omitempty"`
}

// NewIntegrationEnvironmentApiTokens instantiates a new IntegrationEnvironmentApiTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationEnvironmentApiTokens(integration string) *IntegrationEnvironmentApiTokens {
	this := IntegrationEnvironmentApiTokens{}
	this.Integration = integration
	return &this
}

// NewIntegrationEnvironmentApiTokensWithDefaults instantiates a new IntegrationEnvironmentApiTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationEnvironmentApiTokensWithDefaults() *IntegrationEnvironmentApiTokens {
	this := IntegrationEnvironmentApiTokens{}
	return &this
}

// GetIntegration returns the Integration field value
func (o *IntegrationEnvironmentApiTokens) GetIntegration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironmentApiTokens) GetIntegrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value
func (o *IntegrationEnvironmentApiTokens) SetIntegration(v string) {
	o.Integration = v
}

// GetInstallation returns the Installation field value if set, zero value otherwise.
func (o *IntegrationEnvironmentApiTokens) GetInstallation() string {
	if o == nil || IsNil(o.Installation) {
		var ret string
		return ret
	}
	return *o.Installation
}

// GetInstallationOk returns a tuple with the Installation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironmentApiTokens) GetInstallationOk() (*string, bool) {
	if o == nil || IsNil(o.Installation) {
		return nil, false
	}
	return o.Installation, true
}

// HasInstallation returns a boolean if a field has been set.
func (o *IntegrationEnvironmentApiTokens) HasInstallation() bool {
	if o != nil && !IsNil(o.Installation) {
		return true
	}

	return false
}

// SetInstallation gets a reference to the given string and assigns it to the Installation field.
func (o *IntegrationEnvironmentApiTokens) SetInstallation(v string) {
	o.Installation = &v
}

func (o IntegrationEnvironmentApiTokens) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationEnvironmentApiTokens) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["integration"] = o.Integration
	if !IsNil(o.Installation) {
		toSerialize["installation"] = o.Installation
	}
	return toSerialize, nil
}

type NullableIntegrationEnvironmentApiTokens struct {
	value *IntegrationEnvironmentApiTokens
	isSet bool
}

func (v NullableIntegrationEnvironmentApiTokens) Get() *IntegrationEnvironmentApiTokens {
	return v.value
}

func (v *NullableIntegrationEnvironmentApiTokens) Set(val *IntegrationEnvironmentApiTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationEnvironmentApiTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationEnvironmentApiTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationEnvironmentApiTokens(val *IntegrationEnvironmentApiTokens) *NullableIntegrationEnvironmentApiTokens {
	return &NullableIntegrationEnvironmentApiTokens{value: val, isSet: true}
}

func (v NullableIntegrationEnvironmentApiTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationEnvironmentApiTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
