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

// CustomFieldValue - struct for CustomFieldValue
type CustomFieldValue struct {
	ArrayOfString *[]string
	Bool          *bool
	Float32       *float32
	String        *string
}

// []stringAsCustomFieldValue is a convenience function that returns []string wrapped in CustomFieldValue
func ArrayOfStringAsCustomFieldValue(v *[]string) CustomFieldValue {
	return CustomFieldValue{
		ArrayOfString: v,
	}
}

// boolAsCustomFieldValue is a convenience function that returns bool wrapped in CustomFieldValue
func BoolAsCustomFieldValue(v *bool) CustomFieldValue {
	return CustomFieldValue{
		Bool: v,
	}
}

// float32AsCustomFieldValue is a convenience function that returns float32 wrapped in CustomFieldValue
func Float32AsCustomFieldValue(v *float32) CustomFieldValue {
	return CustomFieldValue{
		Float32: v,
	}
}

// stringAsCustomFieldValue is a convenience function that returns string wrapped in CustomFieldValue
func StringAsCustomFieldValue(v *string) CustomFieldValue {
	return CustomFieldValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CustomFieldValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
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
		dst.ArrayOfString = nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CustomFieldValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CustomFieldValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CustomFieldValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CustomFieldValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCustomFieldValue struct {
	value *CustomFieldValue
	isSet bool
}

func (v NullableCustomFieldValue) Get() *CustomFieldValue {
	return v.value
}

func (v *NullableCustomFieldValue) Set(val *CustomFieldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldValue(val *CustomFieldValue) *NullableCustomFieldValue {
	return &NullableCustomFieldValue{value: val, isSet: true}
}

func (v NullableCustomFieldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
