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

// checks if the GetReviewsByChangeRequestId200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReviewsByChangeRequestId200ResponseAllOf{}

// GetReviewsByChangeRequestId200ResponseAllOf struct for GetReviewsByChangeRequestId200ResponseAllOf
type GetReviewsByChangeRequestId200ResponseAllOf struct {
	Items []ChangeRequestReview `json:"items"`
}

// NewGetReviewsByChangeRequestId200ResponseAllOf instantiates a new GetReviewsByChangeRequestId200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReviewsByChangeRequestId200ResponseAllOf(items []ChangeRequestReview) *GetReviewsByChangeRequestId200ResponseAllOf {
	this := GetReviewsByChangeRequestId200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewGetReviewsByChangeRequestId200ResponseAllOfWithDefaults instantiates a new GetReviewsByChangeRequestId200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReviewsByChangeRequestId200ResponseAllOfWithDefaults() *GetReviewsByChangeRequestId200ResponseAllOf {
	this := GetReviewsByChangeRequestId200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *GetReviewsByChangeRequestId200ResponseAllOf) GetItems() []ChangeRequestReview {
	if o == nil {
		var ret []ChangeRequestReview
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetReviewsByChangeRequestId200ResponseAllOf) GetItemsOk() ([]ChangeRequestReview, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GetReviewsByChangeRequestId200ResponseAllOf) SetItems(v []ChangeRequestReview) {
	o.Items = v
}

func (o GetReviewsByChangeRequestId200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReviewsByChangeRequestId200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableGetReviewsByChangeRequestId200ResponseAllOf struct {
	value *GetReviewsByChangeRequestId200ResponseAllOf
	isSet bool
}

func (v NullableGetReviewsByChangeRequestId200ResponseAllOf) Get() *GetReviewsByChangeRequestId200ResponseAllOf {
	return v.value
}

func (v *NullableGetReviewsByChangeRequestId200ResponseAllOf) Set(val *GetReviewsByChangeRequestId200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReviewsByChangeRequestId200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReviewsByChangeRequestId200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReviewsByChangeRequestId200ResponseAllOf(val *GetReviewsByChangeRequestId200ResponseAllOf) *NullableGetReviewsByChangeRequestId200ResponseAllOf {
	return &NullableGetReviewsByChangeRequestId200ResponseAllOf{value: val, isSet: true}
}

func (v NullableGetReviewsByChangeRequestId200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReviewsByChangeRequestId200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
