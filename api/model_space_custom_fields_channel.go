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

// checks if the SpaceCustomFieldsChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceCustomFieldsChannel{}

// SpaceCustomFieldsChannel struct for SpaceCustomFieldsChannel
type SpaceCustomFieldsChannel struct {
	Channel string `json:"channel"`
	Space   string `json:"space"`
}

// NewSpaceCustomFieldsChannel instantiates a new SpaceCustomFieldsChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceCustomFieldsChannel(channel string, space string) *SpaceCustomFieldsChannel {
	this := SpaceCustomFieldsChannel{}
	this.Channel = channel
	this.Space = space
	return &this
}

// NewSpaceCustomFieldsChannelWithDefaults instantiates a new SpaceCustomFieldsChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceCustomFieldsChannelWithDefaults() *SpaceCustomFieldsChannel {
	this := SpaceCustomFieldsChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *SpaceCustomFieldsChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SpaceCustomFieldsChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *SpaceCustomFieldsChannel) SetChannel(v string) {
	o.Channel = v
}

// GetSpace returns the Space field value
func (o *SpaceCustomFieldsChannel) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *SpaceCustomFieldsChannel) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *SpaceCustomFieldsChannel) SetSpace(v string) {
	o.Space = v
}

func (o SpaceCustomFieldsChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceCustomFieldsChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["space"] = o.Space
	return toSerialize, nil
}

type NullableSpaceCustomFieldsChannel struct {
	value *SpaceCustomFieldsChannel
	isSet bool
}

func (v NullableSpaceCustomFieldsChannel) Get() *SpaceCustomFieldsChannel {
	return v.value
}

func (v *NullableSpaceCustomFieldsChannel) Set(val *SpaceCustomFieldsChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceCustomFieldsChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceCustomFieldsChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceCustomFieldsChannel(val *SpaceCustomFieldsChannel) *NullableSpaceCustomFieldsChannel {
	return &NullableSpaceCustomFieldsChannel{value: val, isSet: true}
}

func (v NullableSpaceCustomFieldsChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceCustomFieldsChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
