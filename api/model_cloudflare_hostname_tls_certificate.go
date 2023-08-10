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

// checks if the CloudflareHostnameTLSCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudflareHostnameTLSCertificate{}

// CloudflareHostnameTLSCertificate The Cloudflare Hostname TLS certificate
type CloudflareHostnameTLSCertificate struct {
	Issuer    *string `json:"issuer,omitempty"`
	ExpiresOn *string `json:"expiresOn,omitempty"`
	IssuedOn  *string `json:"issuedOn,omitempty"`
}

// NewCloudflareHostnameTLSCertificate instantiates a new CloudflareHostnameTLSCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudflareHostnameTLSCertificate() *CloudflareHostnameTLSCertificate {
	this := CloudflareHostnameTLSCertificate{}
	return &this
}

// NewCloudflareHostnameTLSCertificateWithDefaults instantiates a new CloudflareHostnameTLSCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudflareHostnameTLSCertificateWithDefaults() *CloudflareHostnameTLSCertificate {
	this := CloudflareHostnameTLSCertificate{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *CloudflareHostnameTLSCertificate) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSCertificate) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *CloudflareHostnameTLSCertificate) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *CloudflareHostnameTLSCertificate) SetIssuer(v string) {
	o.Issuer = &v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *CloudflareHostnameTLSCertificate) GetExpiresOn() string {
	if o == nil || IsNil(o.ExpiresOn) {
		var ret string
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSCertificate) GetExpiresOnOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresOn) {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *CloudflareHostnameTLSCertificate) HasExpiresOn() bool {
	if o != nil && !IsNil(o.ExpiresOn) {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given string and assigns it to the ExpiresOn field.
func (o *CloudflareHostnameTLSCertificate) SetExpiresOn(v string) {
	o.ExpiresOn = &v
}

// GetIssuedOn returns the IssuedOn field value if set, zero value otherwise.
func (o *CloudflareHostnameTLSCertificate) GetIssuedOn() string {
	if o == nil || IsNil(o.IssuedOn) {
		var ret string
		return ret
	}
	return *o.IssuedOn
}

// GetIssuedOnOk returns a tuple with the IssuedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSCertificate) GetIssuedOnOk() (*string, bool) {
	if o == nil || IsNil(o.IssuedOn) {
		return nil, false
	}
	return o.IssuedOn, true
}

// HasIssuedOn returns a boolean if a field has been set.
func (o *CloudflareHostnameTLSCertificate) HasIssuedOn() bool {
	if o != nil && !IsNil(o.IssuedOn) {
		return true
	}

	return false
}

// SetIssuedOn gets a reference to the given string and assigns it to the IssuedOn field.
func (o *CloudflareHostnameTLSCertificate) SetIssuedOn(v string) {
	o.IssuedOn = &v
}

func (o CloudflareHostnameTLSCertificate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudflareHostnameTLSCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.ExpiresOn) {
		toSerialize["expiresOn"] = o.ExpiresOn
	}
	if !IsNil(o.IssuedOn) {
		toSerialize["issuedOn"] = o.IssuedOn
	}
	return toSerialize, nil
}

type NullableCloudflareHostnameTLSCertificate struct {
	value *CloudflareHostnameTLSCertificate
	isSet bool
}

func (v NullableCloudflareHostnameTLSCertificate) Get() *CloudflareHostnameTLSCertificate {
	return v.value
}

func (v *NullableCloudflareHostnameTLSCertificate) Set(val *CloudflareHostnameTLSCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudflareHostnameTLSCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudflareHostnameTLSCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudflareHostnameTLSCertificate(val *CloudflareHostnameTLSCertificate) *NullableCloudflareHostnameTLSCertificate {
	return &NullableCloudflareHostnameTLSCertificate{value: val, isSet: true}
}

func (v NullableCloudflareHostnameTLSCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudflareHostnameTLSCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
