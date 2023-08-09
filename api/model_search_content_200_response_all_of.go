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

// checks if the SearchContent200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchContent200ResponseAllOf{}

// SearchContent200ResponseAllOf struct for SearchContent200ResponseAllOf
type SearchContent200ResponseAllOf struct {
	Items []SearchSpaceResult `json:"items"`
}

// NewSearchContent200ResponseAllOf instantiates a new SearchContent200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchContent200ResponseAllOf(items []SearchSpaceResult) *SearchContent200ResponseAllOf {
	this := SearchContent200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewSearchContent200ResponseAllOfWithDefaults instantiates a new SearchContent200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchContent200ResponseAllOfWithDefaults() *SearchContent200ResponseAllOf {
	this := SearchContent200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *SearchContent200ResponseAllOf) GetItems() []SearchSpaceResult {
	if o == nil {
		var ret []SearchSpaceResult
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SearchContent200ResponseAllOf) GetItemsOk() ([]SearchSpaceResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SearchContent200ResponseAllOf) SetItems(v []SearchSpaceResult) {
	o.Items = v
}

func (o SearchContent200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchContent200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSearchContent200ResponseAllOf struct {
	value *SearchContent200ResponseAllOf
	isSet bool
}

func (v NullableSearchContent200ResponseAllOf) Get() *SearchContent200ResponseAllOf {
	return v.value
}

func (v *NullableSearchContent200ResponseAllOf) Set(val *SearchContent200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchContent200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchContent200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchContent200ResponseAllOf(val *SearchContent200ResponseAllOf) *NullableSearchContent200ResponseAllOf {
	return &NullableSearchContent200ResponseAllOf{value: val, isSet: true}
}

func (v NullableSearchContent200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchContent200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}