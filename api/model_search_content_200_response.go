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

// checks if the SearchContent200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchContent200Response{}

// SearchContent200Response struct for SearchContent200Response
type SearchContent200Response struct {
	Next *ListNext `json:"next,omitempty"`
	// Total count of objects in the list
	Count *float32            `json:"count,omitempty"`
	Items []SearchSpaceResult `json:"items"`
}

// NewSearchContent200Response instantiates a new SearchContent200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchContent200Response(items []SearchSpaceResult) *SearchContent200Response {
	this := SearchContent200Response{}
	this.Items = items
	return &this
}

// NewSearchContent200ResponseWithDefaults instantiates a new SearchContent200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchContent200ResponseWithDefaults() *SearchContent200Response {
	this := SearchContent200Response{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SearchContent200Response) GetNext() ListNext {
	if o == nil || IsNil(o.Next) {
		var ret ListNext
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchContent200Response) GetNextOk() (*ListNext, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SearchContent200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListNext and assigns it to the Next field.
func (o *SearchContent200Response) SetNext(v ListNext) {
	o.Next = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchContent200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchContent200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchContent200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *SearchContent200Response) SetCount(v float32) {
	o.Count = &v
}

// GetItems returns the Items field value
func (o *SearchContent200Response) GetItems() []SearchSpaceResult {
	if o == nil {
		var ret []SearchSpaceResult
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SearchContent200Response) GetItemsOk() ([]SearchSpaceResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SearchContent200Response) SetItems(v []SearchSpaceResult) {
	o.Items = v
}

func (o SearchContent200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchContent200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSearchContent200Response struct {
	value *SearchContent200Response
	isSet bool
}

func (v NullableSearchContent200Response) Get() *SearchContent200Response {
	return v.value
}

func (v *NullableSearchContent200Response) Set(val *SearchContent200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchContent200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchContent200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchContent200Response(val *SearchContent200Response) *NullableSearchContent200Response {
	return &NullableSearchContent200Response{value: val, isSet: true}
}

func (v NullableSearchContent200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchContent200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
