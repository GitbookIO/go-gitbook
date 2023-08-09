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

// checks if the IntegrationExternalLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationExternalLinksInner{}

// IntegrationExternalLinksInner struct for IntegrationExternalLinksInner
type IntegrationExternalLinksInner struct {
	Url   string `json:"url"`
	Label string `json:"label"`
}

// NewIntegrationExternalLinksInner instantiates a new IntegrationExternalLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationExternalLinksInner(url string, label string) *IntegrationExternalLinksInner {
	this := IntegrationExternalLinksInner{}
	this.Url = url
	this.Label = label
	return &this
}

// NewIntegrationExternalLinksInnerWithDefaults instantiates a new IntegrationExternalLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationExternalLinksInnerWithDefaults() *IntegrationExternalLinksInner {
	this := IntegrationExternalLinksInner{}
	return &this
}

// GetUrl returns the Url field value
func (o *IntegrationExternalLinksInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IntegrationExternalLinksInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IntegrationExternalLinksInner) SetUrl(v string) {
	o.Url = v
}

// GetLabel returns the Label field value
func (o *IntegrationExternalLinksInner) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *IntegrationExternalLinksInner) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *IntegrationExternalLinksInner) SetLabel(v string) {
	o.Label = v
}

func (o IntegrationExternalLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationExternalLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableIntegrationExternalLinksInner struct {
	value *IntegrationExternalLinksInner
	isSet bool
}

func (v NullableIntegrationExternalLinksInner) Get() *IntegrationExternalLinksInner {
	return v.value
}

func (v *NullableIntegrationExternalLinksInner) Set(val *IntegrationExternalLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationExternalLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationExternalLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationExternalLinksInner(val *IntegrationExternalLinksInner) *NullableIntegrationExternalLinksInner {
	return &NullableIntegrationExternalLinksInner{value: val, isSet: true}
}

func (v NullableIntegrationExternalLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationExternalLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}