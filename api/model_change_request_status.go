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

// ChangeRequestStatus the model 'ChangeRequestStatus'
type ChangeRequestStatus string

// List of ChangeRequestStatus
const (
	CHANGEREQUESTSTATUS_DRAFT    ChangeRequestStatus = "draft"
	CHANGEREQUESTSTATUS_OPEN     ChangeRequestStatus = "open"
	CHANGEREQUESTSTATUS_ARCHIVED ChangeRequestStatus = "archived"
	CHANGEREQUESTSTATUS_MERGED   ChangeRequestStatus = "merged"
)

// All allowed values of ChangeRequestStatus enum
var AllowedChangeRequestStatusEnumValues = []ChangeRequestStatus{
	"draft",
	"open",
	"archived",
	"merged",
}

func (v *ChangeRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeRequestStatus(value)
	for _, existing := range AllowedChangeRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeRequestStatus", value)
}

// NewChangeRequestStatusFromValue returns a pointer to a valid ChangeRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeRequestStatusFromValue(v string) (*ChangeRequestStatus, error) {
	ev := ChangeRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeRequestStatus: valid values are %v", v, AllowedChangeRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeRequestStatus) IsValid() bool {
	for _, existing := range AllowedChangeRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeRequestStatus value
func (v ChangeRequestStatus) Ptr() *ChangeRequestStatus {
	return &v
}

type NullableChangeRequestStatus struct {
	value *ChangeRequestStatus
	isSet bool
}

func (v NullableChangeRequestStatus) Get() *ChangeRequestStatus {
	return v.value
}

func (v *NullableChangeRequestStatus) Set(val *ChangeRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestStatus(val *ChangeRequestStatus) *NullableChangeRequestStatus {
	return &NullableChangeRequestStatus{value: val, isSet: true}
}

func (v NullableChangeRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
