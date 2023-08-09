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

// checks if the ChangeRequestUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestUrls{}

// ChangeRequestUrls URLs associated with the object
type ChangeRequestUrls struct {
	// URL of the space in the application
	App string `json:"app"`
}

// NewChangeRequestUrls instantiates a new ChangeRequestUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestUrls(app string) *ChangeRequestUrls {
	this := ChangeRequestUrls{}
	this.App = app
	return &this
}

// NewChangeRequestUrlsWithDefaults instantiates a new ChangeRequestUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestUrlsWithDefaults() *ChangeRequestUrls {
	this := ChangeRequestUrls{}
	return &this
}

// GetApp returns the App field value
func (o *ChangeRequestUrls) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestUrls) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *ChangeRequestUrls) SetApp(v string) {
	o.App = v
}

func (o ChangeRequestUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	return toSerialize, nil
}

type NullableChangeRequestUrls struct {
	value *ChangeRequestUrls
	isSet bool
}

func (v NullableChangeRequestUrls) Get() *ChangeRequestUrls {
	return v.value
}

func (v *NullableChangeRequestUrls) Set(val *ChangeRequestUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestUrls(val *ChangeRequestUrls) *NullableChangeRequestUrls {
	return &NullableChangeRequestUrls{value: val, isSet: true}
}

func (v NullableChangeRequestUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}