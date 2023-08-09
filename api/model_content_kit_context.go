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

// ContentKitContext - Object representing the context in which a ContentKit component is rendered.
type ContentKitContext struct {
	ContentKitContextConfigurationAccount *ContentKitContextConfigurationAccount
	ContentKitContextConfigurationSpace   *ContentKitContextConfigurationSpace
	ContentKitContextDocument             *ContentKitContextDocument
}

// ContentKitContextConfigurationAccountAsContentKitContext is a convenience function that returns ContentKitContextConfigurationAccount wrapped in ContentKitContext
func ContentKitContextConfigurationAccountAsContentKitContext(v *ContentKitContextConfigurationAccount) ContentKitContext {
	return ContentKitContext{
		ContentKitContextConfigurationAccount: v,
	}
}

// ContentKitContextConfigurationSpaceAsContentKitContext is a convenience function that returns ContentKitContextConfigurationSpace wrapped in ContentKitContext
func ContentKitContextConfigurationSpaceAsContentKitContext(v *ContentKitContextConfigurationSpace) ContentKitContext {
	return ContentKitContext{
		ContentKitContextConfigurationSpace: v,
	}
}

// ContentKitContextDocumentAsContentKitContext is a convenience function that returns ContentKitContextDocument wrapped in ContentKitContext
func ContentKitContextDocumentAsContentKitContext(v *ContentKitContextDocument) ContentKitContext {
	return ContentKitContext{
		ContentKitContextDocument: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContentKitContext) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContentKitContextConfigurationAccount
	err = newStrictDecoder(data).Decode(&dst.ContentKitContextConfigurationAccount)
	if err == nil {
		jsonContentKitContextConfigurationAccount, _ := json.Marshal(dst.ContentKitContextConfigurationAccount)
		if string(jsonContentKitContextConfigurationAccount) == "{}" { // empty struct
			dst.ContentKitContextConfigurationAccount = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitContextConfigurationAccount = nil
	}

	// try to unmarshal data into ContentKitContextConfigurationSpace
	err = newStrictDecoder(data).Decode(&dst.ContentKitContextConfigurationSpace)
	if err == nil {
		jsonContentKitContextConfigurationSpace, _ := json.Marshal(dst.ContentKitContextConfigurationSpace)
		if string(jsonContentKitContextConfigurationSpace) == "{}" { // empty struct
			dst.ContentKitContextConfigurationSpace = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitContextConfigurationSpace = nil
	}

	// try to unmarshal data into ContentKitContextDocument
	err = newStrictDecoder(data).Decode(&dst.ContentKitContextDocument)
	if err == nil {
		jsonContentKitContextDocument, _ := json.Marshal(dst.ContentKitContextDocument)
		if string(jsonContentKitContextDocument) == "{}" { // empty struct
			dst.ContentKitContextDocument = nil
		} else {
			match++
		}
	} else {
		dst.ContentKitContextDocument = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContentKitContextConfigurationAccount = nil
		dst.ContentKitContextConfigurationSpace = nil
		dst.ContentKitContextDocument = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContentKitContext)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContentKitContext)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentKitContext) MarshalJSON() ([]byte, error) {
	if src.ContentKitContextConfigurationAccount != nil {
		return json.Marshal(&src.ContentKitContextConfigurationAccount)
	}

	if src.ContentKitContextConfigurationSpace != nil {
		return json.Marshal(&src.ContentKitContextConfigurationSpace)
	}

	if src.ContentKitContextDocument != nil {
		return json.Marshal(&src.ContentKitContextDocument)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContentKitContext) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ContentKitContextConfigurationAccount != nil {
		return obj.ContentKitContextConfigurationAccount
	}

	if obj.ContentKitContextConfigurationSpace != nil {
		return obj.ContentKitContextConfigurationSpace
	}

	if obj.ContentKitContextDocument != nil {
		return obj.ContentKitContextDocument
	}

	// all schemas are nil
	return nil
}

type NullableContentKitContext struct {
	value *ContentKitContext
	isSet bool
}

func (v NullableContentKitContext) Get() *ContentKitContext {
	return v.value
}

func (v *NullableContentKitContext) Set(val *ContentKitContext) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitContext(val *ContentKitContext) *NullableContentKitContext {
	return &NullableContentKitContext{value: val, isSet: true}
}

func (v NullableContentKitContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}