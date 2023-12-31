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

// checks if the IntegrationSpaceInstallationUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationSpaceInstallationUrls{}

// IntegrationSpaceInstallationUrls URLs associated with the object
type IntegrationSpaceInstallationUrls struct {
	// Public HTTP endpoint for the integration's space installation
	PublicEndpoint string `json:"publicEndpoint"`
}

// NewIntegrationSpaceInstallationUrls instantiates a new IntegrationSpaceInstallationUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationSpaceInstallationUrls(publicEndpoint string) *IntegrationSpaceInstallationUrls {
	this := IntegrationSpaceInstallationUrls{}
	this.PublicEndpoint = publicEndpoint
	return &this
}

// NewIntegrationSpaceInstallationUrlsWithDefaults instantiates a new IntegrationSpaceInstallationUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationSpaceInstallationUrlsWithDefaults() *IntegrationSpaceInstallationUrls {
	this := IntegrationSpaceInstallationUrls{}
	return &this
}

// GetPublicEndpoint returns the PublicEndpoint field value
func (o *IntegrationSpaceInstallationUrls) GetPublicEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicEndpoint
}

// GetPublicEndpointOk returns a tuple with the PublicEndpoint field value
// and a boolean to check if the value has been set.
func (o *IntegrationSpaceInstallationUrls) GetPublicEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicEndpoint, true
}

// SetPublicEndpoint sets field value
func (o *IntegrationSpaceInstallationUrls) SetPublicEndpoint(v string) {
	o.PublicEndpoint = v
}

func (o IntegrationSpaceInstallationUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationSpaceInstallationUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publicEndpoint"] = o.PublicEndpoint
	return toSerialize, nil
}

type NullableIntegrationSpaceInstallationUrls struct {
	value *IntegrationSpaceInstallationUrls
	isSet bool
}

func (v NullableIntegrationSpaceInstallationUrls) Get() *IntegrationSpaceInstallationUrls {
	return v.value
}

func (v *NullableIntegrationSpaceInstallationUrls) Set(val *IntegrationSpaceInstallationUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationSpaceInstallationUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationSpaceInstallationUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationSpaceInstallationUrls(val *IntegrationSpaceInstallationUrls) *NullableIntegrationSpaceInstallationUrls {
	return &NullableIntegrationSpaceInstallationUrls{value: val, isSet: true}
}

func (v NullableIntegrationSpaceInstallationUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationSpaceInstallationUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
