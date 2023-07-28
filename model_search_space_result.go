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

// checks if the SearchSpaceResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSpaceResult{}

// SearchSpaceResult Search result representing a space.
type SearchSpaceResult struct {
	Id    string             `json:"id"`
	Title string             `json:"title"`
	Pages []SearchPageResult `json:"pages"`
}

// NewSearchSpaceResult instantiates a new SearchSpaceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSpaceResult(id string, title string, pages []SearchPageResult) *SearchSpaceResult {
	this := SearchSpaceResult{}
	this.Id = id
	this.Title = title
	this.Pages = pages
	return &this
}

// NewSearchSpaceResultWithDefaults instantiates a new SearchSpaceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSpaceResultWithDefaults() *SearchSpaceResult {
	this := SearchSpaceResult{}
	return &this
}

// GetId returns the Id field value
func (o *SearchSpaceResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SearchSpaceResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SearchSpaceResult) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *SearchSpaceResult) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SearchSpaceResult) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SearchSpaceResult) SetTitle(v string) {
	o.Title = v
}

// GetPages returns the Pages field value
func (o *SearchSpaceResult) GetPages() []SearchPageResult {
	if o == nil {
		var ret []SearchPageResult
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *SearchSpaceResult) GetPagesOk() ([]SearchPageResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages, true
}

// SetPages sets field value
func (o *SearchSpaceResult) SetPages(v []SearchPageResult) {
	o.Pages = v
}

func (o SearchSpaceResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSpaceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["pages"] = o.Pages
	return toSerialize, nil
}

type NullableSearchSpaceResult struct {
	value *SearchSpaceResult
	isSet bool
}

func (v NullableSearchSpaceResult) Get() *SearchSpaceResult {
	return v.value
}

func (v *NullableSearchSpaceResult) Set(val *SearchSpaceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSpaceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSpaceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSpaceResult(val *SearchSpaceResult) *NullableSearchSpaceResult {
	return &NullableSearchSpaceResult{value: val, isSet: true}
}

func (v NullableSearchSpaceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSpaceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}