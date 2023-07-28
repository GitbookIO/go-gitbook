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

// checks if the UnsplashImageAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnsplashImageAuthor{}

// UnsplashImageAuthor struct for UnsplashImageAuthor
type UnsplashImageAuthor struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// NewUnsplashImageAuthor instantiates a new UnsplashImageAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsplashImageAuthor(name string, url string) *UnsplashImageAuthor {
	this := UnsplashImageAuthor{}
	this.Name = name
	this.Url = url
	return &this
}

// NewUnsplashImageAuthorWithDefaults instantiates a new UnsplashImageAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsplashImageAuthorWithDefaults() *UnsplashImageAuthor {
	this := UnsplashImageAuthor{}
	return &this
}

// GetName returns the Name field value
func (o *UnsplashImageAuthor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UnsplashImageAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UnsplashImageAuthor) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *UnsplashImageAuthor) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UnsplashImageAuthor) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UnsplashImageAuthor) SetUrl(v string) {
	o.Url = v
}

func (o UnsplashImageAuthor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnsplashImageAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableUnsplashImageAuthor struct {
	value *UnsplashImageAuthor
	isSet bool
}

func (v NullableUnsplashImageAuthor) Get() *UnsplashImageAuthor {
	return v.value
}

func (v *NullableUnsplashImageAuthor) Set(val *UnsplashImageAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsplashImageAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsplashImageAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsplashImageAuthor(val *UnsplashImageAuthor) *NullableUnsplashImageAuthor {
	return &NullableUnsplashImageAuthor{value: val, isSet: true}
}

func (v NullableUnsplashImageAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsplashImageAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
