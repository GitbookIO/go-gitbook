/*
GitBook API

The GitBook API

API version: 0.0.1-beta
Contact: support@gitbook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the SyncDirectorySyncGroupsToTeamsRequestToSyncInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncDirectorySyncGroupsToTeamsRequestToSyncInner{}

// SyncDirectorySyncGroupsToTeamsRequestToSyncInner struct for SyncDirectorySyncGroupsToTeamsRequestToSyncInner
type SyncDirectorySyncGroupsToTeamsRequestToSyncInner struct {
	// The unique identifier of the group for the given identity provider.
	GroupId string `json:"group_id"`
	// The unique identifier of the team for the given identity provider.
	TeamId *string `json:"team_id,omitempty"`
}

// NewSyncDirectorySyncGroupsToTeamsRequestToSyncInner instantiates a new SyncDirectorySyncGroupsToTeamsRequestToSyncInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncDirectorySyncGroupsToTeamsRequestToSyncInner(groupId string) *SyncDirectorySyncGroupsToTeamsRequestToSyncInner {
	this := SyncDirectorySyncGroupsToTeamsRequestToSyncInner{}
	this.GroupId = groupId
	return &this
}

// NewSyncDirectorySyncGroupsToTeamsRequestToSyncInnerWithDefaults instantiates a new SyncDirectorySyncGroupsToTeamsRequestToSyncInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncDirectorySyncGroupsToTeamsRequestToSyncInnerWithDefaults() *SyncDirectorySyncGroupsToTeamsRequestToSyncInner {
	this := SyncDirectorySyncGroupsToTeamsRequestToSyncInner{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) SetGroupId(v string) {
	o.GroupId = v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) SetTeamId(v string) {
	o.TeamId = &v
}

func (o SyncDirectorySyncGroupsToTeamsRequestToSyncInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncDirectorySyncGroupsToTeamsRequestToSyncInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	return toSerialize, nil
}

type NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner struct {
	value *SyncDirectorySyncGroupsToTeamsRequestToSyncInner
	isSet bool
}

func (v NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner) Get() *SyncDirectorySyncGroupsToTeamsRequestToSyncInner {
	return v.value
}

func (v *NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner) Set(val *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner(val *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) *NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner {
	return &NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner{value: val, isSet: true}
}

func (v NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncDirectorySyncGroupsToTeamsRequestToSyncInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
