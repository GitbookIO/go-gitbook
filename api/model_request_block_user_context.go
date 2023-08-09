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

// checks if the RequestBlockUserContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestBlockUserContext{}

// RequestBlockUserContext The context to send when blocking/unblocking a user
type RequestBlockUserContext struct {
	Block bool `json:"block"`
}

// NewRequestBlockUserContext instantiates a new RequestBlockUserContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestBlockUserContext(block bool) *RequestBlockUserContext {
	this := RequestBlockUserContext{}
	this.Block = block
	return &this
}

// NewRequestBlockUserContextWithDefaults instantiates a new RequestBlockUserContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestBlockUserContextWithDefaults() *RequestBlockUserContext {
	this := RequestBlockUserContext{}
	return &this
}

// GetBlock returns the Block field value
func (o *RequestBlockUserContext) GetBlock() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Block
}

// GetBlockOk returns a tuple with the Block field value
// and a boolean to check if the value has been set.
func (o *RequestBlockUserContext) GetBlockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block, true
}

// SetBlock sets field value
func (o *RequestBlockUserContext) SetBlock(v bool) {
	o.Block = v
}

func (o RequestBlockUserContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestBlockUserContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block"] = o.Block
	return toSerialize, nil
}

type NullableRequestBlockUserContext struct {
	value *RequestBlockUserContext
	isSet bool
}

func (v NullableRequestBlockUserContext) Get() *RequestBlockUserContext {
	return v.value
}

func (v *NullableRequestBlockUserContext) Set(val *RequestBlockUserContext) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestBlockUserContext) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestBlockUserContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestBlockUserContext(val *RequestBlockUserContext) *NullableRequestBlockUserContext {
	return &NullableRequestBlockUserContext{value: val, isSet: true}
}

func (v NullableRequestBlockUserContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestBlockUserContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}