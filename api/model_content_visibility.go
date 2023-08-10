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

// ContentVisibility * `public`: Anyone can access the content, and the content is indexed by search engines. * `unlisted`: Anyone can access the content, and the content is not indexed by search engines * `share-link`: Anyone with a secret token in the url can access the content. * `visitor-auth`: Anyone authenticated through a JWT token can access the content. * `in-collection`: Anyone who can access the parent collection can access the content.   Only available for spaces in a collection. * `private`: Authorized members can access the content.
type ContentVisibility string

// List of ContentVisibility
const (
	CONTENTVISIBILITY_PUBLIC        ContentVisibility = "public"
	CONTENTVISIBILITY_UNLISTED      ContentVisibility = "unlisted"
	CONTENTVISIBILITY_SHARE_LINK    ContentVisibility = "share-link"
	CONTENTVISIBILITY_VISITOR_AUTH  ContentVisibility = "visitor-auth"
	CONTENTVISIBILITY_IN_COLLECTION ContentVisibility = "in-collection"
	CONTENTVISIBILITY_PRIVATE       ContentVisibility = "private"
)

// All allowed values of ContentVisibility enum
var AllowedContentVisibilityEnumValues = []ContentVisibility{
	"public",
	"unlisted",
	"share-link",
	"visitor-auth",
	"in-collection",
	"private",
}

func (v *ContentVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContentVisibility(value)
	for _, existing := range AllowedContentVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContentVisibility", value)
}

// NewContentVisibilityFromValue returns a pointer to a valid ContentVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContentVisibilityFromValue(v string) (*ContentVisibility, error) {
	ev := ContentVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContentVisibility: valid values are %v", v, AllowedContentVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContentVisibility) IsValid() bool {
	for _, existing := range AllowedContentVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContentVisibility value
func (v ContentVisibility) Ptr() *ContentVisibility {
	return &v
}

type NullableContentVisibility struct {
	value *ContentVisibility
	isSet bool
}

func (v NullableContentVisibility) Get() *ContentVisibility {
	return v.value
}

func (v *NullableContentVisibility) Set(val *ContentVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableContentVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableContentVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentVisibility(val *ContentVisibility) *NullableContentVisibility {
	return &NullableContentVisibility{value: val, isSet: true}
}

func (v NullableContentVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
