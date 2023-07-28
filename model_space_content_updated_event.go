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

// checks if the SpaceContentUpdatedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceContentUpdatedEvent{}

// SpaceContentUpdatedEvent struct for SpaceContentUpdatedEvent
type SpaceContentUpdatedEvent struct {
	// Unique identifier for the event.
	EventId string `json:"eventId"`
	Type    string `json:"type"`
	// ID of the integration installation
	InstallationId string `json:"installationId"`
	// ID of the space
	SpaceId string `json:"spaceId"`
	// Unique identifier of the new content revision
	RevisionId string `json:"revisionId"`
}

// NewSpaceContentUpdatedEvent instantiates a new SpaceContentUpdatedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceContentUpdatedEvent(eventId string, type_ string, installationId string, spaceId string, revisionId string) *SpaceContentUpdatedEvent {
	this := SpaceContentUpdatedEvent{}
	this.EventId = eventId
	this.Type = type_
	this.InstallationId = installationId
	this.SpaceId = spaceId
	this.RevisionId = revisionId
	return &this
}

// NewSpaceContentUpdatedEventWithDefaults instantiates a new SpaceContentUpdatedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceContentUpdatedEventWithDefaults() *SpaceContentUpdatedEvent {
	this := SpaceContentUpdatedEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *SpaceContentUpdatedEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *SpaceContentUpdatedEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *SpaceContentUpdatedEvent) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *SpaceContentUpdatedEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceContentUpdatedEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceContentUpdatedEvent) SetType(v string) {
	o.Type = v
}

// GetInstallationId returns the InstallationId field value
func (o *SpaceContentUpdatedEvent) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *SpaceContentUpdatedEvent) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *SpaceContentUpdatedEvent) SetInstallationId(v string) {
	o.InstallationId = v
}

// GetSpaceId returns the SpaceId field value
func (o *SpaceContentUpdatedEvent) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *SpaceContentUpdatedEvent) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *SpaceContentUpdatedEvent) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetRevisionId returns the RevisionId field value
func (o *SpaceContentUpdatedEvent) GetRevisionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value
// and a boolean to check if the value has been set.
func (o *SpaceContentUpdatedEvent) GetRevisionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionId, true
}

// SetRevisionId sets field value
func (o *SpaceContentUpdatedEvent) SetRevisionId(v string) {
	o.RevisionId = v
}

func (o SpaceContentUpdatedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceContentUpdatedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["installationId"] = o.InstallationId
	toSerialize["spaceId"] = o.SpaceId
	toSerialize["revisionId"] = o.RevisionId
	return toSerialize, nil
}

type NullableSpaceContentUpdatedEvent struct {
	value *SpaceContentUpdatedEvent
	isSet bool
}

func (v NullableSpaceContentUpdatedEvent) Get() *SpaceContentUpdatedEvent {
	return v.value
}

func (v *NullableSpaceContentUpdatedEvent) Set(val *SpaceContentUpdatedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceContentUpdatedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceContentUpdatedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceContentUpdatedEvent(val *SpaceContentUpdatedEvent) *NullableSpaceContentUpdatedEvent {
	return &NullableSpaceContentUpdatedEvent{value: val, isSet: true}
}

func (v NullableSpaceContentUpdatedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceContentUpdatedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
