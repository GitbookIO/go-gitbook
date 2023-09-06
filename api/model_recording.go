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
	"time"
)

// checks if the Recording type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recording{}

// Recording struct for Recording
type Recording struct {
	// Type of Object, always equals to \"recording\"
	Object string `json:"object"`
	// Unique identifier for the recording
	Id string `json:"id"`
	// Optional title describing the recording
	Title   string           `json:"title"`
	Context RecordingContext `json:"context"`
	// ID in the original source of the recording.
	ExternalId *string `json:"externalId,omitempty"`
	// URL of the source from which the recording originated
	ExternalURL  *string                      `json:"externalURL,omitempty"`
	CreatedAt    time.Time                    `json:"createdAt"`
	StoppedAt    *time.Time                   `json:"stoppedAt,omitempty"`
	Events       RecordingEvents              `json:"events"`
	Contributors []RecordingContributorsInner `json:"contributors"`
	Output       *RecordingOutput             `json:"output,omitempty"`
	Urls         RecordingUrls                `json:"urls"`
}

// NewRecording instantiates a new Recording object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecording(object string, id string, title string, context RecordingContext, createdAt time.Time, events RecordingEvents, contributors []RecordingContributorsInner, urls RecordingUrls) *Recording {
	this := Recording{}
	this.Object = object
	this.Id = id
	this.Title = title
	this.Context = context
	this.CreatedAt = createdAt
	this.Events = events
	this.Contributors = contributors
	this.Urls = urls
	return &this
}

// NewRecordingWithDefaults instantiates a new Recording object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingWithDefaults() *Recording {
	this := Recording{}
	return &this
}

// GetObject returns the Object field value
func (o *Recording) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Recording) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Recording) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Recording) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Recording) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Recording) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *Recording) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Recording) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Recording) SetTitle(v string) {
	o.Title = v
}

// GetContext returns the Context field value
func (o *Recording) GetContext() RecordingContext {
	if o == nil {
		var ret RecordingContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *Recording) GetContextOk() (*RecordingContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *Recording) SetContext(v RecordingContext) {
	o.Context = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Recording) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recording) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Recording) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Recording) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalURL returns the ExternalURL field value if set, zero value otherwise.
func (o *Recording) GetExternalURL() string {
	if o == nil || IsNil(o.ExternalURL) {
		var ret string
		return ret
	}
	return *o.ExternalURL
}

// GetExternalURLOk returns a tuple with the ExternalURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recording) GetExternalURLOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalURL) {
		return nil, false
	}
	return o.ExternalURL, true
}

// HasExternalURL returns a boolean if a field has been set.
func (o *Recording) HasExternalURL() bool {
	if o != nil && !IsNil(o.ExternalURL) {
		return true
	}

	return false
}

// SetExternalURL gets a reference to the given string and assigns it to the ExternalURL field.
func (o *Recording) SetExternalURL(v string) {
	o.ExternalURL = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Recording) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Recording) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Recording) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStoppedAt returns the StoppedAt field value if set, zero value otherwise.
func (o *Recording) GetStoppedAt() time.Time {
	if o == nil || IsNil(o.StoppedAt) {
		var ret time.Time
		return ret
	}
	return *o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recording) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StoppedAt) {
		return nil, false
	}
	return o.StoppedAt, true
}

// HasStoppedAt returns a boolean if a field has been set.
func (o *Recording) HasStoppedAt() bool {
	if o != nil && !IsNil(o.StoppedAt) {
		return true
	}

	return false
}

// SetStoppedAt gets a reference to the given time.Time and assigns it to the StoppedAt field.
func (o *Recording) SetStoppedAt(v time.Time) {
	o.StoppedAt = &v
}

// GetEvents returns the Events field value
func (o *Recording) GetEvents() RecordingEvents {
	if o == nil {
		var ret RecordingEvents
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *Recording) GetEventsOk() (*RecordingEvents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *Recording) SetEvents(v RecordingEvents) {
	o.Events = v
}

// GetContributors returns the Contributors field value
func (o *Recording) GetContributors() []RecordingContributorsInner {
	if o == nil {
		var ret []RecordingContributorsInner
		return ret
	}

	return o.Contributors
}

// GetContributorsOk returns a tuple with the Contributors field value
// and a boolean to check if the value has been set.
func (o *Recording) GetContributorsOk() ([]RecordingContributorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contributors, true
}

// SetContributors sets field value
func (o *Recording) SetContributors(v []RecordingContributorsInner) {
	o.Contributors = v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *Recording) GetOutput() RecordingOutput {
	if o == nil || IsNil(o.Output) {
		var ret RecordingOutput
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recording) GetOutputOk() (*RecordingOutput, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *Recording) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given RecordingOutput and assigns it to the Output field.
func (o *Recording) SetOutput(v RecordingOutput) {
	o.Output = &v
}

// GetUrls returns the Urls field value
func (o *Recording) GetUrls() RecordingUrls {
	if o == nil {
		var ret RecordingUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Recording) GetUrlsOk() (*RecordingUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Recording) SetUrls(v RecordingUrls) {
	o.Urls = v
}

func (o Recording) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recording) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["context"] = o.Context
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.ExternalURL) {
		toSerialize["externalURL"] = o.ExternalURL
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.StoppedAt) {
		toSerialize["stoppedAt"] = o.StoppedAt
	}
	toSerialize["events"] = o.Events
	toSerialize["contributors"] = o.Contributors
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableRecording struct {
	value *Recording
	isSet bool
}

func (v NullableRecording) Get() *Recording {
	return v.value
}

func (v *NullableRecording) Set(val *Recording) {
	v.value = val
	v.isSet = true
}

func (v NullableRecording) IsSet() bool {
	return v.isSet
}

func (v *NullableRecording) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecording(val *Recording) *NullableRecording {
	return &NullableRecording{value: val, isSet: true}
}

func (v NullableRecording) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecording) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
