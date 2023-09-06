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
)

// checks if the SpaceGitSyncStartedEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceGitSyncStartedEventAllOf{}

// SpaceGitSyncStartedEventAllOf Event when a GitSync operation has been started.
type SpaceGitSyncStartedEventAllOf struct {
	Type string `json:"type"`
	// Unique identifier of the new content revision
	RevisionId string `json:"revisionId"`
	// Unique identifier for the commit (sha)
	CommitId string `json:"commitId"`
}

// NewSpaceGitSyncStartedEventAllOf instantiates a new SpaceGitSyncStartedEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceGitSyncStartedEventAllOf(type_ string, revisionId string, commitId string) *SpaceGitSyncStartedEventAllOf {
	this := SpaceGitSyncStartedEventAllOf{}
	this.Type = type_
	this.RevisionId = revisionId
	this.CommitId = commitId
	return &this
}

// NewSpaceGitSyncStartedEventAllOfWithDefaults instantiates a new SpaceGitSyncStartedEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceGitSyncStartedEventAllOfWithDefaults() *SpaceGitSyncStartedEventAllOf {
	this := SpaceGitSyncStartedEventAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *SpaceGitSyncStartedEventAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncStartedEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceGitSyncStartedEventAllOf) SetType(v string) {
	o.Type = v
}

// GetRevisionId returns the RevisionId field value
func (o *SpaceGitSyncStartedEventAllOf) GetRevisionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncStartedEventAllOf) GetRevisionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionId, true
}

// SetRevisionId sets field value
func (o *SpaceGitSyncStartedEventAllOf) SetRevisionId(v string) {
	o.RevisionId = v
}

// GetCommitId returns the CommitId field value
func (o *SpaceGitSyncStartedEventAllOf) GetCommitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitId
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncStartedEventAllOf) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitId, true
}

// SetCommitId sets field value
func (o *SpaceGitSyncStartedEventAllOf) SetCommitId(v string) {
	o.CommitId = v
}

func (o SpaceGitSyncStartedEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceGitSyncStartedEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["revisionId"] = o.RevisionId
	toSerialize["commitId"] = o.CommitId
	return toSerialize, nil
}

type NullableSpaceGitSyncStartedEventAllOf struct {
	value *SpaceGitSyncStartedEventAllOf
	isSet bool
}

func (v NullableSpaceGitSyncStartedEventAllOf) Get() *SpaceGitSyncStartedEventAllOf {
	return v.value
}

func (v *NullableSpaceGitSyncStartedEventAllOf) Set(val *SpaceGitSyncStartedEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceGitSyncStartedEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceGitSyncStartedEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceGitSyncStartedEventAllOf(val *SpaceGitSyncStartedEventAllOf) *NullableSpaceGitSyncStartedEventAllOf {
	return &NullableSpaceGitSyncStartedEventAllOf{value: val, isSet: true}
}

func (v NullableSpaceGitSyncStartedEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceGitSyncStartedEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
