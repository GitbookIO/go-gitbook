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

// checks if the RevisionBaseGit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionBaseGit{}

// RevisionBaseGit Metadata about a potential associated git commit.
type RevisionBaseGit struct {
	// SHA of the Git commit.
	Oid string `json:"oid"`
	// Git commit message.
	Message string `json:"message"`
	// Whether not this commit was created by GitBook, while exporting the revision.
	CreatedByGitBook bool `json:"createdByGitBook"`
	// URL of the Git commit.
	Url *string `json:"url,omitempty"`
}

// NewRevisionBaseGit instantiates a new RevisionBaseGit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionBaseGit(oid string, message string, createdByGitBook bool) *RevisionBaseGit {
	this := RevisionBaseGit{}
	this.Oid = oid
	this.Message = message
	this.CreatedByGitBook = createdByGitBook
	return &this
}

// NewRevisionBaseGitWithDefaults instantiates a new RevisionBaseGit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionBaseGitWithDefaults() *RevisionBaseGit {
	this := RevisionBaseGit{}
	return &this
}

// GetOid returns the Oid field value
func (o *RevisionBaseGit) GetOid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Oid
}

// GetOidOk returns a tuple with the Oid field value
// and a boolean to check if the value has been set.
func (o *RevisionBaseGit) GetOidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Oid, true
}

// SetOid sets field value
func (o *RevisionBaseGit) SetOid(v string) {
	o.Oid = v
}

// GetMessage returns the Message field value
func (o *RevisionBaseGit) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RevisionBaseGit) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RevisionBaseGit) SetMessage(v string) {
	o.Message = v
}

// GetCreatedByGitBook returns the CreatedByGitBook field value
func (o *RevisionBaseGit) GetCreatedByGitBook() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreatedByGitBook
}

// GetCreatedByGitBookOk returns a tuple with the CreatedByGitBook field value
// and a boolean to check if the value has been set.
func (o *RevisionBaseGit) GetCreatedByGitBookOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByGitBook, true
}

// SetCreatedByGitBook sets field value
func (o *RevisionBaseGit) SetCreatedByGitBook(v bool) {
	o.CreatedByGitBook = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RevisionBaseGit) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionBaseGit) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RevisionBaseGit) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RevisionBaseGit) SetUrl(v string) {
	o.Url = &v
}

func (o RevisionBaseGit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionBaseGit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oid"] = o.Oid
	toSerialize["message"] = o.Message
	toSerialize["createdByGitBook"] = o.CreatedByGitBook
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableRevisionBaseGit struct {
	value *RevisionBaseGit
	isSet bool
}

func (v NullableRevisionBaseGit) Get() *RevisionBaseGit {
	return v.value
}

func (v *NullableRevisionBaseGit) Set(val *RevisionBaseGit) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionBaseGit) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionBaseGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionBaseGit(val *RevisionBaseGit) *NullableRevisionBaseGit {
	return &NullableRevisionBaseGit{value: val, isSet: true}
}

func (v NullableRevisionBaseGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionBaseGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
