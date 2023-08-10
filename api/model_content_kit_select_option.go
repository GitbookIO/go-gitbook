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

// checks if the ContentKitSelectOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitSelectOption{}

// ContentKitSelectOption An individual option in a select element
type ContentKitSelectOption struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

// NewContentKitSelectOption instantiates a new ContentKitSelectOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitSelectOption(id string, label string) *ContentKitSelectOption {
	this := ContentKitSelectOption{}
	this.Id = id
	this.Label = label
	return &this
}

// NewContentKitSelectOptionWithDefaults instantiates a new ContentKitSelectOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitSelectOptionWithDefaults() *ContentKitSelectOption {
	this := ContentKitSelectOption{}
	return &this
}

// GetId returns the Id field value
func (o *ContentKitSelectOption) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContentKitSelectOption) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContentKitSelectOption) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *ContentKitSelectOption) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ContentKitSelectOption) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ContentKitSelectOption) SetLabel(v string) {
	o.Label = v
}

func (o ContentKitSelectOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitSelectOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableContentKitSelectOption struct {
	value *ContentKitSelectOption
	isSet bool
}

func (v NullableContentKitSelectOption) Get() *ContentKitSelectOption {
	return v.value
}

func (v *NullableContentKitSelectOption) Set(val *ContentKitSelectOption) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitSelectOption) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitSelectOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitSelectOption(val *ContentKitSelectOption) *NullableContentKitSelectOption {
	return &NullableContentKitSelectOption{value: val, isSet: true}
}

func (v NullableContentKitSelectOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitSelectOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
