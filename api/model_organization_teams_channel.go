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

// checks if the OrganizationTeamsChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationTeamsChannel{}

// OrganizationTeamsChannel struct for OrganizationTeamsChannel
type OrganizationTeamsChannel struct {
	Channel      string `json:"channel"`
	Organization string `json:"organization"`
}

// NewOrganizationTeamsChannel instantiates a new OrganizationTeamsChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationTeamsChannel(channel string, organization string) *OrganizationTeamsChannel {
	this := OrganizationTeamsChannel{}
	this.Channel = channel
	this.Organization = organization
	return &this
}

// NewOrganizationTeamsChannelWithDefaults instantiates a new OrganizationTeamsChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationTeamsChannelWithDefaults() *OrganizationTeamsChannel {
	this := OrganizationTeamsChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *OrganizationTeamsChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeamsChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *OrganizationTeamsChannel) SetChannel(v string) {
	o.Channel = v
}

// GetOrganization returns the Organization field value
func (o *OrganizationTeamsChannel) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *OrganizationTeamsChannel) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *OrganizationTeamsChannel) SetOrganization(v string) {
	o.Organization = v
}

func (o OrganizationTeamsChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationTeamsChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["organization"] = o.Organization
	return toSerialize, nil
}

type NullableOrganizationTeamsChannel struct {
	value *OrganizationTeamsChannel
	isSet bool
}

func (v NullableOrganizationTeamsChannel) Get() *OrganizationTeamsChannel {
	return v.value
}

func (v *NullableOrganizationTeamsChannel) Set(val *OrganizationTeamsChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationTeamsChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationTeamsChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationTeamsChannel(val *OrganizationTeamsChannel) *NullableOrganizationTeamsChannel {
	return &NullableOrganizationTeamsChannel{value: val, isSet: true}
}

func (v NullableOrganizationTeamsChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationTeamsChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
