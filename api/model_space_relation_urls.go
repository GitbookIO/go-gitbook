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

// checks if the SpaceRelationUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceRelationUrls{}

// SpaceRelationUrls URLs associated with the object
type SpaceRelationUrls struct {
	// URL of the space relation in the API
	Location string `json:"location"`
}

// NewSpaceRelationUrls instantiates a new SpaceRelationUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceRelationUrls(location string) *SpaceRelationUrls {
	this := SpaceRelationUrls{}
	this.Location = location
	return &this
}

// NewSpaceRelationUrlsWithDefaults instantiates a new SpaceRelationUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceRelationUrlsWithDefaults() *SpaceRelationUrls {
	this := SpaceRelationUrls{}
	return &this
}

// GetLocation returns the Location field value
func (o *SpaceRelationUrls) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationUrls) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *SpaceRelationUrls) SetLocation(v string) {
	o.Location = v
}

func (o SpaceRelationUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceRelationUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

type NullableSpaceRelationUrls struct {
	value *SpaceRelationUrls
	isSet bool
}

func (v NullableSpaceRelationUrls) Get() *SpaceRelationUrls {
	return v.value
}

func (v *NullableSpaceRelationUrls) Set(val *SpaceRelationUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceRelationUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceRelationUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceRelationUrls(val *SpaceRelationUrls) *NullableSpaceRelationUrls {
	return &NullableSpaceRelationUrls{value: val, isSet: true}
}

func (v NullableSpaceRelationUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceRelationUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
