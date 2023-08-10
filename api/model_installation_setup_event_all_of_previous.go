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

// checks if the InstallationSetupEventAllOfPrevious type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallationSetupEventAllOfPrevious{}

// InstallationSetupEventAllOfPrevious The state of the installation at the account level before it was updated.
type InstallationSetupEventAllOfPrevious struct {
	Status IntegrationInstallationStatus `json:"status"`
	// The previous configuration of the installation at the account level.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
}

// NewInstallationSetupEventAllOfPrevious instantiates a new InstallationSetupEventAllOfPrevious object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallationSetupEventAllOfPrevious(status IntegrationInstallationStatus) *InstallationSetupEventAllOfPrevious {
	this := InstallationSetupEventAllOfPrevious{}
	this.Status = status
	return &this
}

// NewInstallationSetupEventAllOfPreviousWithDefaults instantiates a new InstallationSetupEventAllOfPrevious object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallationSetupEventAllOfPreviousWithDefaults() *InstallationSetupEventAllOfPrevious {
	this := InstallationSetupEventAllOfPrevious{}
	return &this
}

// GetStatus returns the Status field value
func (o *InstallationSetupEventAllOfPrevious) GetStatus() IntegrationInstallationStatus {
	if o == nil {
		var ret IntegrationInstallationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InstallationSetupEventAllOfPrevious) GetStatusOk() (*IntegrationInstallationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InstallationSetupEventAllOfPrevious) SetStatus(v IntegrationInstallationStatus) {
	o.Status = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *InstallationSetupEventAllOfPrevious) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationSetupEventAllOfPrevious) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *InstallationSetupEventAllOfPrevious) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *InstallationSetupEventAllOfPrevious) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

func (o InstallationSetupEventAllOfPrevious) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallationSetupEventAllOfPrevious) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableInstallationSetupEventAllOfPrevious struct {
	value *InstallationSetupEventAllOfPrevious
	isSet bool
}

func (v NullableInstallationSetupEventAllOfPrevious) Get() *InstallationSetupEventAllOfPrevious {
	return v.value
}

func (v *NullableInstallationSetupEventAllOfPrevious) Set(val *InstallationSetupEventAllOfPrevious) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationSetupEventAllOfPrevious) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationSetupEventAllOfPrevious) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationSetupEventAllOfPrevious(val *InstallationSetupEventAllOfPrevious) *NullableInstallationSetupEventAllOfPrevious {
	return &NullableInstallationSetupEventAllOfPrevious{value: val, isSet: true}
}

func (v NullableInstallationSetupEventAllOfPrevious) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationSetupEventAllOfPrevious) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
