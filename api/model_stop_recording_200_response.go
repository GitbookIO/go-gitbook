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
)

// checks if the StopRecording200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopRecording200Response{}

// StopRecording200Response struct for StopRecording200Response
type StopRecording200Response struct {
	Recording RecordingOutput `json:"recording"`
	// Example questions that would be answered by the content of this recording.
	FollowupQuestions []string `json:"followupQuestions,omitempty"`
}

// NewStopRecording200Response instantiates a new StopRecording200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopRecording200Response(recording RecordingOutput) *StopRecording200Response {
	this := StopRecording200Response{}
	this.Recording = recording
	return &this
}

// NewStopRecording200ResponseWithDefaults instantiates a new StopRecording200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopRecording200ResponseWithDefaults() *StopRecording200Response {
	this := StopRecording200Response{}
	return &this
}

// GetRecording returns the Recording field value
func (o *StopRecording200Response) GetRecording() RecordingOutput {
	if o == nil {
		var ret RecordingOutput
		return ret
	}

	return o.Recording
}

// GetRecordingOk returns a tuple with the Recording field value
// and a boolean to check if the value has been set.
func (o *StopRecording200Response) GetRecordingOk() (*RecordingOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recording, true
}

// SetRecording sets field value
func (o *StopRecording200Response) SetRecording(v RecordingOutput) {
	o.Recording = v
}

// GetFollowupQuestions returns the FollowupQuestions field value if set, zero value otherwise.
func (o *StopRecording200Response) GetFollowupQuestions() []string {
	if o == nil || IsNil(o.FollowupQuestions) {
		var ret []string
		return ret
	}
	return o.FollowupQuestions
}

// GetFollowupQuestionsOk returns a tuple with the FollowupQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopRecording200Response) GetFollowupQuestionsOk() ([]string, bool) {
	if o == nil || IsNil(o.FollowupQuestions) {
		return nil, false
	}
	return o.FollowupQuestions, true
}

// HasFollowupQuestions returns a boolean if a field has been set.
func (o *StopRecording200Response) HasFollowupQuestions() bool {
	if o != nil && !IsNil(o.FollowupQuestions) {
		return true
	}

	return false
}

// SetFollowupQuestions gets a reference to the given []string and assigns it to the FollowupQuestions field.
func (o *StopRecording200Response) SetFollowupQuestions(v []string) {
	o.FollowupQuestions = v
}

func (o StopRecording200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopRecording200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recording"] = o.Recording
	if !IsNil(o.FollowupQuestions) {
		toSerialize["followupQuestions"] = o.FollowupQuestions
	}
	return toSerialize, nil
}

type NullableStopRecording200Response struct {
	value *StopRecording200Response
	isSet bool
}

func (v NullableStopRecording200Response) Get() *StopRecording200Response {
	return v.value
}

func (v *NullableStopRecording200Response) Set(val *StopRecording200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStopRecording200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStopRecording200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopRecording200Response(val *StopRecording200Response) *NullableStopRecording200Response {
	return &NullableStopRecording200Response{value: val, isSet: true}
}

func (v NullableStopRecording200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopRecording200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}