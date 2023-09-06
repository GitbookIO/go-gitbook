/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the OrganizationSchemasChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSchemasChannel{}

// OrganizationSchemasChannel Subscription channel for all entity schemas in an organization.
type OrganizationSchemasChannel struct {
	Channel      string `json:"channel"`
	Organization string `json:"organization"`
}

// NewOrganizationSchemasChannel instantiates a new OrganizationSchemasChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSchemasChannel(channel string, organization string) *OrganizationSchemasChannel {
	this := OrganizationSchemasChannel{}
	this.Channel = channel
	this.Organization = organization
	return &this
}

// NewOrganizationSchemasChannelWithDefaults instantiates a new OrganizationSchemasChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSchemasChannelWithDefaults() *OrganizationSchemasChannel {
	this := OrganizationSchemasChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *OrganizationSchemasChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *OrganizationSchemasChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *OrganizationSchemasChannel) SetChannel(v string) {
	o.Channel = v
}

// GetOrganization returns the Organization field value
func (o *OrganizationSchemasChannel) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *OrganizationSchemasChannel) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *OrganizationSchemasChannel) SetOrganization(v string) {
	o.Organization = v
}

func (o OrganizationSchemasChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSchemasChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["organization"] = o.Organization
	return toSerialize, nil
}

type NullableOrganizationSchemasChannel struct {
	value *OrganizationSchemasChannel
	isSet bool
}

func (v NullableOrganizationSchemasChannel) Get() *OrganizationSchemasChannel {
	return v.value
}

func (v *NullableOrganizationSchemasChannel) Set(val *OrganizationSchemasChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSchemasChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSchemasChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSchemasChannel(val *OrganizationSchemasChannel) *NullableOrganizationSchemasChannel {
	return &NullableOrganizationSchemasChannel{value: val, isSet: true}
}

func (v NullableOrganizationSchemasChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSchemasChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}