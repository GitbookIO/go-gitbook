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

// checks if the CreateOrganizationTeamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationTeamRequest{}

// CreateOrganizationTeamRequest struct for CreateOrganizationTeamRequest
type CreateOrganizationTeamRequest struct {
	Title   string   `json:"title"`
	Members []string `json:"members,omitempty"`
}

// NewCreateOrganizationTeamRequest instantiates a new CreateOrganizationTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationTeamRequest(title string) *CreateOrganizationTeamRequest {
	this := CreateOrganizationTeamRequest{}
	this.Title = title
	return &this
}

// NewCreateOrganizationTeamRequestWithDefaults instantiates a new CreateOrganizationTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationTeamRequestWithDefaults() *CreateOrganizationTeamRequest {
	this := CreateOrganizationTeamRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *CreateOrganizationTeamRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationTeamRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateOrganizationTeamRequest) SetTitle(v string) {
	o.Title = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *CreateOrganizationTeamRequest) GetMembers() []string {
	if o == nil || IsNil(o.Members) {
		var ret []string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationTeamRequest) GetMembersOk() ([]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CreateOrganizationTeamRequest) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *CreateOrganizationTeamRequest) SetMembers(v []string) {
	o.Members = v
}

func (o CreateOrganizationTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationTeamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullableCreateOrganizationTeamRequest struct {
	value *CreateOrganizationTeamRequest
	isSet bool
}

func (v NullableCreateOrganizationTeamRequest) Get() *CreateOrganizationTeamRequest {
	return v.value
}

func (v *NullableCreateOrganizationTeamRequest) Set(val *CreateOrganizationTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationTeamRequest(val *CreateOrganizationTeamRequest) *NullableCreateOrganizationTeamRequest {
	return &NullableCreateOrganizationTeamRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
