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

// checks if the SpaceGitSyncCompletedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceGitSyncCompletedEvent{}

// SpaceGitSyncCompletedEvent struct for SpaceGitSyncCompletedEvent
type SpaceGitSyncCompletedEvent struct {
	// Unique identifier for the event.
	EventId string `json:"eventId"`
	Type    string `json:"type"`
	// ID of the integration installation
	InstallationId string `json:"installationId"`
	// ID of the space
	SpaceId string `json:"spaceId"`
	State   string `json:"state"`
	// Unique identifier of the new content revision
	RevisionId string `json:"revisionId"`
	// Unique identifier for the commit (sha)
	CommitId string `json:"commitId"`
}

// NewSpaceGitSyncCompletedEvent instantiates a new SpaceGitSyncCompletedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceGitSyncCompletedEvent(eventId string, type_ string, installationId string, spaceId string, state string, revisionId string, commitId string) *SpaceGitSyncCompletedEvent {
	this := SpaceGitSyncCompletedEvent{}
	this.EventId = eventId
	this.Type = type_
	this.InstallationId = installationId
	this.SpaceId = spaceId
	this.State = state
	this.RevisionId = revisionId
	this.CommitId = commitId
	return &this
}

// NewSpaceGitSyncCompletedEventWithDefaults instantiates a new SpaceGitSyncCompletedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceGitSyncCompletedEventWithDefaults() *SpaceGitSyncCompletedEvent {
	this := SpaceGitSyncCompletedEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *SpaceGitSyncCompletedEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncCompletedEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *SpaceGitSyncCompletedEvent) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *SpaceGitSyncCompletedEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncCompletedEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceGitSyncCompletedEvent) SetType(v string) {
	o.Type = v
}

// GetInstallationId returns the InstallationId field value
func (o *SpaceGitSyncCompletedEvent) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncCompletedEvent) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *SpaceGitSyncCompletedEvent) SetInstallationId(v string) {
	o.InstallationId = v
}

// GetSpaceId returns the SpaceId field value
func (o *SpaceGitSyncCompletedEvent) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncCompletedEvent) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *SpaceGitSyncCompletedEvent) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetState returns the State field value
func (o *SpaceGitSyncCompletedEvent) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncCompletedEvent) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SpaceGitSyncCompletedEvent) SetState(v string) {
	o.State = v
}

// GetRevisionId returns the RevisionId field value
func (o *SpaceGitSyncCompletedEvent) GetRevisionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncCompletedEvent) GetRevisionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionId, true
}

// SetRevisionId sets field value
func (o *SpaceGitSyncCompletedEvent) SetRevisionId(v string) {
	o.RevisionId = v
}

// GetCommitId returns the CommitId field value
func (o *SpaceGitSyncCompletedEvent) GetCommitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitId
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
func (o *SpaceGitSyncCompletedEvent) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitId, true
}

// SetCommitId sets field value
func (o *SpaceGitSyncCompletedEvent) SetCommitId(v string) {
	o.CommitId = v
}

func (o SpaceGitSyncCompletedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceGitSyncCompletedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["installationId"] = o.InstallationId
	toSerialize["spaceId"] = o.SpaceId
	toSerialize["state"] = o.State
	toSerialize["revisionId"] = o.RevisionId
	toSerialize["commitId"] = o.CommitId
	return toSerialize, nil
}

type NullableSpaceGitSyncCompletedEvent struct {
	value *SpaceGitSyncCompletedEvent
	isSet bool
}

func (v NullableSpaceGitSyncCompletedEvent) Get() *SpaceGitSyncCompletedEvent {
	return v.value
}

func (v *NullableSpaceGitSyncCompletedEvent) Set(val *SpaceGitSyncCompletedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceGitSyncCompletedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceGitSyncCompletedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceGitSyncCompletedEvent(val *SpaceGitSyncCompletedEvent) *NullableSpaceGitSyncCompletedEvent {
	return &NullableSpaceGitSyncCompletedEvent{value: val, isSet: true}
}

func (v NullableSpaceGitSyncCompletedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceGitSyncCompletedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
