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

// checks if the UserAPITokensChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAPITokensChannel{}

// UserAPITokensChannel struct for UserAPITokensChannel
type UserAPITokensChannel struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// NewUserAPITokensChannel instantiates a new UserAPITokensChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAPITokensChannel(channel string, user string) *UserAPITokensChannel {
	this := UserAPITokensChannel{}
	this.Channel = channel
	this.User = user
	return &this
}

// NewUserAPITokensChannelWithDefaults instantiates a new UserAPITokensChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAPITokensChannelWithDefaults() *UserAPITokensChannel {
	this := UserAPITokensChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *UserAPITokensChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *UserAPITokensChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *UserAPITokensChannel) SetChannel(v string) {
	o.Channel = v
}

// GetUser returns the User field value
func (o *UserAPITokensChannel) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserAPITokensChannel) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserAPITokensChannel) SetUser(v string) {
	o.User = v
}

func (o UserAPITokensChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAPITokensChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableUserAPITokensChannel struct {
	value *UserAPITokensChannel
	isSet bool
}

func (v NullableUserAPITokensChannel) Get() *UserAPITokensChannel {
	return v.value
}

func (v *NullableUserAPITokensChannel) Set(val *UserAPITokensChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAPITokensChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAPITokensChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAPITokensChannel(val *UserAPITokensChannel) *NullableUserAPITokensChannel {
	return &NullableUserAPITokensChannel{value: val, isSet: true}
}

func (v NullableUserAPITokensChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAPITokensChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
