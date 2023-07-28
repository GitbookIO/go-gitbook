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

// checks if the ContentKitWebFrameSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitWebFrameSource{}

// ContentKitWebFrameSource Content to load in the frame.
type ContentKitWebFrameSource struct {
	Url string `json:"url"`
}

// NewContentKitWebFrameSource instantiates a new ContentKitWebFrameSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitWebFrameSource(url string) *ContentKitWebFrameSource {
	this := ContentKitWebFrameSource{}
	this.Url = url
	return &this
}

// NewContentKitWebFrameSourceWithDefaults instantiates a new ContentKitWebFrameSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitWebFrameSourceWithDefaults() *ContentKitWebFrameSource {
	this := ContentKitWebFrameSource{}
	return &this
}

// GetUrl returns the Url field value
func (o *ContentKitWebFrameSource) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ContentKitWebFrameSource) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ContentKitWebFrameSource) SetUrl(v string) {
	o.Url = v
}

func (o ContentKitWebFrameSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitWebFrameSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableContentKitWebFrameSource struct {
	value *ContentKitWebFrameSource
	isSet bool
}

func (v NullableContentKitWebFrameSource) Get() *ContentKitWebFrameSource {
	return v.value
}

func (v *NullableContentKitWebFrameSource) Set(val *ContentKitWebFrameSource) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitWebFrameSource) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitWebFrameSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitWebFrameSource(val *ContentKitWebFrameSource) *NullableContentKitWebFrameSource {
	return &NullableContentKitWebFrameSource{value: val, isSet: true}
}

func (v NullableContentKitWebFrameSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitWebFrameSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}