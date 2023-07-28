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

// checks if the UserContributor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserContributor{}

// UserContributor Contributor towards content.
type UserContributor struct {
	UpdatedAt string `json:"updatedAt"`
	Count     int32  `json:"count"`
	User      User   `json:"user"`
}

// NewUserContributor instantiates a new UserContributor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserContributor(updatedAt string, count int32, user User) *UserContributor {
	this := UserContributor{}
	this.UpdatedAt = updatedAt
	this.Count = count
	this.User = user
	return &this
}

// NewUserContributorWithDefaults instantiates a new UserContributor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserContributorWithDefaults() *UserContributor {
	this := UserContributor{}
	return &this
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserContributor) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserContributor) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserContributor) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetCount returns the Count field value
func (o *UserContributor) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *UserContributor) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *UserContributor) SetCount(v int32) {
	o.Count = v
}

// GetUser returns the User field value
func (o *UserContributor) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserContributor) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserContributor) SetUser(v User) {
	o.User = v
}

func (o UserContributor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserContributor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["count"] = o.Count
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableUserContributor struct {
	value *UserContributor
	isSet bool
}

func (v NullableUserContributor) Get() *UserContributor {
	return v.value
}

func (v *NullableUserContributor) Set(val *UserContributor) {
	v.value = val
	v.isSet = true
}

func (v NullableUserContributor) IsSet() bool {
	return v.isSet
}

func (v *NullableUserContributor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserContributor(val *UserContributor) *NullableUserContributor {
	return &NullableUserContributor{value: val, isSet: true}
}

func (v NullableUserContributor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserContributor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
