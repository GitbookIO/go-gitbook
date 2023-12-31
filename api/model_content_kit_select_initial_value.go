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

// ContentKitSelectInitialValue - Value to initialize the select with.
type ContentKitSelectInitialValue struct {
	ArrayOfString *[]string
	String        *string
}

// []stringAsContentKitSelectInitialValue is a convenience function that returns []string wrapped in ContentKitSelectInitialValue
func ArrayOfStringAsContentKitSelectInitialValue(v *[]string) ContentKitSelectInitialValue {
	return ContentKitSelectInitialValue{
		ArrayOfString: v,
	}
}

// stringAsContentKitSelectInitialValue is a convenience function that returns string wrapped in ContentKitSelectInitialValue
func StringAsContentKitSelectInitialValue(v *string) ContentKitSelectInitialValue {
	return ContentKitSelectInitialValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitSelectInitialValue) UnmarshalJSON(data []byte) error {
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
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitSelectInitialValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitSelectInitialValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitSelectInitialValue) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitSelectInitialValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableContentKitSelectInitialValue struct {
	value *ContentKitSelectInitialValue
	isSet bool
}

func (v NullableContentKitSelectInitialValue) Get() *ContentKitSelectInitialValue {
	return v.value
}

func (v *NullableContentKitSelectInitialValue) Set(val *ContentKitSelectInitialValue) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitSelectInitialValue) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitSelectInitialValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitSelectInitialValue(val *ContentKitSelectInitialValue) *NullableContentKitSelectInitialValue {
	return &NullableContentKitSelectInitialValue{value: val, isSet: true}
}

func (v NullableContentKitSelectInitialValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitSelectInitialValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
