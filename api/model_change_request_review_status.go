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
	"fmt"
)

// ChangeRequestReviewStatus Status of a change request review.
type ChangeRequestReviewStatus string

// List of ChangeRequestReviewStatus
const (
	CHANGEREQUESTREVIEWSTATUS_CHANGES_REQUESTED ChangeRequestReviewStatus = "changes-requested"
	CHANGEREQUESTREVIEWSTATUS_APPROVED          ChangeRequestReviewStatus = "approved"
)

// All allowed values of ChangeRequestReviewStatus enum
var AllowedChangeRequestReviewStatusEnumValues = []ChangeRequestReviewStatus{
	"changes-requested",
	"approved",
}

func (v *ChangeRequestReviewStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeRequestReviewStatus(value)
	for _, existing := range AllowedChangeRequestReviewStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeRequestReviewStatus", value)
}

// NewChangeRequestReviewStatusFromValue returns a pointer to a valid ChangeRequestReviewStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeRequestReviewStatusFromValue(v string) (*ChangeRequestReviewStatus, error) {
	ev := ChangeRequestReviewStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeRequestReviewStatus: valid values are %v", v, AllowedChangeRequestReviewStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeRequestReviewStatus) IsValid() bool {
	for _, existing := range AllowedChangeRequestReviewStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeRequestReviewStatus value
func (v ChangeRequestReviewStatus) Ptr() *ChangeRequestReviewStatus {
	return &v
}

type NullableChangeRequestReviewStatus struct {
	value *ChangeRequestReviewStatus
	isSet bool
}

func (v NullableChangeRequestReviewStatus) Get() *ChangeRequestReviewStatus {
	return v.value
}

func (v *NullableChangeRequestReviewStatus) Set(val *ChangeRequestReviewStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestReviewStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestReviewStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestReviewStatus(val *ChangeRequestReviewStatus) *NullableChangeRequestReviewStatus {
	return &NullableChangeRequestReviewStatus{value: val, isSet: true}
}

func (v NullableChangeRequestReviewStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestReviewStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
