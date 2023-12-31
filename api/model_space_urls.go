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

// checks if the SpaceUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceUrls{}

// SpaceUrls URLs associated with the object
type SpaceUrls struct {
	// URL of the space in the API
	Location string `json:"location"`
	// URL of the space in the application
	App string `json:"app"`
	// URL of the published version of the space. Only defined when visibility is not \"private.\"
	Published *string `json:"published,omitempty"`
	// URL of the public version of the space. Only defined when visibility is \"public\".
	Public *string `json:"public,omitempty"`
}

// NewSpaceUrls instantiates a new SpaceUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceUrls(location string, app string) *SpaceUrls {
	this := SpaceUrls{}
	this.Location = location
	this.App = app
	return &this
}

// NewSpaceUrlsWithDefaults instantiates a new SpaceUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceUrlsWithDefaults() *SpaceUrls {
	this := SpaceUrls{}
	return &this
}

// GetLocation returns the Location field value
func (o *SpaceUrls) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *SpaceUrls) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *SpaceUrls) SetLocation(v string) {
	o.Location = v
}

// GetApp returns the App field value
func (o *SpaceUrls) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *SpaceUrls) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *SpaceUrls) SetApp(v string) {
	o.App = v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *SpaceUrls) GetPublished() string {
	if o == nil || IsNil(o.Published) {
		var ret string
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceUrls) GetPublishedOk() (*string, bool) {
	if o == nil || IsNil(o.Published) {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *SpaceUrls) HasPublished() bool {
	if o != nil && !IsNil(o.Published) {
		return true
	}

	return false
}

// SetPublished gets a reference to the given string and assigns it to the Published field.
func (o *SpaceUrls) SetPublished(v string) {
	o.Published = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *SpaceUrls) GetPublic() string {
	if o == nil || IsNil(o.Public) {
		var ret string
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceUrls) GetPublicOk() (*string, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *SpaceUrls) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given string and assigns it to the Public field.
func (o *SpaceUrls) SetPublic(v string) {
	o.Public = &v
}

func (o SpaceUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	toSerialize["app"] = o.App
	if !IsNil(o.Published) {
		toSerialize["published"] = o.Published
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	return toSerialize, nil
}

type NullableSpaceUrls struct {
	value *SpaceUrls
	isSet bool
}

func (v NullableSpaceUrls) Get() *SpaceUrls {
	return v.value
}

func (v *NullableSpaceUrls) Set(val *SpaceUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceUrls(val *SpaceUrls) *NullableSpaceUrls {
	return &NullableSpaceUrls{value: val, isSet: true}
}

func (v NullableSpaceUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
