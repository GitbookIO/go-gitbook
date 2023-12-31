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

// checks if the EntitySchemaAllOfUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitySchemaAllOfUrls{}

// EntitySchemaAllOfUrls URLs associated with the object
type EntitySchemaAllOfUrls struct {
	// URL of the entity schema in the API
	Location string `json:"location"`
}

// NewEntitySchemaAllOfUrls instantiates a new EntitySchemaAllOfUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySchemaAllOfUrls(location string) *EntitySchemaAllOfUrls {
	this := EntitySchemaAllOfUrls{}
	this.Location = location
	return &this
}

// NewEntitySchemaAllOfUrlsWithDefaults instantiates a new EntitySchemaAllOfUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySchemaAllOfUrlsWithDefaults() *EntitySchemaAllOfUrls {
	this := EntitySchemaAllOfUrls{}
	return &this
}

// GetLocation returns the Location field value
func (o *EntitySchemaAllOfUrls) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOfUrls) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *EntitySchemaAllOfUrls) SetLocation(v string) {
	o.Location = v
}

func (o EntitySchemaAllOfUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitySchemaAllOfUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

type NullableEntitySchemaAllOfUrls struct {
	value *EntitySchemaAllOfUrls
	isSet bool
}

func (v NullableEntitySchemaAllOfUrls) Get() *EntitySchemaAllOfUrls {
	return v.value
}

func (v *NullableEntitySchemaAllOfUrls) Set(val *EntitySchemaAllOfUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySchemaAllOfUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySchemaAllOfUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySchemaAllOfUrls(val *EntitySchemaAllOfUrls) *NullableEntitySchemaAllOfUrls {
	return &NullableEntitySchemaAllOfUrls{value: val, isSet: true}
}

func (v NullableEntitySchemaAllOfUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySchemaAllOfUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
