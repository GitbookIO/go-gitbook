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

// checks if the ContentKitActionAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitActionAnyOf{}

// ContentKitActionAnyOf Custom action to re-render the block.
type ContentKitActionAnyOf struct {
	Action               string `json:"action"`
	AdditionalProperties map[string]interface{}
}

type _ContentKitActionAnyOf ContentKitActionAnyOf

// NewContentKitActionAnyOf instantiates a new ContentKitActionAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitActionAnyOf(action string) *ContentKitActionAnyOf {
	this := ContentKitActionAnyOf{}
	this.Action = action
	return &this
}

// NewContentKitActionAnyOfWithDefaults instantiates a new ContentKitActionAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitActionAnyOfWithDefaults() *ContentKitActionAnyOf {
	this := ContentKitActionAnyOf{}
	return &this
}

// GetAction returns the Action field value
func (o *ContentKitActionAnyOf) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContentKitActionAnyOf) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContentKitActionAnyOf) SetAction(v string) {
	o.Action = v
}

func (o ContentKitActionAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitActionAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContentKitActionAnyOf) UnmarshalJSON(bytes []byte) (err error) {
	varContentKitActionAnyOf := _ContentKitActionAnyOf{}

	if err = json.Unmarshal(bytes, &varContentKitActionAnyOf); err == nil {
		*o = ContentKitActionAnyOf(varContentKitActionAnyOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContentKitActionAnyOf struct {
	value *ContentKitActionAnyOf
	isSet bool
}

func (v NullableContentKitActionAnyOf) Get() *ContentKitActionAnyOf {
	return v.value
}

func (v *NullableContentKitActionAnyOf) Set(val *ContentKitActionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitActionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitActionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitActionAnyOf(val *ContentKitActionAnyOf) *NullableContentKitActionAnyOf {
	return &NullableContentKitActionAnyOf{value: val, isSet: true}
}

func (v NullableContentKitActionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitActionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}