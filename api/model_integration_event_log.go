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

// checks if the IntegrationEventLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationEventLog{}

// IntegrationEventLog struct for IntegrationEventLog
type IntegrationEventLog struct {
	// The message of the log entry.
	Message   *string `json:"message,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	// The level of the log entry.
	Level *string `json:"level,omitempty"`
}

// NewIntegrationEventLog instantiates a new IntegrationEventLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationEventLog() *IntegrationEventLog {
	this := IntegrationEventLog{}
	return &this
}

// NewIntegrationEventLogWithDefaults instantiates a new IntegrationEventLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationEventLogWithDefaults() *IntegrationEventLog {
	this := IntegrationEventLog{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IntegrationEventLog) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEventLog) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IntegrationEventLog) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IntegrationEventLog) SetMessage(v string) {
	o.Message = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *IntegrationEventLog) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEventLog) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *IntegrationEventLog) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *IntegrationEventLog) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *IntegrationEventLog) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEventLog) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *IntegrationEventLog) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *IntegrationEventLog) SetLevel(v string) {
	o.Level = &v
}

func (o IntegrationEventLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationEventLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return toSerialize, nil
}

type NullableIntegrationEventLog struct {
	value *IntegrationEventLog
	isSet bool
}

func (v NullableIntegrationEventLog) Get() *IntegrationEventLog {
	return v.value
}

func (v *NullableIntegrationEventLog) Set(val *IntegrationEventLog) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationEventLog) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationEventLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationEventLog(val *IntegrationEventLog) *NullableIntegrationEventLog {
	return &NullableIntegrationEventLog{value: val, isSet: true}
}

func (v NullableIntegrationEventLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationEventLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
