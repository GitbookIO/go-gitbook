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

// checks if the UpdateMemberInOrganizationByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMemberInOrganizationByIdRequest{}

// UpdateMemberInOrganizationByIdRequest struct for UpdateMemberInOrganizationByIdRequest
type UpdateMemberInOrganizationByIdRequest struct {
	Role *MemberRoleOrGuest `json:"role,omitempty"`
}

// NewUpdateMemberInOrganizationByIdRequest instantiates a new UpdateMemberInOrganizationByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMemberInOrganizationByIdRequest() *UpdateMemberInOrganizationByIdRequest {
	this := UpdateMemberInOrganizationByIdRequest{}
	return &this
}

// NewUpdateMemberInOrganizationByIdRequestWithDefaults instantiates a new UpdateMemberInOrganizationByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMemberInOrganizationByIdRequestWithDefaults() *UpdateMemberInOrganizationByIdRequest {
	this := UpdateMemberInOrganizationByIdRequest{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UpdateMemberInOrganizationByIdRequest) GetRole() MemberRoleOrGuest {
	if o == nil || IsNil(o.Role) {
		var ret MemberRoleOrGuest
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMemberInOrganizationByIdRequest) GetRoleOk() (*MemberRoleOrGuest, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UpdateMemberInOrganizationByIdRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given MemberRoleOrGuest and assigns it to the Role field.
func (o *UpdateMemberInOrganizationByIdRequest) SetRole(v MemberRoleOrGuest) {
	o.Role = &v
}

func (o UpdateMemberInOrganizationByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMemberInOrganizationByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableUpdateMemberInOrganizationByIdRequest struct {
	value *UpdateMemberInOrganizationByIdRequest
	isSet bool
}

func (v NullableUpdateMemberInOrganizationByIdRequest) Get() *UpdateMemberInOrganizationByIdRequest {
	return v.value
}

func (v *NullableUpdateMemberInOrganizationByIdRequest) Set(val *UpdateMemberInOrganizationByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMemberInOrganizationByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMemberInOrganizationByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMemberInOrganizationByIdRequest(val *UpdateMemberInOrganizationByIdRequest) *NullableUpdateMemberInOrganizationByIdRequest {
	return &NullableUpdateMemberInOrganizationByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateMemberInOrganizationByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMemberInOrganizationByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
