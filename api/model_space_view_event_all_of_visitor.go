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

// checks if the SpaceViewEventAllOfVisitor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceViewEventAllOfVisitor{}

// SpaceViewEventAllOfVisitor Analytics info on the GitBook's content visitor.
type SpaceViewEventAllOfVisitor struct {
	// GitBook's unique identifier of the visitor.
	AnonymousId string `json:"anonymousId"`
	// The visitors cookies.
	Cookies map[string]string `json:"cookies"`
	// User-agent of the visitor.
	UserAgent string `json:"userAgent"`
	// IP address of the visitor.
	Ip string `json:"ip"`
}

// NewSpaceViewEventAllOfVisitor instantiates a new SpaceViewEventAllOfVisitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceViewEventAllOfVisitor(anonymousId string, cookies map[string]string, userAgent string, ip string) *SpaceViewEventAllOfVisitor {
	this := SpaceViewEventAllOfVisitor{}
	this.AnonymousId = anonymousId
	this.Cookies = cookies
	this.UserAgent = userAgent
	this.Ip = ip
	return &this
}

// NewSpaceViewEventAllOfVisitorWithDefaults instantiates a new SpaceViewEventAllOfVisitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceViewEventAllOfVisitorWithDefaults() *SpaceViewEventAllOfVisitor {
	this := SpaceViewEventAllOfVisitor{}
	return &this
}

// GetAnonymousId returns the AnonymousId field value
func (o *SpaceViewEventAllOfVisitor) GetAnonymousId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnonymousId
}

// GetAnonymousIdOk returns a tuple with the AnonymousId field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOfVisitor) GetAnonymousIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnonymousId, true
}

// SetAnonymousId sets field value
func (o *SpaceViewEventAllOfVisitor) SetAnonymousId(v string) {
	o.AnonymousId = v
}

// GetCookies returns the Cookies field value
func (o *SpaceViewEventAllOfVisitor) GetCookies() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Cookies
}

// GetCookiesOk returns a tuple with the Cookies field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOfVisitor) GetCookiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cookies, true
}

// SetCookies sets field value
func (o *SpaceViewEventAllOfVisitor) SetCookies(v map[string]string) {
	o.Cookies = v
}

// GetUserAgent returns the UserAgent field value
func (o *SpaceViewEventAllOfVisitor) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOfVisitor) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *SpaceViewEventAllOfVisitor) SetUserAgent(v string) {
	o.UserAgent = v
}

// GetIp returns the Ip field value
func (o *SpaceViewEventAllOfVisitor) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOfVisitor) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *SpaceViewEventAllOfVisitor) SetIp(v string) {
	o.Ip = v
}

func (o SpaceViewEventAllOfVisitor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceViewEventAllOfVisitor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["anonymousId"] = o.AnonymousId
	toSerialize["cookies"] = o.Cookies
	toSerialize["userAgent"] = o.UserAgent
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

type NullableSpaceViewEventAllOfVisitor struct {
	value *SpaceViewEventAllOfVisitor
	isSet bool
}

func (v NullableSpaceViewEventAllOfVisitor) Get() *SpaceViewEventAllOfVisitor {
	return v.value
}

func (v *NullableSpaceViewEventAllOfVisitor) Set(val *SpaceViewEventAllOfVisitor) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceViewEventAllOfVisitor) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceViewEventAllOfVisitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceViewEventAllOfVisitor(val *SpaceViewEventAllOfVisitor) *NullableSpaceViewEventAllOfVisitor {
	return &NullableSpaceViewEventAllOfVisitor{value: val, isSet: true}
}

func (v NullableSpaceViewEventAllOfVisitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceViewEventAllOfVisitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
