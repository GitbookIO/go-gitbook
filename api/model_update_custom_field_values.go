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

// checks if the UpdateCustomFieldValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomFieldValues{}

// UpdateCustomFieldValues struct for UpdateCustomFieldValues
type UpdateCustomFieldValues struct {
	Values map[string]UpdateCustomFieldValuesValuesValue `json:"values"`
}

// NewUpdateCustomFieldValues instantiates a new UpdateCustomFieldValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomFieldValues(values map[string]UpdateCustomFieldValuesValuesValue) *UpdateCustomFieldValues {
	this := UpdateCustomFieldValues{}
	this.Values = values
	return &this
}

// NewUpdateCustomFieldValuesWithDefaults instantiates a new UpdateCustomFieldValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomFieldValuesWithDefaults() *UpdateCustomFieldValues {
	this := UpdateCustomFieldValues{}
	return &this
}

// GetValues returns the Values field value
func (o *UpdateCustomFieldValues) GetValues() map[string]UpdateCustomFieldValuesValuesValue {
	if o == nil {
		var ret map[string]UpdateCustomFieldValuesValuesValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomFieldValues) GetValuesOk() (*map[string]UpdateCustomFieldValuesValuesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *UpdateCustomFieldValues) SetValues(v map[string]UpdateCustomFieldValuesValuesValue) {
	o.Values = v
}

func (o UpdateCustomFieldValues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomFieldValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableUpdateCustomFieldValues struct {
	value *UpdateCustomFieldValues
	isSet bool
}

func (v NullableUpdateCustomFieldValues) Get() *UpdateCustomFieldValues {
	return v.value
}

func (v *NullableUpdateCustomFieldValues) Set(val *UpdateCustomFieldValues) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomFieldValues) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomFieldValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomFieldValues(val *UpdateCustomFieldValues) *NullableUpdateCustomFieldValues {
	return &NullableUpdateCustomFieldValues{value: val, isSet: true}
}

func (v NullableUpdateCustomFieldValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomFieldValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
