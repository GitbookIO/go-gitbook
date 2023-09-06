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

// checks if the BackofficeUserInfoChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackofficeUserInfoChannel{}

// BackofficeUserInfoChannel struct for BackofficeUserInfoChannel
type BackofficeUserInfoChannel struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// NewBackofficeUserInfoChannel instantiates a new BackofficeUserInfoChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackofficeUserInfoChannel(channel string, user string) *BackofficeUserInfoChannel {
	this := BackofficeUserInfoChannel{}
	this.Channel = channel
	this.User = user
	return &this
}

// NewBackofficeUserInfoChannelWithDefaults instantiates a new BackofficeUserInfoChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackofficeUserInfoChannelWithDefaults() *BackofficeUserInfoChannel {
	this := BackofficeUserInfoChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *BackofficeUserInfoChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *BackofficeUserInfoChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *BackofficeUserInfoChannel) SetChannel(v string) {
	o.Channel = v
}

// GetUser returns the User field value
func (o *BackofficeUserInfoChannel) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *BackofficeUserInfoChannel) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *BackofficeUserInfoChannel) SetUser(v string) {
	o.User = v
}

func (o BackofficeUserInfoChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackofficeUserInfoChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableBackofficeUserInfoChannel struct {
	value *BackofficeUserInfoChannel
	isSet bool
}

func (v NullableBackofficeUserInfoChannel) Get() *BackofficeUserInfoChannel {
	return v.value
}

func (v *NullableBackofficeUserInfoChannel) Set(val *BackofficeUserInfoChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackofficeUserInfoChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackofficeUserInfoChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackofficeUserInfoChannel(val *BackofficeUserInfoChannel) *NullableBackofficeUserInfoChannel {
	return &NullableBackofficeUserInfoChannel{value: val, isSet: true}
}

func (v NullableBackofficeUserInfoChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackofficeUserInfoChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
