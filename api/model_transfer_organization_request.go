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

// checks if the TransferOrganizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferOrganizationRequest{}

// TransferOrganizationRequest struct for TransferOrganizationRequest
type TransferOrganizationRequest struct {
	// The unique id of the organization to transfer into the current one.
	Source         string                                     `json:"source"`
	DefaultOrgRole *TransferOrganizationRequestDefaultOrgRole `json:"defaultOrgRole,omitempty"`
}

// NewTransferOrganizationRequest instantiates a new TransferOrganizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOrganizationRequest(source string) *TransferOrganizationRequest {
	this := TransferOrganizationRequest{}
	this.Source = source
	return &this
}

// NewTransferOrganizationRequestWithDefaults instantiates a new TransferOrganizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOrganizationRequestWithDefaults() *TransferOrganizationRequest {
	this := TransferOrganizationRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *TransferOrganizationRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TransferOrganizationRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *TransferOrganizationRequest) SetSource(v string) {
	o.Source = v
}

// GetDefaultOrgRole returns the DefaultOrgRole field value if set, zero value otherwise.
func (o *TransferOrganizationRequest) GetDefaultOrgRole() TransferOrganizationRequestDefaultOrgRole {
	if o == nil || IsNil(o.DefaultOrgRole) {
		var ret TransferOrganizationRequestDefaultOrgRole
		return ret
	}
	return *o.DefaultOrgRole
}

// GetDefaultOrgRoleOk returns a tuple with the DefaultOrgRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOrganizationRequest) GetDefaultOrgRoleOk() (*TransferOrganizationRequestDefaultOrgRole, bool) {
	if o == nil || IsNil(o.DefaultOrgRole) {
		return nil, false
	}
	return o.DefaultOrgRole, true
}

// HasDefaultOrgRole returns a boolean if a field has been set.
func (o *TransferOrganizationRequest) HasDefaultOrgRole() bool {
	if o != nil && !IsNil(o.DefaultOrgRole) {
		return true
	}

	return false
}

// SetDefaultOrgRole gets a reference to the given TransferOrganizationRequestDefaultOrgRole and assigns it to the DefaultOrgRole field.
func (o *TransferOrganizationRequest) SetDefaultOrgRole(v TransferOrganizationRequestDefaultOrgRole) {
	o.DefaultOrgRole = &v
}

func (o TransferOrganizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferOrganizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	if !IsNil(o.DefaultOrgRole) {
		toSerialize["defaultOrgRole"] = o.DefaultOrgRole
	}
	return toSerialize, nil
}

type NullableTransferOrganizationRequest struct {
	value *TransferOrganizationRequest
	isSet bool
}

func (v NullableTransferOrganizationRequest) Get() *TransferOrganizationRequest {
	return v.value
}

func (v *NullableTransferOrganizationRequest) Set(val *TransferOrganizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOrganizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOrganizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOrganizationRequest(val *TransferOrganizationRequest) *NullableTransferOrganizationRequest {
	return &NullableTransferOrganizationRequest{value: val, isSet: true}
}

func (v NullableTransferOrganizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOrganizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}