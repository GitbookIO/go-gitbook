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

// checks if the RequestCreateChangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestCreateChangeRequest{}

// RequestCreateChangeRequest struct for RequestCreateChangeRequest
type RequestCreateChangeRequest struct {
	// Subject of the change-request
	Subject *string `json:"subject,omitempty"`
}

// NewRequestCreateChangeRequest instantiates a new RequestCreateChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestCreateChangeRequest() *RequestCreateChangeRequest {
	this := RequestCreateChangeRequest{}
	return &this
}

// NewRequestCreateChangeRequestWithDefaults instantiates a new RequestCreateChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestCreateChangeRequestWithDefaults() *RequestCreateChangeRequest {
	this := RequestCreateChangeRequest{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *RequestCreateChangeRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestCreateChangeRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *RequestCreateChangeRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *RequestCreateChangeRequest) SetSubject(v string) {
	o.Subject = &v
}

func (o RequestCreateChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestCreateChangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableRequestCreateChangeRequest struct {
	value *RequestCreateChangeRequest
	isSet bool
}

func (v NullableRequestCreateChangeRequest) Get() *RequestCreateChangeRequest {
	return v.value
}

func (v *NullableRequestCreateChangeRequest) Set(val *RequestCreateChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestCreateChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestCreateChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestCreateChangeRequest(val *RequestCreateChangeRequest) *NullableRequestCreateChangeRequest {
	return &NullableRequestCreateChangeRequest{value: val, isSet: true}
}

func (v NullableRequestCreateChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestCreateChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
