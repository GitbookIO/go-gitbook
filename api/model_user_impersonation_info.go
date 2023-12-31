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

// checks if the UserImpersonationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImpersonationInfo{}

// UserImpersonationInfo The GitBook User impersonation info.
type UserImpersonationInfo struct {
	AuthURL        string `json:"authURL"`
	ImpersonatorId string `json:"impersonatorId"`
}

// NewUserImpersonationInfo instantiates a new UserImpersonationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImpersonationInfo(authURL string, impersonatorId string) *UserImpersonationInfo {
	this := UserImpersonationInfo{}
	this.AuthURL = authURL
	this.ImpersonatorId = impersonatorId
	return &this
}

// NewUserImpersonationInfoWithDefaults instantiates a new UserImpersonationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImpersonationInfoWithDefaults() *UserImpersonationInfo {
	this := UserImpersonationInfo{}
	return &this
}

// GetAuthURL returns the AuthURL field value
func (o *UserImpersonationInfo) GetAuthURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthURL
}

// GetAuthURLOk returns a tuple with the AuthURL field value
// and a boolean to check if the value has been set.
func (o *UserImpersonationInfo) GetAuthURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthURL, true
}

// SetAuthURL sets field value
func (o *UserImpersonationInfo) SetAuthURL(v string) {
	o.AuthURL = v
}

// GetImpersonatorId returns the ImpersonatorId field value
func (o *UserImpersonationInfo) GetImpersonatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImpersonatorId
}

// GetImpersonatorIdOk returns a tuple with the ImpersonatorId field value
// and a boolean to check if the value has been set.
func (o *UserImpersonationInfo) GetImpersonatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpersonatorId, true
}

// SetImpersonatorId sets field value
func (o *UserImpersonationInfo) SetImpersonatorId(v string) {
	o.ImpersonatorId = v
}

func (o UserImpersonationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImpersonationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authURL"] = o.AuthURL
	toSerialize["impersonatorId"] = o.ImpersonatorId
	return toSerialize, nil
}

type NullableUserImpersonationInfo struct {
	value *UserImpersonationInfo
	isSet bool
}

func (v NullableUserImpersonationInfo) Get() *UserImpersonationInfo {
	return v.value
}

func (v *NullableUserImpersonationInfo) Set(val *UserImpersonationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImpersonationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImpersonationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImpersonationInfo(val *UserImpersonationInfo) *NullableUserImpersonationInfo {
	return &NullableUserImpersonationInfo{value: val, isSet: true}
}

func (v NullableUserImpersonationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImpersonationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
