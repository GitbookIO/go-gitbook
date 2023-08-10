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

// checks if the CreateOrganizationCustomFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationCustomFieldRequest{}

// CreateOrganizationCustomFieldRequest struct for CreateOrganizationCustomFieldRequest
type CreateOrganizationCustomFieldRequest struct {
	Name        string          `json:"name"`
	Type        CustomFieldType `json:"type"`
	Title       *string         `json:"title,omitempty"`
	Description *string         `json:"description,omitempty"`
}

// NewCreateOrganizationCustomFieldRequest instantiates a new CreateOrganizationCustomFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationCustomFieldRequest(name string, type_ CustomFieldType) *CreateOrganizationCustomFieldRequest {
	this := CreateOrganizationCustomFieldRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewCreateOrganizationCustomFieldRequestWithDefaults instantiates a new CreateOrganizationCustomFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationCustomFieldRequestWithDefaults() *CreateOrganizationCustomFieldRequest {
	this := CreateOrganizationCustomFieldRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationCustomFieldRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCustomFieldRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationCustomFieldRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CreateOrganizationCustomFieldRequest) GetType() CustomFieldType {
	if o == nil {
		var ret CustomFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCustomFieldRequest) GetTypeOk() (*CustomFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOrganizationCustomFieldRequest) SetType(v CustomFieldType) {
	o.Type = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateOrganizationCustomFieldRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCustomFieldRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateOrganizationCustomFieldRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateOrganizationCustomFieldRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrganizationCustomFieldRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCustomFieldRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrganizationCustomFieldRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrganizationCustomFieldRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CreateOrganizationCustomFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationCustomFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCreateOrganizationCustomFieldRequest struct {
	value *CreateOrganizationCustomFieldRequest
	isSet bool
}

func (v NullableCreateOrganizationCustomFieldRequest) Get() *CreateOrganizationCustomFieldRequest {
	return v.value
}

func (v *NullableCreateOrganizationCustomFieldRequest) Set(val *CreateOrganizationCustomFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationCustomFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationCustomFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationCustomFieldRequest(val *CreateOrganizationCustomFieldRequest) *NullableCreateOrganizationCustomFieldRequest {
	return &NullableCreateOrganizationCustomFieldRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationCustomFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationCustomFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
