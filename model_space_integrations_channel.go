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

// checks if the SpaceIntegrationsChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceIntegrationsChannel{}

// SpaceIntegrationsChannel struct for SpaceIntegrationsChannel
type SpaceIntegrationsChannel struct {
	Channel string `json:"channel"`
	Space   string `json:"space"`
}

// NewSpaceIntegrationsChannel instantiates a new SpaceIntegrationsChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceIntegrationsChannel(channel string, space string) *SpaceIntegrationsChannel {
	this := SpaceIntegrationsChannel{}
	this.Channel = channel
	this.Space = space
	return &this
}

// NewSpaceIntegrationsChannelWithDefaults instantiates a new SpaceIntegrationsChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceIntegrationsChannelWithDefaults() *SpaceIntegrationsChannel {
	this := SpaceIntegrationsChannel{}
	return &this
}

// GetChannel returns the Channel field value
func (o *SpaceIntegrationsChannel) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SpaceIntegrationsChannel) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *SpaceIntegrationsChannel) SetChannel(v string) {
	o.Channel = v
}

// GetSpace returns the Space field value
func (o *SpaceIntegrationsChannel) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *SpaceIntegrationsChannel) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *SpaceIntegrationsChannel) SetSpace(v string) {
	o.Space = v
}

func (o SpaceIntegrationsChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceIntegrationsChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["space"] = o.Space
	return toSerialize, nil
}

type NullableSpaceIntegrationsChannel struct {
	value *SpaceIntegrationsChannel
	isSet bool
}

func (v NullableSpaceIntegrationsChannel) Get() *SpaceIntegrationsChannel {
	return v.value
}

func (v *NullableSpaceIntegrationsChannel) Set(val *SpaceIntegrationsChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceIntegrationsChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceIntegrationsChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceIntegrationsChannel(val *SpaceIntegrationsChannel) *NullableSpaceIntegrationsChannel {
	return &NullableSpaceIntegrationsChannel{value: val, isSet: true}
}

func (v NullableSpaceIntegrationsChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceIntegrationsChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
