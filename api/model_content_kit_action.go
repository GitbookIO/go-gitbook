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

// ContentKitAction struct for ContentKitAction
type ContentKitAction struct {
	ContentKitActionAnyOf   *ContentKitActionAnyOf
	ContentKitDefaultAction *ContentKitDefaultAction
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ContentKitAction) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ContentKitActionAnyOf
	err = json.Unmarshal(data, &dst.ContentKitActionAnyOf)
	if err == nil {
		jsonContentKitActionAnyOf, _ := json.Marshal(dst.ContentKitActionAnyOf)
		if string(jsonContentKitActionAnyOf) == "{}" { // empty struct
			dst.ContentKitActionAnyOf = nil
		} else {
			return nil // data stored in dst.ContentKitActionAnyOf, return on the first match
		}
	} else {
		dst.ContentKitActionAnyOf = nil
	}

	// try to unmarshal JSON data into ContentKitDefaultAction
	err = json.Unmarshal(data, &dst.ContentKitDefaultAction)
	if err == nil {
		jsonContentKitDefaultAction, _ := json.Marshal(dst.ContentKitDefaultAction)
		if string(jsonContentKitDefaultAction) == "{}" { // empty struct
			dst.ContentKitDefaultAction = nil
		} else {
			return nil // data stored in dst.ContentKitDefaultAction, return on the first match
		}
	} else {
		dst.ContentKitDefaultAction = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ContentKitAction)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ContentKitAction) MarshalJSON() ([]byte, error) {
	if src.ContentKitActionAnyOf != nil {
		return json.Marshal(&src.ContentKitActionAnyOf)
	}

	if src.ContentKitDefaultAction != nil {
		return json.Marshal(&src.ContentKitDefaultAction)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableContentKitAction struct {
	value *ContentKitAction
	isSet bool
}

func (v NullableContentKitAction) Get() *ContentKitAction {
	return v.value
}

func (v *NullableContentKitAction) Set(val *ContentKitAction) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitAction) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitAction(val *ContentKitAction) *NullableContentKitAction {
	return &NullableContentKitAction{value: val, isSet: true}
}

func (v NullableContentKitAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
