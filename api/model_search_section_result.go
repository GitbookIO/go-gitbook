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

// checks if the SearchSectionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSectionResult{}

// SearchSectionResult Search result representing a section in a page.
type SearchSectionResult struct {
	Id    string                  `json:"id"`
	Title string                  `json:"title"`
	Path  string                  `json:"path"`
	Body  string                  `json:"body"`
	Urls  SearchSectionResultUrls `json:"urls"`
}

// NewSearchSectionResult instantiates a new SearchSectionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSectionResult(id string, title string, path string, body string, urls SearchSectionResultUrls) *SearchSectionResult {
	this := SearchSectionResult{}
	this.Id = id
	this.Title = title
	this.Path = path
	this.Body = body
	this.Urls = urls
	return &this
}

// NewSearchSectionResultWithDefaults instantiates a new SearchSectionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSectionResultWithDefaults() *SearchSectionResult {
	this := SearchSectionResult{}
	return &this
}

// GetId returns the Id field value
func (o *SearchSectionResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SearchSectionResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SearchSectionResult) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *SearchSectionResult) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SearchSectionResult) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SearchSectionResult) SetTitle(v string) {
	o.Title = v
}

// GetPath returns the Path field value
func (o *SearchSectionResult) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *SearchSectionResult) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *SearchSectionResult) SetPath(v string) {
	o.Path = v
}

// GetBody returns the Body field value
func (o *SearchSectionResult) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *SearchSectionResult) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *SearchSectionResult) SetBody(v string) {
	o.Body = v
}

// GetUrls returns the Urls field value
func (o *SearchSectionResult) GetUrls() SearchSectionResultUrls {
	if o == nil {
		var ret SearchSectionResultUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *SearchSectionResult) GetUrlsOk() (*SearchSectionResultUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *SearchSectionResult) SetUrls(v SearchSectionResultUrls) {
	o.Urls = v
}

func (o SearchSectionResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSectionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["path"] = o.Path
	toSerialize["body"] = o.Body
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableSearchSectionResult struct {
	value *SearchSectionResult
	isSet bool
}

func (v NullableSearchSectionResult) Get() *SearchSectionResult {
	return v.value
}

func (v *NullableSearchSectionResult) Set(val *SearchSectionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSectionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSectionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSectionResult(val *SearchSectionResult) *NullableSearchSectionResult {
	return &NullableSearchSectionResult{value: val, isSet: true}
}

func (v NullableSearchSectionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSectionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
