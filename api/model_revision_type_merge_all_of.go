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

// checks if the RevisionTypeMergeAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionTypeMergeAllOf{}

// RevisionTypeMergeAllOf struct for RevisionTypeMergeAllOf
type RevisionTypeMergeAllOf struct {
	// Revision created when merging a change request with primary.
	Type       string        `json:"type"`
	MergedFrom ChangeRequest `json:"mergedFrom"`
}

// NewRevisionTypeMergeAllOf instantiates a new RevisionTypeMergeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionTypeMergeAllOf(type_ string, mergedFrom ChangeRequest) *RevisionTypeMergeAllOf {
	this := RevisionTypeMergeAllOf{}
	this.Type = type_
	this.MergedFrom = mergedFrom
	return &this
}

// NewRevisionTypeMergeAllOfWithDefaults instantiates a new RevisionTypeMergeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionTypeMergeAllOfWithDefaults() *RevisionTypeMergeAllOf {
	this := RevisionTypeMergeAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *RevisionTypeMergeAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeMergeAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RevisionTypeMergeAllOf) SetType(v string) {
	o.Type = v
}

// GetMergedFrom returns the MergedFrom field value
func (o *RevisionTypeMergeAllOf) GetMergedFrom() ChangeRequest {
	if o == nil {
		var ret ChangeRequest
		return ret
	}

	return o.MergedFrom
}

// GetMergedFromOk returns a tuple with the MergedFrom field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeMergeAllOf) GetMergedFromOk() (*ChangeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergedFrom, true
}

// SetMergedFrom sets field value
func (o *RevisionTypeMergeAllOf) SetMergedFrom(v ChangeRequest) {
	o.MergedFrom = v
}

func (o RevisionTypeMergeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionTypeMergeAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["mergedFrom"] = o.MergedFrom
	return toSerialize, nil
}

type NullableRevisionTypeMergeAllOf struct {
	value *RevisionTypeMergeAllOf
	isSet bool
}

func (v NullableRevisionTypeMergeAllOf) Get() *RevisionTypeMergeAllOf {
	return v.value
}

func (v *NullableRevisionTypeMergeAllOf) Set(val *RevisionTypeMergeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionTypeMergeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionTypeMergeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionTypeMergeAllOf(val *RevisionTypeMergeAllOf) *NullableRevisionTypeMergeAllOf {
	return &NullableRevisionTypeMergeAllOf{value: val, isSet: true}
}

func (v NullableRevisionTypeMergeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionTypeMergeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}