/*
GitBook API

The GitBook API

API version: 0.0.1-beta
Contact: support@gitbook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
	"fmt"
)

// ContentKitCodeBlockLineNumbers - struct for ContentKitCodeBlockLineNumbers
type ContentKitCodeBlockLineNumbers struct {
	Bool    *bool
	Float32 *float32
}

// boolAsContentKitCodeBlockLineNumbers is a convenience function that returns bool wrapped in ContentKitCodeBlockLineNumbers
func BoolAsContentKitCodeBlockLineNumbers(v *bool) ContentKitCodeBlockLineNumbers {
	return ContentKitCodeBlockLineNumbers{
		Bool: v,
	}
}

// float32AsContentKitCodeBlockLineNumbers is a convenience function that returns float32 wrapped in ContentKitCodeBlockLineNumbers
func Float32AsContentKitCodeBlockLineNumbers(v *float32) ContentKitCodeBlockLineNumbers {
	return ContentKitCodeBlockLineNumbers{
		Float32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitCodeBlockLineNumbers) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.Float32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitCodeBlockLineNumbers)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitCodeBlockLineNumbers)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitCodeBlockLineNumbers) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitCodeBlockLineNumbers) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	// all schemas are nil
	return nil
}

type NullableContentKitCodeBlockLineNumbers struct {
	value *ContentKitCodeBlockLineNumbers
	isSet bool
}

func (v NullableContentKitCodeBlockLineNumbers) Get() *ContentKitCodeBlockLineNumbers {
	return v.value
}

func (v *NullableContentKitCodeBlockLineNumbers) Set(val *ContentKitCodeBlockLineNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitCodeBlockLineNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitCodeBlockLineNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitCodeBlockLineNumbers(val *ContentKitCodeBlockLineNumbers) *NullableContentKitCodeBlockLineNumbers {
	return &NullableContentKitCodeBlockLineNumbers{value: val, isSet: true}
}

func (v NullableContentKitCodeBlockLineNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitCodeBlockLineNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
