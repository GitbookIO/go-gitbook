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

// checks if the ContentKitDefaultActionOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitDefaultActionOneOf1{}

// ContentKitDefaultActionOneOf1 Action when a modal overlay is closed, with a return value to the higher level component in the stack. This action will be triggered on the parent component instance.
type ContentKitDefaultActionOneOf1 struct {
	Action      string                 `json:"action"`
	ReturnValue map[string]interface{} `json:"returnValue"`
}

// NewContentKitDefaultActionOneOf1 instantiates a new ContentKitDefaultActionOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitDefaultActionOneOf1(action string, returnValue map[string]interface{}) *ContentKitDefaultActionOneOf1 {
	this := ContentKitDefaultActionOneOf1{}
	this.Action = action
	this.ReturnValue = returnValue
	return &this
}

// NewContentKitDefaultActionOneOf1WithDefaults instantiates a new ContentKitDefaultActionOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitDefaultActionOneOf1WithDefaults() *ContentKitDefaultActionOneOf1 {
	this := ContentKitDefaultActionOneOf1{}
	return &this
}

// GetAction returns the Action field value
func (o *ContentKitDefaultActionOneOf1) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf1) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContentKitDefaultActionOneOf1) SetAction(v string) {
	o.Action = v
}

// GetReturnValue returns the ReturnValue field value
func (o *ContentKitDefaultActionOneOf1) GetReturnValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ReturnValue
}

// GetReturnValueOk returns a tuple with the ReturnValue field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf1) GetReturnValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ReturnValue, true
}

// SetReturnValue sets field value
func (o *ContentKitDefaultActionOneOf1) SetReturnValue(v map[string]interface{}) {
	o.ReturnValue = v
}

func (o ContentKitDefaultActionOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitDefaultActionOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["returnValue"] = o.ReturnValue
	return toSerialize, nil
}

type NullableContentKitDefaultActionOneOf1 struct {
	value *ContentKitDefaultActionOneOf1
	isSet bool
}

func (v NullableContentKitDefaultActionOneOf1) Get() *ContentKitDefaultActionOneOf1 {
	return v.value
}

func (v *NullableContentKitDefaultActionOneOf1) Set(val *ContentKitDefaultActionOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitDefaultActionOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitDefaultActionOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitDefaultActionOneOf1(val *ContentKitDefaultActionOneOf1) *NullableContentKitDefaultActionOneOf1 {
	return &NullableContentKitDefaultActionOneOf1{value: val, isSet: true}
}

func (v NullableContentKitDefaultActionOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitDefaultActionOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
