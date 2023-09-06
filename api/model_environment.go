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

// checks if the Environment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Environment{}

// Environment An environment
type Environment struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	// Whether or not an environment is considered to have elevated responsibilities over other environments. Useful for distinguishing a production environment from a staging environment. Multiple primary environments are allowed. Your organization must have at least one primary environment.
	Primary bool            `json:"primary"`
	Urls    EnvironmentUrls `json:"urls"`
}

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment(id string, name string, title string, primary bool, urls EnvironmentUrls) *Environment {
	this := Environment{}
	this.Id = id
	this.Name = name
	this.Title = title
	this.Primary = primary
	this.Urls = urls
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	return &this
}

// GetId returns the Id field value
func (o *Environment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Environment) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Environment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Environment) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *Environment) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Environment) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Environment) SetTitle(v string) {
	o.Title = v
}

// GetPrimary returns the Primary field value
func (o *Environment) GetPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *Environment) GetPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *Environment) SetPrimary(v bool) {
	o.Primary = v
}

// GetUrls returns the Urls field value
func (o *Environment) GetUrls() EnvironmentUrls {
	if o == nil {
		var ret EnvironmentUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Environment) GetUrlsOk() (*EnvironmentUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Environment) SetUrls(v EnvironmentUrls) {
	o.Urls = v
}

func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Environment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["primary"] = o.Primary
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
