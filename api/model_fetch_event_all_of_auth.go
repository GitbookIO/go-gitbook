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

// checks if the FetchEventAllOfAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchEventAllOfAuth{}

// FetchEventAllOfAuth struct for FetchEventAllOfAuth
type FetchEventAllOfAuth struct {
	// The user's ID.
	UserId string `json:"userId"`
}

// NewFetchEventAllOfAuth instantiates a new FetchEventAllOfAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchEventAllOfAuth(userId string) *FetchEventAllOfAuth {
	this := FetchEventAllOfAuth{}
	this.UserId = userId
	return &this
}

// NewFetchEventAllOfAuthWithDefaults instantiates a new FetchEventAllOfAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchEventAllOfAuthWithDefaults() *FetchEventAllOfAuth {
	this := FetchEventAllOfAuth{}
	return &this
}

// GetUserId returns the UserId field value
func (o *FetchEventAllOfAuth) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FetchEventAllOfAuth) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FetchEventAllOfAuth) SetUserId(v string) {
	o.UserId = v
}

func (o FetchEventAllOfAuth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchEventAllOfAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

type NullableFetchEventAllOfAuth struct {
	value *FetchEventAllOfAuth
	isSet bool
}

func (v NullableFetchEventAllOfAuth) Get() *FetchEventAllOfAuth {
	return v.value
}

func (v *NullableFetchEventAllOfAuth) Set(val *FetchEventAllOfAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchEventAllOfAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchEventAllOfAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchEventAllOfAuth(val *FetchEventAllOfAuth) *NullableFetchEventAllOfAuth {
	return &NullableFetchEventAllOfAuth{value: val, isSet: true}
}

func (v NullableFetchEventAllOfAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchEventAllOfAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
