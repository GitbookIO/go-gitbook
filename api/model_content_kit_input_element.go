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

// ContentKitInputElement - struct for ContentKitInputElement
type ContentKitInputElement struct {
	ContentKitButton    *ContentKitButton
	ContentKitCheckbox  *ContentKitCheckbox
	ContentKitRadio     *ContentKitRadio
	ContentKitSelect    *ContentKitSelect
	ContentKitSwitch    *ContentKitSwitch
	ContentKitTextInput *ContentKitTextInput
}

// ContentKitButtonAsContentKitInputElement is a convenience function that returns ContentKitButton wrapped in ContentKitInputElement
func ContentKitButtonAsContentKitInputElement(v *ContentKitButton) ContentKitInputElement {
	return ContentKitInputElement{
		ContentKitButton: v,
	}
}

// ContentKitCheckboxAsContentKitInputElement is a convenience function that returns ContentKitCheckbox wrapped in ContentKitInputElement
func ContentKitCheckboxAsContentKitInputElement(v *ContentKitCheckbox) ContentKitInputElement {
	return ContentKitInputElement{
		ContentKitCheckbox: v,
	}
}

// ContentKitRadioAsContentKitInputElement is a convenience function that returns ContentKitRadio wrapped in ContentKitInputElement
func ContentKitRadioAsContentKitInputElement(v *ContentKitRadio) ContentKitInputElement {
	return ContentKitInputElement{
		ContentKitRadio: v,
	}
}

// ContentKitSelectAsContentKitInputElement is a convenience function that returns ContentKitSelect wrapped in ContentKitInputElement
func ContentKitSelectAsContentKitInputElement(v *ContentKitSelect) ContentKitInputElement {
	return ContentKitInputElement{
		ContentKitSelect: v,
	}
}

// ContentKitSwitchAsContentKitInputElement is a convenience function that returns ContentKitSwitch wrapped in ContentKitInputElement
func ContentKitSwitchAsContentKitInputElement(v *ContentKitSwitch) ContentKitInputElement {
	return ContentKitInputElement{
		ContentKitSwitch: v,
	}
}

// ContentKitTextInputAsContentKitInputElement is a convenience function that returns ContentKitTextInput wrapped in ContentKitInputElement
func ContentKitTextInputAsContentKitInputElement(v *ContentKitTextInput) ContentKitInputElement {
	return ContentKitInputElement{
		ContentKitTextInput: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitInputElement) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentKitButton
	err = newStrictDecoder(data).Decode(&dst.ContentKitButton)
	if err == nil {
		jsonContentKitButton, _ := json.Marshal(dst.ContentKitButton)
		if string(jsonContentKitButton) == "{}" { // empty struct
			dst.ContentKitButton = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitButton = nil
	}

	// try to unmarshal data into ContentKitCheckbox
	err = newStrictDecoder(data).Decode(&dst.ContentKitCheckbox)
	if err == nil {
		jsonContentKitCheckbox, _ := json.Marshal(dst.ContentKitCheckbox)
		if string(jsonContentKitCheckbox) == "{}" { // empty struct
			dst.ContentKitCheckbox = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitCheckbox = nil
	}

	// try to unmarshal data into ContentKitRadio
	err = newStrictDecoder(data).Decode(&dst.ContentKitRadio)
	if err == nil {
		jsonContentKitRadio, _ := json.Marshal(dst.ContentKitRadio)
		if string(jsonContentKitRadio) == "{}" { // empty struct
			dst.ContentKitRadio = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitRadio = nil
	}

	// try to unmarshal data into ContentKitSelect
	err = newStrictDecoder(data).Decode(&dst.ContentKitSelect)
	if err == nil {
		jsonContentKitSelect, _ := json.Marshal(dst.ContentKitSelect)
		if string(jsonContentKitSelect) == "{}" { // empty struct
			dst.ContentKitSelect = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitSelect = nil
	}

	// try to unmarshal data into ContentKitSwitch
	err = newStrictDecoder(data).Decode(&dst.ContentKitSwitch)
	if err == nil {
		jsonContentKitSwitch, _ := json.Marshal(dst.ContentKitSwitch)
		if string(jsonContentKitSwitch) == "{}" { // empty struct
			dst.ContentKitSwitch = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitSwitch = nil
	}

	// try to unmarshal data into ContentKitTextInput
	err = newStrictDecoder(data).Decode(&dst.ContentKitTextInput)
	if err == nil {
		jsonContentKitTextInput, _ := json.Marshal(dst.ContentKitTextInput)
		if string(jsonContentKitTextInput) == "{}" { // empty struct
			dst.ContentKitTextInput = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitTextInput = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContentKitButton = nil
		dst.ContentKitCheckbox = nil
		dst.ContentKitRadio = nil
		dst.ContentKitSelect = nil
		dst.ContentKitSwitch = nil
		dst.ContentKitTextInput = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitInputElement)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitInputElement)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitInputElement) MarshalJSON() ([]byte, error) {
	if src.ContentKitButton != nil {
		return json.Marshal(&src.ContentKitButton)
	}

	if src.ContentKitCheckbox != nil {
		return json.Marshal(&src.ContentKitCheckbox)
	}

	if src.ContentKitRadio != nil {
		return json.Marshal(&src.ContentKitRadio)
	}

	if src.ContentKitSelect != nil {
		return json.Marshal(&src.ContentKitSelect)
	}

	if src.ContentKitSwitch != nil {
		return json.Marshal(&src.ContentKitSwitch)
	}

	if src.ContentKitTextInput != nil {
		return json.Marshal(&src.ContentKitTextInput)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitInputElement) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentKitButton != nil {
		return obj.ContentKitButton
	}

	if obj.ContentKitCheckbox != nil {
		return obj.ContentKitCheckbox
	}

	if obj.ContentKitRadio != nil {
		return obj.ContentKitRadio
	}

	if obj.ContentKitSelect != nil {
		return obj.ContentKitSelect
	}

	if obj.ContentKitSwitch != nil {
		return obj.ContentKitSwitch
	}

	if obj.ContentKitTextInput != nil {
		return obj.ContentKitTextInput
	}

	// all schemas are nil
	return nil
}

type NullableContentKitInputElement struct {
	value *ContentKitInputElement
	isSet bool
}

func (v NullableContentKitInputElement) Get() *ContentKitInputElement {
	return v.value
}

func (v *NullableContentKitInputElement) Set(val *ContentKitInputElement) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitInputElement) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitInputElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitInputElement(val *ContentKitInputElement) *NullableContentKitInputElement {
	return &NullableContentKitInputElement{value: val, isSet: true}
}

func (v NullableContentKitInputElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitInputElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
