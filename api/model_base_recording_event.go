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

// checks if the BaseRecordingEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseRecordingEvent{}

// BaseRecordingEvent struct for BaseRecordingEvent
type BaseRecordingEvent struct {
	// The type of event
	Type string `json:"type"`
	// When the event happened
	Timestamp time.Time `json:"timestamp"`
	// Optionally, provide the source of the event. GitBook may use this to improve the generated content.
	Source *string                  `json:"source,omitempty"`
	Actor  *BaseRecordingEventActor `json:"actor,omitempty"`
}

// NewBaseRecordingEvent instantiates a new BaseRecordingEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseRecordingEvent(type_ string, timestamp time.Time) *BaseRecordingEvent {
	this := BaseRecordingEvent{}
	this.Type = type_
	this.Timestamp = timestamp
	return &this
}

// NewBaseRecordingEventWithDefaults instantiates a new BaseRecordingEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseRecordingEventWithDefaults() *BaseRecordingEvent {
	this := BaseRecordingEvent{}
	return &this
}

// GetType returns the Type field value
func (o *BaseRecordingEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseRecordingEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseRecordingEvent) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *BaseRecordingEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BaseRecordingEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BaseRecordingEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BaseRecordingEvent) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseRecordingEvent) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BaseRecordingEvent) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *BaseRecordingEvent) SetSource(v string) {
	o.Source = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *BaseRecordingEvent) GetActor() BaseRecordingEventActor {
	if o == nil || IsNil(o.Actor) {
		var ret BaseRecordingEventActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseRecordingEvent) GetActorOk() (*BaseRecordingEventActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *BaseRecordingEvent) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given BaseRecordingEventActor and assigns it to the Actor field.
func (o *BaseRecordingEvent) SetActor(v BaseRecordingEventActor) {
	o.Actor = &v
}

func (o BaseRecordingEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseRecordingEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	return toSerialize, nil
}

type NullableBaseRecordingEvent struct {
	value *BaseRecordingEvent
	isSet bool
}

func (v NullableBaseRecordingEvent) Get() *BaseRecordingEvent {
	return v.value
}

func (v *NullableBaseRecordingEvent) Set(val *BaseRecordingEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseRecordingEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseRecordingEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseRecordingEvent(val *BaseRecordingEvent) *NullableBaseRecordingEvent {
	return &NullableBaseRecordingEvent{value: val, isSet: true}
}

func (v NullableBaseRecordingEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseRecordingEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
