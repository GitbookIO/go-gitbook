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

// checks if the InviteUsersToOrganization200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteUsersToOrganization200Response{}

// InviteUsersToOrganization200Response struct for InviteUsersToOrganization200Response
type InviteUsersToOrganization200Response struct {
	Users []string `json:"users"`
	// The number of users who were added to the organization
	Invited         float32  `json:"invited"`
	FailedSSOEmails []string `json:"failedSSOEmails,omitempty"`
}

// NewInviteUsersToOrganization200Response instantiates a new InviteUsersToOrganization200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteUsersToOrganization200Response(users []string, invited float32) *InviteUsersToOrganization200Response {
	this := InviteUsersToOrganization200Response{}
	this.Users = users
	this.Invited = invited
	return &this
}

// NewInviteUsersToOrganization200ResponseWithDefaults instantiates a new InviteUsersToOrganization200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteUsersToOrganization200ResponseWithDefaults() *InviteUsersToOrganization200Response {
	this := InviteUsersToOrganization200Response{}
	return &this
}

// GetUsers returns the Users field value
func (o *InviteUsersToOrganization200Response) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *InviteUsersToOrganization200Response) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *InviteUsersToOrganization200Response) SetUsers(v []string) {
	o.Users = v
}

// GetInvited returns the Invited field value
func (o *InviteUsersToOrganization200Response) GetInvited() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Invited
}

// GetInvitedOk returns a tuple with the Invited field value
// and a boolean to check if the value has been set.
func (o *InviteUsersToOrganization200Response) GetInvitedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invited, true
}

// SetInvited sets field value
func (o *InviteUsersToOrganization200Response) SetInvited(v float32) {
	o.Invited = v
}

// GetFailedSSOEmails returns the FailedSSOEmails field value if set, zero value otherwise.
func (o *InviteUsersToOrganization200Response) GetFailedSSOEmails() []string {
	if o == nil || IsNil(o.FailedSSOEmails) {
		var ret []string
		return ret
	}
	return o.FailedSSOEmails
}

// GetFailedSSOEmailsOk returns a tuple with the FailedSSOEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteUsersToOrganization200Response) GetFailedSSOEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.FailedSSOEmails) {
		return nil, false
	}
	return o.FailedSSOEmails, true
}

// HasFailedSSOEmails returns a boolean if a field has been set.
func (o *InviteUsersToOrganization200Response) HasFailedSSOEmails() bool {
	if o != nil && !IsNil(o.FailedSSOEmails) {
		return true
	}

	return false
}

// SetFailedSSOEmails gets a reference to the given []string and assigns it to the FailedSSOEmails field.
func (o *InviteUsersToOrganization200Response) SetFailedSSOEmails(v []string) {
	o.FailedSSOEmails = v
}

func (o InviteUsersToOrganization200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteUsersToOrganization200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["invited"] = o.Invited
	if !IsNil(o.FailedSSOEmails) {
		toSerialize["failedSSOEmails"] = o.FailedSSOEmails
	}
	return toSerialize, nil
}

type NullableInviteUsersToOrganization200Response struct {
	value *InviteUsersToOrganization200Response
	isSet bool
}

func (v NullableInviteUsersToOrganization200Response) Get() *InviteUsersToOrganization200Response {
	return v.value
}

func (v *NullableInviteUsersToOrganization200Response) Set(val *InviteUsersToOrganization200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteUsersToOrganization200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteUsersToOrganization200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteUsersToOrganization200Response(val *InviteUsersToOrganization200Response) *NullableInviteUsersToOrganization200Response {
	return &NullableInviteUsersToOrganization200Response{value: val, isSet: true}
}

func (v NullableInviteUsersToOrganization200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteUsersToOrganization200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}