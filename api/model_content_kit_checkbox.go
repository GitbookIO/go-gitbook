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

// checks if the ContentKitCheckbox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitCheckbox{}

// ContentKitCheckbox struct for ContentKitCheckbox
type ContentKitCheckbox struct {
	Type string `json:"type"`
	// State binding. The value of the input will be stored as a property in the state named after this ID.
	State   string                  `json:"state"`
	Value   ContentKitCheckboxValue `json:"value"`
	Confirm *ContentKitConfirm      `json:"confirm,omitempty"`
}

// NewContentKitCheckbox instantiates a new ContentKitCheckbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitCheckbox(type_ string, state string, value ContentKitCheckboxValue) *ContentKitCheckbox {
	this := ContentKitCheckbox{}
	this.Type = type_
	this.State = state
	this.Value = value
	return &this
}

// NewContentKitCheckboxWithDefaults instantiates a new ContentKitCheckbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitCheckboxWithDefaults() *ContentKitCheckbox {
	this := ContentKitCheckbox{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitCheckbox) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitCheckbox) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitCheckbox) SetType(v string) {
	o.Type = v
}

// GetState returns the State field value
func (o *ContentKitCheckbox) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ContentKitCheckbox) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ContentKitCheckbox) SetState(v string) {
	o.State = v
}

// GetValue returns the Value field value
func (o *ContentKitCheckbox) GetValue() ContentKitCheckboxValue {
	if o == nil {
		var ret ContentKitCheckboxValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ContentKitCheckbox) GetValueOk() (*ContentKitCheckboxValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ContentKitCheckbox) SetValue(v ContentKitCheckboxValue) {
	o.Value = v
}

// GetConfirm returns the Confirm field value if set, zero value otherwise.
func (o *ContentKitCheckbox) GetConfirm() ContentKitConfirm {
	if o == nil || IsNil(o.Confirm) {
		var ret ContentKitConfirm
		return ret
	}
	return *o.Confirm
}

// GetConfirmOk returns a tuple with the Confirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCheckbox) GetConfirmOk() (*ContentKitConfirm, bool) {
	if o == nil || IsNil(o.Confirm) {
		return nil, false
	}
	return o.Confirm, true
}

// HasConfirm returns a boolean if a field has been set.
func (o *ContentKitCheckbox) HasConfirm() bool {
	if o != nil && !IsNil(o.Confirm) {
		return true
	}

	return false
}

// SetConfirm gets a reference to the given ContentKitConfirm and assigns it to the Confirm field.
func (o *ContentKitCheckbox) SetConfirm(v ContentKitConfirm) {
	o.Confirm = &v
}

func (o ContentKitCheckbox) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitCheckbox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["state"] = o.State
	toSerialize["value"] = o.Value
	if !IsNil(o.Confirm) {
		toSerialize["confirm"] = o.Confirm
	}
	return toSerialize, nil
}

type NullableContentKitCheckbox struct {
	value *ContentKitCheckbox
	isSet bool
}

func (v NullableContentKitCheckbox) Get() *ContentKitCheckbox {
	return v.value
}

func (v *NullableContentKitCheckbox) Set(val *ContentKitCheckbox) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitCheckbox) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitCheckbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitCheckbox(val *ContentKitCheckbox) *NullableContentKitCheckbox {
	return &NullableContentKitCheckbox{value: val, isSet: true}
}

func (v NullableContentKitCheckbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitCheckbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
