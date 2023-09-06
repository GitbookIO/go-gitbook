// Copyright 2023 GitBook, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the GetContentByUrl200ResponseOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContentByUrl200ResponseOneOf1{}

// GetContentByUrl200ResponseOneOf1 URL resolved to the content of a space
type GetContentByUrl200ResponseOneOf1 struct {
	Space         Space                     `json:"space"`
	ChangeRequest *ChangeRequest            `json:"changeRequest,omitempty"`
	Page          *GetPageByPath200Response `json:"page,omitempty"`
}

// NewGetContentByUrl200ResponseOneOf1 instantiates a new GetContentByUrl200ResponseOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContentByUrl200ResponseOneOf1(space Space) *GetContentByUrl200ResponseOneOf1 {
	this := GetContentByUrl200ResponseOneOf1{}
	this.Space = space
	return &this
}

// NewGetContentByUrl200ResponseOneOf1WithDefaults instantiates a new GetContentByUrl200ResponseOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContentByUrl200ResponseOneOf1WithDefaults() *GetContentByUrl200ResponseOneOf1 {
	this := GetContentByUrl200ResponseOneOf1{}
	return &this
}

// GetSpace returns the Space field value
func (o *GetContentByUrl200ResponseOneOf1) GetSpace() Space {
	if o == nil {
		var ret Space
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *GetContentByUrl200ResponseOneOf1) GetSpaceOk() (*Space, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *GetContentByUrl200ResponseOneOf1) SetSpace(v Space) {
	o.Space = v
}

// GetChangeRequest returns the ChangeRequest field value if set, zero value otherwise.
func (o *GetContentByUrl200ResponseOneOf1) GetChangeRequest() ChangeRequest {
	if o == nil || IsNil(o.ChangeRequest) {
		var ret ChangeRequest
		return ret
	}
	return *o.ChangeRequest
}

// GetChangeRequestOk returns a tuple with the ChangeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContentByUrl200ResponseOneOf1) GetChangeRequestOk() (*ChangeRequest, bool) {
	if o == nil || IsNil(o.ChangeRequest) {
		return nil, false
	}
	return o.ChangeRequest, true
}

// HasChangeRequest returns a boolean if a field has been set.
func (o *GetContentByUrl200ResponseOneOf1) HasChangeRequest() bool {
	if o != nil && !IsNil(o.ChangeRequest) {
		return true
	}

	return false
}

// SetChangeRequest gets a reference to the given ChangeRequest and assigns it to the ChangeRequest field.
func (o *GetContentByUrl200ResponseOneOf1) SetChangeRequest(v ChangeRequest) {
	o.ChangeRequest = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetContentByUrl200ResponseOneOf1) GetPage() GetPageByPath200Response {
	if o == nil || IsNil(o.Page) {
		var ret GetPageByPath200Response
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContentByUrl200ResponseOneOf1) GetPageOk() (*GetPageByPath200Response, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetContentByUrl200ResponseOneOf1) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given GetPageByPath200Response and assigns it to the Page field.
func (o *GetContentByUrl200ResponseOneOf1) SetPage(v GetPageByPath200Response) {
	o.Page = &v
}

func (o GetContentByUrl200ResponseOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContentByUrl200ResponseOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["space"] = o.Space
	if !IsNil(o.ChangeRequest) {
		toSerialize["changeRequest"] = o.ChangeRequest
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

type NullableGetContentByUrl200ResponseOneOf1 struct {
	value *GetContentByUrl200ResponseOneOf1
	isSet bool
}

func (v NullableGetContentByUrl200ResponseOneOf1) Get() *GetContentByUrl200ResponseOneOf1 {
	return v.value
}

func (v *NullableGetContentByUrl200ResponseOneOf1) Set(val *GetContentByUrl200ResponseOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContentByUrl200ResponseOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContentByUrl200ResponseOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContentByUrl200ResponseOneOf1(val *GetContentByUrl200ResponseOneOf1) *NullableGetContentByUrl200ResponseOneOf1 {
	return &NullableGetContentByUrl200ResponseOneOf1{value: val, isSet: true}
}

func (v NullableGetContentByUrl200ResponseOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContentByUrl200ResponseOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
