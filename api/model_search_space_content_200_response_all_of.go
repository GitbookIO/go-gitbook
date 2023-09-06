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

// checks if the SearchSpaceContent200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSpaceContent200ResponseAllOf{}

// SearchSpaceContent200ResponseAllOf struct for SearchSpaceContent200ResponseAllOf
type SearchSpaceContent200ResponseAllOf struct {
	Items []SearchPageResult `json:"items"`
}

// NewSearchSpaceContent200ResponseAllOf instantiates a new SearchSpaceContent200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSpaceContent200ResponseAllOf(items []SearchPageResult) *SearchSpaceContent200ResponseAllOf {
	this := SearchSpaceContent200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewSearchSpaceContent200ResponseAllOfWithDefaults instantiates a new SearchSpaceContent200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSpaceContent200ResponseAllOfWithDefaults() *SearchSpaceContent200ResponseAllOf {
	this := SearchSpaceContent200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *SearchSpaceContent200ResponseAllOf) GetItems() []SearchPageResult {
	if o == nil {
		var ret []SearchPageResult
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SearchSpaceContent200ResponseAllOf) GetItemsOk() ([]SearchPageResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SearchSpaceContent200ResponseAllOf) SetItems(v []SearchPageResult) {
	o.Items = v
}

func (o SearchSpaceContent200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSpaceContent200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSearchSpaceContent200ResponseAllOf struct {
	value *SearchSpaceContent200ResponseAllOf
	isSet bool
}

func (v NullableSearchSpaceContent200ResponseAllOf) Get() *SearchSpaceContent200ResponseAllOf {
	return v.value
}

func (v *NullableSearchSpaceContent200ResponseAllOf) Set(val *SearchSpaceContent200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSpaceContent200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSpaceContent200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSpaceContent200ResponseAllOf(val *SearchSpaceContent200ResponseAllOf) *NullableSearchSpaceContent200ResponseAllOf {
	return &NullableSearchSpaceContent200ResponseAllOf{value: val, isSet: true}
}

func (v NullableSearchSpaceContent200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSpaceContent200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
