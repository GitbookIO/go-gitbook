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

// checks if the SearchPageResultUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPageResultUrls{}

// SearchPageResultUrls URLs associated with the object
type SearchPageResultUrls struct {
	// URL of the page in the application
	App string `json:"app"`
}

// NewSearchPageResultUrls instantiates a new SearchPageResultUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPageResultUrls(app string) *SearchPageResultUrls {
	this := SearchPageResultUrls{}
	this.App = app
	return &this
}

// NewSearchPageResultUrlsWithDefaults instantiates a new SearchPageResultUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPageResultUrlsWithDefaults() *SearchPageResultUrls {
	this := SearchPageResultUrls{}
	return &this
}

// GetApp returns the App field value
func (o *SearchPageResultUrls) GetApp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *SearchPageResultUrls) GetAppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *SearchPageResultUrls) SetApp(v string) {
	o.App = v
}

func (o SearchPageResultUrls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPageResultUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	return toSerialize, nil
}

type NullableSearchPageResultUrls struct {
	value *SearchPageResultUrls
	isSet bool
}

func (v NullableSearchPageResultUrls) Get() *SearchPageResultUrls {
	return v.value
}

func (v *NullableSearchPageResultUrls) Set(val *SearchPageResultUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPageResultUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPageResultUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPageResultUrls(val *SearchPageResultUrls) *NullableSearchPageResultUrls {
	return &NullableSearchPageResultUrls{value: val, isSet: true}
}

func (v NullableSearchPageResultUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPageResultUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
