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

// ContentKitCardHint - struct for ContentKitCardHint
type ContentKitCardHint struct {
	ArrayOfContentKitInlineElement *[]ContentKitInlineElement
	String                         *string
}

// []ContentKitInlineElementAsContentKitCardHint is a convenience function that returns []ContentKitInlineElement wrapped in ContentKitCardHint
func ArrayOfContentKitInlineElementAsContentKitCardHint(v *[]ContentKitInlineElement) ContentKitCardHint {
	return ContentKitCardHint{
		ArrayOfContentKitInlineElement: v,
	}
}

// stringAsContentKitCardHint is a convenience function that returns string wrapped in ContentKitCardHint
func StringAsContentKitCardHint(v *string) ContentKitCardHint {
	return ContentKitCardHint{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitCardHint) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfContentKitInlineElement
	err = newStrictDecoder(data).Decode(&dst.ArrayOfContentKitInlineElement)
	if err == nil {
		jsonArrayOfContentKitInlineElement, _ := json.Marshal(dst.ArrayOfContentKitInlineElement)
		if string(jsonArrayOfContentKitInlineElement) == "{}" { // empty struct
			dst.ArrayOfContentKitInlineElement = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfContentKitInlineElement = nil
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
		dst.ArrayOfContentKitInlineElement = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitCardHint)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitCardHint)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitCardHint) MarshalJSON() ([]byte, error) {
	if src.ArrayOfContentKitInlineElement != nil {
		return json.Marshal(&src.ArrayOfContentKitInlineElement)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitCardHint) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfContentKitInlineElement != nil {
		return obj.ArrayOfContentKitInlineElement
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableContentKitCardHint struct {
	value *ContentKitCardHint
	isSet bool
}

func (v NullableContentKitCardHint) Get() *ContentKitCardHint {
	return v.value
}

func (v *NullableContentKitCardHint) Set(val *ContentKitCardHint) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitCardHint) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitCardHint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitCardHint(val *ContentKitCardHint) *NullableContentKitCardHint {
	return &NullableContentKitCardHint{value: val, isSet: true}
}

func (v NullableContentKitCardHint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitCardHint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
