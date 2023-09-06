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
	"fmt"
)

// ListMembersInOrganizationByIdRoleParameter - struct for ListMembersInOrganizationByIdRoleParameter
type ListMembersInOrganizationByIdRoleParameter struct {
	MemberRole *MemberRole
	String     *string
}

// MemberRoleAsListMembersInOrganizationByIdRoleParameter is a convenience function that returns MemberRole wrapped in ListMembersInOrganizationByIdRoleParameter
func MemberRoleAsListMembersInOrganizationByIdRoleParameter(v *MemberRole) ListMembersInOrganizationByIdRoleParameter {
	return ListMembersInOrganizationByIdRoleParameter{
		MemberRole: v,
	}
}

// stringAsListMembersInOrganizationByIdRoleParameter is a convenience function that returns string wrapped in ListMembersInOrganizationByIdRoleParameter
func StringAsListMembersInOrganizationByIdRoleParameter(v *string) ListMembersInOrganizationByIdRoleParameter {
	return ListMembersInOrganizationByIdRoleParameter{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListMembersInOrganizationByIdRoleParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MemberRole
	err = newStrictDecoder(data).Decode(&dst.MemberRole)
	if err == nil {
		jsonMemberRole, _ := json.Marshal(dst.MemberRole)
		if string(jsonMemberRole) == "{}" { // empty struct
			dst.MemberRole = nil
		} else {
			match++
		}
	} else {
		dst.MemberRole = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MemberRole = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListMembersInOrganizationByIdRoleParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListMembersInOrganizationByIdRoleParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListMembersInOrganizationByIdRoleParameter) MarshalJSON() ([]byte, error) {
	if src.MemberRole != nil {
		return json.Marshal(&src.MemberRole)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListMembersInOrganizationByIdRoleParameter) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MemberRole != nil {
		return obj.MemberRole
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableListMembersInOrganizationByIdRoleParameter struct {
	value *ListMembersInOrganizationByIdRoleParameter
	isSet bool
}

func (v NullableListMembersInOrganizationByIdRoleParameter) Get() *ListMembersInOrganizationByIdRoleParameter {
	return v.value
}

func (v *NullableListMembersInOrganizationByIdRoleParameter) Set(val *ListMembersInOrganizationByIdRoleParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableListMembersInOrganizationByIdRoleParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableListMembersInOrganizationByIdRoleParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMembersInOrganizationByIdRoleParameter(val *ListMembersInOrganizationByIdRoleParameter) *NullableListMembersInOrganizationByIdRoleParameter {
	return &NullableListMembersInOrganizationByIdRoleParameter{value: val, isSet: true}
}

func (v NullableListMembersInOrganizationByIdRoleParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMembersInOrganizationByIdRoleParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
