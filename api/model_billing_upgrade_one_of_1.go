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

// checks if the BillingUpgradeOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingUpgradeOneOf1{}

// BillingUpgradeOneOf1 struct for BillingUpgradeOneOf1
type BillingUpgradeOneOf1 struct {
	Result  string                `json:"result"`
	Invoice BillingInvoicePreview `json:"invoice"`
}

// NewBillingUpgradeOneOf1 instantiates a new BillingUpgradeOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingUpgradeOneOf1(result string, invoice BillingInvoicePreview) *BillingUpgradeOneOf1 {
	this := BillingUpgradeOneOf1{}
	this.Result = result
	this.Invoice = invoice
	return &this
}

// NewBillingUpgradeOneOf1WithDefaults instantiates a new BillingUpgradeOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingUpgradeOneOf1WithDefaults() *BillingUpgradeOneOf1 {
	this := BillingUpgradeOneOf1{}
	return &this
}

// GetResult returns the Result field value
func (o *BillingUpgradeOneOf1) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *BillingUpgradeOneOf1) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *BillingUpgradeOneOf1) SetResult(v string) {
	o.Result = v
}

// GetInvoice returns the Invoice field value
func (o *BillingUpgradeOneOf1) GetInvoice() BillingInvoicePreview {
	if o == nil {
		var ret BillingInvoicePreview
		return ret
	}

	return o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value
// and a boolean to check if the value has been set.
func (o *BillingUpgradeOneOf1) GetInvoiceOk() (*BillingInvoicePreview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoice, true
}

// SetInvoice sets field value
func (o *BillingUpgradeOneOf1) SetInvoice(v BillingInvoicePreview) {
	o.Invoice = v
}

func (o BillingUpgradeOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingUpgradeOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["invoice"] = o.Invoice
	return toSerialize, nil
}

type NullableBillingUpgradeOneOf1 struct {
	value *BillingUpgradeOneOf1
	isSet bool
}

func (v NullableBillingUpgradeOneOf1) Get() *BillingUpgradeOneOf1 {
	return v.value
}

func (v *NullableBillingUpgradeOneOf1) Set(val *BillingUpgradeOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingUpgradeOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingUpgradeOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingUpgradeOneOf1(val *BillingUpgradeOneOf1) *NullableBillingUpgradeOneOf1 {
	return &NullableBillingUpgradeOneOf1{value: val, isSet: true}
}

func (v NullableBillingUpgradeOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingUpgradeOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
