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

// UpsertEntityPropertiesValue - struct for UpsertEntityPropertiesValue
type UpsertEntityPropertiesValue struct {
	UpsertEntityPropertiesValueOneOf *UpsertEntityPropertiesValueOneOf
	Bool                             *bool
	Float32                          *float32
	String                           *string
}

// UpsertEntityPropertiesValueOneOfAsUpsertEntityPropertiesValue is a convenience function that returns UpsertEntityPropertiesValueOneOf wrapped in UpsertEntityPropertiesValue
func UpsertEntityPropertiesValueOneOfAsUpsertEntityPropertiesValue(v *UpsertEntityPropertiesValueOneOf) UpsertEntityPropertiesValue {
	return UpsertEntityPropertiesValue{
		UpsertEntityPropertiesValueOneOf: v,
	}
}

// boolAsUpsertEntityPropertiesValue is a convenience function that returns bool wrapped in UpsertEntityPropertiesValue
func BoolAsUpsertEntityPropertiesValue(v *bool) UpsertEntityPropertiesValue {
	return UpsertEntityPropertiesValue{
		Bool: v,
	}
}

// float32AsUpsertEntityPropertiesValue is a convenience function that returns float32 wrapped in UpsertEntityPropertiesValue
func Float32AsUpsertEntityPropertiesValue(v *float32) UpsertEntityPropertiesValue {
	return UpsertEntityPropertiesValue{
		Float32: v,
	}
}

// stringAsUpsertEntityPropertiesValue is a convenience function that returns string wrapped in UpsertEntityPropertiesValue
func StringAsUpsertEntityPropertiesValue(v *string) UpsertEntityPropertiesValue {
	return UpsertEntityPropertiesValue{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpsertEntityPropertiesValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpsertEntityPropertiesValueOneOf
	err = newStrictDecoder(data).Decode(&dst.UpsertEntityPropertiesValueOneOf)
	if err == nil {
		jsonUpsertEntityPropertiesValueOneOf, _ := json.Marshal(dst.UpsertEntityPropertiesValueOneOf)
		if string(jsonUpsertEntityPropertiesValueOneOf) == "{}" { // empty struct
			dst.UpsertEntityPropertiesValueOneOf = nil
		} else {
			match++
		}
	} else {
		dst.UpsertEntityPropertiesValueOneOf = nil
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
		dst.UpsertEntityPropertiesValueOneOf = nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpsertEntityPropertiesValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpsertEntityPropertiesValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpsertEntityPropertiesValue) MarshalJSON() ([]byte, error) {
	if src.UpsertEntityPropertiesValueOneOf != nil {
		return json.Marshal(&src.UpsertEntityPropertiesValueOneOf)
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
func (obj *UpsertEntityPropertiesValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpsertEntityPropertiesValueOneOf != nil {
		return obj.UpsertEntityPropertiesValueOneOf
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

type NullableUpsertEntityPropertiesValue struct {
	value *UpsertEntityPropertiesValue
	isSet bool
}

func (v NullableUpsertEntityPropertiesValue) Get() *UpsertEntityPropertiesValue {
	return v.value
}

func (v *NullableUpsertEntityPropertiesValue) Set(val *UpsertEntityPropertiesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertEntityPropertiesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertEntityPropertiesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertEntityPropertiesValue(val *UpsertEntityPropertiesValue) *NullableUpsertEntityPropertiesValue {
	return &NullableUpsertEntityPropertiesValue{value: val, isSet: true}
}

func (v NullableUpsertEntityPropertiesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertEntityPropertiesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
