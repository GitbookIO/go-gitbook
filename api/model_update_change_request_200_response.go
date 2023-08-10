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

// checks if the UpdateChangeRequest200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateChangeRequest200Response{}

// UpdateChangeRequest200Response struct for UpdateChangeRequest200Response
type UpdateChangeRequest200Response struct {
	// ID of the resulting revision
	Revision string `json:"revision"`
	Result   string `json:"result"`
}

// NewUpdateChangeRequest200Response instantiates a new UpdateChangeRequest200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateChangeRequest200Response(revision string, result string) *UpdateChangeRequest200Response {
	this := UpdateChangeRequest200Response{}
	this.Revision = revision
	this.Result = result
	return &this
}

// NewUpdateChangeRequest200ResponseWithDefaults instantiates a new UpdateChangeRequest200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateChangeRequest200ResponseWithDefaults() *UpdateChangeRequest200Response {
	this := UpdateChangeRequest200Response{}
	return &this
}

// GetRevision returns the Revision field value
func (o *UpdateChangeRequest200Response) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *UpdateChangeRequest200Response) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *UpdateChangeRequest200Response) SetRevision(v string) {
	o.Revision = v
}

// GetResult returns the Result field value
func (o *UpdateChangeRequest200Response) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *UpdateChangeRequest200Response) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *UpdateChangeRequest200Response) SetResult(v string) {
	o.Result = v
}

func (o UpdateChangeRequest200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateChangeRequest200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revision"] = o.Revision
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullableUpdateChangeRequest200Response struct {
	value *UpdateChangeRequest200Response
	isSet bool
}

func (v NullableUpdateChangeRequest200Response) Get() *UpdateChangeRequest200Response {
	return v.value
}

func (v *NullableUpdateChangeRequest200Response) Set(val *UpdateChangeRequest200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateChangeRequest200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateChangeRequest200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateChangeRequest200Response(val *UpdateChangeRequest200Response) *NullableUpdateChangeRequest200Response {
	return &NullableUpdateChangeRequest200Response{value: val, isSet: true}
}

func (v NullableUpdateChangeRequest200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateChangeRequest200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
