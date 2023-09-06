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

// checks if the ChangeRequestReviewsChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestReviewsChannel{}

// ChangeRequestReviewsChannel struct for ChangeRequestReviewsChannel
type ChangeRequestReviewsChannel struct {
	Channel       string `json:"channel"`
	Space         string `json:"space"`
	ChangeRequest string `json:"changeRequest"`
}

// NewChangeRequestReviewsChannel instantiates a new ChangeRequestReviewsChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestReviewsChannel(channel string, space string, changeRequest string) *ChangeRequestReviewsChannel {
	this := ChangeRequestReviewsChannel{}
	this.Channel = channel
	this.Space = space
	this.ChangeRequest = changeRequest
	return &this
}

// NewChangeRequestReviewsChannelWithDefaults instantiates a new ChangeRequestReviewsChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestReviewsChannelWithDefaults() *ChangeRequestReviewsChannel {
	this := ChangeRequestReviewsChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *ChangeRequestReviewsChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestReviewsChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *ChangeRequestReviewsChannel) SetChannel(v string) {
	o.Channel = v
}

// GetSpace returns the Space field value
func (o *ChangeRequestReviewsChannel) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestReviewsChannel) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *ChangeRequestReviewsChannel) SetSpace(v string) {
	o.Space = v
}

// GetChangeRequest returns the ChangeRequest field value
func (o *ChangeRequestReviewsChannel) GetChangeRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangeRequest
}

// GetChangeRequestOk returns a tuple with the ChangeRequest field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestReviewsChannel) GetChangeRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequest, true
}

// SetChangeRequest sets field value
func (o *ChangeRequestReviewsChannel) SetChangeRequest(v string) {
	o.ChangeRequest = v
}

func (o ChangeRequestReviewsChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestReviewsChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["space"] = o.Space
	toSerialize["changeRequest"] = o.ChangeRequest
	return toSerialize, nil
}

type NullableChangeRequestReviewsChannel struct {
	value *ChangeRequestReviewsChannel
	isSet bool
}

func (v NullableChangeRequestReviewsChannel) Get() *ChangeRequestReviewsChannel {
	return v.value
}

func (v *NullableChangeRequestReviewsChannel) Set(val *ChangeRequestReviewsChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestReviewsChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestReviewsChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestReviewsChannel(val *ChangeRequestReviewsChannel) *NullableChangeRequestReviewsChannel {
	return &NullableChangeRequestReviewsChannel{value: val, isSet: true}
}

func (v NullableChangeRequestReviewsChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestReviewsChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
