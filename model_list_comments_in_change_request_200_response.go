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

// checks if the ListCommentsInChangeRequest200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCommentsInChangeRequest200Response{}

// ListCommentsInChangeRequest200Response struct for ListCommentsInChangeRequest200Response
type ListCommentsInChangeRequest200Response struct {
	Next *ListNext `json:"next,omitempty"`
	// Total count of objects in the list
	Count *float32  `json:"count,omitempty"`
	Items []Comment `json:"items"`
}

// NewListCommentsInChangeRequest200Response instantiates a new ListCommentsInChangeRequest200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCommentsInChangeRequest200Response(items []Comment) *ListCommentsInChangeRequest200Response {
	this := ListCommentsInChangeRequest200Response{}
	this.Items = items
	return &this
}

// NewListCommentsInChangeRequest200ResponseWithDefaults instantiates a new ListCommentsInChangeRequest200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCommentsInChangeRequest200ResponseWithDefaults() *ListCommentsInChangeRequest200Response {
	this := ListCommentsInChangeRequest200Response{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListCommentsInChangeRequest200Response) GetNext() ListNext {
	if o == nil || IsNil(o.Next) {
		var ret ListNext
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommentsInChangeRequest200Response) GetNextOk() (*ListNext, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListCommentsInChangeRequest200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListNext and assigns it to the Next field.
func (o *ListCommentsInChangeRequest200Response) SetNext(v ListNext) {
	o.Next = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListCommentsInChangeRequest200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCommentsInChangeRequest200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListCommentsInChangeRequest200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ListCommentsInChangeRequest200Response) SetCount(v float32) {
	o.Count = &v
}

// GetItems returns the Items field value
func (o *ListCommentsInChangeRequest200Response) GetItems() []Comment {
	if o == nil {
		var ret []Comment
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListCommentsInChangeRequest200Response) GetItemsOk() ([]Comment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListCommentsInChangeRequest200Response) SetItems(v []Comment) {
	o.Items = v
}

func (o ListCommentsInChangeRequest200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCommentsInChangeRequest200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListCommentsInChangeRequest200Response struct {
	value *ListCommentsInChangeRequest200Response
	isSet bool
}

func (v NullableListCommentsInChangeRequest200Response) Get() *ListCommentsInChangeRequest200Response {
	return v.value
}

func (v *NullableListCommentsInChangeRequest200Response) Set(val *ListCommentsInChangeRequest200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListCommentsInChangeRequest200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListCommentsInChangeRequest200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCommentsInChangeRequest200Response(val *ListCommentsInChangeRequest200Response) *NullableListCommentsInChangeRequest200Response {
	return &NullableListCommentsInChangeRequest200Response{value: val, isSet: true}
}

func (v NullableListCommentsInChangeRequest200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCommentsInChangeRequest200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
