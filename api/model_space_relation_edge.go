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

// checks if the SpaceRelationEdge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceRelationEdge{}

// SpaceRelationEdge Edge of a space relation (target or source)
type SpaceRelationEdge struct {
	// ID of the space
	Id string `json:"id"`
}

// NewSpaceRelationEdge instantiates a new SpaceRelationEdge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceRelationEdge(id string) *SpaceRelationEdge {
	this := SpaceRelationEdge{}
	this.Id = id
	return &this
}

// NewSpaceRelationEdgeWithDefaults instantiates a new SpaceRelationEdge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceRelationEdgeWithDefaults() *SpaceRelationEdge {
	this := SpaceRelationEdge{}
	return &this
}

// GetId returns the Id field value
func (o *SpaceRelationEdge) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationEdge) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SpaceRelationEdge) SetId(v string) {
	o.Id = v
}

func (o SpaceRelationEdge) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceRelationEdge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSpaceRelationEdge struct {
	value *SpaceRelationEdge
	isSet bool
}

func (v NullableSpaceRelationEdge) Get() *SpaceRelationEdge {
	return v.value
}

func (v *NullableSpaceRelationEdge) Set(val *SpaceRelationEdge) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceRelationEdge) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceRelationEdge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceRelationEdge(val *SpaceRelationEdge) *NullableSpaceRelationEdge {
	return &NullableSpaceRelationEdge{value: val, isSet: true}
}

func (v NullableSpaceRelationEdge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceRelationEdge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
