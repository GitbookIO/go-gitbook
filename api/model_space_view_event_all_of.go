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

// checks if the SpaceViewEventAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceViewEventAllOf{}

// SpaceViewEventAllOf Event received when a page has been visited.
type SpaceViewEventAllOf struct {
	Type string `json:"type"`
	// Unique identifier of the visited page.
	PageId  *string                    `json:"pageId,omitempty"`
	Visitor SpaceViewEventAllOfVisitor `json:"visitor"`
	// The GitBook content's URL visited (including URL params).
	Url string `json:"url"`
	// The URL of referrer that linked to the page.
	Referrer string `json:"referrer"`
}

// NewSpaceViewEventAllOf instantiates a new SpaceViewEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceViewEventAllOf(type_ string, visitor SpaceViewEventAllOfVisitor, url string, referrer string) *SpaceViewEventAllOf {
	this := SpaceViewEventAllOf{}
	this.Type = type_
	this.Visitor = visitor
	this.Url = url
	this.Referrer = referrer
	return &this
}

// NewSpaceViewEventAllOfWithDefaults instantiates a new SpaceViewEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceViewEventAllOfWithDefaults() *SpaceViewEventAllOf {
	this := SpaceViewEventAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *SpaceViewEventAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceViewEventAllOf) SetType(v string) {
	o.Type = v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *SpaceViewEventAllOf) GetPageId() string {
	if o == nil || IsNil(o.PageId) {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOf) GetPageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PageId) {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *SpaceViewEventAllOf) HasPageId() bool {
	if o != nil && !IsNil(o.PageId) {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *SpaceViewEventAllOf) SetPageId(v string) {
	o.PageId = &v
}

// GetVisitor returns the Visitor field value
func (o *SpaceViewEventAllOf) GetVisitor() SpaceViewEventAllOfVisitor {
	if o == nil {
		var ret SpaceViewEventAllOfVisitor
		return ret
	}

	return o.Visitor
}

// GetVisitorOk returns a tuple with the Visitor field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOf) GetVisitorOk() (*SpaceViewEventAllOfVisitor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visitor, true
}

// SetVisitor sets field value
func (o *SpaceViewEventAllOf) SetVisitor(v SpaceViewEventAllOfVisitor) {
	o.Visitor = v
}

// GetUrl returns the Url field value
func (o *SpaceViewEventAllOf) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOf) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SpaceViewEventAllOf) SetUrl(v string) {
	o.Url = v
}

// GetReferrer returns the Referrer field value
func (o *SpaceViewEventAllOf) GetReferrer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Referrer
}

// GetReferrerOk returns a tuple with the Referrer field value
// and a boolean to check if the value has been set.
func (o *SpaceViewEventAllOf) GetReferrerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Referrer, true
}

// SetReferrer sets field value
func (o *SpaceViewEventAllOf) SetReferrer(v string) {
	o.Referrer = v
}

func (o SpaceViewEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceViewEventAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.PageId) {
		toSerialize["pageId"] = o.PageId
	}
	toSerialize["visitor"] = o.Visitor
	toSerialize["url"] = o.Url
	toSerialize["referrer"] = o.Referrer
	return toSerialize, nil
}

type NullableSpaceViewEventAllOf struct {
	value *SpaceViewEventAllOf
	isSet bool
}

func (v NullableSpaceViewEventAllOf) Get() *SpaceViewEventAllOf {
	return v.value
}

func (v *NullableSpaceViewEventAllOf) Set(val *SpaceViewEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceViewEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceViewEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceViewEventAllOf(val *SpaceViewEventAllOf) *NullableSpaceViewEventAllOf {
	return &NullableSpaceViewEventAllOf{value: val, isSet: true}
}

func (v NullableSpaceViewEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceViewEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
