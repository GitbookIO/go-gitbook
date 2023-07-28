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

// checks if the GetChangeRequestById404ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChangeRequestById404ResponseError{}

// GetChangeRequestById404ResponseError struct for GetChangeRequestById404ResponseError
type GetChangeRequestById404ResponseError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NewGetChangeRequestById404ResponseError instantiates a new GetChangeRequestById404ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChangeRequestById404ResponseError(code int32, message string) *GetChangeRequestById404ResponseError {
	this := GetChangeRequestById404ResponseError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewGetChangeRequestById404ResponseErrorWithDefaults instantiates a new GetChangeRequestById404ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChangeRequestById404ResponseErrorWithDefaults() *GetChangeRequestById404ResponseError {
	this := GetChangeRequestById404ResponseError{}
	return &this
}

// GetCode returns the Code field value
func (o *GetChangeRequestById404ResponseError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetChangeRequestById404ResponseError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GetChangeRequestById404ResponseError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *GetChangeRequestById404ResponseError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetChangeRequestById404ResponseError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GetChangeRequestById404ResponseError) SetMessage(v string) {
	o.Message = v
}

func (o GetChangeRequestById404ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChangeRequestById404ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableGetChangeRequestById404ResponseError struct {
	value *GetChangeRequestById404ResponseError
	isSet bool
}

func (v NullableGetChangeRequestById404ResponseError) Get() *GetChangeRequestById404ResponseError {
	return v.value
}

func (v *NullableGetChangeRequestById404ResponseError) Set(val *GetChangeRequestById404ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChangeRequestById404ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChangeRequestById404ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChangeRequestById404ResponseError(val *GetChangeRequestById404ResponseError) *NullableGetChangeRequestById404ResponseError {
	return &NullableGetChangeRequestById404ResponseError{value: val, isSet: true}
}

func (v NullableGetChangeRequestById404ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChangeRequestById404ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
