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

// checks if the GetContributorsByChangeRequestId200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContributorsByChangeRequestId200ResponseAllOf{}

// GetContributorsByChangeRequestId200ResponseAllOf struct for GetContributorsByChangeRequestId200ResponseAllOf
type GetContributorsByChangeRequestId200ResponseAllOf struct {
	Items []UserContributor `json:"items"`
}

// NewGetContributorsByChangeRequestId200ResponseAllOf instantiates a new GetContributorsByChangeRequestId200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContributorsByChangeRequestId200ResponseAllOf(items []UserContributor) *GetContributorsByChangeRequestId200ResponseAllOf {
	this := GetContributorsByChangeRequestId200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewGetContributorsByChangeRequestId200ResponseAllOfWithDefaults instantiates a new GetContributorsByChangeRequestId200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContributorsByChangeRequestId200ResponseAllOfWithDefaults() *GetContributorsByChangeRequestId200ResponseAllOf {
	this := GetContributorsByChangeRequestId200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *GetContributorsByChangeRequestId200ResponseAllOf) GetItems() []UserContributor {
	if o == nil {
		var ret []UserContributor
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetContributorsByChangeRequestId200ResponseAllOf) GetItemsOk() ([]UserContributor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GetContributorsByChangeRequestId200ResponseAllOf) SetItems(v []UserContributor) {
	o.Items = v
}

func (o GetContributorsByChangeRequestId200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContributorsByChangeRequestId200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableGetContributorsByChangeRequestId200ResponseAllOf struct {
	value *GetContributorsByChangeRequestId200ResponseAllOf
	isSet bool
}

func (v NullableGetContributorsByChangeRequestId200ResponseAllOf) Get() *GetContributorsByChangeRequestId200ResponseAllOf {
	return v.value
}

func (v *NullableGetContributorsByChangeRequestId200ResponseAllOf) Set(val *GetContributorsByChangeRequestId200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContributorsByChangeRequestId200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContributorsByChangeRequestId200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContributorsByChangeRequestId200ResponseAllOf(val *GetContributorsByChangeRequestId200ResponseAllOf) *NullableGetContributorsByChangeRequestId200ResponseAllOf {
	return &NullableGetContributorsByChangeRequestId200ResponseAllOf{value: val, isSet: true}
}

func (v NullableGetContributorsByChangeRequestId200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContributorsByChangeRequestId200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
