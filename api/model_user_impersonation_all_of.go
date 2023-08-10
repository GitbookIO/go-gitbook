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

// checks if the UserImpersonationAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImpersonationAllOf{}

// UserImpersonationAllOf struct for UserImpersonationAllOf
type UserImpersonationAllOf struct {
	Impersonation UserImpersonationInfo `json:"impersonation"`
}

// NewUserImpersonationAllOf instantiates a new UserImpersonationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImpersonationAllOf(impersonation UserImpersonationInfo) *UserImpersonationAllOf {
	this := UserImpersonationAllOf{}
	this.Impersonation = impersonation
	return &this
}

// NewUserImpersonationAllOfWithDefaults instantiates a new UserImpersonationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImpersonationAllOfWithDefaults() *UserImpersonationAllOf {
	this := UserImpersonationAllOf{}
	return &this
}

// GetImpersonation returns the Impersonation field value
func (o *UserImpersonationAllOf) GetImpersonation() UserImpersonationInfo {
	if o == nil {
		var ret UserImpersonationInfo
		return ret
	}

	return o.Impersonation
}

// GetImpersonationOk returns a tuple with the Impersonation field value
// and a boolean to check if the value has been set.
func (o *UserImpersonationAllOf) GetImpersonationOk() (*UserImpersonationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impersonation, true
}

// SetImpersonation sets field value
func (o *UserImpersonationAllOf) SetImpersonation(v UserImpersonationInfo) {
	o.Impersonation = v
}

func (o UserImpersonationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImpersonationAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["impersonation"] = o.Impersonation
	return toSerialize, nil
}

type NullableUserImpersonationAllOf struct {
	value *UserImpersonationAllOf
	isSet bool
}

func (v NullableUserImpersonationAllOf) Get() *UserImpersonationAllOf {
	return v.value
}

func (v *NullableUserImpersonationAllOf) Set(val *UserImpersonationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImpersonationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImpersonationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImpersonationAllOf(val *UserImpersonationAllOf) *NullableUserImpersonationAllOf {
	return &NullableUserImpersonationAllOf{value: val, isSet: true}
}

func (v NullableUserImpersonationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImpersonationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
