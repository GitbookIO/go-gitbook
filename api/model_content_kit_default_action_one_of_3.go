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

// checks if the ContentKitDefaultActionOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitDefaultActionOneOf3{}

// ContentKitDefaultActionOneOf3 Action when a link is being unfurled into a block.
type ContentKitDefaultActionOneOf3 struct {
	Action string `json:"action"`
	Url    string `json:"url"`
}

// NewContentKitDefaultActionOneOf3 instantiates a new ContentKitDefaultActionOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitDefaultActionOneOf3(action string, url string) *ContentKitDefaultActionOneOf3 {
	this := ContentKitDefaultActionOneOf3{}
	this.Action = action
	this.Url = url
	return &this
}

// NewContentKitDefaultActionOneOf3WithDefaults instantiates a new ContentKitDefaultActionOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitDefaultActionOneOf3WithDefaults() *ContentKitDefaultActionOneOf3 {
	this := ContentKitDefaultActionOneOf3{}
	return &this
}

// GetAction returns the Action field value
func (o *ContentKitDefaultActionOneOf3) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf3) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContentKitDefaultActionOneOf3) SetAction(v string) {
	o.Action = v
}

// GetUrl returns the Url field value
func (o *ContentKitDefaultActionOneOf3) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf3) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ContentKitDefaultActionOneOf3) SetUrl(v string) {
	o.Url = v
}

func (o ContentKitDefaultActionOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitDefaultActionOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableContentKitDefaultActionOneOf3 struct {
	value *ContentKitDefaultActionOneOf3
	isSet bool
}

func (v NullableContentKitDefaultActionOneOf3) Get() *ContentKitDefaultActionOneOf3 {
	return v.value
}

func (v *NullableContentKitDefaultActionOneOf3) Set(val *ContentKitDefaultActionOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitDefaultActionOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitDefaultActionOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitDefaultActionOneOf3(val *ContentKitDefaultActionOneOf3) *NullableContentKitDefaultActionOneOf3 {
	return &NullableContentKitDefaultActionOneOf3{value: val, isSet: true}
}

func (v NullableContentKitDefaultActionOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitDefaultActionOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
