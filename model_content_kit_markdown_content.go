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

// ContentKitMarkdownContent - struct for ContentKitMarkdownContent
type ContentKitMarkdownContent struct {
	ContentKitDynamicBinding *ContentKitDynamicBinding
	String                   *string
}

// ContentKitDynamicBindingAsContentKitMarkdownContent is a convenience function that returns ContentKitDynamicBinding wrapped in ContentKitMarkdownContent
func ContentKitDynamicBindingAsContentKitMarkdownContent(v *ContentKitDynamicBinding) ContentKitMarkdownContent {
	return ContentKitMarkdownContent{
		ContentKitDynamicBinding: v,
	}
}

// stringAsContentKitMarkdownContent is a convenience function that returns string wrapped in ContentKitMarkdownContent
func StringAsContentKitMarkdownContent(v *string) ContentKitMarkdownContent {
	return ContentKitMarkdownContent{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitMarkdownContent) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitMarkdownContent)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitMarkdownContent)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitMarkdownContent) MarshalJSON() ([]byte, error) {
	if src.ContentKitDynamicBinding != nil {
		return json.Marshal(&src.ContentKitDynamicBinding)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitMarkdownContent) GetActualInstance() interface{} {
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

type NullableContentKitMarkdownContent struct {
	value *ContentKitMarkdownContent
	isSet bool
}

func (v NullableContentKitMarkdownContent) Get() *ContentKitMarkdownContent {
	return v.value
}

func (v *NullableContentKitMarkdownContent) Set(val *ContentKitMarkdownContent) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitMarkdownContent) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitMarkdownContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitMarkdownContent(val *ContentKitMarkdownContent) *NullableContentKitMarkdownContent {
	return &NullableContentKitMarkdownContent{value: val, isSet: true}
}

func (v NullableContentKitMarkdownContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitMarkdownContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
