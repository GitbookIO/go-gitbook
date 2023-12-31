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

// GetContentByUrl200Response - struct for GetContentByUrl200Response
type GetContentByUrl200Response struct {
	GetContentByUrl200ResponseOneOf  *GetContentByUrl200ResponseOneOf
	GetContentByUrl200ResponseOneOf1 *GetContentByUrl200ResponseOneOf1
}

// GetContentByUrl200ResponseOneOfAsGetContentByUrl200Response is a convenience function that returns GetContentByUrl200ResponseOneOf wrapped in GetContentByUrl200Response
func GetContentByUrl200ResponseOneOfAsGetContentByUrl200Response(v *GetContentByUrl200ResponseOneOf) GetContentByUrl200Response {
	return GetContentByUrl200Response{
		GetContentByUrl200ResponseOneOf: v,
	}
}

// GetContentByUrl200ResponseOneOf1AsGetContentByUrl200Response is a convenience function that returns GetContentByUrl200ResponseOneOf1 wrapped in GetContentByUrl200Response
func GetContentByUrl200ResponseOneOf1AsGetContentByUrl200Response(v *GetContentByUrl200ResponseOneOf1) GetContentByUrl200Response {
	return GetContentByUrl200Response{
		GetContentByUrl200ResponseOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetContentByUrl200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetContentByUrl200ResponseOneOf
	err = newStrictDecoder(data).Decode(&dst.GetContentByUrl200ResponseOneOf)
	if err == nil {
		jsonGetContentByUrl200ResponseOneOf, _ := json.Marshal(dst.GetContentByUrl200ResponseOneOf)
		if string(jsonGetContentByUrl200ResponseOneOf) == "{}" { // empty struct
			dst.GetContentByUrl200ResponseOneOf = nil
		} else {
			match++
		}
	} else {
		dst.GetContentByUrl200ResponseOneOf = nil
	}

	// try to unmarshal data into GetContentByUrl200ResponseOneOf1
	err = newStrictDecoder(data).Decode(&dst.GetContentByUrl200ResponseOneOf1)
	if err == nil {
		jsonGetContentByUrl200ResponseOneOf1, _ := json.Marshal(dst.GetContentByUrl200ResponseOneOf1)
		if string(jsonGetContentByUrl200ResponseOneOf1) == "{}" { // empty struct
			dst.GetContentByUrl200ResponseOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.GetContentByUrl200ResponseOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetContentByUrl200ResponseOneOf = nil
		dst.GetContentByUrl200ResponseOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetContentByUrl200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetContentByUrl200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetContentByUrl200Response) MarshalJSON() ([]byte, error) {
	if src.GetContentByUrl200ResponseOneOf != nil {
		return json.Marshal(&src.GetContentByUrl200ResponseOneOf)
	}

	if src.GetContentByUrl200ResponseOneOf1 != nil {
		return json.Marshal(&src.GetContentByUrl200ResponseOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetContentByUrl200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetContentByUrl200ResponseOneOf != nil {
		return obj.GetContentByUrl200ResponseOneOf
	}

	if obj.GetContentByUrl200ResponseOneOf1 != nil {
		return obj.GetContentByUrl200ResponseOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableGetContentByUrl200Response struct {
	value *GetContentByUrl200Response
	isSet bool
}

func (v NullableGetContentByUrl200Response) Get() *GetContentByUrl200Response {
	return v.value
}

func (v *NullableGetContentByUrl200Response) Set(val *GetContentByUrl200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContentByUrl200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContentByUrl200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContentByUrl200Response(val *GetContentByUrl200Response) *NullableGetContentByUrl200Response {
	return &NullableGetContentByUrl200Response{value: val, isSet: true}
}

func (v NullableGetContentByUrl200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContentByUrl200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
