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

// checks if the IntegrationInstallationUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationInstallationUrls{}

// IntegrationInstallationUrls URLs associated with the object
type IntegrationInstallationUrls struct {
	// URL of the integration's installation in the application
	App string `json:"app"`
	// Public HTTP endpoint for the integration's installation
	PublicEndpoint string `json:"publicEndpoint"`
}

// NewIntegrationInstallationUrls instantiates a new IntegrationInstallationUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationInstallationUrls(app string, publicEndpoint string) *IntegrationInstallationUrls {
	this := IntegrationInstallationUrls{}
	this.App = app
	this.PublicEndpoint = publicEndpoint
	return &this
}

// NewIntegrationInstallationUrlsWithDefaults instantiates a new IntegrationInstallationUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationInstallationUrlsWithDefaults() *IntegrationInstallationUrls {
	this := IntegrationInstallationUrls{}
	return &this
}

// GetApp returns the App field value
func (o *IntegrationInstallationUrls) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallationUrls) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *IntegrationInstallationUrls) SetApp(v string) {
	o.App = v
}

// GetPublicEndpoint returns the PublicEndpoint field value
func (o *IntegrationInstallationUrls) GetPublicEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicEndpoint
}

// GetPublicEndpointOk returns a tuple with the PublicEndpoint field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallationUrls) GetPublicEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicEndpoint, true
}

// SetPublicEndpoint sets field value
func (o *IntegrationInstallationUrls) SetPublicEndpoint(v string) {
	o.PublicEndpoint = v
}

func (o IntegrationInstallationUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationInstallationUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	toSerialize["publicEndpoint"] = o.PublicEndpoint
	return toSerialize, nil
}

type NullableIntegrationInstallationUrls struct {
	value *IntegrationInstallationUrls
	isSet bool
}

func (v NullableIntegrationInstallationUrls) Get() *IntegrationInstallationUrls {
	return v.value
}

func (v *NullableIntegrationInstallationUrls) Set(val *IntegrationInstallationUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationInstallationUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationInstallationUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationInstallationUrls(val *IntegrationInstallationUrls) *NullableIntegrationInstallationUrls {
	return &NullableIntegrationInstallationUrls{value: val, isSet: true}
}

func (v NullableIntegrationInstallationUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationInstallationUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
