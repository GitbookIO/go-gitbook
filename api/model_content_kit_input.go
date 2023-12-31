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

// checks if the ContentKitInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitInput{}

// ContentKitInput Field for an input.
type ContentKitInput struct {
	Type string `json:"type"`
	// Text label displayed next to the input.
	Label   string                 `json:"label"`
	Hint    *ContentKitInputHint   `json:"hint,omitempty"`
	Element ContentKitInputElement `json:"element"`
}

// NewContentKitInput instantiates a new ContentKitInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitInput(type_ string, label string, element ContentKitInputElement) *ContentKitInput {
	this := ContentKitInput{}
	this.Type = type_
	this.Label = label
	this.Element = element
	return &this
}

// NewContentKitInputWithDefaults instantiates a new ContentKitInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitInputWithDefaults() *ContentKitInput {
	this := ContentKitInput{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitInput) SetType(v string) {
	o.Type = v
}

// GetLabel returns the Label field value
func (o *ContentKitInput) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ContentKitInput) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ContentKitInput) SetLabel(v string) {
	o.Label = v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *ContentKitInput) GetHint() ContentKitInputHint {
	if o == nil || IsNil(o.Hint) {
		var ret ContentKitInputHint
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitInput) GetHintOk() (*ContentKitInputHint, bool) {
	if o == nil || IsNil(o.Hint) {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *ContentKitInput) HasHint() bool {
	if o != nil && !IsNil(o.Hint) {
		return true
	}

	return false
}

// SetHint gets a reference to the given ContentKitInputHint and assigns it to the Hint field.
func (o *ContentKitInput) SetHint(v ContentKitInputHint) {
	o.Hint = &v
}

// GetElement returns the Element field value
func (o *ContentKitInput) GetElement() ContentKitInputElement {
	if o == nil {
		var ret ContentKitInputElement
		return ret
	}

	return o.Element
}

// GetElementOk returns a tuple with the Element field value
// and a boolean to check if the value has been set.
func (o *ContentKitInput) GetElementOk() (*ContentKitInputElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Element, true
}

// SetElement sets field value
func (o *ContentKitInput) SetElement(v ContentKitInputElement) {
	o.Element = v
}

func (o ContentKitInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["label"] = o.Label
	if !IsNil(o.Hint) {
		toSerialize["hint"] = o.Hint
	}
	toSerialize["element"] = o.Element
	return toSerialize, nil
}

type NullableContentKitInput struct {
	value *ContentKitInput
	isSet bool
}

func (v NullableContentKitInput) Get() *ContentKitInput {
	return v.value
}

func (v *NullableContentKitInput) Set(val *ContentKitInput) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitInput) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitInput(val *ContentKitInput) *NullableContentKitInput {
	return &NullableContentKitInput{value: val, isSet: true}
}

func (v NullableContentKitInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
