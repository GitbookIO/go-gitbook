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
)

// checks if the SpaceContentUpdatedEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceContentUpdatedEventAllOf{}

// SpaceContentUpdatedEventAllOf Event when the primary content of a space has been updated.
type SpaceContentUpdatedEventAllOf struct {
	Type string `json:"type"`
	// Unique identifier of the new content revision
	RevisionId string `json:"revisionId"`
}

// NewSpaceContentUpdatedEventAllOf instantiates a new SpaceContentUpdatedEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceContentUpdatedEventAllOf(type_ string, revisionId string) *SpaceContentUpdatedEventAllOf {
	this := SpaceContentUpdatedEventAllOf{}
	this.Type = type_
	this.RevisionId = revisionId
	return &this
}

// NewSpaceContentUpdatedEventAllOfWithDefaults instantiates a new SpaceContentUpdatedEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceContentUpdatedEventAllOfWithDefaults() *SpaceContentUpdatedEventAllOf {
	this := SpaceContentUpdatedEventAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *SpaceContentUpdatedEventAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceContentUpdatedEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceContentUpdatedEventAllOf) SetType(v string) {
	o.Type = v
}

// GetRevisionId returns the RevisionId field value
func (o *SpaceContentUpdatedEventAllOf) GetRevisionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value
// and a boolean to check if the value has been set.
func (o *SpaceContentUpdatedEventAllOf) GetRevisionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionId, true
}

// SetRevisionId sets field value
func (o *SpaceContentUpdatedEventAllOf) SetRevisionId(v string) {
	o.RevisionId = v
}

func (o SpaceContentUpdatedEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceContentUpdatedEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["revisionId"] = o.RevisionId
	return toSerialize, nil
}

type NullableSpaceContentUpdatedEventAllOf struct {
	value *SpaceContentUpdatedEventAllOf
	isSet bool
}

func (v NullableSpaceContentUpdatedEventAllOf) Get() *SpaceContentUpdatedEventAllOf {
	return v.value
}

func (v *NullableSpaceContentUpdatedEventAllOf) Set(val *SpaceContentUpdatedEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceContentUpdatedEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceContentUpdatedEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceContentUpdatedEventAllOf(val *SpaceContentUpdatedEventAllOf) *NullableSpaceContentUpdatedEventAllOf {
	return &NullableSpaceContentUpdatedEventAllOf{value: val, isSet: true}
}

func (v NullableSpaceContentUpdatedEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceContentUpdatedEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
