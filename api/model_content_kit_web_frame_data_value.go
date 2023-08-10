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

// ContentKitWebFrameDataValue - struct for ContentKitWebFrameDataValue
type ContentKitWebFrameDataValue struct {
	ContentKitDynamicBinding *ContentKitDynamicBinding
	String                   *string
}

// ContentKitDynamicBindingAsContentKitWebFrameDataValue is a convenience function that returns ContentKitDynamicBinding wrapped in ContentKitWebFrameDataValue
func ContentKitDynamicBindingAsContentKitWebFrameDataValue(v *ContentKitDynamicBinding) ContentKitWebFrameDataValue {
	return ContentKitWebFrameDataValue{
		ContentKitDynamicBinding: v,
	}
}

// stringAsContentKitWebFrameDataValue is a convenience function that returns string wrapped in ContentKitWebFrameDataValue
func StringAsContentKitWebFrameDataValue(v *string) ContentKitWebFrameDataValue {
	return ContentKitWebFrameDataValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitWebFrameDataValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentKitDynamicBinding
	err = newStrictDecoder(data).Decode(&dst.ContentKitDynamicBinding)
	if err == nil {
		jsonContentKitDynamicBinding, _ := json.Marshal(dst.ContentKitDynamicBinding)
		if string(jsonContentKitDynamicBinding) == "{}" { // empty struct
			dst.ContentKitDynamicBinding = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitDynamicBinding = nil
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
		dst.ContentKitDynamicBinding = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitWebFrameDataValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitWebFrameDataValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitWebFrameDataValue) MarshalJSON() ([]byte, error) {
	if src.ContentKitDynamicBinding != nil {
		return json.Marshal(&src.ContentKitDynamicBinding)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitWebFrameDataValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentKitDynamicBinding != nil {
		return obj.ContentKitDynamicBinding
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableContentKitWebFrameDataValue struct {
	value *ContentKitWebFrameDataValue
	isSet bool
}

func (v NullableContentKitWebFrameDataValue) Get() *ContentKitWebFrameDataValue {
	return v.value
}

func (v *NullableContentKitWebFrameDataValue) Set(val *ContentKitWebFrameDataValue) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitWebFrameDataValue) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitWebFrameDataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitWebFrameDataValue(val *ContentKitWebFrameDataValue) *NullableContentKitWebFrameDataValue {
	return &NullableContentKitWebFrameDataValue{value: val, isSet: true}
}

func (v NullableContentKitWebFrameDataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitWebFrameDataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
