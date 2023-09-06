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

// checks if the SpaceInfoChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceInfoChannel{}

// SpaceInfoChannel struct for SpaceInfoChannel
type SpaceInfoChannel struct {
	Channel string `json:"channel"`
	Space   string `json:"space"`
}

// NewSpaceInfoChannel instantiates a new SpaceInfoChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceInfoChannel(channel string, space string) *SpaceInfoChannel {
	this := SpaceInfoChannel{}
	this.Channel = channel
	this.Space = space
	return &this
}

// NewSpaceInfoChannelWithDefaults instantiates a new SpaceInfoChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceInfoChannelWithDefaults() *SpaceInfoChannel {
	this := SpaceInfoChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *SpaceInfoChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SpaceInfoChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *SpaceInfoChannel) SetChannel(v string) {
	o.Channel = v
}

// GetSpace returns the Space field value
func (o *SpaceInfoChannel) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *SpaceInfoChannel) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *SpaceInfoChannel) SetSpace(v string) {
	o.Space = v
}

func (o SpaceInfoChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceInfoChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["space"] = o.Space
	return toSerialize, nil
}

type NullableSpaceInfoChannel struct {
	value *SpaceInfoChannel
	isSet bool
}

func (v NullableSpaceInfoChannel) Get() *SpaceInfoChannel {
	return v.value
}

func (v *NullableSpaceInfoChannel) Set(val *SpaceInfoChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceInfoChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceInfoChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceInfoChannel(val *SpaceInfoChannel) *NullableSpaceInfoChannel {
	return &NullableSpaceInfoChannel{value: val, isSet: true}
}

func (v NullableSpaceInfoChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceInfoChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
