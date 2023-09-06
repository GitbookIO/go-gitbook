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

// checks if the BillingPortal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPortal{}

// BillingPortal struct for BillingPortal
type BillingPortal struct {
	// URL to the billing portal for an organization
	Url string `json:"url"`
}

// NewBillingPortal instantiates a new BillingPortal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPortal(url string) *BillingPortal {
	this := BillingPortal{}
	this.Url = url
	return &this
}

// NewBillingPortalWithDefaults instantiates a new BillingPortal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPortalWithDefaults() *BillingPortal {
	this := BillingPortal{}
	return &this
}

// GetUrl returns the Url field value
func (o *BillingPortal) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BillingPortal) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BillingPortal) SetUrl(v string) {
	o.Url = v
}

func (o BillingPortal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPortal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableBillingPortal struct {
	value *BillingPortal
	isSet bool
}

func (v NullableBillingPortal) Get() *BillingPortal {
	return v.value
}

func (v *NullableBillingPortal) Set(val *BillingPortal) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPortal) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPortal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPortal(val *BillingPortal) *NullableBillingPortal {
	return &NullableBillingPortal{value: val, isSet: true}
}

func (v NullableBillingPortal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPortal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
