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

// checks if the UnsplashImageUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnsplashImageUrls{}

// UnsplashImageUrls struct for UnsplashImageUrls
type UnsplashImageUrls struct {
	Full  string `json:"full"`
	Small string `json:"small"`
}

// NewUnsplashImageUrls instantiates a new UnsplashImageUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsplashImageUrls(full string, small string) *UnsplashImageUrls {
	this := UnsplashImageUrls{}
	this.Full = full
	this.Small = small
	return &this
}

// NewUnsplashImageUrlsWithDefaults instantiates a new UnsplashImageUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsplashImageUrlsWithDefaults() *UnsplashImageUrls {
	this := UnsplashImageUrls{}
	return &this
}

// GetFull returns the Full field value
func (o *UnsplashImageUrls) GetFull() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Full
}

// GetFullOk returns a tuple with the Full field value
// and a boolean to check if the value has been set.
func (o *UnsplashImageUrls) GetFullOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Full, true
}

// SetFull sets field value
func (o *UnsplashImageUrls) SetFull(v string) {
	o.Full = v
}

// GetSmall returns the Small field value
func (o *UnsplashImageUrls) GetSmall() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Small
}

// GetSmallOk returns a tuple with the Small field value
// and a boolean to check if the value has been set.
func (o *UnsplashImageUrls) GetSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Small, true
}

// SetSmall sets field value
func (o *UnsplashImageUrls) SetSmall(v string) {
	o.Small = v
}

func (o UnsplashImageUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnsplashImageUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["full"] = o.Full
	toSerialize["small"] = o.Small
	return toSerialize, nil
}

type NullableUnsplashImageUrls struct {
	value *UnsplashImageUrls
	isSet bool
}

func (v NullableUnsplashImageUrls) Get() *UnsplashImageUrls {
	return v.value
}

func (v *NullableUnsplashImageUrls) Set(val *UnsplashImageUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsplashImageUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsplashImageUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsplashImageUrls(val *UnsplashImageUrls) *NullableUnsplashImageUrls {
	return &NullableUnsplashImageUrls{value: val, isSet: true}
}

func (v NullableUnsplashImageUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsplashImageUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
