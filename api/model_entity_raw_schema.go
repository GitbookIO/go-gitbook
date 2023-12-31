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

// checks if the EntityRawSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityRawSchema{}

// EntityRawSchema Schema for a type of entities
type EntityRawSchema struct {
	// Type of an entity
	Type  string               `json:"type"`
	Title EntityRawSchemaTitle `json:"title"`
	// Ordered list of all properties stored in entities.
	Properties []EntityPropertySchema `json:"properties"`
}

// NewEntityRawSchema instantiates a new EntityRawSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRawSchema(type_ string, title EntityRawSchemaTitle, properties []EntityPropertySchema) *EntityRawSchema {
	this := EntityRawSchema{}
	this.Type = type_
	this.Title = title
	this.Properties = properties
	return &this
}

// NewEntityRawSchemaWithDefaults instantiates a new EntityRawSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRawSchemaWithDefaults() *EntityRawSchema {
	this := EntityRawSchema{}
	return &this
}

// GetType returns the Type field value
func (o *EntityRawSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntityRawSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntityRawSchema) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *EntityRawSchema) GetTitle() EntityRawSchemaTitle {
	if o == nil {
		var ret EntityRawSchemaTitle
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EntityRawSchema) GetTitleOk() (*EntityRawSchemaTitle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *EntityRawSchema) SetTitle(v EntityRawSchemaTitle) {
	o.Title = v
}

// GetProperties returns the Properties field value
func (o *EntityRawSchema) GetProperties() []EntityPropertySchema {
	if o == nil {
		var ret []EntityPropertySchema
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *EntityRawSchema) GetPropertiesOk() ([]EntityPropertySchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *EntityRawSchema) SetProperties(v []EntityPropertySchema) {
	o.Properties = v
}

func (o EntityRawSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityRawSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableEntityRawSchema struct {
	value *EntityRawSchema
	isSet bool
}

func (v NullableEntityRawSchema) Get() *EntityRawSchema {
	return v.value
}

func (v *NullableEntityRawSchema) Set(val *EntityRawSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRawSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRawSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRawSchema(val *EntityRawSchema) *NullableEntityRawSchema {
	return &NullableEntityRawSchema{value: val, isSet: true}
}

func (v NullableEntityRawSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRawSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
