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

// checks if the SpaceVisibilityUpdatedEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceVisibilityUpdatedEventAllOf{}

// SpaceVisibilityUpdatedEventAllOf Event when the visibility of the space has been changed.
type SpaceVisibilityUpdatedEventAllOf struct {
	Type               string            `json:"type"`
	PreviousVisibility ContentVisibility `json:"previousVisibility"`
	Visibility         ContentVisibility `json:"visibility"`
}

// NewSpaceVisibilityUpdatedEventAllOf instantiates a new SpaceVisibilityUpdatedEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceVisibilityUpdatedEventAllOf(type_ string, previousVisibility ContentVisibility, visibility ContentVisibility) *SpaceVisibilityUpdatedEventAllOf {
	this := SpaceVisibilityUpdatedEventAllOf{}
	this.Type = type_
	this.PreviousVisibility = previousVisibility
	this.Visibility = visibility
	return &this
}

// NewSpaceVisibilityUpdatedEventAllOfWithDefaults instantiates a new SpaceVisibilityUpdatedEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceVisibilityUpdatedEventAllOfWithDefaults() *SpaceVisibilityUpdatedEventAllOf {
	this := SpaceVisibilityUpdatedEventAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *SpaceVisibilityUpdatedEventAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceVisibilityUpdatedEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceVisibilityUpdatedEventAllOf) SetType(v string) {
	o.Type = v
}

// GetPreviousVisibility returns the PreviousVisibility field value
func (o *SpaceVisibilityUpdatedEventAllOf) GetPreviousVisibility() ContentVisibility {
	if o == nil {
		var ret ContentVisibility
		return ret
	}

	return o.PreviousVisibility
}

// GetPreviousVisibilityOk returns a tuple with the PreviousVisibility field value
// and a boolean to check if the value has been set.
func (o *SpaceVisibilityUpdatedEventAllOf) GetPreviousVisibilityOk() (*ContentVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousVisibility, true
}

// SetPreviousVisibility sets field value
func (o *SpaceVisibilityUpdatedEventAllOf) SetPreviousVisibility(v ContentVisibility) {
	o.PreviousVisibility = v
}

// GetVisibility returns the Visibility field value
func (o *SpaceVisibilityUpdatedEventAllOf) GetVisibility() ContentVisibility {
	if o == nil {
		var ret ContentVisibility
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *SpaceVisibilityUpdatedEventAllOf) GetVisibilityOk() (*ContentVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *SpaceVisibilityUpdatedEventAllOf) SetVisibility(v ContentVisibility) {
	o.Visibility = v
}

func (o SpaceVisibilityUpdatedEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceVisibilityUpdatedEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["previousVisibility"] = o.PreviousVisibility
	toSerialize["visibility"] = o.Visibility
	return toSerialize, nil
}

type NullableSpaceVisibilityUpdatedEventAllOf struct {
	value *SpaceVisibilityUpdatedEventAllOf
	isSet bool
}

func (v NullableSpaceVisibilityUpdatedEventAllOf) Get() *SpaceVisibilityUpdatedEventAllOf {
	return v.value
}

func (v *NullableSpaceVisibilityUpdatedEventAllOf) Set(val *SpaceVisibilityUpdatedEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceVisibilityUpdatedEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceVisibilityUpdatedEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceVisibilityUpdatedEventAllOf(val *SpaceVisibilityUpdatedEventAllOf) *NullableSpaceVisibilityUpdatedEventAllOf {
	return &NullableSpaceVisibilityUpdatedEventAllOf{value: val, isSet: true}
}

func (v NullableSpaceVisibilityUpdatedEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceVisibilityUpdatedEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
