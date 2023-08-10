/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the ErrorError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorError{}

// ErrorError struct for ErrorError
type ErrorError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NewErrorError instantiates a new ErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorError(code int32, message string) *ErrorError {
	this := ErrorError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorErrorWithDefaults instantiates a new ErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorErrorWithDefaults() *ErrorError {
	this := ErrorError{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorError) SetMessage(v string) {
	o.Message = v
}

func (o ErrorError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableErrorError struct {
	value *ErrorError
	isSet bool
}

func (v NullableErrorError) Get() *ErrorError {
	return v.value
}

func (v *NullableErrorError) Set(val *ErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorError(val *ErrorError) *NullableErrorError {
	return &NullableErrorError{value: val, isSet: true}
}

func (v NullableErrorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
