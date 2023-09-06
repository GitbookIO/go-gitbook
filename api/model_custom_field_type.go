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

// CustomFieldType the model 'CustomFieldType'
type CustomFieldType string

// List of CustomFieldType
const (
	CUSTOMFIELDTYPE_TEXT         CustomFieldType = "text"
	CUSTOMFIELDTYPE_NUMBER       CustomFieldType = "number"
	CUSTOMFIELDTYPE_BOOLEAN      CustomFieldType = "boolean"
	CUSTOMFIELDTYPE_TAGS         CustomFieldType = "tags"
	CUSTOMFIELDTYPE_SELECTMULTI  CustomFieldType = "select:multi"
	CUSTOMFIELDTYPE_SELECTSINGLE CustomFieldType = "select:single"
)

// All allowed values of CustomFieldType enum
var AllowedCustomFieldTypeEnumValues = []CustomFieldType{
	"text",
	"number",
	"boolean",
	"tags",
	"select:multi",
	"select:single",
}

func (v *CustomFieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomFieldType(value)
	for _, existing := range AllowedCustomFieldTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomFieldType", value)
}

// NewCustomFieldTypeFromValue returns a pointer to a valid CustomFieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomFieldTypeFromValue(v string) (*CustomFieldType, error) {
	ev := CustomFieldType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomFieldType: valid values are %v", v, AllowedCustomFieldTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomFieldType) IsValid() bool {
	for _, existing := range AllowedCustomFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomFieldType value
func (v CustomFieldType) Ptr() *CustomFieldType {
	return &v
}

type NullableCustomFieldType struct {
	value *CustomFieldType
	isSet bool
}

func (v NullableCustomFieldType) Get() *CustomFieldType {
	return v.value
}

func (v *NullableCustomFieldType) Set(val *CustomFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldType(val *CustomFieldType) *NullableCustomFieldType {
	return &NullableCustomFieldType{value: val, isSet: true}
}

func (v NullableCustomFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
