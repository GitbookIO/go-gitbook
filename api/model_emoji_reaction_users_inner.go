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

// checks if the EmojiReactionUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmojiReactionUsersInner{}

// EmojiReactionUsersInner struct for EmojiReactionUsersInner
type EmojiReactionUsersInner struct {
	User      User   `json:"user"`
	ReactedAt string `json:"reactedAt"`
}

// NewEmojiReactionUsersInner instantiates a new EmojiReactionUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmojiReactionUsersInner(user User, reactedAt string) *EmojiReactionUsersInner {
	this := EmojiReactionUsersInner{}
	this.User = user
	this.ReactedAt = reactedAt
	return &this
}

// NewEmojiReactionUsersInnerWithDefaults instantiates a new EmojiReactionUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmojiReactionUsersInnerWithDefaults() *EmojiReactionUsersInner {
	this := EmojiReactionUsersInner{}
	return &this
}

// GetUser returns the User field value
func (o *EmojiReactionUsersInner) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *EmojiReactionUsersInner) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *EmojiReactionUsersInner) SetUser(v User) {
	o.User = v
}

// GetReactedAt returns the ReactedAt field value
func (o *EmojiReactionUsersInner) GetReactedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReactedAt
}

// GetReactedAtOk returns a tuple with the ReactedAt field value
// and a boolean to check if the value has been set.
func (o *EmojiReactionUsersInner) GetReactedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReactedAt, true
}

// SetReactedAt sets field value
func (o *EmojiReactionUsersInner) SetReactedAt(v string) {
	o.ReactedAt = v
}

func (o EmojiReactionUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmojiReactionUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	toSerialize["reactedAt"] = o.ReactedAt
	return toSerialize, nil
}

type NullableEmojiReactionUsersInner struct {
	value *EmojiReactionUsersInner
	isSet bool
}

func (v NullableEmojiReactionUsersInner) Get() *EmojiReactionUsersInner {
	return v.value
}

func (v *NullableEmojiReactionUsersInner) Set(val *EmojiReactionUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEmojiReactionUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEmojiReactionUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmojiReactionUsersInner(val *EmojiReactionUsersInner) *NullableEmojiReactionUsersInner {
	return &NullableEmojiReactionUsersInner{value: val, isSet: true}
}

func (v NullableEmojiReactionUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmojiReactionUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
