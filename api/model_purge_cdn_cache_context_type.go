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

// PurgeCDNCacheContextType The type of purge, e.g by tags or hosts
type PurgeCDNCacheContextType string

// List of PurgeCDNCacheContextType
const (
	PURGECDNCACHECONTEXTTYPE_TAGS  PurgeCDNCacheContextType = "tags"
	PURGECDNCACHECONTEXTTYPE_HOSTS PurgeCDNCacheContextType = "hosts"
)

// All allowed values of PurgeCDNCacheContextType enum
var AllowedPurgeCDNCacheContextTypeEnumValues = []PurgeCDNCacheContextType{
	"tags",
	"hosts",
}

func (v *PurgeCDNCacheContextType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PurgeCDNCacheContextType(value)
	for _, existing := range AllowedPurgeCDNCacheContextTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PurgeCDNCacheContextType", value)
}

// NewPurgeCDNCacheContextTypeFromValue returns a pointer to a valid PurgeCDNCacheContextType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPurgeCDNCacheContextTypeFromValue(v string) (*PurgeCDNCacheContextType, error) {
	ev := PurgeCDNCacheContextType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PurgeCDNCacheContextType: valid values are %v", v, AllowedPurgeCDNCacheContextTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PurgeCDNCacheContextType) IsValid() bool {
	for _, existing := range AllowedPurgeCDNCacheContextTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PurgeCDNCacheContextType value
func (v PurgeCDNCacheContextType) Ptr() *PurgeCDNCacheContextType {
	return &v
}

type NullablePurgeCDNCacheContextType struct {
	value *PurgeCDNCacheContextType
	isSet bool
}

func (v NullablePurgeCDNCacheContextType) Get() *PurgeCDNCacheContextType {
	return v.value
}

func (v *NullablePurgeCDNCacheContextType) Set(val *PurgeCDNCacheContextType) {
	v.value = val
	v.isSet = true
}

func (v NullablePurgeCDNCacheContextType) IsSet() bool {
	return v.isSet
}

func (v *NullablePurgeCDNCacheContextType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurgeCDNCacheContextType(val *PurgeCDNCacheContextType) *NullablePurgeCDNCacheContextType {
	return &NullablePurgeCDNCacheContextType{value: val, isSet: true}
}

func (v NullablePurgeCDNCacheContextType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurgeCDNCacheContextType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
