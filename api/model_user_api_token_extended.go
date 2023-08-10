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

// checks if the UserAPITokenExtended type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAPITokenExtended{}

// UserAPITokenExtended The API token details, including the token itself.
type UserAPITokenExtended struct {
	// The API token ID.
	Id string `json:"id"`
	// The API token name.
	Label     string `json:"label"`
	CreatedAt string `json:"createdAt"`
	// The actual token value.
	Token string `json:"token"`
}

// NewUserAPITokenExtended instantiates a new UserAPITokenExtended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAPITokenExtended(id string, label string, createdAt string, token string) *UserAPITokenExtended {
	this := UserAPITokenExtended{}
	this.Id = id
	this.Label = label
	this.CreatedAt = createdAt
	this.Token = token
	return &this
}

// NewUserAPITokenExtendedWithDefaults instantiates a new UserAPITokenExtended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAPITokenExtendedWithDefaults() *UserAPITokenExtended {
	this := UserAPITokenExtended{}
	return &this
}

// GetId returns the Id field value
func (o *UserAPITokenExtended) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserAPITokenExtended) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserAPITokenExtended) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *UserAPITokenExtended) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *UserAPITokenExtended) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *UserAPITokenExtended) SetLabel(v string) {
	o.Label = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserAPITokenExtended) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserAPITokenExtended) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserAPITokenExtended) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetToken returns the Token field value
func (o *UserAPITokenExtended) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserAPITokenExtended) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserAPITokenExtended) SetToken(v string) {
	o.Token = v
}

func (o UserAPITokenExtended) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAPITokenExtended) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableUserAPITokenExtended struct {
	value *UserAPITokenExtended
	isSet bool
}

func (v NullableUserAPITokenExtended) Get() *UserAPITokenExtended {
	return v.value
}

func (v *NullableUserAPITokenExtended) Set(val *UserAPITokenExtended) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAPITokenExtended) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAPITokenExtended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAPITokenExtended(val *UserAPITokenExtended) *NullableUserAPITokenExtended {
	return &NullableUserAPITokenExtended{value: val, isSet: true}
}

func (v NullableUserAPITokenExtended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAPITokenExtended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
