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

// checks if the SpaceInstallationDeletedEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceInstallationDeletedEventAllOf{}

// SpaceInstallationDeletedEventAllOf Event received when integration has been uninstalled from a space.
type SpaceInstallationDeletedEventAllOf struct {
	Type     string                                     `json:"type"`
	Previous SpaceInstallationDeletedEventAllOfPrevious `json:"previous"`
}

// NewSpaceInstallationDeletedEventAllOf instantiates a new SpaceInstallationDeletedEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceInstallationDeletedEventAllOf(type_ string, previous SpaceInstallationDeletedEventAllOfPrevious) *SpaceInstallationDeletedEventAllOf {
	this := SpaceInstallationDeletedEventAllOf{}
	this.Type = type_
	this.Previous = previous
	return &this
}

// NewSpaceInstallationDeletedEventAllOfWithDefaults instantiates a new SpaceInstallationDeletedEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceInstallationDeletedEventAllOfWithDefaults() *SpaceInstallationDeletedEventAllOf {
	this := SpaceInstallationDeletedEventAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *SpaceInstallationDeletedEventAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceInstallationDeletedEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceInstallationDeletedEventAllOf) SetType(v string) {
	o.Type = v
}

// GetPrevious returns the Previous field value
func (o *SpaceInstallationDeletedEventAllOf) GetPrevious() SpaceInstallationDeletedEventAllOfPrevious {
	if o == nil {
		var ret SpaceInstallationDeletedEventAllOfPrevious
		return ret
	}

	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
func (o *SpaceInstallationDeletedEventAllOf) GetPreviousOk() (*SpaceInstallationDeletedEventAllOfPrevious, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Previous, true
}

// SetPrevious sets field value
func (o *SpaceInstallationDeletedEventAllOf) SetPrevious(v SpaceInstallationDeletedEventAllOfPrevious) {
	o.Previous = v
}

func (o SpaceInstallationDeletedEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceInstallationDeletedEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["previous"] = o.Previous
	return toSerialize, nil
}

type NullableSpaceInstallationDeletedEventAllOf struct {
	value *SpaceInstallationDeletedEventAllOf
	isSet bool
}

func (v NullableSpaceInstallationDeletedEventAllOf) Get() *SpaceInstallationDeletedEventAllOf {
	return v.value
}

func (v *NullableSpaceInstallationDeletedEventAllOf) Set(val *SpaceInstallationDeletedEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceInstallationDeletedEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceInstallationDeletedEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceInstallationDeletedEventAllOf(val *SpaceInstallationDeletedEventAllOf) *NullableSpaceInstallationDeletedEventAllOf {
	return &NullableSpaceInstallationDeletedEventAllOf{value: val, isSet: true}
}

func (v NullableSpaceInstallationDeletedEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceInstallationDeletedEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
