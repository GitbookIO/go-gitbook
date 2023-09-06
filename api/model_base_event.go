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

// checks if the BaseEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEvent{}

// BaseEvent Common properties for all events.
type BaseEvent struct {
	// Unique identifier for the event.
	EventId string `json:"eventId"`
	// Type of the event.
	Type string `json:"type"`
}

// NewBaseEvent instantiates a new BaseEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEvent(eventId string, type_ string) *BaseEvent {
	this := BaseEvent{}
	this.EventId = eventId
	this.Type = type_
	return &this
}

// NewBaseEventWithDefaults instantiates a new BaseEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventWithDefaults() *BaseEvent {
	this := BaseEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *BaseEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *BaseEvent) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *BaseEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseEvent) SetType(v string) {
	o.Type = v
}

func (o BaseEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBaseEvent struct {
	value *BaseEvent
	isSet bool
}

func (v NullableBaseEvent) Get() *BaseEvent {
	return v.value
}

func (v *NullableBaseEvent) Set(val *BaseEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEvent(val *BaseEvent) *NullableBaseEvent {
	return &NullableBaseEvent{value: val, isSet: true}
}

func (v NullableBaseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
