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

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization struct for Organization
type Organization struct {
	// Type of Object, always equals to \"organization\"
	Object string `json:"object"`
	// Unique identifier for the organization
	Id string `json:"id"`
	// Name of the organization
	Title         string                     `json:"title"`
	EmailDomains  []string                   `json:"emailDomains"`
	Type          OrganizationType           `json:"type"`
	UseCase       *OrganizationUseCase       `json:"useCase,omitempty"`
	CommunityType *OrganizationCommunityType `json:"communityType,omitempty"`
	Urls          OrganizationUrls           `json:"urls"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(object string, id string, title string, emailDomains []string, type_ OrganizationType, urls OrganizationUrls) *Organization {
	this := Organization{}
	this.Object = object
	this.Id = id
	this.Title = title
	this.EmailDomains = emailDomains
	this.Type = type_
	this.Urls = urls
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetObject returns the Object field value
func (o *Organization) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Organization) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Organization) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Organization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Organization) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *Organization) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Organization) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Organization) SetTitle(v string) {
	o.Title = v
}

// GetEmailDomains returns the EmailDomains field value
func (o *Organization) GetEmailDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EmailDomains
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value
// and a boolean to check if the value has been set.
func (o *Organization) GetEmailDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailDomains, true
}

// SetEmailDomains sets field value
func (o *Organization) SetEmailDomains(v []string) {
	o.EmailDomains = v
}

// GetType returns the Type field value
func (o *Organization) GetType() OrganizationType {
	if o == nil {
		var ret OrganizationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Organization) GetTypeOk() (*OrganizationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Organization) SetType(v OrganizationType) {
	o.Type = v
}

// GetUseCase returns the UseCase field value if set, zero value otherwise.
func (o *Organization) GetUseCase() OrganizationUseCase {
	if o == nil || IsNil(o.UseCase) {
		var ret OrganizationUseCase
		return ret
	}
	return *o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetUseCaseOk() (*OrganizationUseCase, bool) {
	if o == nil || IsNil(o.UseCase) {
		return nil, false
	}
	return o.UseCase, true
}

// HasUseCase returns a boolean if a field has been set.
func (o *Organization) HasUseCase() bool {
	if o != nil && !IsNil(o.UseCase) {
		return true
	}

	return false
}

// SetUseCase gets a reference to the given OrganizationUseCase and assigns it to the UseCase field.
func (o *Organization) SetUseCase(v OrganizationUseCase) {
	o.UseCase = &v
}

// GetCommunityType returns the CommunityType field value if set, zero value otherwise.
func (o *Organization) GetCommunityType() OrganizationCommunityType {
	if o == nil || IsNil(o.CommunityType) {
		var ret OrganizationCommunityType
		return ret
	}
	return *o.CommunityType
}

// GetCommunityTypeOk returns a tuple with the CommunityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCommunityTypeOk() (*OrganizationCommunityType, bool) {
	if o == nil || IsNil(o.CommunityType) {
		return nil, false
	}
	return o.CommunityType, true
}

// HasCommunityType returns a boolean if a field has been set.
func (o *Organization) HasCommunityType() bool {
	if o != nil && !IsNil(o.CommunityType) {
		return true
	}

	return false
}

// SetCommunityType gets a reference to the given OrganizationCommunityType and assigns it to the CommunityType field.
func (o *Organization) SetCommunityType(v OrganizationCommunityType) {
	o.CommunityType = &v
}

// GetUrls returns the Urls field value
func (o *Organization) GetUrls() OrganizationUrls {
	if o == nil {
		var ret OrganizationUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Organization) GetUrlsOk() (*OrganizationUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Organization) SetUrls(v OrganizationUrls) {
	o.Urls = v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["emailDomains"] = o.EmailDomains
	toSerialize["type"] = o.Type
	if !IsNil(o.UseCase) {
		toSerialize["useCase"] = o.UseCase
	}
	if !IsNil(o.CommunityType) {
		toSerialize["communityType"] = o.CommunityType
	}
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
