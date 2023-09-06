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

// checks if the OrganizationTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationTeam{}

// OrganizationTeam struct for OrganizationTeam
type OrganizationTeam struct {
	// Type of Object, always equals to \"team\"
	Object string `json:"object"`
	// Unique identifier for the team.
	Id string `json:"id"`
	// Title of the team.
	Title string `json:"title"`
	// Count of members in this team.
	Members int32 `json:"members"`
	// Count of spaces this team has access to.
	Spaces    float32 `json:"spaces"`
	CreatedAt string  `json:"createdAt"`
}

// NewOrganizationTeam instantiates a new OrganizationTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeam(object string, id string, title string, members int32, spaces float32, createdAt string) *OrganizationTeam {
	this := OrganizationTeam{}
	this.Object = object
	this.Id = id
	this.Title = title
	this.Members = members
	this.Spaces = spaces
	this.CreatedAt = createdAt
	return &this
}

// NewOrganizationTeamWithDefaults instantiates a new OrganizationTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamWithDefaults() *OrganizationTeam {
	this := OrganizationTeam{}
	return &this
}

// GetObject returns the Object field value
func (o *OrganizationTeam) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *OrganizationTeam) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *OrganizationTeam) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationTeam) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *OrganizationTeam) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *OrganizationTeam) SetTitle(v string) {
	o.Title = v
}

// GetMembers returns the Members field value
func (o *OrganizationTeam) GetMembers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetMembersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value
func (o *OrganizationTeam) SetMembers(v int32) {
	o.Members = v
}

// GetSpaces returns the Spaces field value
func (o *OrganizationTeam) GetSpaces() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Spaces
}

// GetSpacesOk returns a tuple with the Spaces field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetSpacesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spaces, true
}

// SetSpaces sets field value
func (o *OrganizationTeam) SetSpaces(v float32) {
	o.Spaces = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationTeam) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeam) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationTeam) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o OrganizationTeam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["members"] = o.Members
	toSerialize["spaces"] = o.Spaces
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

type NullableOrganizationTeam struct {
	value *OrganizationTeam
	isSet bool
}

func (v NullableOrganizationTeam) Get() *OrganizationTeam {
	return v.value
}

func (v *NullableOrganizationTeam) Set(val *OrganizationTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeam(val *OrganizationTeam) *NullableOrganizationTeam {
	return &NullableOrganizationTeam{value: val, isSet: true}
}

func (v NullableOrganizationTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
