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

// checks if the SearchAIQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAIQuery{}

// SearchAIQuery struct for SearchAIQuery
type SearchAIQuery struct {
	Query           string   `json:"query"`
	PreviousQueries []string `json:"previousQueries,omitempty"`
}

// NewSearchAIQuery instantiates a new SearchAIQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAIQuery(query string) *SearchAIQuery {
	this := SearchAIQuery{}
	this.Query = query
	return &this
}

// NewSearchAIQueryWithDefaults instantiates a new SearchAIQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAIQueryWithDefaults() *SearchAIQuery {
	this := SearchAIQuery{}
	return &this
}

// GetQuery returns the Query field value
func (o *SearchAIQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SearchAIQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SearchAIQuery) SetQuery(v string) {
	o.Query = v
}

// GetPreviousQueries returns the PreviousQueries field value if set, zero value otherwise.
func (o *SearchAIQuery) GetPreviousQueries() []string {
	if o == nil || IsNil(o.PreviousQueries) {
		var ret []string
		return ret
	}
	return o.PreviousQueries
}

// GetPreviousQueriesOk returns a tuple with the PreviousQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAIQuery) GetPreviousQueriesOk() ([]string, bool) {
	if o == nil || IsNil(o.PreviousQueries) {
		return nil, false
	}
	return o.PreviousQueries, true
}

// HasPreviousQueries returns a boolean if a field has been set.
func (o *SearchAIQuery) HasPreviousQueries() bool {
	if o != nil && !IsNil(o.PreviousQueries) {
		return true
	}

	return false
}

// SetPreviousQueries gets a reference to the given []string and assigns it to the PreviousQueries field.
func (o *SearchAIQuery) SetPreviousQueries(v []string) {
	o.PreviousQueries = v
}

func (o SearchAIQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAIQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	if !IsNil(o.PreviousQueries) {
		toSerialize["previousQueries"] = o.PreviousQueries
	}
	return toSerialize, nil
}

type NullableSearchAIQuery struct {
	value *SearchAIQuery
	isSet bool
}

func (v NullableSearchAIQuery) Get() *SearchAIQuery {
	return v.value
}

func (v *NullableSearchAIQuery) Set(val *SearchAIQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAIQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAIQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAIQuery(val *SearchAIQuery) *NullableSearchAIQuery {
	return &NullableSearchAIQuery{value: val, isSet: true}
}

func (v NullableSearchAIQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAIQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}