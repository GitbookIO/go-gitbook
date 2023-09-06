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

// checks if the TransferOrganization200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferOrganization200Response{}

// TransferOrganization200Response struct for TransferOrganization200Response
type TransferOrganization200Response struct {
	// The unique id of the collection created in the target organization containing the content of the source organization.
	Collection string `json:"collection"`
	// The new hostname if the source organization needed to change hostname.
	NewSourceHostname *string `json:"newSourceHostname,omitempty"`
}

// NewTransferOrganization200Response instantiates a new TransferOrganization200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOrganization200Response(collection string) *TransferOrganization200Response {
	this := TransferOrganization200Response{}
	this.Collection = collection
	return &this
}

// NewTransferOrganization200ResponseWithDefaults instantiates a new TransferOrganization200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOrganization200ResponseWithDefaults() *TransferOrganization200Response {
	this := TransferOrganization200Response{}
	return &this
}

// GetCollection returns the Collection field value
func (o *TransferOrganization200Response) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *TransferOrganization200Response) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *TransferOrganization200Response) SetCollection(v string) {
	o.Collection = v
}

// GetNewSourceHostname returns the NewSourceHostname field value if set, zero value otherwise.
func (o *TransferOrganization200Response) GetNewSourceHostname() string {
	if o == nil || IsNil(o.NewSourceHostname) {
		var ret string
		return ret
	}
	return *o.NewSourceHostname
}

// GetNewSourceHostnameOk returns a tuple with the NewSourceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOrganization200Response) GetNewSourceHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.NewSourceHostname) {
		return nil, false
	}
	return o.NewSourceHostname, true
}

// HasNewSourceHostname returns a boolean if a field has been set.
func (o *TransferOrganization200Response) HasNewSourceHostname() bool {
	if o != nil && !IsNil(o.NewSourceHostname) {
		return true
	}

	return false
}

// SetNewSourceHostname gets a reference to the given string and assigns it to the NewSourceHostname field.
func (o *TransferOrganization200Response) SetNewSourceHostname(v string) {
	o.NewSourceHostname = &v
}

func (o TransferOrganization200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferOrganization200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection"] = o.Collection
	if !IsNil(o.NewSourceHostname) {
		toSerialize["newSourceHostname"] = o.NewSourceHostname
	}
	return toSerialize, nil
}

type NullableTransferOrganization200Response struct {
	value *TransferOrganization200Response
	isSet bool
}

func (v NullableTransferOrganization200Response) Get() *TransferOrganization200Response {
	return v.value
}

func (v *NullableTransferOrganization200Response) Set(val *TransferOrganization200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOrganization200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOrganization200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOrganization200Response(val *TransferOrganization200Response) *NullableTransferOrganization200Response {
	return &NullableTransferOrganization200Response{value: val, isSet: true}
}

func (v NullableTransferOrganization200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOrganization200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
