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

// checks if the ContentKitSelect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitSelect{}

// ContentKitSelect Creates a drop down menu with a list of options for a user to choose.
type ContentKitSelect struct {
	Type string `json:"type"`
	// State binding. The value of the input will be stored as a property in the state named after this ID.
	State         string                        `json:"state"`
	InitialValue  *ContentKitSelectInitialValue `json:"initialValue,omitempty"`
	OnValueChange *ContentKitAction             `json:"onValueChange,omitempty"`
	// Text that appears in the form control when it has no value set
	Placeholder *string `json:"placeholder,omitempty"`
	// Should the select accept the selection of multiple options. If true, the state will be an array.
	Multiple *bool `json:"multiple,omitempty"`
	// Should the filter input be allowed to be selected as an option.
	AcceptInput *bool                   `json:"acceptInput,omitempty"`
	Options     ContentKitSelectOptions `json:"options"`
}

// NewContentKitSelect instantiates a new ContentKitSelect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitSelect(type_ string, state string, options ContentKitSelectOptions) *ContentKitSelect {
	this := ContentKitSelect{}
	this.Type = type_
	this.State = state
	this.Options = options
	return &this
}

// NewContentKitSelectWithDefaults instantiates a new ContentKitSelect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitSelectWithDefaults() *ContentKitSelect {
	this := ContentKitSelect{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitSelect) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitSelect) SetType(v string) {
	o.Type = v
}

// GetState returns the State field value
func (o *ContentKitSelect) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ContentKitSelect) SetState(v string) {
	o.State = v
}

// GetInitialValue returns the InitialValue field value if set, zero value otherwise.
func (o *ContentKitSelect) GetInitialValue() ContentKitSelectInitialValue {
	if o == nil || IsNil(o.InitialValue) {
		var ret ContentKitSelectInitialValue
		return ret
	}
	return *o.InitialValue
}

// GetInitialValueOk returns a tuple with the InitialValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetInitialValueOk() (*ContentKitSelectInitialValue, bool) {
	if o == nil || IsNil(o.InitialValue) {
		return nil, false
	}
	return o.InitialValue, true
}

// HasInitialValue returns a boolean if a field has been set.
func (o *ContentKitSelect) HasInitialValue() bool {
	if o != nil && !IsNil(o.InitialValue) {
		return true
	}

	return false
}

// SetInitialValue gets a reference to the given ContentKitSelectInitialValue and assigns it to the InitialValue field.
func (o *ContentKitSelect) SetInitialValue(v ContentKitSelectInitialValue) {
	o.InitialValue = &v
}

// GetOnValueChange returns the OnValueChange field value if set, zero value otherwise.
func (o *ContentKitSelect) GetOnValueChange() ContentKitAction {
	if o == nil || IsNil(o.OnValueChange) {
		var ret ContentKitAction
		return ret
	}
	return *o.OnValueChange
}

// GetOnValueChangeOk returns a tuple with the OnValueChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetOnValueChangeOk() (*ContentKitAction, bool) {
	if o == nil || IsNil(o.OnValueChange) {
		return nil, false
	}
	return o.OnValueChange, true
}

// HasOnValueChange returns a boolean if a field has been set.
func (o *ContentKitSelect) HasOnValueChange() bool {
	if o != nil && !IsNil(o.OnValueChange) {
		return true
	}

	return false
}

// SetOnValueChange gets a reference to the given ContentKitAction and assigns it to the OnValueChange field.
func (o *ContentKitSelect) SetOnValueChange(v ContentKitAction) {
	o.OnValueChange = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *ContentKitSelect) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *ContentKitSelect) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *ContentKitSelect) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetMultiple returns the Multiple field value if set, zero value otherwise.
func (o *ContentKitSelect) GetMultiple() bool {
	if o == nil || IsNil(o.Multiple) {
		var ret bool
		return ret
	}
	return *o.Multiple
}

// GetMultipleOk returns a tuple with the Multiple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetMultipleOk() (*bool, bool) {
	if o == nil || IsNil(o.Multiple) {
		return nil, false
	}
	return o.Multiple, true
}

// HasMultiple returns a boolean if a field has been set.
func (o *ContentKitSelect) HasMultiple() bool {
	if o != nil && !IsNil(o.Multiple) {
		return true
	}

	return false
}

// SetMultiple gets a reference to the given bool and assigns it to the Multiple field.
func (o *ContentKitSelect) SetMultiple(v bool) {
	o.Multiple = &v
}

// GetAcceptInput returns the AcceptInput field value if set, zero value otherwise.
func (o *ContentKitSelect) GetAcceptInput() bool {
	if o == nil || IsNil(o.AcceptInput) {
		var ret bool
		return ret
	}
	return *o.AcceptInput
}

// GetAcceptInputOk returns a tuple with the AcceptInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetAcceptInputOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptInput) {
		return nil, false
	}
	return o.AcceptInput, true
}

// HasAcceptInput returns a boolean if a field has been set.
func (o *ContentKitSelect) HasAcceptInput() bool {
	if o != nil && !IsNil(o.AcceptInput) {
		return true
	}

	return false
}

// SetAcceptInput gets a reference to the given bool and assigns it to the AcceptInput field.
func (o *ContentKitSelect) SetAcceptInput(v bool) {
	o.AcceptInput = &v
}

// GetOptions returns the Options field value
func (o *ContentKitSelect) GetOptions() ContentKitSelectOptions {
	if o == nil {
		var ret ContentKitSelectOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ContentKitSelect) GetOptionsOk() (*ContentKitSelectOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *ContentKitSelect) SetOptions(v ContentKitSelectOptions) {
	o.Options = v
}

func (o ContentKitSelect) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitSelect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["state"] = o.State
	if !IsNil(o.InitialValue) {
		toSerialize["initialValue"] = o.InitialValue
	}
	if !IsNil(o.OnValueChange) {
		toSerialize["onValueChange"] = o.OnValueChange
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !IsNil(o.Multiple) {
		toSerialize["multiple"] = o.Multiple
	}
	if !IsNil(o.AcceptInput) {
		toSerialize["acceptInput"] = o.AcceptInput
	}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

type NullableContentKitSelect struct {
	value *ContentKitSelect
	isSet bool
}

func (v NullableContentKitSelect) Get() *ContentKitSelect {
	return v.value
}

func (v *NullableContentKitSelect) Set(val *ContentKitSelect) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitSelect) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitSelect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitSelect(val *ContentKitSelect) *NullableContentKitSelect {
	return &NullableContentKitSelect{value: val, isSet: true}
}

func (v NullableContentKitSelect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitSelect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
