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

// checks if the ListIntegrations200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIntegrations200ResponseAllOf{}

// ListIntegrations200ResponseAllOf struct for ListIntegrations200ResponseAllOf
type ListIntegrations200ResponseAllOf struct {
	Items []Integration `json:"items"`
}

// NewListIntegrations200ResponseAllOf instantiates a new ListIntegrations200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIntegrations200ResponseAllOf(items []Integration) *ListIntegrations200ResponseAllOf {
	this := ListIntegrations200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewListIntegrations200ResponseAllOfWithDefaults instantiates a new ListIntegrations200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIntegrations200ResponseAllOfWithDefaults() *ListIntegrations200ResponseAllOf {
	this := ListIntegrations200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *ListIntegrations200ResponseAllOf) GetItems() []Integration {
	if o == nil {
		var ret []Integration
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAllOf) GetItemsOk() ([]Integration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListIntegrations200ResponseAllOf) SetItems(v []Integration) {
	o.Items = v
}

func (o ListIntegrations200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIntegrations200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListIntegrations200ResponseAllOf struct {
	value *ListIntegrations200ResponseAllOf
	isSet bool
}

func (v NullableListIntegrations200ResponseAllOf) Get() *ListIntegrations200ResponseAllOf {
	return v.value
}

func (v *NullableListIntegrations200ResponseAllOf) Set(val *ListIntegrations200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListIntegrations200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListIntegrations200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIntegrations200ResponseAllOf(val *ListIntegrations200ResponseAllOf) *NullableListIntegrations200ResponseAllOf {
	return &NullableListIntegrations200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListIntegrations200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIntegrations200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
