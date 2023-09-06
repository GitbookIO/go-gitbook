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

// checks if the ContentKitSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitSwitch{}

// ContentKitSwitch Renders a boolean input.
type ContentKitSwitch struct {
	Type string `json:"type"`
	// State binding. The value of the input will be stored as a property in the state named after this ID.
	State string `json:"state"`
	// Value to initialize the switch with.
	InitialValue  *bool              `json:"initialValue,omitempty"`
	OnValueChange *ContentKitAction  `json:"onValueChange,omitempty"`
	Confirm       *ContentKitConfirm `json:"confirm,omitempty"`
}

// NewContentKitSwitch instantiates a new ContentKitSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitSwitch(type_ string, state string) *ContentKitSwitch {
	this := ContentKitSwitch{}
	this.Type = type_
	this.State = state
	return &this
}

// NewContentKitSwitchWithDefaults instantiates a new ContentKitSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitSwitchWithDefaults() *ContentKitSwitch {
	this := ContentKitSwitch{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitSwitch) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitSwitch) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitSwitch) SetType(v string) {
	o.Type = v
}

// GetState returns the State field value
func (o *ContentKitSwitch) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ContentKitSwitch) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ContentKitSwitch) SetState(v string) {
	o.State = v
}

// GetInitialValue returns the InitialValue field value if set, zero value otherwise.
func (o *ContentKitSwitch) GetInitialValue() bool {
	if o == nil || IsNil(o.InitialValue) {
		var ret bool
		return ret
	}
	return *o.InitialValue
}

// GetInitialValueOk returns a tuple with the InitialValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSwitch) GetInitialValueOk() (*bool, bool) {
	if o == nil || IsNil(o.InitialValue) {
		return nil, false
	}
	return o.InitialValue, true
}

// HasInitialValue returns a boolean if a field has been set.
func (o *ContentKitSwitch) HasInitialValue() bool {
	if o != nil && !IsNil(o.InitialValue) {
		return true
	}

	return false
}

// SetInitialValue gets a reference to the given bool and assigns it to the InitialValue field.
func (o *ContentKitSwitch) SetInitialValue(v bool) {
	o.InitialValue = &v
}

// GetOnValueChange returns the OnValueChange field value if set, zero value otherwise.
func (o *ContentKitSwitch) GetOnValueChange() ContentKitAction {
	if o == nil || IsNil(o.OnValueChange) {
		var ret ContentKitAction
		return ret
	}
	return *o.OnValueChange
}

// GetOnValueChangeOk returns a tuple with the OnValueChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSwitch) GetOnValueChangeOk() (*ContentKitAction, bool) {
	if o == nil || IsNil(o.OnValueChange) {
		return nil, false
	}
	return o.OnValueChange, true
}

// HasOnValueChange returns a boolean if a field has been set.
func (o *ContentKitSwitch) HasOnValueChange() bool {
	if o != nil && !IsNil(o.OnValueChange) {
		return true
	}

	return false
}

// SetOnValueChange gets a reference to the given ContentKitAction and assigns it to the OnValueChange field.
func (o *ContentKitSwitch) SetOnValueChange(v ContentKitAction) {
	o.OnValueChange = &v
}

// GetConfirm returns the Confirm field value if set, zero value otherwise.
func (o *ContentKitSwitch) GetConfirm() ContentKitConfirm {
	if o == nil || IsNil(o.Confirm) {
		var ret ContentKitConfirm
		return ret
	}
	return *o.Confirm
}

// GetConfirmOk returns a tuple with the Confirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSwitch) GetConfirmOk() (*ContentKitConfirm, bool) {
	if o == nil || IsNil(o.Confirm) {
		return nil, false
	}
	return o.Confirm, true
}

// HasConfirm returns a boolean if a field has been set.
func (o *ContentKitSwitch) HasConfirm() bool {
	if o != nil && !IsNil(o.Confirm) {
		return true
	}

	return false
}

// SetConfirm gets a reference to the given ContentKitConfirm and assigns it to the Confirm field.
func (o *ContentKitSwitch) SetConfirm(v ContentKitConfirm) {
	o.Confirm = &v
}

func (o ContentKitSwitch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitSwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["state"] = o.State
	if !IsNil(o.InitialValue) {
		toSerialize["initialValue"] = o.InitialValue
	}
	if !IsNil(o.OnValueChange) {
		toSerialize["onValueChange"] = o.OnValueChange
	}
	if !IsNil(o.Confirm) {
		toSerialize["confirm"] = o.Confirm
	}
	return toSerialize, nil
}

type NullableContentKitSwitch struct {
	value *ContentKitSwitch
	isSet bool
}

func (v NullableContentKitSwitch) Get() *ContentKitSwitch {
	return v.value
}

func (v *NullableContentKitSwitch) Set(val *ContentKitSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitSwitch(val *ContentKitSwitch) *NullableContentKitSwitch {
	return &NullableContentKitSwitch{value: val, isSet: true}
}

func (v NullableContentKitSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
