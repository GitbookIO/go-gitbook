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

// checks if the SetupDirectorySync200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetupDirectorySync200Response{}

// SetupDirectorySync200Response struct for SetupDirectorySync200Response
type SetupDirectorySync200Response struct {
	// The URL to navigate to to continue Directory Sync setup.
	SetupUrl string `json:"setupUrl"`
}

// NewSetupDirectorySync200Response instantiates a new SetupDirectorySync200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetupDirectorySync200Response(setupUrl string) *SetupDirectorySync200Response {
	this := SetupDirectorySync200Response{}
	this.SetupUrl = setupUrl
	return &this
}

// NewSetupDirectorySync200ResponseWithDefaults instantiates a new SetupDirectorySync200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetupDirectorySync200ResponseWithDefaults() *SetupDirectorySync200Response {
	this := SetupDirectorySync200Response{}
	return &this
}

// GetSetupUrl returns the SetupUrl field value
func (o *SetupDirectorySync200Response) GetSetupUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SetupUrl
}

// GetSetupUrlOk returns a tuple with the SetupUrl field value
// and a boolean to check if the value has been set.
func (o *SetupDirectorySync200Response) GetSetupUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetupUrl, true
}

// SetSetupUrl sets field value
func (o *SetupDirectorySync200Response) SetSetupUrl(v string) {
	o.SetupUrl = v
}

func (o SetupDirectorySync200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetupDirectorySync200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["setupUrl"] = o.SetupUrl
	return toSerialize, nil
}

type NullableSetupDirectorySync200Response struct {
	value *SetupDirectorySync200Response
	isSet bool
}

func (v NullableSetupDirectorySync200Response) Get() *SetupDirectorySync200Response {
	return v.value
}

func (v *NullableSetupDirectorySync200Response) Set(val *SetupDirectorySync200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetupDirectorySync200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetupDirectorySync200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetupDirectorySync200Response(val *SetupDirectorySync200Response) *NullableSetupDirectorySync200Response {
	return &NullableSetupDirectorySync200Response{value: val, isSet: true}
}

func (v NullableSetupDirectorySync200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetupDirectorySync200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
