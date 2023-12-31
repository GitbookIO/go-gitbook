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

// RequestInviteUsersToOrganizationEmailsInner - struct for RequestInviteUsersToOrganizationEmailsInner
type RequestInviteUsersToOrganizationEmailsInner struct {
	RequestInviteUsersToOrganizationEmailsInnerOneOf *RequestInviteUsersToOrganizationEmailsInnerOneOf
	String                                           *string
}

// RequestInviteUsersToOrganizationEmailsInnerOneOfAsRequestInviteUsersToOrganizationEmailsInner is a convenience function that returns RequestInviteUsersToOrganizationEmailsInnerOneOf wrapped in RequestInviteUsersToOrganizationEmailsInner
func RequestInviteUsersToOrganizationEmailsInnerOneOfAsRequestInviteUsersToOrganizationEmailsInner(v *RequestInviteUsersToOrganizationEmailsInnerOneOf) RequestInviteUsersToOrganizationEmailsInner {
	return RequestInviteUsersToOrganizationEmailsInner{
		RequestInviteUsersToOrganizationEmailsInnerOneOf: v,
	}
}

// stringAsRequestInviteUsersToOrganizationEmailsInner is a convenience function that returns string wrapped in RequestInviteUsersToOrganizationEmailsInner
func StringAsRequestInviteUsersToOrganizationEmailsInner(v *string) RequestInviteUsersToOrganizationEmailsInner {
	return RequestInviteUsersToOrganizationEmailsInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestInviteUsersToOrganizationEmailsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RequestInviteUsersToOrganizationEmailsInnerOneOf
	err = newStrictDecoder(data).Decode(&dst.RequestInviteUsersToOrganizationEmailsInnerOneOf)
	if err == nil {
		jsonRequestInviteUsersToOrganizationEmailsInnerOneOf, _ := json.Marshal(dst.RequestInviteUsersToOrganizationEmailsInnerOneOf)
		if string(jsonRequestInviteUsersToOrganizationEmailsInnerOneOf) == "{}" { // empty struct
			dst.RequestInviteUsersToOrganizationEmailsInnerOneOf = nil
		} else {
			match++
		}
	} else {
		dst.RequestInviteUsersToOrganizationEmailsInnerOneOf = nil
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
		dst.RequestInviteUsersToOrganizationEmailsInnerOneOf = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RequestInviteUsersToOrganizationEmailsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestInviteUsersToOrganizationEmailsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestInviteUsersToOrganizationEmailsInner) MarshalJSON() ([]byte, error) {
	if src.RequestInviteUsersToOrganizationEmailsInnerOneOf != nil {
		return json.Marshal(&src.RequestInviteUsersToOrganizationEmailsInnerOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestInviteUsersToOrganizationEmailsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RequestInviteUsersToOrganizationEmailsInnerOneOf != nil {
		return obj.RequestInviteUsersToOrganizationEmailsInnerOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRequestInviteUsersToOrganizationEmailsInner struct {
	value *RequestInviteUsersToOrganizationEmailsInner
	isSet bool
}

func (v NullableRequestInviteUsersToOrganizationEmailsInner) Get() *RequestInviteUsersToOrganizationEmailsInner {
	return v.value
}

func (v *NullableRequestInviteUsersToOrganizationEmailsInner) Set(val *RequestInviteUsersToOrganizationEmailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestInviteUsersToOrganizationEmailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestInviteUsersToOrganizationEmailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestInviteUsersToOrganizationEmailsInner(val *RequestInviteUsersToOrganizationEmailsInner) *NullableRequestInviteUsersToOrganizationEmailsInner {
	return &NullableRequestInviteUsersToOrganizationEmailsInner{value: val, isSet: true}
}

func (v NullableRequestInviteUsersToOrganizationEmailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestInviteUsersToOrganizationEmailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
