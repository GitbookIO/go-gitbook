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

// checks if the ChangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequest{}

// ChangeRequest struct for ChangeRequest
type ChangeRequest struct {
	// Type of Object, always equals to \"change-request\"
	Object string `json:"object"`
	// Unique identifier for the change request
	Id string `json:"id"`
	// Incremental identifier of the change request
	Number float32             `json:"number"`
	Status ChangeRequestStatus `json:"status"`
	// Subject of the change request
	Subject   string            `json:"subject"`
	CreatedBy User              `json:"createdBy"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt string            `json:"updatedAt"`
	Urls      ChangeRequestUrls `json:"urls"`
}

// NewChangeRequest instantiates a new ChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequest(object string, id string, number float32, status ChangeRequestStatus, subject string, createdBy User, createdAt string, updatedAt string, urls ChangeRequestUrls) *ChangeRequest {
	this := ChangeRequest{}
	this.Object = object
	this.Id = id
	this.Number = number
	this.Status = status
	this.Subject = subject
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Urls = urls
	return &this
}

// NewChangeRequestWithDefaults instantiates a new ChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestWithDefaults() *ChangeRequest {
	this := ChangeRequest{}
	return &this
}

// GetObject returns the Object field value
func (o *ChangeRequest) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ChangeRequest) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *ChangeRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChangeRequest) SetId(v string) {
	o.Id = v
}

// GetNumber returns the Number field value
func (o *ChangeRequest) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *ChangeRequest) SetNumber(v float32) {
	o.Number = v
}

// GetStatus returns the Status field value
func (o *ChangeRequest) GetStatus() ChangeRequestStatus {
	if o == nil {
		var ret ChangeRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetStatusOk() (*ChangeRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChangeRequest) SetStatus(v ChangeRequestStatus) {
	o.Status = v
}

// GetSubject returns the Subject field value
func (o *ChangeRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *ChangeRequest) SetSubject(v string) {
	o.Subject = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ChangeRequest) GetCreatedBy() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetCreatedByOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ChangeRequest) SetCreatedBy(v User) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChangeRequest) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChangeRequest) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ChangeRequest) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ChangeRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUrls returns the Urls field value
func (o *ChangeRequest) GetUrls() ChangeRequestUrls {
	if o == nil {
		var ret ChangeRequestUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *ChangeRequest) GetUrlsOk() (*ChangeRequestUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *ChangeRequest) SetUrls(v ChangeRequestUrls) {
	o.Urls = v
}

func (o ChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["number"] = o.Number
	toSerialize["status"] = o.Status
	toSerialize["subject"] = o.Subject
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableChangeRequest struct {
	value *ChangeRequest
	isSet bool
}

func (v NullableChangeRequest) Get() *ChangeRequest {
	return v.value
}

func (v *NullableChangeRequest) Set(val *ChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequest(val *ChangeRequest) *NullableChangeRequest {
	return &NullableChangeRequest{value: val, isSet: true}
}

func (v NullableChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
