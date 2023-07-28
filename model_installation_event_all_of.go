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

// checks if the InstallationEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallationEventAllOf{}

// InstallationEventAllOf Common properties for all events related to an installation
type InstallationEventAllOf struct {
	// ID of the integration installation
	InstallationId string `json:"installationId"`
}

// NewInstallationEventAllOf instantiates a new InstallationEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallationEventAllOf(installationId string) *InstallationEventAllOf {
	this := InstallationEventAllOf{}
	this.InstallationId = installationId
	return &this
}

// NewInstallationEventAllOfWithDefaults instantiates a new InstallationEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallationEventAllOfWithDefaults() *InstallationEventAllOf {
	this := InstallationEventAllOf{}
	return &this
}

// GetInstallationId returns the InstallationId field value
func (o *InstallationEventAllOf) GetInstallationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
func (o *InstallationEventAllOf) GetInstallationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallationId, true
}

// SetInstallationId sets field value
func (o *InstallationEventAllOf) SetInstallationId(v string) {
	o.InstallationId = v
}

func (o InstallationEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallationEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["installationId"] = o.InstallationId
	return toSerialize, nil
}

type NullableInstallationEventAllOf struct {
	value *InstallationEventAllOf
	isSet bool
}

func (v NullableInstallationEventAllOf) Get() *InstallationEventAllOf {
	return v.value
}

func (v *NullableInstallationEventAllOf) Set(val *InstallationEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationEventAllOf(val *InstallationEventAllOf) *NullableInstallationEventAllOf {
	return &NullableInstallationEventAllOf{value: val, isSet: true}
}

func (v NullableInstallationEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}