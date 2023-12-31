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

// checks if the RecordingThreadMessageEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingThreadMessageEvent{}

// RecordingThreadMessageEvent struct for RecordingThreadMessageEvent
type RecordingThreadMessageEvent struct {
	Type string `json:"type"`
	// When the event happened
	Timestamp time.Time `json:"timestamp"`
	// Optionally, provide the source of the event. GitBook may use this to improve the generated content.
	Source  *string                  `json:"source,omitempty"`
	Actor   *BaseRecordingEventActor `json:"actor,omitempty"`
	IsFirst *bool                    `json:"isFirst,omitempty"`
	Text    string                   `json:"text"`
}

// NewRecordingThreadMessageEvent instantiates a new RecordingThreadMessageEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingThreadMessageEvent(type_ string, timestamp time.Time, text string) *RecordingThreadMessageEvent {
	this := RecordingThreadMessageEvent{}
	this.Type = type_
	this.Timestamp = timestamp
	this.Text = text
	return &this
}

// NewRecordingThreadMessageEventWithDefaults instantiates a new RecordingThreadMessageEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingThreadMessageEventWithDefaults() *RecordingThreadMessageEvent {
	this := RecordingThreadMessageEvent{}
	return &this
}

// GetType returns the Type field value
func (o *RecordingThreadMessageEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecordingThreadMessageEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecordingThreadMessageEvent) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *RecordingThreadMessageEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *RecordingThreadMessageEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *RecordingThreadMessageEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RecordingThreadMessageEvent) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingThreadMessageEvent) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RecordingThreadMessageEvent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RecordingThreadMessageEvent) SetSource(v string) {
	o.Source = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *RecordingThreadMessageEvent) GetActor() BaseRecordingEventActor {
	if o == nil || IsNil(o.Actor) {
		var ret BaseRecordingEventActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingThreadMessageEvent) GetActorOk() (*BaseRecordingEventActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *RecordingThreadMessageEvent) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given BaseRecordingEventActor and assigns it to the Actor field.
func (o *RecordingThreadMessageEvent) SetActor(v BaseRecordingEventActor) {
	o.Actor = &v
}

// GetIsFirst returns the IsFirst field value if set, zero value otherwise.
func (o *RecordingThreadMessageEvent) GetIsFirst() bool {
	if o == nil || IsNil(o.IsFirst) {
		var ret bool
		return ret
	}
	return *o.IsFirst
}

// GetIsFirstOk returns a tuple with the IsFirst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingThreadMessageEvent) GetIsFirstOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFirst) {
		return nil, false
	}
	return o.IsFirst, true
}

// HasIsFirst returns a boolean if a field has been set.
func (o *RecordingThreadMessageEvent) HasIsFirst() bool {
	if o != nil && !IsNil(o.IsFirst) {
		return true
	}

	return false
}

// SetIsFirst gets a reference to the given bool and assigns it to the IsFirst field.
func (o *RecordingThreadMessageEvent) SetIsFirst(v bool) {
	o.IsFirst = &v
}

// GetText returns the Text field value
func (o *RecordingThreadMessageEvent) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *RecordingThreadMessageEvent) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *RecordingThreadMessageEvent) SetText(v string) {
	o.Text = v
}

func (o RecordingThreadMessageEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingThreadMessageEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.IsFirst) {
		toSerialize["isFirst"] = o.IsFirst
	}
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableRecordingThreadMessageEvent struct {
	value *RecordingThreadMessageEvent
	isSet bool
}

func (v NullableRecordingThreadMessageEvent) Get() *RecordingThreadMessageEvent {
	return v.value
}

func (v *NullableRecordingThreadMessageEvent) Set(val *RecordingThreadMessageEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingThreadMessageEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingThreadMessageEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingThreadMessageEvent(val *RecordingThreadMessageEvent) *NullableRecordingThreadMessageEvent {
	return &NullableRecordingThreadMessageEvent{value: val, isSet: true}
}

func (v NullableRecordingThreadMessageEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingThreadMessageEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
