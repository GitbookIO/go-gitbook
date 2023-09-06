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

// checks if the EntityPropertySchemaAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityPropertySchemaAllOf{}

// EntityPropertySchemaAllOf struct for EntityPropertySchemaAllOf
type EntityPropertySchemaAllOf struct {
	// Name of the property in the object
	Name string `json:"name"`
	// Title displayed to the users
	Title string `json:"title"`
	// Description of the property
	Description *string `json:"description,omitempty"`
	// If true, the property is no longer required and not taken into consideration
	Deprecated *bool `json:"deprecated,omitempty"`
}

// NewEntityPropertySchemaAllOf instantiates a new EntityPropertySchemaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityPropertySchemaAllOf(name string, title string) *EntityPropertySchemaAllOf {
	this := EntityPropertySchemaAllOf{}
	this.Name = name
	this.Title = title
	return &this
}

// NewEntityPropertySchemaAllOfWithDefaults instantiates a new EntityPropertySchemaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityPropertySchemaAllOfWithDefaults() *EntityPropertySchemaAllOf {
	this := EntityPropertySchemaAllOf{}
	return &this
}

// GetName returns the Name field value
func (o *EntityPropertySchemaAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntityPropertySchemaAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntityPropertySchemaAllOf) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *EntityPropertySchemaAllOf) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EntityPropertySchemaAllOf) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *EntityPropertySchemaAllOf) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntityPropertySchemaAllOf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySchemaAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntityPropertySchemaAllOf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntityPropertySchemaAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *EntityPropertySchemaAllOf) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityPropertySchemaAllOf) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *EntityPropertySchemaAllOf) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *EntityPropertySchemaAllOf) SetDeprecated(v bool) {
	o.Deprecated = &v
}

func (o EntityPropertySchemaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityPropertySchemaAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	return toSerialize, nil
}

type NullableEntityPropertySchemaAllOf struct {
	value *EntityPropertySchemaAllOf
	isSet bool
}

func (v NullableEntityPropertySchemaAllOf) Get() *EntityPropertySchemaAllOf {
	return v.value
}

func (v *NullableEntityPropertySchemaAllOf) Set(val *EntityPropertySchemaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityPropertySchemaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityPropertySchemaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityPropertySchemaAllOf(val *EntityPropertySchemaAllOf) *NullableEntityPropertySchemaAllOf {
	return &NullableEntityPropertySchemaAllOf{value: val, isSet: true}
}

func (v NullableEntityPropertySchemaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityPropertySchemaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
