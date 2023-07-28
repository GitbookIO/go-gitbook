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

// checks if the InstallationEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallationEvent{}

// InstallationEvent struct for InstallationEvent
type InstallationEvent struct {
	// Unique identifier for the event.
	EventId string `json:"eventId"`
	// Type of the event.
	Type string `json:"type"`
	// ID of the integration installation
	InstallationId string `json:"installationId"`
}

// NewInstallationEvent instantiates a new InstallationEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallationEvent(eventId string, type_ string, installationId string) *InstallationEvent {
	this := InstallationEvent{}
	this.EventId = eventId
	this.Type = type_
	this.InstallationId = installationId
	return &this
}

// NewInstallationEventWithDefaults instantiates a new InstallationEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallationEventWithDefaults() *InstallationEvent {
	this := InstallationEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *InstallationEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *InstallationEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *InstallationEvent) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *InstallationEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InstallationEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InstallationEvent) SetType(v string) {
	o.Type = v
}

// GetInstallationId returns the InstallationId field value
func (o *InstallationEvent) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *InstallationEvent) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *InstallationEvent) SetInstallationId(v string) {
	o.InstallationId = v
}

func (o InstallationEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallationEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["installationId"] = o.InstallationId
	return toSerialize, nil
}

type NullableInstallationEvent struct {
	value *InstallationEvent
	isSet bool
}

func (v NullableInstallationEvent) Get() *InstallationEvent {
	return v.value
}

func (v *NullableInstallationEvent) Set(val *InstallationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationEvent(val *InstallationEvent) *NullableInstallationEvent {
	return &NullableInstallationEvent{value: val, isSet: true}
}

func (v NullableInstallationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
