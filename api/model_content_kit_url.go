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

// checks if the ContentKitURL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitURL{}

// ContentKitURL Specification for an URL in ContentKit.
type ContentKitURL struct {
	// Hostname of the URL along with the port number if required.
	Host string `json:"host"`
	// Path of the URL prefixed with a `/`.
	Pathname string                                  `json:"pathname"`
	Query    *map[string]ContentKitWebFrameDataValue `json:"query,omitempty"`
}

// NewContentKitURL instantiates a new ContentKitURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitURL(host string, pathname string) *ContentKitURL {
	this := ContentKitURL{}
	this.Host = host
	this.Pathname = pathname
	return &this
}

// NewContentKitURLWithDefaults instantiates a new ContentKitURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitURLWithDefaults() *ContentKitURL {
	this := ContentKitURL{}
	return &this
}

// GetHost returns the Host field value
func (o *ContentKitURL) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ContentKitURL) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ContentKitURL) SetHost(v string) {
	o.Host = v
}

// GetPathname returns the Pathname field value
func (o *ContentKitURL) GetPathname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pathname
}

// GetPathnameOk returns a tuple with the Pathname field value
// and a boolean to check if the value has been set.
func (o *ContentKitURL) GetPathnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pathname, true
}

// SetPathname sets field value
func (o *ContentKitURL) SetPathname(v string) {
	o.Pathname = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ContentKitURL) GetQuery() map[string]ContentKitWebFrameDataValue {
	if o == nil || IsNil(o.Query) {
		var ret map[string]ContentKitWebFrameDataValue
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitURL) GetQueryOk() (*map[string]ContentKitWebFrameDataValue, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ContentKitURL) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given map[string]ContentKitWebFrameDataValue and assigns it to the Query field.
func (o *ContentKitURL) SetQuery(v map[string]ContentKitWebFrameDataValue) {
	o.Query = &v
}

func (o ContentKitURL) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitURL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["pathname"] = o.Pathname
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableContentKitURL struct {
	value *ContentKitURL
	isSet bool
}

func (v NullableContentKitURL) Get() *ContentKitURL {
	return v.value
}

func (v *NullableContentKitURL) Set(val *ContentKitURL) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitURL) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitURL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitURL(val *ContentKitURL) *NullableContentKitURL {
	return &NullableContentKitURL{value: val, isSet: true}
}

func (v NullableContentKitURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
