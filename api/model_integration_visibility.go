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

// IntegrationVisibility the model 'IntegrationVisibility'
type IntegrationVisibility string

// List of IntegrationVisibility
const (
	INTEGRATIONVISIBILITY_PUBLIC   IntegrationVisibility = "public"
	INTEGRATIONVISIBILITY_PRIVATE  IntegrationVisibility = "private"
	INTEGRATIONVISIBILITY_UNLISTED IntegrationVisibility = "unlisted"
)

// All allowed values of IntegrationVisibility enum
var AllowedIntegrationVisibilityEnumValues = []IntegrationVisibility{
	"public",
	"private",
	"unlisted",
}

func (v *IntegrationVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationVisibility(value)
	for _, existing := range AllowedIntegrationVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationVisibility", value)
}

// NewIntegrationVisibilityFromValue returns a pointer to a valid IntegrationVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationVisibilityFromValue(v string) (*IntegrationVisibility, error) {
	ev := IntegrationVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationVisibility: valid values are %v", v, AllowedIntegrationVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationVisibility) IsValid() bool {
	for _, existing := range AllowedIntegrationVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationVisibility value
func (v IntegrationVisibility) Ptr() *IntegrationVisibility {
	return &v
}

type NullableIntegrationVisibility struct {
	value *IntegrationVisibility
	isSet bool
}

func (v NullableIntegrationVisibility) Get() *IntegrationVisibility {
	return v.value
}

func (v *NullableIntegrationVisibility) Set(val *IntegrationVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVisibility(val *IntegrationVisibility) *NullableIntegrationVisibility {
	return &NullableIntegrationVisibility{value: val, isSet: true}
}

func (v NullableIntegrationVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
