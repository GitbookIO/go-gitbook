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

// checks if the CustomFieldValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldValuesInner{}

// CustomFieldValuesInner struct for CustomFieldValuesInner
type CustomFieldValuesInner struct {
	CustomField CustomField       `json:"customField"`
	Value       *CustomFieldValue `json:"value,omitempty"`
}

// NewCustomFieldValuesInner instantiates a new CustomFieldValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldValuesInner(customField CustomField) *CustomFieldValuesInner {
	this := CustomFieldValuesInner{}
	this.CustomField = customField
	return &this
}

// NewCustomFieldValuesInnerWithDefaults instantiates a new CustomFieldValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldValuesInnerWithDefaults() *CustomFieldValuesInner {
	this := CustomFieldValuesInner{}
	return &this
}

// GetCustomField returns the CustomField field value
func (o *CustomFieldValuesInner) GetCustomField() CustomField {
	if o == nil {
		var ret CustomField
		return ret
	}

	return o.CustomField
}

// GetCustomFieldOk returns a tuple with the CustomField field value
// and a boolean to check if the value has been set.
func (o *CustomFieldValuesInner) GetCustomFieldOk() (*CustomField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomField, true
}

// SetCustomField sets field value
func (o *CustomFieldValuesInner) SetCustomField(v CustomField) {
	o.CustomField = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomFieldValuesInner) GetValue() CustomFieldValue {
	if o == nil || IsNil(o.Value) {
		var ret CustomFieldValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldValuesInner) GetValueOk() (*CustomFieldValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldValuesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CustomFieldValue and assigns it to the Value field.
func (o *CustomFieldValuesInner) SetValue(v CustomFieldValue) {
	o.Value = &v
}

func (o CustomFieldValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customField"] = o.CustomField
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCustomFieldValuesInner struct {
	value *CustomFieldValuesInner
	isSet bool
}

func (v NullableCustomFieldValuesInner) Get() *CustomFieldValuesInner {
	return v.value
}

func (v *NullableCustomFieldValuesInner) Set(val *CustomFieldValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldValuesInner(val *CustomFieldValuesInner) *NullableCustomFieldValuesInner {
	return &NullableCustomFieldValuesInner{value: val, isSet: true}
}

func (v NullableCustomFieldValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
