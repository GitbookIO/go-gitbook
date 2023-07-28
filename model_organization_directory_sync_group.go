/*
GitBook API

The GitBook API

API version: 0.0.1-beta
Contact: support@gitbook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the OrganizationDirectorySyncGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationDirectorySyncGroup{}

// OrganizationDirectorySyncGroup struct for OrganizationDirectorySyncGroup
type OrganizationDirectorySyncGroup struct {
	// The unique identifier of this group in WorkOS. Not the unique ID from GitBook.
	Id string `json:"id"`
	// The identity provider's unique ID for this group, should be used to generate the team's unique ID when syncing the groups.
	IdpId string `json:"idp_id"`
	// The unique ID of the directory this group is owned by in WorkOS. Is not a unique ID from our database.
	DirectoryId string `json:"directory_id"`
	// The name of the group from the identity provider, it should always be set according to the WorkOS documentation.
	Name string `json:"name"`
	// The unique ID of the GitBook team already synced to this group, if applicable.
	TeamKey *string `json:"teamKey,omitempty"`
}

// NewOrganizationDirectorySyncGroup instantiates a new OrganizationDirectorySyncGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDirectorySyncGroup(id string, idpId string, directoryId string, name string) *OrganizationDirectorySyncGroup {
	this := OrganizationDirectorySyncGroup{}
	this.Id = id
	this.IdpId = idpId
	this.DirectoryId = directoryId
	this.Name = name
	return &this
}

// NewOrganizationDirectorySyncGroupWithDefaults instantiates a new OrganizationDirectorySyncGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDirectorySyncGroupWithDefaults() *OrganizationDirectorySyncGroup {
	this := OrganizationDirectorySyncGroup{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationDirectorySyncGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationDirectorySyncGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationDirectorySyncGroup) SetId(v string) {
	o.Id = v
}

// GetIdpId returns the IdpId field value
func (o *OrganizationDirectorySyncGroup) GetIdpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpId
}

// GetIdpIdOk returns a tuple with the IdpId field value
// and a boolean to check if the value has been set.
func (o *OrganizationDirectorySyncGroup) GetIdpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpId, true
}

// SetIdpId sets field value
func (o *OrganizationDirectorySyncGroup) SetIdpId(v string) {
	o.IdpId = v
}

// GetDirectoryId returns the DirectoryId field value
func (o *OrganizationDirectorySyncGroup) GetDirectoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectoryId
}

// GetDirectoryIdOk returns a tuple with the DirectoryId field value
// and a boolean to check if the value has been set.
func (o *OrganizationDirectorySyncGroup) GetDirectoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectoryId, true
}

// SetDirectoryId sets field value
func (o *OrganizationDirectorySyncGroup) SetDirectoryId(v string) {
	o.DirectoryId = v
}

// GetName returns the Name field value
func (o *OrganizationDirectorySyncGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationDirectorySyncGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationDirectorySyncGroup) SetName(v string) {
	o.Name = v
}

// GetTeamKey returns the TeamKey field value if set, zero value otherwise.
func (o *OrganizationDirectorySyncGroup) GetTeamKey() string {
	if o == nil || IsNil(o.TeamKey) {
		var ret string
		return ret
	}
	return *o.TeamKey
}

// GetTeamKeyOk returns a tuple with the TeamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDirectorySyncGroup) GetTeamKeyOk() (*string, bool) {
	if o == nil || IsNil(o.TeamKey) {
		return nil, false
	}
	return o.TeamKey, true
}

// HasTeamKey returns a boolean if a field has been set.
func (o *OrganizationDirectorySyncGroup) HasTeamKey() bool {
	if o != nil && !IsNil(o.TeamKey) {
		return true
	}

	return false
}

// SetTeamKey gets a reference to the given string and assigns it to the TeamKey field.
func (o *OrganizationDirectorySyncGroup) SetTeamKey(v string) {
	o.TeamKey = &v
}

func (o OrganizationDirectorySyncGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationDirectorySyncGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["idp_id"] = o.IdpId
	toSerialize["directory_id"] = o.DirectoryId
	toSerialize["name"] = o.Name
	if !IsNil(o.TeamKey) {
		toSerialize["teamKey"] = o.TeamKey
	}
	return toSerialize, nil
}

type NullableOrganizationDirectorySyncGroup struct {
	value *OrganizationDirectorySyncGroup
	isSet bool
}

func (v NullableOrganizationDirectorySyncGroup) Get() *OrganizationDirectorySyncGroup {
	return v.value
}

func (v *NullableOrganizationDirectorySyncGroup) Set(val *OrganizationDirectorySyncGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDirectorySyncGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDirectorySyncGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDirectorySyncGroup(val *OrganizationDirectorySyncGroup) *NullableOrganizationDirectorySyncGroup {
	return &NullableOrganizationDirectorySyncGroup{value: val, isSet: true}
}

func (v NullableOrganizationDirectorySyncGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDirectorySyncGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
