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

// checks if the RecordingFileRemovedEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingFileRemovedEventAllOf{}

// RecordingFileRemovedEventAllOf struct for RecordingFileRemovedEventAllOf
type RecordingFileRemovedEventAllOf struct {
	Type         string `json:"type"`
	Filename     string `json:"filename"`
	FileSnapshot string `json:"fileSnapshot"`
}

// NewRecordingFileRemovedEventAllOf instantiates a new RecordingFileRemovedEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingFileRemovedEventAllOf(type_ string, filename string, fileSnapshot string) *RecordingFileRemovedEventAllOf {
	this := RecordingFileRemovedEventAllOf{}
	this.Type = type_
	this.Filename = filename
	this.FileSnapshot = fileSnapshot
	return &this
}

// NewRecordingFileRemovedEventAllOfWithDefaults instantiates a new RecordingFileRemovedEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingFileRemovedEventAllOfWithDefaults() *RecordingFileRemovedEventAllOf {
	this := RecordingFileRemovedEventAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *RecordingFileRemovedEventAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecordingFileRemovedEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecordingFileRemovedEventAllOf) SetType(v string) {
	o.Type = v
}

// GetFilename returns the Filename field value
func (o *RecordingFileRemovedEventAllOf) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *RecordingFileRemovedEventAllOf) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *RecordingFileRemovedEventAllOf) SetFilename(v string) {
	o.Filename = v
}

// GetFileSnapshot returns the FileSnapshot field value
func (o *RecordingFileRemovedEventAllOf) GetFileSnapshot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileSnapshot
}

// GetFileSnapshotOk returns a tuple with the FileSnapshot field value
// and a boolean to check if the value has been set.
func (o *RecordingFileRemovedEventAllOf) GetFileSnapshotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSnapshot, true
}

// SetFileSnapshot sets field value
func (o *RecordingFileRemovedEventAllOf) SetFileSnapshot(v string) {
	o.FileSnapshot = v
}

func (o RecordingFileRemovedEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingFileRemovedEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["filename"] = o.Filename
	toSerialize["fileSnapshot"] = o.FileSnapshot
	return toSerialize, nil
}

type NullableRecordingFileRemovedEventAllOf struct {
	value *RecordingFileRemovedEventAllOf
	isSet bool
}

func (v NullableRecordingFileRemovedEventAllOf) Get() *RecordingFileRemovedEventAllOf {
	return v.value
}

func (v *NullableRecordingFileRemovedEventAllOf) Set(val *RecordingFileRemovedEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingFileRemovedEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingFileRemovedEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingFileRemovedEventAllOf(val *RecordingFileRemovedEventAllOf) *NullableRecordingFileRemovedEventAllOf {
	return &NullableRecordingFileRemovedEventAllOf{value: val, isSet: true}
}

func (v NullableRecordingFileRemovedEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingFileRemovedEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
