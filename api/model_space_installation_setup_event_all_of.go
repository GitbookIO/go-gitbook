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

// checks if the SpaceInstallationSetupEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceInstallationSetupEventAllOf{}

// SpaceInstallationSetupEventAllOf Event received when integration has been installed or updated on a space.
type SpaceInstallationSetupEventAllOf struct {
	Type     string                                    `json:"type"`
	Status   IntegrationInstallationStatus             `json:"status"`
	Previous *SpaceInstallationSetupEventAllOfPrevious `json:"previous,omitempty"`
}

// NewSpaceInstallationSetupEventAllOf instantiates a new SpaceInstallationSetupEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceInstallationSetupEventAllOf(type_ string, status IntegrationInstallationStatus) *SpaceInstallationSetupEventAllOf {
	this := SpaceInstallationSetupEventAllOf{}
	this.Type = type_
	this.Status = status
	return &this
}

// NewSpaceInstallationSetupEventAllOfWithDefaults instantiates a new SpaceInstallationSetupEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceInstallationSetupEventAllOfWithDefaults() *SpaceInstallationSetupEventAllOf {
	this := SpaceInstallationSetupEventAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *SpaceInstallationSetupEventAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceInstallationSetupEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceInstallationSetupEventAllOf) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *SpaceInstallationSetupEventAllOf) GetStatus() IntegrationInstallationStatus {
	if o == nil {
		var ret IntegrationInstallationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SpaceInstallationSetupEventAllOf) GetStatusOk() (*IntegrationInstallationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SpaceInstallationSetupEventAllOf) SetStatus(v IntegrationInstallationStatus) {
	o.Status = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *SpaceInstallationSetupEventAllOf) GetPrevious() SpaceInstallationSetupEventAllOfPrevious {
	if o == nil || IsNil(o.Previous) {
		var ret SpaceInstallationSetupEventAllOfPrevious
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceInstallationSetupEventAllOf) GetPreviousOk() (*SpaceInstallationSetupEventAllOfPrevious, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *SpaceInstallationSetupEventAllOf) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given SpaceInstallationSetupEventAllOfPrevious and assigns it to the Previous field.
func (o *SpaceInstallationSetupEventAllOf) SetPrevious(v SpaceInstallationSetupEventAllOfPrevious) {
	o.Previous = &v
}

func (o SpaceInstallationSetupEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceInstallationSetupEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	return toSerialize, nil
}

type NullableSpaceInstallationSetupEventAllOf struct {
	value *SpaceInstallationSetupEventAllOf
	isSet bool
}

func (v NullableSpaceInstallationSetupEventAllOf) Get() *SpaceInstallationSetupEventAllOf {
	return v.value
}

func (v *NullableSpaceInstallationSetupEventAllOf) Set(val *SpaceInstallationSetupEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceInstallationSetupEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceInstallationSetupEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceInstallationSetupEventAllOf(val *SpaceInstallationSetupEventAllOf) *NullableSpaceInstallationSetupEventAllOf {
	return &NullableSpaceInstallationSetupEventAllOf{value: val, isSet: true}
}

func (v NullableSpaceInstallationSetupEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceInstallationSetupEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
