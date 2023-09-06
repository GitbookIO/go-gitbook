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

// checks if the RecordingSpeechEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingSpeechEvent{}

// RecordingSpeechEvent struct for RecordingSpeechEvent
type RecordingSpeechEvent struct {
	Type string `json:"type"`
	// When the event happened
	Timestamp time.Time `json:"timestamp"`
	// Optionally, provide the source of the event. GitBook may use this to improve the generated content.
	Source *string                  `json:"source,omitempty"`
	Actor  *BaseRecordingEventActor `json:"actor,omitempty"`
	// WAV audio file, encoded as base64
	Audio string `json:"audio"`
	// Transcript of the speech
	Transcript string `json:"transcript"`
}

// NewRecordingSpeechEvent instantiates a new RecordingSpeechEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingSpeechEvent(type_ string, timestamp time.Time, audio string, transcript string) *RecordingSpeechEvent {
	this := RecordingSpeechEvent{}
	this.Type = type_
	this.Timestamp = timestamp
	this.Audio = audio
	this.Transcript = transcript
	return &this
}

// NewRecordingSpeechEventWithDefaults instantiates a new RecordingSpeechEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingSpeechEventWithDefaults() *RecordingSpeechEvent {
	this := RecordingSpeechEvent{}
	return &this
}

// GetType returns the Type field value
func (o *RecordingSpeechEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecordingSpeechEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecordingSpeechEvent) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *RecordingSpeechEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *RecordingSpeechEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *RecordingSpeechEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RecordingSpeechEvent) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingSpeechEvent) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RecordingSpeechEvent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RecordingSpeechEvent) SetSource(v string) {
	o.Source = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *RecordingSpeechEvent) GetActor() BaseRecordingEventActor {
	if o == nil || IsNil(o.Actor) {
		var ret BaseRecordingEventActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingSpeechEvent) GetActorOk() (*BaseRecordingEventActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *RecordingSpeechEvent) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given BaseRecordingEventActor and assigns it to the Actor field.
func (o *RecordingSpeechEvent) SetActor(v BaseRecordingEventActor) {
	o.Actor = &v
}

// GetAudio returns the Audio field value
func (o *RecordingSpeechEvent) GetAudio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audio
}

// GetAudioOk returns a tuple with the Audio field value
// and a boolean to check if the value has been set.
func (o *RecordingSpeechEvent) GetAudioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audio, true
}

// SetAudio sets field value
func (o *RecordingSpeechEvent) SetAudio(v string) {
	o.Audio = v
}

// GetTranscript returns the Transcript field value
func (o *RecordingSpeechEvent) GetTranscript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transcript
}

// GetTranscriptOk returns a tuple with the Transcript field value
// and a boolean to check if the value has been set.
func (o *RecordingSpeechEvent) GetTranscriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transcript, true
}

// SetTranscript sets field value
func (o *RecordingSpeechEvent) SetTranscript(v string) {
	o.Transcript = v
}

func (o RecordingSpeechEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingSpeechEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	toSerialize["audio"] = o.Audio
	toSerialize["transcript"] = o.Transcript
	return toSerialize, nil
}

type NullableRecordingSpeechEvent struct {
	value *RecordingSpeechEvent
	isSet bool
}

func (v NullableRecordingSpeechEvent) Get() *RecordingSpeechEvent {
	return v.value
}

func (v *NullableRecordingSpeechEvent) Set(val *RecordingSpeechEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingSpeechEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingSpeechEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingSpeechEvent(val *RecordingSpeechEvent) *NullableRecordingSpeechEvent {
	return &NullableRecordingSpeechEvent{value: val, isSet: true}
}

func (v NullableRecordingSpeechEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingSpeechEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
