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

// checks if the RevisionTypeRollbackAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionTypeRollbackAllOf{}

// RevisionTypeRollbackAllOf struct for RevisionTypeRollbackAllOf
type RevisionTypeRollbackAllOf struct {
	// Revision created as a rollback of a previous revision.
	Type string `json:"type"`
}

// NewRevisionTypeRollbackAllOf instantiates a new RevisionTypeRollbackAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionTypeRollbackAllOf(type_ string) *RevisionTypeRollbackAllOf {
	this := RevisionTypeRollbackAllOf{}
	this.Type = type_
	return &this
}

// NewRevisionTypeRollbackAllOfWithDefaults instantiates a new RevisionTypeRollbackAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionTypeRollbackAllOfWithDefaults() *RevisionTypeRollbackAllOf {
	this := RevisionTypeRollbackAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *RevisionTypeRollbackAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeRollbackAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RevisionTypeRollbackAllOf) SetType(v string) {
	o.Type = v
}

func (o RevisionTypeRollbackAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionTypeRollbackAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableRevisionTypeRollbackAllOf struct {
	value *RevisionTypeRollbackAllOf
	isSet bool
}

func (v NullableRevisionTypeRollbackAllOf) Get() *RevisionTypeRollbackAllOf {
	return v.value
}

func (v *NullableRevisionTypeRollbackAllOf) Set(val *RevisionTypeRollbackAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionTypeRollbackAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionTypeRollbackAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionTypeRollbackAllOf(val *RevisionTypeRollbackAllOf) *NullableRevisionTypeRollbackAllOf {
	return &NullableRevisionTypeRollbackAllOf{value: val, isSet: true}
}

func (v NullableRevisionTypeRollbackAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionTypeRollbackAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}