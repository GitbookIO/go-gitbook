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

// checks if the ListSchemaEntities200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSchemaEntities200ResponseAllOf{}

// ListSchemaEntities200ResponseAllOf struct for ListSchemaEntities200ResponseAllOf
type ListSchemaEntities200ResponseAllOf struct {
	Items []Entity `json:"items"`
}

// NewListSchemaEntities200ResponseAllOf instantiates a new ListSchemaEntities200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSchemaEntities200ResponseAllOf(items []Entity) *ListSchemaEntities200ResponseAllOf {
	this := ListSchemaEntities200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewListSchemaEntities200ResponseAllOfWithDefaults instantiates a new ListSchemaEntities200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSchemaEntities200ResponseAllOfWithDefaults() *ListSchemaEntities200ResponseAllOf {
	this := ListSchemaEntities200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *ListSchemaEntities200ResponseAllOf) GetItems() []Entity {
	if o == nil {
		var ret []Entity
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListSchemaEntities200ResponseAllOf) GetItemsOk() ([]Entity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListSchemaEntities200ResponseAllOf) SetItems(v []Entity) {
	o.Items = v
}

func (o ListSchemaEntities200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSchemaEntities200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListSchemaEntities200ResponseAllOf struct {
	value *ListSchemaEntities200ResponseAllOf
	isSet bool
}

func (v NullableListSchemaEntities200ResponseAllOf) Get() *ListSchemaEntities200ResponseAllOf {
	return v.value
}

func (v *NullableListSchemaEntities200ResponseAllOf) Set(val *ListSchemaEntities200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListSchemaEntities200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListSchemaEntities200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSchemaEntities200ResponseAllOf(val *ListSchemaEntities200ResponseAllOf) *NullableListSchemaEntities200ResponseAllOf {
	return &NullableListSchemaEntities200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListSchemaEntities200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSchemaEntities200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
