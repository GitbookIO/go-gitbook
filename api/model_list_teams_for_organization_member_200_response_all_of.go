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

// checks if the ListTeamsForOrganizationMember200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTeamsForOrganizationMember200ResponseAllOf{}

// ListTeamsForOrganizationMember200ResponseAllOf struct for ListTeamsForOrganizationMember200ResponseAllOf
type ListTeamsForOrganizationMember200ResponseAllOf struct {
	Items []ListTeamsForOrganizationMember200ResponseAllOfItemsInner `json:"items"`
}

// NewListTeamsForOrganizationMember200ResponseAllOf instantiates a new ListTeamsForOrganizationMember200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTeamsForOrganizationMember200ResponseAllOf(items []ListTeamsForOrganizationMember200ResponseAllOfItemsInner) *ListTeamsForOrganizationMember200ResponseAllOf {
	this := ListTeamsForOrganizationMember200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewListTeamsForOrganizationMember200ResponseAllOfWithDefaults instantiates a new ListTeamsForOrganizationMember200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTeamsForOrganizationMember200ResponseAllOfWithDefaults() *ListTeamsForOrganizationMember200ResponseAllOf {
	this := ListTeamsForOrganizationMember200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *ListTeamsForOrganizationMember200ResponseAllOf) GetItems() []ListTeamsForOrganizationMember200ResponseAllOfItemsInner {
	if o == nil {
		var ret []ListTeamsForOrganizationMember200ResponseAllOfItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListTeamsForOrganizationMember200ResponseAllOf) GetItemsOk() ([]ListTeamsForOrganizationMember200ResponseAllOfItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListTeamsForOrganizationMember200ResponseAllOf) SetItems(v []ListTeamsForOrganizationMember200ResponseAllOfItemsInner) {
	o.Items = v
}

func (o ListTeamsForOrganizationMember200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTeamsForOrganizationMember200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListTeamsForOrganizationMember200ResponseAllOf struct {
	value *ListTeamsForOrganizationMember200ResponseAllOf
	isSet bool
}

func (v NullableListTeamsForOrganizationMember200ResponseAllOf) Get() *ListTeamsForOrganizationMember200ResponseAllOf {
	return v.value
}

func (v *NullableListTeamsForOrganizationMember200ResponseAllOf) Set(val *ListTeamsForOrganizationMember200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListTeamsForOrganizationMember200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListTeamsForOrganizationMember200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTeamsForOrganizationMember200ResponseAllOf(val *ListTeamsForOrganizationMember200ResponseAllOf) *NullableListTeamsForOrganizationMember200ResponseAllOf {
	return &NullableListTeamsForOrganizationMember200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListTeamsForOrganizationMember200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTeamsForOrganizationMember200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
