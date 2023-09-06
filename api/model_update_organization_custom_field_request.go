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

// checks if the UpdateOrganizationCustomFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationCustomFieldRequest{}

// UpdateOrganizationCustomFieldRequest struct for UpdateOrganizationCustomFieldRequest
type UpdateOrganizationCustomFieldRequest struct {
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Placeholder *string  `json:"placeholder,omitempty"`
	Options     []string `json:"options,omitempty"`
}

// NewUpdateOrganizationCustomFieldRequest instantiates a new UpdateOrganizationCustomFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationCustomFieldRequest() *UpdateOrganizationCustomFieldRequest {
	this := UpdateOrganizationCustomFieldRequest{}
	return &this
}

// NewUpdateOrganizationCustomFieldRequestWithDefaults instantiates a new UpdateOrganizationCustomFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationCustomFieldRequestWithDefaults() *UpdateOrganizationCustomFieldRequest {
	this := UpdateOrganizationCustomFieldRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateOrganizationCustomFieldRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCustomFieldRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateOrganizationCustomFieldRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateOrganizationCustomFieldRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateOrganizationCustomFieldRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCustomFieldRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateOrganizationCustomFieldRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateOrganizationCustomFieldRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *UpdateOrganizationCustomFieldRequest) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCustomFieldRequest) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *UpdateOrganizationCustomFieldRequest) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *UpdateOrganizationCustomFieldRequest) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *UpdateOrganizationCustomFieldRequest) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCustomFieldRequest) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *UpdateOrganizationCustomFieldRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *UpdateOrganizationCustomFieldRequest) SetOptions(v []string) {
	o.Options = v
}

func (o UpdateOrganizationCustomFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationCustomFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationCustomFieldRequest struct {
	value *UpdateOrganizationCustomFieldRequest
	isSet bool
}

func (v NullableUpdateOrganizationCustomFieldRequest) Get() *UpdateOrganizationCustomFieldRequest {
	return v.value
}

func (v *NullableUpdateOrganizationCustomFieldRequest) Set(val *UpdateOrganizationCustomFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationCustomFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationCustomFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationCustomFieldRequest(val *UpdateOrganizationCustomFieldRequest) *NullableUpdateOrganizationCustomFieldRequest {
	return &NullableUpdateOrganizationCustomFieldRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationCustomFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationCustomFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
