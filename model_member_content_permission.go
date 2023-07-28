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

// checks if the MemberContentPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberContentPermission{}

// MemberContentPermission Permission of a member in a content.
type MemberContentPermission struct {
	Permission MemberRole `json:"permission"`
	Space      Space      `json:"space"`
}

// NewMemberContentPermission instantiates a new MemberContentPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberContentPermission(permission MemberRole, space Space) *MemberContentPermission {
	this := MemberContentPermission{}
	this.Permission = permission
	this.Space = space
	return &this
}

// NewMemberContentPermissionWithDefaults instantiates a new MemberContentPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberContentPermissionWithDefaults() *MemberContentPermission {
	this := MemberContentPermission{}
	return &this
}

// GetPermission returns the Permission field value
func (o *MemberContentPermission) GetPermission() MemberRole {
	if o == nil {
		var ret MemberRole
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *MemberContentPermission) GetPermissionOk() (*MemberRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *MemberContentPermission) SetPermission(v MemberRole) {
	o.Permission = v
}

// GetSpace returns the Space field value
func (o *MemberContentPermission) GetSpace() Space {
	if o == nil {
		var ret Space
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *MemberContentPermission) GetSpaceOk() (*Space, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *MemberContentPermission) SetSpace(v Space) {
	o.Space = v
}

func (o MemberContentPermission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberContentPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permission"] = o.Permission
	toSerialize["space"] = o.Space
	return toSerialize, nil
}

type NullableMemberContentPermission struct {
	value *MemberContentPermission
	isSet bool
}

func (v NullableMemberContentPermission) Get() *MemberContentPermission {
	return v.value
}

func (v *NullableMemberContentPermission) Set(val *MemberContentPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberContentPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberContentPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberContentPermission(val *MemberContentPermission) *NullableMemberContentPermission {
	return &NullableMemberContentPermission{value: val, isSet: true}
}

func (v NullableMemberContentPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberContentPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
