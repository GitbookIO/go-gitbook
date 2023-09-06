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

// checks if the RecordingEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingEvents{}

// RecordingEvents Count of events recorded.
type RecordingEvents struct {
	TerminalCommand *int32 `json:"terminal.command,omitempty"`
	Speech          *int32 `json:"speech,omitempty"`
	ThreadMessage   *int32 `json:"thread.message,omitempty"`
}

// NewRecordingEvents instantiates a new RecordingEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingEvents() *RecordingEvents {
	this := RecordingEvents{}
	return &this
}

// NewRecordingEventsWithDefaults instantiates a new RecordingEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingEventsWithDefaults() *RecordingEvents {
	this := RecordingEvents{}
	return &this
}

// GetTerminalCommand returns the TerminalCommand field value if set, zero value otherwise.
func (o *RecordingEvents) GetTerminalCommand() int32 {
	if o == nil || IsNil(o.TerminalCommand) {
		var ret int32
		return ret
	}
	return *o.TerminalCommand
}

// GetTerminalCommandOk returns a tuple with the TerminalCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingEvents) GetTerminalCommandOk() (*int32, bool) {
	if o == nil || IsNil(o.TerminalCommand) {
		return nil, false
	}
	return o.TerminalCommand, true
}

// HasTerminalCommand returns a boolean if a field has been set.
func (o *RecordingEvents) HasTerminalCommand() bool {
	if o != nil && !IsNil(o.TerminalCommand) {
		return true
	}

	return false
}

// SetTerminalCommand gets a reference to the given int32 and assigns it to the TerminalCommand field.
func (o *RecordingEvents) SetTerminalCommand(v int32) {
	o.TerminalCommand = &v
}

// GetSpeech returns the Speech field value if set, zero value otherwise.
func (o *RecordingEvents) GetSpeech() int32 {
	if o == nil || IsNil(o.Speech) {
		var ret int32
		return ret
	}
	return *o.Speech
}

// GetSpeechOk returns a tuple with the Speech field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingEvents) GetSpeechOk() (*int32, bool) {
	if o == nil || IsNil(o.Speech) {
		return nil, false
	}
	return o.Speech, true
}

// HasSpeech returns a boolean if a field has been set.
func (o *RecordingEvents) HasSpeech() bool {
	if o != nil && !IsNil(o.Speech) {
		return true
	}

	return false
}

// SetSpeech gets a reference to the given int32 and assigns it to the Speech field.
func (o *RecordingEvents) SetSpeech(v int32) {
	o.Speech = &v
}

// GetThreadMessage returns the ThreadMessage field value if set, zero value otherwise.
func (o *RecordingEvents) GetThreadMessage() int32 {
	if o == nil || IsNil(o.ThreadMessage) {
		var ret int32
		return ret
	}
	return *o.ThreadMessage
}

// GetThreadMessageOk returns a tuple with the ThreadMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingEvents) GetThreadMessageOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreadMessage) {
		return nil, false
	}
	return o.ThreadMessage, true
}

// HasThreadMessage returns a boolean if a field has been set.
func (o *RecordingEvents) HasThreadMessage() bool {
	if o != nil && !IsNil(o.ThreadMessage) {
		return true
	}

	return false
}

// SetThreadMessage gets a reference to the given int32 and assigns it to the ThreadMessage field.
func (o *RecordingEvents) SetThreadMessage(v int32) {
	o.ThreadMessage = &v
}

func (o RecordingEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TerminalCommand) {
		toSerialize["terminal.command"] = o.TerminalCommand
	}
	if !IsNil(o.Speech) {
		toSerialize["speech"] = o.Speech
	}
	if !IsNil(o.ThreadMessage) {
		toSerialize["thread.message"] = o.ThreadMessage
	}
	return toSerialize, nil
}

type NullableRecordingEvents struct {
	value *RecordingEvents
	isSet bool
}

func (v NullableRecordingEvents) Get() *RecordingEvents {
	return v.value
}

func (v *NullableRecordingEvents) Set(val *RecordingEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingEvents(val *RecordingEvents) *NullableRecordingEvents {
	return &NullableRecordingEvents{value: val, isSet: true}
}

func (v NullableRecordingEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
