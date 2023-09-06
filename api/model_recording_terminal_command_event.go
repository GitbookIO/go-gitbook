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
	"time"
)

// checks if the RecordingTerminalCommandEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingTerminalCommandEvent{}

// RecordingTerminalCommandEvent struct for RecordingTerminalCommandEvent
type RecordingTerminalCommandEvent struct {
	Type string `json:"type"`
	// When the event happened
	Timestamp time.Time `json:"timestamp"`
	// Optionally, provide the source of the event. GitBook may use this to improve the generated content.
	Source  *string                  `json:"source,omitempty"`
	Actor   *BaseRecordingEventActor `json:"actor,omitempty"`
	Command string                   `json:"command"`
	Stdout  string                   `json:"stdout"`
}

// NewRecordingTerminalCommandEvent instantiates a new RecordingTerminalCommandEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingTerminalCommandEvent(type_ string, timestamp time.Time, command string, stdout string) *RecordingTerminalCommandEvent {
	this := RecordingTerminalCommandEvent{}
	this.Type = type_
	this.Timestamp = timestamp
	this.Command = command
	this.Stdout = stdout
	return &this
}

// NewRecordingTerminalCommandEventWithDefaults instantiates a new RecordingTerminalCommandEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingTerminalCommandEventWithDefaults() *RecordingTerminalCommandEvent {
	this := RecordingTerminalCommandEvent{}
	return &this
}

// GetType returns the Type field value
func (o *RecordingTerminalCommandEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecordingTerminalCommandEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecordingTerminalCommandEvent) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *RecordingTerminalCommandEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *RecordingTerminalCommandEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *RecordingTerminalCommandEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RecordingTerminalCommandEvent) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingTerminalCommandEvent) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RecordingTerminalCommandEvent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RecordingTerminalCommandEvent) SetSource(v string) {
	o.Source = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *RecordingTerminalCommandEvent) GetActor() BaseRecordingEventActor {
	if o == nil || IsNil(o.Actor) {
		var ret BaseRecordingEventActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingTerminalCommandEvent) GetActorOk() (*BaseRecordingEventActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *RecordingTerminalCommandEvent) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given BaseRecordingEventActor and assigns it to the Actor field.
func (o *RecordingTerminalCommandEvent) SetActor(v BaseRecordingEventActor) {
	o.Actor = &v
}

// GetCommand returns the Command field value
func (o *RecordingTerminalCommandEvent) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *RecordingTerminalCommandEvent) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *RecordingTerminalCommandEvent) SetCommand(v string) {
	o.Command = v
}

// GetStdout returns the Stdout field value
func (o *RecordingTerminalCommandEvent) GetStdout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stdout
}

// GetStdoutOk returns a tuple with the Stdout field value
// and a boolean to check if the value has been set.
func (o *RecordingTerminalCommandEvent) GetStdoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stdout, true
}

// SetStdout sets field value
func (o *RecordingTerminalCommandEvent) SetStdout(v string) {
	o.Stdout = v
}

func (o RecordingTerminalCommandEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingTerminalCommandEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	toSerialize["command"] = o.Command
	toSerialize["stdout"] = o.Stdout
	return toSerialize, nil
}

type NullableRecordingTerminalCommandEvent struct {
	value *RecordingTerminalCommandEvent
	isSet bool
}

func (v NullableRecordingTerminalCommandEvent) Get() *RecordingTerminalCommandEvent {
	return v.value
}

func (v *NullableRecordingTerminalCommandEvent) Set(val *RecordingTerminalCommandEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingTerminalCommandEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingTerminalCommandEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingTerminalCommandEvent(val *RecordingTerminalCommandEvent) *NullableRecordingTerminalCommandEvent {
	return &NullableRecordingTerminalCommandEvent{value: val, isSet: true}
}

func (v NullableRecordingTerminalCommandEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingTerminalCommandEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}