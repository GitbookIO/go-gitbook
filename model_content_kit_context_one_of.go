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

// checks if the ContentKitContextOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitContextOneOf{}

// ContentKitContextOneOf struct for ContentKitContextOneOf
type ContentKitContextOneOf struct {
	Type     string `json:"type"`
	Editable bool   `json:"editable"`
	Theme    string `json:"theme"`
}

// NewContentKitContextOneOf instantiates a new ContentKitContextOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitContextOneOf(type_ string, editable bool, theme string) *ContentKitContextOneOf {
	this := ContentKitContextOneOf{}
	this.Type = type_
	this.Editable = editable
	this.Theme = theme
	return &this
}

// NewContentKitContextOneOfWithDefaults instantiates a new ContentKitContextOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitContextOneOfWithDefaults() *ContentKitContextOneOf {
	this := ContentKitContextOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitContextOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitContextOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitContextOneOf) SetType(v string) {
	o.Type = v
}

// GetEditable returns the Editable field value
func (o *ContentKitContextOneOf) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *ContentKitContextOneOf) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value
func (o *ContentKitContextOneOf) SetEditable(v bool) {
	o.Editable = v
}

// GetTheme returns the Theme field value
func (o *ContentKitContextOneOf) GetTheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *ContentKitContextOneOf) GetThemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value
func (o *ContentKitContextOneOf) SetTheme(v string) {
	o.Theme = v
}

func (o ContentKitContextOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitContextOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["editable"] = o.Editable
	toSerialize["theme"] = o.Theme
	return toSerialize, nil
}

type NullableContentKitContextOneOf struct {
	value *ContentKitContextOneOf
	isSet bool
}

func (v NullableContentKitContextOneOf) Get() *ContentKitContextOneOf {
	return v.value
}

func (v *NullableContentKitContextOneOf) Set(val *ContentKitContextOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitContextOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitContextOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitContextOneOf(val *ContentKitContextOneOf) *NullableContentKitContextOneOf {
	return &NullableContentKitContextOneOf{value: val, isSet: true}
}

func (v NullableContentKitContextOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitContextOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
