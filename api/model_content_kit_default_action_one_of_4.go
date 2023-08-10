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

// checks if the ContentKitDefaultActionOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitDefaultActionOneOf4{}

// ContentKitDefaultActionOneOf4 Action to update the properties stored in the related node.
type ContentKitDefaultActionOneOf4 struct {
	Action string                 `json:"action"`
	Props  map[string]interface{} `json:"props"`
}

// NewContentKitDefaultActionOneOf4 instantiates a new ContentKitDefaultActionOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitDefaultActionOneOf4(action string, props map[string]interface{}) *ContentKitDefaultActionOneOf4 {
	this := ContentKitDefaultActionOneOf4{}
	this.Action = action
	this.Props = props
	return &this
}

// NewContentKitDefaultActionOneOf4WithDefaults instantiates a new ContentKitDefaultActionOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitDefaultActionOneOf4WithDefaults() *ContentKitDefaultActionOneOf4 {
	this := ContentKitDefaultActionOneOf4{}
	return &this
}

// GetAction returns the Action field value
func (o *ContentKitDefaultActionOneOf4) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf4) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContentKitDefaultActionOneOf4) SetAction(v string) {
	o.Action = v
}

// GetProps returns the Props field value
func (o *ContentKitDefaultActionOneOf4) GetProps() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Props
}

// GetPropsOk returns a tuple with the Props field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf4) GetPropsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Props, true
}

// SetProps sets field value
func (o *ContentKitDefaultActionOneOf4) SetProps(v map[string]interface{}) {
	o.Props = v
}

func (o ContentKitDefaultActionOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitDefaultActionOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["props"] = o.Props
	return toSerialize, nil
}

type NullableContentKitDefaultActionOneOf4 struct {
	value *ContentKitDefaultActionOneOf4
	isSet bool
}

func (v NullableContentKitDefaultActionOneOf4) Get() *ContentKitDefaultActionOneOf4 {
	return v.value
}

func (v *NullableContentKitDefaultActionOneOf4) Set(val *ContentKitDefaultActionOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitDefaultActionOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitDefaultActionOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitDefaultActionOneOf4(val *ContentKitDefaultActionOneOf4) *NullableContentKitDefaultActionOneOf4 {
	return &NullableContentKitDefaultActionOneOf4{value: val, isSet: true}
}

func (v NullableContentKitDefaultActionOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitDefaultActionOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
