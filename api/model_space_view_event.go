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

// checks if the SpaceViewEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceViewEvent{}

// SpaceViewEvent struct for SpaceViewEvent
type SpaceViewEvent struct {
	// Unique identifier for the event.
	EventId string `json:"eventId"`
	Type    string `json:"type"`
	// ID of the integration installation
	InstallationId string `json:"installationId"`
	// ID of the space
	SpaceId string `json:"spaceId"`
	// Unique identifier of the visited page.
	PageId  *string                    `json:"pageId,omitempty"`
	Visitor SpaceViewEventAllOfVisitor `json:"visitor"`
	// The GitBook content's URL visited (including URL params).
	Url string `json:"url"`
	// The URL of referrer that linked to the page.
	Referrer string `json:"referrer"`
}

// NewSpaceViewEvent instantiates a new SpaceViewEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceViewEvent(eventId string, type_ string, installationId string, spaceId string, visitor SpaceViewEventAllOfVisitor, url string, referrer string) *SpaceViewEvent {
	this := SpaceViewEvent{}
	this.EventId = eventId
	this.Type = type_
	this.InstallationId = installationId
	this.SpaceId = spaceId
	this.Visitor = visitor
	this.Url = url
	this.Referrer = referrer
	return &this
}

// NewSpaceViewEventWithDefaults instantiates a new SpaceViewEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceViewEventWithDefaults() *SpaceViewEvent {
	this := SpaceViewEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *SpaceViewEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *SpaceViewEvent) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value
func (o *SpaceViewEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceViewEvent) SetType(v string) {
	o.Type = v
}

// GetInstallationId returns the InstallationId field value
func (o *SpaceViewEvent) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *SpaceViewEvent) SetInstallationId(v string) {
	o.InstallationId = v
}

// GetSpaceId returns the SpaceId field value
func (o *SpaceViewEvent) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *SpaceViewEvent) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *SpaceViewEvent) GetPageId() string {
	if o == nil || IsNil(o.PageId) {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetPageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PageId) {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *SpaceViewEvent) HasPageId() bool {
	if o != nil && !IsNil(o.PageId) {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *SpaceViewEvent) SetPageId(v string) {
	o.PageId = &v
}

// GetVisitor returns the Visitor field value
func (o *SpaceViewEvent) GetVisitor() SpaceViewEventAllOfVisitor {
	if o == nil {
		var ret SpaceViewEventAllOfVisitor
		return ret
	}

	return o.Visitor
}

// GetVisitorOk returns a tuple with the Visitor field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetVisitorOk() (*SpaceViewEventAllOfVisitor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visitor, true
}

// SetVisitor sets field value
func (o *SpaceViewEvent) SetVisitor(v SpaceViewEventAllOfVisitor) {
	o.Visitor = v
}

// GetUrl returns the Url field value
func (o *SpaceViewEvent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SpaceViewEvent) SetUrl(v string) {
	o.Url = v
}

// GetReferrer returns the Referrer field value
func (o *SpaceViewEvent) GetReferrer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Referrer
}

// GetReferrerOk returns a tuple with the Referrer field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEvent) GetReferrerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Referrer, true
}

// SetReferrer sets field value
func (o *SpaceViewEvent) SetReferrer(v string) {
	o.Referrer = v
}

func (o SpaceViewEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceViewEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["type"] = o.Type
	toSerialize["installationId"] = o.InstallationId
	toSerialize["spaceId"] = o.SpaceId
	if !IsNil(o.PageId) {
		toSerialize["pageId"] = o.PageId
	}
	toSerialize["visitor"] = o.Visitor
	toSerialize["url"] = o.Url
	toSerialize["referrer"] = o.Referrer
	return toSerialize, nil
}

type NullableSpaceViewEvent struct {
	value *SpaceViewEvent
	isSet bool
}

func (v NullableSpaceViewEvent) Get() *SpaceViewEvent {
	return v.value
}

func (v *NullableSpaceViewEvent) Set(val *SpaceViewEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceViewEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceViewEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceViewEvent(val *SpaceViewEvent) *NullableSpaceViewEvent {
	return &NullableSpaceViewEvent{value: val, isSet: true}
}

func (v NullableSpaceViewEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceViewEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
