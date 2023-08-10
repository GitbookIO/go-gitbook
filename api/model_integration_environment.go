/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the IntegrationEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationEnvironment{}

// IntegrationEnvironment Runtime environment provided during the execution of integration's code.
type IntegrationEnvironment struct {
	// URL of the HTTP API
	ApiEndpoint string `json:"apiEndpoint"`
	// Authentication token to use with the HTTP API
	AuthToken         *string                       `json:"authToken,omitempty"`
	Integration       Integration                   `json:"integration"`
	Installation      *IntegrationInstallation      `json:"installation,omitempty"`
	SpaceInstallation *IntegrationSpaceInstallation `json:"spaceInstallation,omitempty"`
	// Secrets stored on the integration and passed at runtime.
	Secrets map[string]string `json:"secrets"`
}

// NewIntegrationEnvironment instantiates a new IntegrationEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationEnvironment(apiEndpoint string, integration Integration, secrets map[string]string) *IntegrationEnvironment {
	this := IntegrationEnvironment{}
	this.ApiEndpoint = apiEndpoint
	this.Integration = integration
	this.Secrets = secrets
	return &this
}

// NewIntegrationEnvironmentWithDefaults instantiates a new IntegrationEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationEnvironmentWithDefaults() *IntegrationEnvironment {
	this := IntegrationEnvironment{}
	return &this
}

// GetApiEndpoint returns the ApiEndpoint field value
func (o *IntegrationEnvironment) GetApiEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiEndpoint
}

// GetApiEndpointOk returns a tuple with the ApiEndpoint field value
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironment) GetApiEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiEndpoint, true
}

// SetApiEndpoint sets field value
func (o *IntegrationEnvironment) SetApiEndpoint(v string) {
	o.ApiEndpoint = v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *IntegrationEnvironment) GetAuthToken() string {
	if o == nil || IsNil(o.AuthToken) {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironment) GetAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AuthToken) {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *IntegrationEnvironment) HasAuthToken() bool {
	if o != nil && !IsNil(o.AuthToken) {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *IntegrationEnvironment) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetIntegration returns the Integration field value
func (o *IntegrationEnvironment) GetIntegration() Integration {
	if o == nil {
		var ret Integration
		return ret
	}

	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironment) GetIntegrationOk() (*Integration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value
func (o *IntegrationEnvironment) SetIntegration(v Integration) {
	o.Integration = v
}

// GetInstallation returns the Installation field value if set, zero value otherwise.
func (o *IntegrationEnvironment) GetInstallation() IntegrationInstallation {
	if o == nil || IsNil(o.Installation) {
		var ret IntegrationInstallation
		return ret
	}
	return *o.Installation
}

// GetInstallationOk returns a tuple with the Installation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironment) GetInstallationOk() (*IntegrationInstallation, bool) {
	if o == nil || IsNil(o.Installation) {
		return nil, false
	}
	return o.Installation, true
}

// HasInstallation returns a boolean if a field has been set.
func (o *IntegrationEnvironment) HasInstallation() bool {
	if o != nil && !IsNil(o.Installation) {
		return true
	}

	return false
}

// SetInstallation gets a reference to the given IntegrationInstallation and assigns it to the Installation field.
func (o *IntegrationEnvironment) SetInstallation(v IntegrationInstallation) {
	o.Installation = &v
}

// GetSpaceInstallation returns the SpaceInstallation field value if set, zero value otherwise.
func (o *IntegrationEnvironment) GetSpaceInstallation() IntegrationSpaceInstallation {
	if o == nil || IsNil(o.SpaceInstallation) {
		var ret IntegrationSpaceInstallation
		return ret
	}
	return *o.SpaceInstallation
}

// GetSpaceInstallationOk returns a tuple with the SpaceInstallation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironment) GetSpaceInstallationOk() (*IntegrationSpaceInstallation, bool) {
	if o == nil || IsNil(o.SpaceInstallation) {
		return nil, false
	}
	return o.SpaceInstallation, true
}

// HasSpaceInstallation returns a boolean if a field has been set.
func (o *IntegrationEnvironment) HasSpaceInstallation() bool {
	if o != nil && !IsNil(o.SpaceInstallation) {
		return true
	}

	return false
}

// SetSpaceInstallation gets a reference to the given IntegrationSpaceInstallation and assigns it to the SpaceInstallation field.
func (o *IntegrationEnvironment) SetSpaceInstallation(v IntegrationSpaceInstallation) {
	o.SpaceInstallation = &v
}

// GetSecrets returns the Secrets field value
func (o *IntegrationEnvironment) GetSecrets() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
func (o *IntegrationEnvironment) GetSecretsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secrets, true
}

// SetSecrets sets field value
func (o *IntegrationEnvironment) SetSecrets(v map[string]string) {
	o.Secrets = v
}

func (o IntegrationEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiEndpoint"] = o.ApiEndpoint
	if !IsNil(o.AuthToken) {
		toSerialize["authToken"] = o.AuthToken
	}
	toSerialize["integration"] = o.Integration
	if !IsNil(o.Installation) {
		toSerialize["installation"] = o.Installation
	}
	if !IsNil(o.SpaceInstallation) {
		toSerialize["spaceInstallation"] = o.SpaceInstallation
	}
	toSerialize["secrets"] = o.Secrets
	return toSerialize, nil
}

type NullableIntegrationEnvironment struct {
	value *IntegrationEnvironment
	isSet bool
}

func (v NullableIntegrationEnvironment) Get() *IntegrationEnvironment {
	return v.value
}

func (v *NullableIntegrationEnvironment) Set(val *IntegrationEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationEnvironment(val *IntegrationEnvironment) *NullableIntegrationEnvironment {
	return &NullableIntegrationEnvironment{value: val, isSet: true}
}

func (v NullableIntegrationEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
