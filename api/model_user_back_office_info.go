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

// checks if the UserBackOfficeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserBackOfficeInfo{}

// UserBackOfficeInfo The GitBook User info shown in the BackOffice.
type UserBackOfficeInfo struct {
	Id             string             `json:"id"`
	RiskEvaluation UserRiskEvaluation `json:"riskEvaluation"`
	AuthProviders  []FirebaseUserInfo `json:"authProviders"`
	CreatedAt      string             `json:"createdAt"`
	LastSignInAt   string             `json:"lastSignInAt"`
	Disabled       bool               `json:"disabled"`
}

// NewUserBackOfficeInfo instantiates a new UserBackOfficeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserBackOfficeInfo(id string, riskEvaluation UserRiskEvaluation, authProviders []FirebaseUserInfo, createdAt string, lastSignInAt string, disabled bool) *UserBackOfficeInfo {
	this := UserBackOfficeInfo{}
	this.Id = id
	this.RiskEvaluation = riskEvaluation
	this.AuthProviders = authProviders
	this.CreatedAt = createdAt
	this.LastSignInAt = lastSignInAt
	this.Disabled = disabled
	return &this
}

// NewUserBackOfficeInfoWithDefaults instantiates a new UserBackOfficeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserBackOfficeInfoWithDefaults() *UserBackOfficeInfo {
	this := UserBackOfficeInfo{}
	return &this
}

// GetId returns the Id field value
func (o *UserBackOfficeInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserBackOfficeInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserBackOfficeInfo) SetId(v string) {
	o.Id = v
}

// GetRiskEvaluation returns the RiskEvaluation field value
func (o *UserBackOfficeInfo) GetRiskEvaluation() UserRiskEvaluation {
	if o == nil {
		var ret UserRiskEvaluation
		return ret
	}

	return o.RiskEvaluation
}

// GetRiskEvaluationOk returns a tuple with the RiskEvaluation field value
// and a boolean to check if the value has been set.
func (o *UserBackOfficeInfo) GetRiskEvaluationOk() (*UserRiskEvaluation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskEvaluation, true
}

// SetRiskEvaluation sets field value
func (o *UserBackOfficeInfo) SetRiskEvaluation(v UserRiskEvaluation) {
	o.RiskEvaluation = v
}

// GetAuthProviders returns the AuthProviders field value
func (o *UserBackOfficeInfo) GetAuthProviders() []FirebaseUserInfo {
	if o == nil {
		var ret []FirebaseUserInfo
		return ret
	}

	return o.AuthProviders
}

// GetAuthProvidersOk returns a tuple with the AuthProviders field value
// and a boolean to check if the value has been set.
func (o *UserBackOfficeInfo) GetAuthProvidersOk() ([]FirebaseUserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthProviders, true
}

// SetAuthProviders sets field value
func (o *UserBackOfficeInfo) SetAuthProviders(v []FirebaseUserInfo) {
	o.AuthProviders = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserBackOfficeInfo) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserBackOfficeInfo) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserBackOfficeInfo) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetLastSignInAt returns the LastSignInAt field value
func (o *UserBackOfficeInfo) GetLastSignInAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastSignInAt
}

// GetLastSignInAtOk returns a tuple with the LastSignInAt field value
// and a boolean to check if the value has been set.
func (o *UserBackOfficeInfo) GetLastSignInAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSignInAt, true
}

// SetLastSignInAt sets field value
func (o *UserBackOfficeInfo) SetLastSignInAt(v string) {
	o.LastSignInAt = v
}

// GetDisabled returns the Disabled field value
func (o *UserBackOfficeInfo) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *UserBackOfficeInfo) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *UserBackOfficeInfo) SetDisabled(v bool) {
	o.Disabled = v
}

func (o UserBackOfficeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserBackOfficeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["riskEvaluation"] = o.RiskEvaluation
	toSerialize["authProviders"] = o.AuthProviders
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["lastSignInAt"] = o.LastSignInAt
	toSerialize["disabled"] = o.Disabled
	return toSerialize, nil
}

type NullableUserBackOfficeInfo struct {
	value *UserBackOfficeInfo
	isSet bool
}

func (v NullableUserBackOfficeInfo) Get() *UserBackOfficeInfo {
	return v.value
}

func (v *NullableUserBackOfficeInfo) Set(val *UserBackOfficeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserBackOfficeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserBackOfficeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserBackOfficeInfo(val *UserBackOfficeInfo) *NullableUserBackOfficeInfo {
	return &NullableUserBackOfficeInfo{value: val, isSet: true}
}

func (v NullableUserBackOfficeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserBackOfficeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
