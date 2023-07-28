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

// checks if the OrganizationTransferResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationTransferResponse{}

// OrganizationTransferResponse struct for OrganizationTransferResponse
type OrganizationTransferResponse struct {
	// The unique id of the collection created in the target organization containing the content of the source organization.
	Collection string `json:"collection"`
	// The new hostname if the source organization needed to change hostname.
	NewSourceHostname *string `json:"newSourceHostname,omitempty"`
}

// NewOrganizationTransferResponse instantiates a new OrganizationTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTransferResponse(collection string) *OrganizationTransferResponse {
	this := OrganizationTransferResponse{}
	this.Collection = collection
	return &this
}

// NewOrganizationTransferResponseWithDefaults instantiates a new OrganizationTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTransferResponseWithDefaults() *OrganizationTransferResponse {
	this := OrganizationTransferResponse{}
	return &this
}

// GetCollection returns the Collection field value
func (o *OrganizationTransferResponse) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *OrganizationTransferResponse) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *OrganizationTransferResponse) SetCollection(v string) {
	o.Collection = v
}

// GetNewSourceHostname returns the NewSourceHostname field value if set, zero value otherwise.
func (o *OrganizationTransferResponse) GetNewSourceHostname() string {
	if o == nil || IsNil(o.NewSourceHostname) {
		var ret string
		return ret
	}
	return *o.NewSourceHostname
}

// GetNewSourceHostnameOk returns a tuple with the NewSourceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationTransferResponse) GetNewSourceHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.NewSourceHostname) {
		return nil, false
	}
	return o.NewSourceHostname, true
}

// HasNewSourceHostname returns a boolean if a field has been set.
func (o *OrganizationTransferResponse) HasNewSourceHostname() bool {
	if o != nil && !IsNil(o.NewSourceHostname) {
		return true
	}

	return false
}

// SetNewSourceHostname gets a reference to the given string and assigns it to the NewSourceHostname field.
func (o *OrganizationTransferResponse) SetNewSourceHostname(v string) {
	o.NewSourceHostname = &v
}

func (o OrganizationTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection"] = o.Collection
	if !IsNil(o.NewSourceHostname) {
		toSerialize["newSourceHostname"] = o.NewSourceHostname
	}
	return toSerialize, nil
}

type NullableOrganizationTransferResponse struct {
	value *OrganizationTransferResponse
	isSet bool
}

func (v NullableOrganizationTransferResponse) Get() *OrganizationTransferResponse {
	return v.value
}

func (v *NullableOrganizationTransferResponse) Set(val *OrganizationTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTransferResponse(val *OrganizationTransferResponse) *NullableOrganizationTransferResponse {
	return &NullableOrganizationTransferResponse{value: val, isSet: true}
}

func (v NullableOrganizationTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}