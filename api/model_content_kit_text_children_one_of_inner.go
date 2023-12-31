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

// ContentKitTextChildrenOneOfInner - struct for ContentKitTextChildrenOneOfInner
type ContentKitTextChildrenOneOfInner struct {
	ContentKitLink *ContentKitLink
	ContentKitText *ContentKitText
	String         *string
}

// ContentKitLinkAsContentKitTextChildrenOneOfInner is a convenience function that returns ContentKitLink wrapped in ContentKitTextChildrenOneOfInner
func ContentKitLinkAsContentKitTextChildrenOneOfInner(v *ContentKitLink) ContentKitTextChildrenOneOfInner {
	return ContentKitTextChildrenOneOfInner{
		ContentKitLink: v,
	}
}

// ContentKitTextAsContentKitTextChildrenOneOfInner is a convenience function that returns ContentKitText wrapped in ContentKitTextChildrenOneOfInner
func ContentKitTextAsContentKitTextChildrenOneOfInner(v *ContentKitText) ContentKitTextChildrenOneOfInner {
	return ContentKitTextChildrenOneOfInner{
		ContentKitText: v,
	}
}

// stringAsContentKitTextChildrenOneOfInner is a convenience function that returns string wrapped in ContentKitTextChildrenOneOfInner
func StringAsContentKitTextChildrenOneOfInner(v *string) ContentKitTextChildrenOneOfInner {
	return ContentKitTextChildrenOneOfInner{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitTextChildrenOneOfInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentKitLink
	err = newStrictDecoder(data).Decode(&dst.ContentKitLink)
	if err == nil {
		jsonContentKitLink, _ := json.Marshal(dst.ContentKitLink)
		if string(jsonContentKitLink) == "{}" { // empty struct
			dst.ContentKitLink = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitLink = nil
	}

	// try to unmarshal data into ContentKitText
	err = newStrictDecoder(data).Decode(&dst.ContentKitText)
	if err == nil {
		jsonContentKitText, _ := json.Marshal(dst.ContentKitText)
		if string(jsonContentKitText) == "{}" { // empty struct
			dst.ContentKitText = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitText = nil
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
		dst.ContentKitLink = nil
		dst.ContentKitText = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitTextChildrenOneOfInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitTextChildrenOneOfInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitTextChildrenOneOfInner) MarshalJSON() ([]byte, error) {
	if src.ContentKitLink != nil {
		return json.Marshal(&src.ContentKitLink)
	}

	if src.ContentKitText != nil {
		return json.Marshal(&src.ContentKitText)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitTextChildrenOneOfInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentKitLink != nil {
		return obj.ContentKitLink
	}

	if obj.ContentKitText != nil {
		return obj.ContentKitText
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableContentKitTextChildrenOneOfInner struct {
	value *ContentKitTextChildrenOneOfInner
	isSet bool
}

func (v NullableContentKitTextChildrenOneOfInner) Get() *ContentKitTextChildrenOneOfInner {
	return v.value
}

func (v *NullableContentKitTextChildrenOneOfInner) Set(val *ContentKitTextChildrenOneOfInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitTextChildrenOneOfInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitTextChildrenOneOfInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitTextChildrenOneOfInner(val *ContentKitTextChildrenOneOfInner) *NullableContentKitTextChildrenOneOfInner {
	return &NullableContentKitTextChildrenOneOfInner{value: val, isSet: true}
}

func (v NullableContentKitTextChildrenOneOfInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitTextChildrenOneOfInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
