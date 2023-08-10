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

// checks if the FetchPublishedScriptEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchPublishedScriptEvent{}

// FetchPublishedScriptEvent struct for FetchPublishedScriptEvent
type FetchPublishedScriptEvent struct {
	// Unique identifier for the event.
	EventId string `json:"eventId"`
	Type    string `json:"type"`
	// ID of the integration installation
	InstallationId string `json:"installationId"`
	// ID of the space
	SpaceId string `json:"spaceId"`
}

// NewFetchPublishedScriptEvent instantiates a new FetchPublishedScriptEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchPublishedScriptEvent(eventId string, type_ string, installationId string, spaceId string) *FetchPublishedScriptEvent {
	this := FetchPublishedScriptEvent{}
	this.EventId = eventId
	this.Type = type_
	this.InstallationId = installationId
	this.SpaceId = spaceId
	return &this
}

// NewFetchPublishedScriptEventWithDefaults instantiates a new FetchPublishedScriptEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchPublishedScriptEventWithDefaults() *FetchPublishedScriptEvent {
	this := FetchPublishedScriptEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *FetchPublishedScriptEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *FetchPublishedScriptEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *FetchPublishedScriptEvent) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *FetchPublishedScriptEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FetchPublishedScriptEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FetchPublishedScriptEvent) SetType(v string) {
	o.Type = v
}

// GetInstallationId returns the InstallationId field value
func (o *FetchPublishedScriptEvent) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *FetchPublishedScriptEvent) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *FetchPublishedScriptEvent) SetInstallationId(v string) {
	o.InstallationId = v
}

// GetSpaceId returns the SpaceId field value
func (o *FetchPublishedScriptEvent) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *FetchPublishedScriptEvent) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *FetchPublishedScriptEvent) SetSpaceId(v string) {
	o.SpaceId = v
}

func (o FetchPublishedScriptEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchPublishedScriptEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["installationId"] = o.InstallationId
	toSerialize["spaceId"] = o.SpaceId
	return toSerialize, nil
}

type NullableFetchPublishedScriptEvent struct {
	value *FetchPublishedScriptEvent
	isSet bool
}

func (v NullableFetchPublishedScriptEvent) Get() *FetchPublishedScriptEvent {
	return v.value
}

func (v *NullableFetchPublishedScriptEvent) Set(val *FetchPublishedScriptEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchPublishedScriptEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchPublishedScriptEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchPublishedScriptEvent(val *FetchPublishedScriptEvent) *NullableFetchPublishedScriptEvent {
	return &NullableFetchPublishedScriptEvent{value: val, isSet: true}
}

func (v NullableFetchPublishedScriptEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchPublishedScriptEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
