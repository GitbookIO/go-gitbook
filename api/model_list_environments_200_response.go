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

// checks if the ListEnvironments200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEnvironments200Response{}

// ListEnvironments200Response struct for ListEnvironments200Response
type ListEnvironments200Response struct {
	Next *ListNext `json:"next,omitempty"`
	// Total count of objects in the list
	Count *float32      `json:"count,omitempty"`
	Items []Environment `json:"items,omitempty"`
}

// NewListEnvironments200Response instantiates a new ListEnvironments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEnvironments200Response() *ListEnvironments200Response {
	this := ListEnvironments200Response{}
	return &this
}

// NewListEnvironments200ResponseWithDefaults instantiates a new ListEnvironments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEnvironments200ResponseWithDefaults() *ListEnvironments200Response {
	this := ListEnvironments200Response{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListEnvironments200Response) GetNext() ListNext {
	if o == nil || IsNil(o.Next) {
		var ret ListNext
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnvironments200Response) GetNextOk() (*ListNext, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListEnvironments200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListNext and assigns it to the Next field.
func (o *ListEnvironments200Response) SetNext(v ListNext) {
	o.Next = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListEnvironments200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnvironments200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListEnvironments200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ListEnvironments200Response) SetCount(v float32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListEnvironments200Response) GetItems() []Environment {
	if o == nil || IsNil(o.Items) {
		var ret []Environment
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnvironments200Response) GetItemsOk() ([]Environment, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListEnvironments200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Environment and assigns it to the Items field.
func (o *ListEnvironments200Response) SetItems(v []Environment) {
	o.Items = v
}

func (o ListEnvironments200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEnvironments200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListEnvironments200Response struct {
	value *ListEnvironments200Response
	isSet bool
}

func (v NullableListEnvironments200Response) Get() *ListEnvironments200Response {
	return v.value
}

func (v *NullableListEnvironments200Response) Set(val *ListEnvironments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListEnvironments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListEnvironments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEnvironments200Response(val *ListEnvironments200Response) *NullableListEnvironments200Response {
	return &NullableListEnvironments200Response{value: val, isSet: true}
}

func (v NullableListEnvironments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEnvironments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
