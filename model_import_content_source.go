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

// ImportContentSource the model 'ImportContentSource'
type ImportContentSource string

// List of ImportContentSource
const (
	IMPORTCONTENTSOURCE_WEBSITE       ImportContentSource = "website"
	IMPORTCONTENTSOURCE_DOCX          ImportContentSource = "docx"
	IMPORTCONTENTSOURCE_MARKDOWN      ImportContentSource = "markdown"
	IMPORTCONTENTSOURCE_HTML          ImportContentSource = "html"
	IMPORTCONTENTSOURCE_ZIP           ImportContentSource = "zip"
	IMPORTCONTENTSOURCE_CONFLUENCE    ImportContentSource = "confluence"
	IMPORTCONTENTSOURCE_GITHUB_WIKI   ImportContentSource = "github-wiki"
	IMPORTCONTENTSOURCE_DROPBOX_PAPER ImportContentSource = "dropbox-paper"
	IMPORTCONTENTSOURCE_NOTION        ImportContentSource = "notion"
	IMPORTCONTENTSOURCE_QUIP          ImportContentSource = "quip"
	IMPORTCONTENTSOURCE_GOOGLE_DOCS   ImportContentSource = "google-docs"
	IMPORTCONTENTSOURCE_OPEN_API      ImportContentSource = "open-api"
)

// All allowed values of ImportContentSource enum
var AllowedImportContentSourceEnumValues = []ImportContentSource{
	"website",
	"docx",
	"markdown",
	"html",
	"zip",
	"confluence",
	"github-wiki",
	"dropbox-paper",
	"notion",
	"quip",
	"google-docs",
	"open-api",
}

func (v *ImportContentSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImportContentSource(value)
	for _, existing := range AllowedImportContentSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImportContentSource", value)
}

// NewImportContentSourceFromValue returns a pointer to a valid ImportContentSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImportContentSourceFromValue(v string) (*ImportContentSource, error) {
	ev := ImportContentSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImportContentSource: valid values are %v", v, AllowedImportContentSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImportContentSource) IsValid() bool {
	for _, existing := range AllowedImportContentSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportContentSource value
func (v ImportContentSource) Ptr() *ImportContentSource {
	return &v
}

type NullableImportContentSource struct {
	value *ImportContentSource
	isSet bool
}

func (v NullableImportContentSource) Get() *ImportContentSource {
	return v.value
}

func (v *NullableImportContentSource) Set(val *ImportContentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportContentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportContentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportContentSource(val *ImportContentSource) *NullableImportContentSource {
	return &NullableImportContentSource{value: val, isSet: true}
}

func (v NullableImportContentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportContentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
