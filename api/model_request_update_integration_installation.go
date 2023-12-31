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

// checks if the RequestUpdateIntegrationInstallation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestUpdateIntegrationInstallation{}

// RequestUpdateIntegrationInstallation struct for RequestUpdateIntegrationInstallation
type RequestUpdateIntegrationInstallation struct {
	// External IDs assigned by the integration.
	ExternalIds []string `json:"externalIds,omitempty"`
	// Configuration of the integration at the account level
	Configuration  map[string]interface{}                 `json:"configuration,omitempty"`
	SpaceSelection *IntegrationInstallationSpaceSelection `json:"space_selection,omitempty"`
}

// NewRequestUpdateIntegrationInstallation instantiates a new RequestUpdateIntegrationInstallation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestUpdateIntegrationInstallation() *RequestUpdateIntegrationInstallation {
	this := RequestUpdateIntegrationInstallation{}
	return &this
}

// NewRequestUpdateIntegrationInstallationWithDefaults instantiates a new RequestUpdateIntegrationInstallation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestUpdateIntegrationInstallationWithDefaults() *RequestUpdateIntegrationInstallation {
	this := RequestUpdateIntegrationInstallation{}
	return &this
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *RequestUpdateIntegrationInstallation) GetExternalIds() []string {
	if o == nil || IsNil(o.ExternalIds) {
		var ret []string
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUpdateIntegrationInstallation) GetExternalIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *RequestUpdateIntegrationInstallation) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []string and assigns it to the ExternalIds field.
func (o *RequestUpdateIntegrationInstallation) SetExternalIds(v []string) {
	o.ExternalIds = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *RequestUpdateIntegrationInstallation) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUpdateIntegrationInstallation) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *RequestUpdateIntegrationInstallation) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *RequestUpdateIntegrationInstallation) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetSpaceSelection returns the SpaceSelection field value if set, zero value otherwise.
func (o *RequestUpdateIntegrationInstallation) GetSpaceSelection() IntegrationInstallationSpaceSelection {
	if o == nil || IsNil(o.SpaceSelection) {
		var ret IntegrationInstallationSpaceSelection
		return ret
	}
	return *o.SpaceSelection
}

// GetSpaceSelectionOk returns a tuple with the SpaceSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUpdateIntegrationInstallation) GetSpaceSelectionOk() (*IntegrationInstallationSpaceSelection, bool) {
	if o == nil || IsNil(o.SpaceSelection) {
		return nil, false
	}
	return o.SpaceSelection, true
}

// HasSpaceSelection returns a boolean if a field has been set.
func (o *RequestUpdateIntegrationInstallation) HasSpaceSelection() bool {
	if o != nil && !IsNil(o.SpaceSelection) {
		return true
	}

	return false
}

// SetSpaceSelection gets a reference to the given IntegrationInstallationSpaceSelection and assigns it to the SpaceSelection field.
func (o *RequestUpdateIntegrationInstallation) SetSpaceSelection(v IntegrationInstallationSpaceSelection) {
	o.SpaceSelection = &v
}

func (o RequestUpdateIntegrationInstallation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestUpdateIntegrationInstallation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.SpaceSelection) {
		toSerialize["space_selection"] = o.SpaceSelection
	}
	return toSerialize, nil
}

type NullableRequestUpdateIntegrationInstallation struct {
	value *RequestUpdateIntegrationInstallation
	isSet bool
}

func (v NullableRequestUpdateIntegrationInstallation) Get() *RequestUpdateIntegrationInstallation {
	return v.value
}

func (v *NullableRequestUpdateIntegrationInstallation) Set(val *RequestUpdateIntegrationInstallation) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestUpdateIntegrationInstallation) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestUpdateIntegrationInstallation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestUpdateIntegrationInstallation(val *RequestUpdateIntegrationInstallation) *NullableRequestUpdateIntegrationInstallation {
	return &NullableRequestUpdateIntegrationInstallation{value: val, isSet: true}
}

func (v NullableRequestUpdateIntegrationInstallation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestUpdateIntegrationInstallation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
