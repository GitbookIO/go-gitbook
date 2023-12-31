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
	"time"
)

// checks if the RecordingFileAddedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingFileAddedEvent{}

// RecordingFileAddedEvent struct for RecordingFileAddedEvent
type RecordingFileAddedEvent struct {
	Type string `json:"type"`
	// When the event happened
	Timestamp time.Time `json:"timestamp"`
	// Optionally, provide the source of the event. GitBook may use this to improve the generated content.
	Source       *string                  `json:"source,omitempty"`
	Actor        *BaseRecordingEventActor `json:"actor,omitempty"`
	Filename     string                   `json:"filename"`
	FileSnapshot string                   `json:"fileSnapshot"`
}

// NewRecordingFileAddedEvent instantiates a new RecordingFileAddedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingFileAddedEvent(type_ string, timestamp time.Time, filename string, fileSnapshot string) *RecordingFileAddedEvent {
	this := RecordingFileAddedEvent{}
	this.Type = type_
	this.Timestamp = timestamp
	this.Filename = filename
	this.FileSnapshot = fileSnapshot
	return &this
}

// NewRecordingFileAddedEventWithDefaults instantiates a new RecordingFileAddedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingFileAddedEventWithDefaults() *RecordingFileAddedEvent {
	this := RecordingFileAddedEvent{}
	return &this
}

// GetType returns the Type field value
func (o *RecordingFileAddedEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecordingFileAddedEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecordingFileAddedEvent) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *RecordingFileAddedEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *RecordingFileAddedEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *RecordingFileAddedEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RecordingFileAddedEvent) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFileAddedEvent) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RecordingFileAddedEvent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RecordingFileAddedEvent) SetSource(v string) {
	o.Source = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *RecordingFileAddedEvent) GetActor() BaseRecordingEventActor {
	if o == nil || IsNil(o.Actor) {
		var ret BaseRecordingEventActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingFileAddedEvent) GetActorOk() (*BaseRecordingEventActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *RecordingFileAddedEvent) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given BaseRecordingEventActor and assigns it to the Actor field.
func (o *RecordingFileAddedEvent) SetActor(v BaseRecordingEventActor) {
	o.Actor = &v
}

// GetFilename returns the Filename field value
func (o *RecordingFileAddedEvent) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *RecordingFileAddedEvent) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *RecordingFileAddedEvent) SetFilename(v string) {
	o.Filename = v
}

// GetFileSnapshot returns the FileSnapshot field value
func (o *RecordingFileAddedEvent) GetFileSnapshot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileSnapshot
}

// GetFileSnapshotOk returns a tuple with the FileSnapshot field value
// and a boolean to check if the value has been set.
func (o *RecordingFileAddedEvent) GetFileSnapshotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSnapshot, true
}

// SetFileSnapshot sets field value
func (o *RecordingFileAddedEvent) SetFileSnapshot(v string) {
	o.FileSnapshot = v
}

func (o RecordingFileAddedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingFileAddedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	toSerialize["filename"] = o.Filename
	toSerialize["fileSnapshot"] = o.FileSnapshot
	return toSerialize, nil
}

type NullableRecordingFileAddedEvent struct {
	value *RecordingFileAddedEvent
	isSet bool
}

func (v NullableRecordingFileAddedEvent) Get() *RecordingFileAddedEvent {
	return v.value
}

func (v *NullableRecordingFileAddedEvent) Set(val *RecordingFileAddedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingFileAddedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingFileAddedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingFileAddedEvent(val *RecordingFileAddedEvent) *NullableRecordingFileAddedEvent {
	return &NullableRecordingFileAddedEvent{value: val, isSet: true}
}

func (v NullableRecordingFileAddedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingFileAddedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
