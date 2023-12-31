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

// checks if the BillingUpgradeOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingUpgradeOneOf3{}

// BillingUpgradeOneOf3 struct for BillingUpgradeOneOf3
type BillingUpgradeOneOf3 struct {
	Result string `json:"result"`
}

// NewBillingUpgradeOneOf3 instantiates a new BillingUpgradeOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingUpgradeOneOf3(result string) *BillingUpgradeOneOf3 {
	this := BillingUpgradeOneOf3{}
	this.Result = result
	return &this
}

// NewBillingUpgradeOneOf3WithDefaults instantiates a new BillingUpgradeOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingUpgradeOneOf3WithDefaults() *BillingUpgradeOneOf3 {
	this := BillingUpgradeOneOf3{}
	return &this
}

// GetResult returns the Result field value
func (o *BillingUpgradeOneOf3) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *BillingUpgradeOneOf3) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *BillingUpgradeOneOf3) SetResult(v string) {
	o.Result = v
}

func (o BillingUpgradeOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingUpgradeOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableBillingUpgradeOneOf3 struct {
	value *BillingUpgradeOneOf3
	isSet bool
}

func (v NullableBillingUpgradeOneOf3) Get() *BillingUpgradeOneOf3 {
	return v.value
}

func (v *NullableBillingUpgradeOneOf3) Set(val *BillingUpgradeOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingUpgradeOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingUpgradeOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingUpgradeOneOf3(val *BillingUpgradeOneOf3) *NullableBillingUpgradeOneOf3 {
	return &NullableBillingUpgradeOneOf3{value: val, isSet: true}
}

func (v NullableBillingUpgradeOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingUpgradeOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
