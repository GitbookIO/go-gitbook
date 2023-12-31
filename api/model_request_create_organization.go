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

// checks if the RequestCreateOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestCreateOrganization{}

// RequestCreateOrganization struct for RequestCreateOrganization
type RequestCreateOrganization struct {
	// Name of the organization
	Title        string               `json:"title"`
	EmailDomains []string             `json:"emailDomains,omitempty"`
	Type         *OrganizationType    `json:"type,omitempty"`
	UseCase      *OrganizationUseCase `json:"useCase,omitempty"`
}

// NewRequestCreateOrganization instantiates a new RequestCreateOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestCreateOrganization(title string) *RequestCreateOrganization {
	this := RequestCreateOrganization{}
	this.Title = title
	return &this
}

// NewRequestCreateOrganizationWithDefaults instantiates a new RequestCreateOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestCreateOrganizationWithDefaults() *RequestCreateOrganization {
	this := RequestCreateOrganization{}
	return &this
}

// GetTitle returns the Title field value
func (o *RequestCreateOrganization) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *RequestCreateOrganization) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *RequestCreateOrganization) SetTitle(v string) {
	o.Title = v
}

// GetEmailDomains returns the EmailDomains field value if set, zero value otherwise.
func (o *RequestCreateOrganization) GetEmailDomains() []string {
	if o == nil || IsNil(o.EmailDomains) {
		var ret []string
		return ret
	}
	return o.EmailDomains
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreateOrganization) GetEmailDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailDomains) {
		return nil, false
	}
	return o.EmailDomains, true
}

// HasEmailDomains returns a boolean if a field has been set.
func (o *RequestCreateOrganization) HasEmailDomains() bool {
	if o != nil && !IsNil(o.EmailDomains) {
		return true
	}

	return false
}

// SetEmailDomains gets a reference to the given []string and assigns it to the EmailDomains field.
func (o *RequestCreateOrganization) SetEmailDomains(v []string) {
	o.EmailDomains = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RequestCreateOrganization) GetType() OrganizationType {
	if o == nil || IsNil(o.Type) {
		var ret OrganizationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreateOrganization) GetTypeOk() (*OrganizationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RequestCreateOrganization) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OrganizationType and assigns it to the Type field.
func (o *RequestCreateOrganization) SetType(v OrganizationType) {
	o.Type = &v
}

// GetUseCase returns the UseCase field value if set, zero value otherwise.
func (o *RequestCreateOrganization) GetUseCase() OrganizationUseCase {
	if o == nil || IsNil(o.UseCase) {
		var ret OrganizationUseCase
		return ret
	}
	return *o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreateOrganization) GetUseCaseOk() (*OrganizationUseCase, bool) {
	if o == nil || IsNil(o.UseCase) {
		return nil, false
	}
	return o.UseCase, true
}

// HasUseCase returns a boolean if a field has been set.
func (o *RequestCreateOrganization) HasUseCase() bool {
	if o != nil && !IsNil(o.UseCase) {
		return true
	}

	return false
}

// SetUseCase gets a reference to the given OrganizationUseCase and assigns it to the UseCase field.
func (o *RequestCreateOrganization) SetUseCase(v OrganizationUseCase) {
	o.UseCase = &v
}

func (o RequestCreateOrganization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestCreateOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.EmailDomains) {
		toSerialize["emailDomains"] = o.EmailDomains
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UseCase) {
		toSerialize["useCase"] = o.UseCase
	}
	return toSerialize, nil
}

type NullableRequestCreateOrganization struct {
	value *RequestCreateOrganization
	isSet bool
}

func (v NullableRequestCreateOrganization) Get() *RequestCreateOrganization {
	return v.value
}

func (v *NullableRequestCreateOrganization) Set(val *RequestCreateOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCreateOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCreateOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCreateOrganization(val *RequestCreateOrganization) *NullableRequestCreateOrganization {
	return &NullableRequestCreateOrganization{value: val, isSet: true}
}

func (v NullableRequestCreateOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCreateOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
