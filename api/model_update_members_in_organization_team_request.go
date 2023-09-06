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

// checks if the UpdateMembersInOrganizationTeamRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMembersInOrganizationTeamRequest{}

// UpdateMembersInOrganizationTeamRequest struct for UpdateMembersInOrganizationTeamRequest
type UpdateMembersInOrganizationTeamRequest struct {
	Add         []string               `json:"add,omitempty"`
	Memberships *map[string]TeamMember `json:"memberships,omitempty"`
	Remove      []string               `json:"remove,omitempty"`
}

// NewUpdateMembersInOrganizationTeamRequest instantiates a new UpdateMembersInOrganizationTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMembersInOrganizationTeamRequest() *UpdateMembersInOrganizationTeamRequest {
	this := UpdateMembersInOrganizationTeamRequest{}
	return &this
}

// NewUpdateMembersInOrganizationTeamRequestWithDefaults instantiates a new UpdateMembersInOrganizationTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMembersInOrganizationTeamRequestWithDefaults() *UpdateMembersInOrganizationTeamRequest {
	this := UpdateMembersInOrganizationTeamRequest{}
	return &this
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *UpdateMembersInOrganizationTeamRequest) GetAdd() []string {
	if o == nil || IsNil(o.Add) {
		var ret []string
		return ret
	}
	return o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMembersInOrganizationTeamRequest) GetAddOk() ([]string, bool) {
	if o == nil || IsNil(o.Add) {
		return nil, false
	}
	return o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *UpdateMembersInOrganizationTeamRequest) HasAdd() bool {
	if o != nil && !IsNil(o.Add) {
		return true
	}

	return false
}

// SetAdd gets a reference to the given []string and assigns it to the Add field.
func (o *UpdateMembersInOrganizationTeamRequest) SetAdd(v []string) {
	o.Add = v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *UpdateMembersInOrganizationTeamRequest) GetMemberships() map[string]TeamMember {
	if o == nil || IsNil(o.Memberships) {
		var ret map[string]TeamMember
		return ret
	}
	return *o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMembersInOrganizationTeamRequest) GetMembershipsOk() (*map[string]TeamMember, bool) {
	if o == nil || IsNil(o.Memberships) {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *UpdateMembersInOrganizationTeamRequest) HasMemberships() bool {
	if o != nil && !IsNil(o.Memberships) {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given map[string]TeamMember and assigns it to the Memberships field.
func (o *UpdateMembersInOrganizationTeamRequest) SetMemberships(v map[string]TeamMember) {
	o.Memberships = &v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *UpdateMembersInOrganizationTeamRequest) GetRemove() []string {
	if o == nil || IsNil(o.Remove) {
		var ret []string
		return ret
	}
	return o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMembersInOrganizationTeamRequest) GetRemoveOk() ([]string, bool) {
	if o == nil || IsNil(o.Remove) {
		return nil, false
	}
	return o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *UpdateMembersInOrganizationTeamRequest) HasRemove() bool {
	if o != nil && !IsNil(o.Remove) {
		return true
	}

	return false
}

// SetRemove gets a reference to the given []string and assigns it to the Remove field.
func (o *UpdateMembersInOrganizationTeamRequest) SetRemove(v []string) {
	o.Remove = v
}

func (o UpdateMembersInOrganizationTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMembersInOrganizationTeamRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Add) {
		toSerialize["add"] = o.Add
	}
	if !IsNil(o.Memberships) {
		toSerialize["memberships"] = o.Memberships
	}
	if !IsNil(o.Remove) {
		toSerialize["remove"] = o.Remove
	}
	return toSerialize, nil
}

type NullableUpdateMembersInOrganizationTeamRequest struct {
	value *UpdateMembersInOrganizationTeamRequest
	isSet bool
}

func (v NullableUpdateMembersInOrganizationTeamRequest) Get() *UpdateMembersInOrganizationTeamRequest {
	return v.value
}

func (v *NullableUpdateMembersInOrganizationTeamRequest) Set(val *UpdateMembersInOrganizationTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMembersInOrganizationTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMembersInOrganizationTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMembersInOrganizationTeamRequest(val *UpdateMembersInOrganizationTeamRequest) *NullableUpdateMembersInOrganizationTeamRequest {
	return &NullableUpdateMembersInOrganizationTeamRequest{value: val, isSet: true}
}

func (v NullableUpdateMembersInOrganizationTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMembersInOrganizationTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
