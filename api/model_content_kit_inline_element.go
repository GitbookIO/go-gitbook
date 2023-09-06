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

// ContentKitInlineElement - Any element that is inline.
type ContentKitInlineElement struct {
	ContentKitImage *ContentKitImage
	ContentKitLink  *ContentKitLink
	ContentKitText  *ContentKitText
}

// ContentKitImageAsContentKitInlineElement is a convenience function that returns ContentKitImage wrapped in ContentKitInlineElement
func ContentKitImageAsContentKitInlineElement(v *ContentKitImage) ContentKitInlineElement {
	return ContentKitInlineElement{
		ContentKitImage: v,
	}
}

// ContentKitLinkAsContentKitInlineElement is a convenience function that returns ContentKitLink wrapped in ContentKitInlineElement
func ContentKitLinkAsContentKitInlineElement(v *ContentKitLink) ContentKitInlineElement {
	return ContentKitInlineElement{
		ContentKitLink: v,
	}
}

// ContentKitTextAsContentKitInlineElement is a convenience function that returns ContentKitText wrapped in ContentKitInlineElement
func ContentKitTextAsContentKitInlineElement(v *ContentKitText) ContentKitInlineElement {
	return ContentKitInlineElement{
		ContentKitText: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitInlineElement) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentKitImage
	err = newStrictDecoder(data).Decode(&dst.ContentKitImage)
	if err == nil {
		jsonContentKitImage, _ := json.Marshal(dst.ContentKitImage)
		if string(jsonContentKitImage) == "{}" { // empty struct
			dst.ContentKitImage = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitImage = nil
	}

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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContentKitImage = nil
		dst.ContentKitLink = nil
		dst.ContentKitText = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitInlineElement)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitInlineElement)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitInlineElement) MarshalJSON() ([]byte, error) {
	if src.ContentKitImage != nil {
		return json.Marshal(&src.ContentKitImage)
	}

	if src.ContentKitLink != nil {
		return json.Marshal(&src.ContentKitLink)
	}

	if src.ContentKitText != nil {
		return json.Marshal(&src.ContentKitText)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitInlineElement) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentKitImage != nil {
		return obj.ContentKitImage
	}

	if obj.ContentKitLink != nil {
		return obj.ContentKitLink
	}

	if obj.ContentKitText != nil {
		return obj.ContentKitText
	}

	// all schemas are nil
	return nil
}

type NullableContentKitInlineElement struct {
	value *ContentKitInlineElement
	isSet bool
}

func (v NullableContentKitInlineElement) Get() *ContentKitInlineElement {
	return v.value
}

func (v *NullableContentKitInlineElement) Set(val *ContentKitInlineElement) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitInlineElement) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitInlineElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitInlineElement(val *ContentKitInlineElement) *NullableContentKitInlineElement {
	return &NullableContentKitInlineElement{value: val, isSet: true}
}

func (v NullableContentKitInlineElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitInlineElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
