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

// checks if the SubmitChangeRequestReviewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitChangeRequestReviewRequest{}

// SubmitChangeRequestReviewRequest struct for SubmitChangeRequestReviewRequest
type SubmitChangeRequestReviewRequest struct {
	Status  ChangeRequestReviewStatus `json:"status"`
	Comment *Document                 `json:"comment,omitempty"`
}

// NewSubmitChangeRequestReviewRequest instantiates a new SubmitChangeRequestReviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitChangeRequestReviewRequest(status ChangeRequestReviewStatus) *SubmitChangeRequestReviewRequest {
	this := SubmitChangeRequestReviewRequest{}
	this.Status = status
	return &this
}

// NewSubmitChangeRequestReviewRequestWithDefaults instantiates a new SubmitChangeRequestReviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitChangeRequestReviewRequestWithDefaults() *SubmitChangeRequestReviewRequest {
	this := SubmitChangeRequestReviewRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *SubmitChangeRequestReviewRequest) GetStatus() ChangeRequestReviewStatus {
	if o == nil {
		var ret ChangeRequestReviewStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubmitChangeRequestReviewRequest) GetStatusOk() (*ChangeRequestReviewStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubmitChangeRequestReviewRequest) SetStatus(v ChangeRequestReviewStatus) {
	o.Status = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SubmitChangeRequestReviewRequest) GetComment() Document {
	if o == nil || IsNil(o.Comment) {
		var ret Document
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitChangeRequestReviewRequest) GetCommentOk() (*Document, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SubmitChangeRequestReviewRequest) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given Document and assigns it to the Comment field.
func (o *SubmitChangeRequestReviewRequest) SetComment(v Document) {
	o.Comment = &v
}

func (o SubmitChangeRequestReviewRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitChangeRequestReviewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableSubmitChangeRequestReviewRequest struct {
	value *SubmitChangeRequestReviewRequest
	isSet bool
}

func (v NullableSubmitChangeRequestReviewRequest) Get() *SubmitChangeRequestReviewRequest {
	return v.value
}

func (v *NullableSubmitChangeRequestReviewRequest) Set(val *SubmitChangeRequestReviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitChangeRequestReviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitChangeRequestReviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitChangeRequestReviewRequest(val *SubmitChangeRequestReviewRequest) *NullableSubmitChangeRequestReviewRequest {
	return &NullableSubmitChangeRequestReviewRequest{value: val, isSet: true}
}

func (v NullableSubmitChangeRequestReviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitChangeRequestReviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
