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

// ContentKitRootElement - Element used as root
type ContentKitRootElement struct {
	ContentKitBlock *ContentKitBlock
	ContentKitModal *ContentKitModal
}

// ContentKitBlockAsContentKitRootElement is a convenience function that returns ContentKitBlock wrapped in ContentKitRootElement
func ContentKitBlockAsContentKitRootElement(v *ContentKitBlock) ContentKitRootElement {
	return ContentKitRootElement{
		ContentKitBlock: v,
	}
}

// ContentKitModalAsContentKitRootElement is a convenience function that returns ContentKitModal wrapped in ContentKitRootElement
func ContentKitModalAsContentKitRootElement(v *ContentKitModal) ContentKitRootElement {
	return ContentKitRootElement{
		ContentKitModal: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitRootElement) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentKitBlock
	err = newStrictDecoder(data).Decode(&dst.ContentKitBlock)
	if err == nil {
		jsonContentKitBlock, _ := json.Marshal(dst.ContentKitBlock)
		if string(jsonContentKitBlock) == "{}" { // empty struct
			dst.ContentKitBlock = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitBlock = nil
	}

	// try to unmarshal data into ContentKitModal
	err = newStrictDecoder(data).Decode(&dst.ContentKitModal)
	if err == nil {
		jsonContentKitModal, _ := json.Marshal(dst.ContentKitModal)
		if string(jsonContentKitModal) == "{}" { // empty struct
			dst.ContentKitModal = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitModal = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContentKitBlock = nil
		dst.ContentKitModal = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitRootElement)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitRootElement)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitRootElement) MarshalJSON() ([]byte, error) {
	if src.ContentKitBlock != nil {
		return json.Marshal(&src.ContentKitBlock)
	}

	if src.ContentKitModal != nil {
		return json.Marshal(&src.ContentKitModal)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitRootElement) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentKitBlock != nil {
		return obj.ContentKitBlock
	}

	if obj.ContentKitModal != nil {
		return obj.ContentKitModal
	}

	// all schemas are nil
	return nil
}

type NullableContentKitRootElement struct {
	value *ContentKitRootElement
	isSet bool
}

func (v NullableContentKitRootElement) Get() *ContentKitRootElement {
	return v.value
}

func (v *NullableContentKitRootElement) Set(val *ContentKitRootElement) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitRootElement) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitRootElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitRootElement(val *ContentKitRootElement) *NullableContentKitRootElement {
	return &NullableContentKitRootElement{value: val, isSet: true}
}

func (v NullableContentKitRootElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitRootElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
