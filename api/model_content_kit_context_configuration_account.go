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

// checks if the ContentKitContextConfigurationAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitContextConfigurationAccount{}

// ContentKitContextConfigurationAccount struct for ContentKitContextConfigurationAccount
type ContentKitContextConfigurationAccount struct {
	Theme string `json:"theme"`
	Type  string `json:"type"`
	// ID of the organization the account installation configuration is in.
	OrganizationId string `json:"organizationId"`
}

// NewContentKitContextConfigurationAccount instantiates a new ContentKitContextConfigurationAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitContextConfigurationAccount(theme string, type_ string, organizationId string) *ContentKitContextConfigurationAccount {
	this := ContentKitContextConfigurationAccount{}
	this.Theme = theme
	this.Type = type_
	this.OrganizationId = organizationId
	return &this
}

// NewContentKitContextConfigurationAccountWithDefaults instantiates a new ContentKitContextConfigurationAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitContextConfigurationAccountWithDefaults() *ContentKitContextConfigurationAccount {
	this := ContentKitContextConfigurationAccount{}
	return &this
}

// GetTheme returns the Theme field value
func (o *ContentKitContextConfigurationAccount) GetTheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *ContentKitContextConfigurationAccount) GetThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *ContentKitContextConfigurationAccount) SetTheme(v string) {
	o.Theme = v
}

// GetType returns the Type field value
func (o *ContentKitContextConfigurationAccount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitContextConfigurationAccount) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitContextConfigurationAccount) SetType(v string) {
	o.Type = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *ContentKitContextConfigurationAccount) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *ContentKitContextConfigurationAccount) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *ContentKitContextConfigurationAccount) SetOrganizationId(v string) {
	o.OrganizationId = v
}

func (o ContentKitContextConfigurationAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitContextConfigurationAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["theme"] = o.Theme
	toSerialize["type"] = o.Type
	toSerialize["organizationId"] = o.OrganizationId
	return toSerialize, nil
}

type NullableContentKitContextConfigurationAccount struct {
	value *ContentKitContextConfigurationAccount
	isSet bool
}

func (v NullableContentKitContextConfigurationAccount) Get() *ContentKitContextConfigurationAccount {
	return v.value
}

func (v *NullableContentKitContextConfigurationAccount) Set(val *ContentKitContextConfigurationAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitContextConfigurationAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitContextConfigurationAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitContextConfigurationAccount(val *ContentKitContextConfigurationAccount) *NullableContentKitContextConfigurationAccount {
	return &NullableContentKitContextConfigurationAccount{value: val, isSet: true}
}

func (v NullableContentKitContextConfigurationAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitContextConfigurationAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
