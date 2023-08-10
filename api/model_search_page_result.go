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

// checks if the SearchPageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPageResult{}

// SearchPageResult Search result representing a page in a space.
type SearchPageResult struct {
	Id       string                `json:"id"`
	Title    string                `json:"title"`
	Path     string                `json:"path"`
	Sections []SearchSectionResult `json:"sections,omitempty"`
	Urls     SearchPageResultUrls  `json:"urls"`
}

// NewSearchPageResult instantiates a new SearchPageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPageResult(id string, title string, path string, urls SearchPageResultUrls) *SearchPageResult {
	this := SearchPageResult{}
	this.Id = id
	this.Title = title
	this.Path = path
	this.Urls = urls
	return &this
}

// NewSearchPageResultWithDefaults instantiates a new SearchPageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPageResultWithDefaults() *SearchPageResult {
	this := SearchPageResult{}
	return &this
}

// GetId returns the Id field value
func (o *SearchPageResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SearchPageResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SearchPageResult) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *SearchPageResult) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SearchPageResult) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SearchPageResult) SetTitle(v string) {
	o.Title = v
}

// GetPath returns the Path field value
func (o *SearchPageResult) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *SearchPageResult) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *SearchPageResult) SetPath(v string) {
	o.Path = v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *SearchPageResult) GetSections() []SearchSectionResult {
	if o == nil || IsNil(o.Sections) {
		var ret []SearchSectionResult
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPageResult) GetSectionsOk() ([]SearchSectionResult, bool) {
	if o == nil || IsNil(o.Sections) {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *SearchPageResult) HasSections() bool {
	if o != nil && !IsNil(o.Sections) {
		return true
	}

	return false
}

// SetSections gets a reference to the given []SearchSectionResult and assigns it to the Sections field.
func (o *SearchPageResult) SetSections(v []SearchSectionResult) {
	o.Sections = v
}

// GetUrls returns the Urls field value
func (o *SearchPageResult) GetUrls() SearchPageResultUrls {
	if o == nil {
		var ret SearchPageResultUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *SearchPageResult) GetUrlsOk() (*SearchPageResultUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *SearchPageResult) SetUrls(v SearchPageResultUrls) {
	o.Urls = v
}

func (o SearchPageResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["path"] = o.Path
	if !IsNil(o.Sections) {
		toSerialize["sections"] = o.Sections
	}
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableSearchPageResult struct {
	value *SearchPageResult
	isSet bool
}

func (v NullableSearchPageResult) Get() *SearchPageResult {
	return v.value
}

func (v *NullableSearchPageResult) Set(val *SearchPageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPageResult(val *SearchPageResult) *NullableSearchPageResult {
	return &NullableSearchPageResult{value: val, isSet: true}
}

func (v NullableSearchPageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
