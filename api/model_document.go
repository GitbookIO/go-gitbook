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

// Document - struct for Document
type Document struct {
	JSONDocument     *JSONDocument
	MarkdownDocument *MarkdownDocument
	Properties       *map[string]interface{}
}

// JSONDocumentAsDocument is a convenience function that returns JSONDocument wrapped in Document
func JSONDocumentAsDocument(v *JSONDocument) Document {
	return Document{
		JSONDocument: v,
	}
}

// MarkdownDocumentAsDocument is a convenience function that returns MarkdownDocument wrapped in Document
func MarkdownDocumentAsDocument(v *MarkdownDocument) Document {
	return Document{
		MarkdownDocument: v,
	}
}

// map[string]interface{}AsDocument is a convenience function that returns map[string]interface{} wrapped in Document
func PropertiesAsDocument(v *map[string]interface{}) Document {
	return Document{
		Properties: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Document) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into JSONDocument
	err = newStrictDecoder(data).Decode(&dst.JSONDocument)
	if err == nil {
		jsonJSONDocument, _ := json.Marshal(dst.JSONDocument)
		if string(jsonJSONDocument) == "{}" { // empty struct
			dst.JSONDocument = nil
		} else {
			match++
		}
	} else {
		dst.JSONDocument = nil
	}

	// try to unmarshal data into MarkdownDocument
	err = newStrictDecoder(data).Decode(&dst.MarkdownDocument)
	if err == nil {
		jsonMarkdownDocument, _ := json.Marshal(dst.MarkdownDocument)
		if string(jsonMarkdownDocument) == "{}" { // empty struct
			dst.MarkdownDocument = nil
		} else {
			match++
		}
	} else {
		dst.MarkdownDocument = nil
	}

	// try to unmarshal data into Properties
	err = newStrictDecoder(data).Decode(&dst.Properties)
	if err == nil {
		jsonProperties, _ := json.Marshal(dst.Properties)
		if string(jsonProperties) == "{}" { // empty struct
			dst.Properties = nil
		} else {
			match++
		}
	} else {
		dst.Properties = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.JSONDocument = nil
		dst.MarkdownDocument = nil
		dst.Properties = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Document)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Document)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Document) MarshalJSON() ([]byte, error) {
	if src.JSONDocument != nil {
		return json.Marshal(&src.JSONDocument)
	}

	if src.MarkdownDocument != nil {
		return json.Marshal(&src.MarkdownDocument)
	}

	if src.Properties != nil {
		return json.Marshal(&src.Properties)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Document) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.JSONDocument != nil {
		return obj.JSONDocument
	}

	if obj.MarkdownDocument != nil {
		return obj.MarkdownDocument
	}

	if obj.Properties != nil {
		return obj.Properties
	}

	// all schemas are nil
	return nil
}

type NullableDocument struct {
	value *Document
	isSet bool
}

func (v NullableDocument) Get() *Document {
	return v.value
}

func (v *NullableDocument) Set(val *Document) {
	v.value = val
	v.isSet = true
}

func (v NullableDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocument(val *Document) *NullableDocument {
	return &NullableDocument{value: val, isSet: true}
}

func (v NullableDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
