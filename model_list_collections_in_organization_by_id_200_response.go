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

// checks if the ListCollectionsInOrganizationById200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCollectionsInOrganizationById200Response{}

// ListCollectionsInOrganizationById200Response struct for ListCollectionsInOrganizationById200Response
type ListCollectionsInOrganizationById200Response struct {
	Next *ListNext `json:"next,omitempty"`
	// Total count of objects in the list
	Count *float32     `json:"count,omitempty"`
	Items []Collection `json:"items"`
}

// NewListCollectionsInOrganizationById200Response instantiates a new ListCollectionsInOrganizationById200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCollectionsInOrganizationById200Response(items []Collection) *ListCollectionsInOrganizationById200Response {
	this := ListCollectionsInOrganizationById200Response{}
	this.Items = items
	return &this
}

// NewListCollectionsInOrganizationById200ResponseWithDefaults instantiates a new ListCollectionsInOrganizationById200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCollectionsInOrganizationById200ResponseWithDefaults() *ListCollectionsInOrganizationById200Response {
	this := ListCollectionsInOrganizationById200Response{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListCollectionsInOrganizationById200Response) GetNext() ListNext {
	if o == nil || IsNil(o.Next) {
		var ret ListNext
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCollectionsInOrganizationById200Response) GetNextOk() (*ListNext, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListCollectionsInOrganizationById200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListNext and assigns it to the Next field.
func (o *ListCollectionsInOrganizationById200Response) SetNext(v ListNext) {
	o.Next = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListCollectionsInOrganizationById200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCollectionsInOrganizationById200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListCollectionsInOrganizationById200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ListCollectionsInOrganizationById200Response) SetCount(v float32) {
	o.Count = &v
}

// GetItems returns the Items field value
func (o *ListCollectionsInOrganizationById200Response) GetItems() []Collection {
	if o == nil {
		var ret []Collection
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListCollectionsInOrganizationById200Response) GetItemsOk() ([]Collection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListCollectionsInOrganizationById200Response) SetItems(v []Collection) {
	o.Items = v
}

func (o ListCollectionsInOrganizationById200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCollectionsInOrganizationById200Response) ToMap() (map[string]interface{}, error) {
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

type NullableListCollectionsInOrganizationById200Response struct {
	value *ListCollectionsInOrganizationById200Response
	isSet bool
}

func (v NullableListCollectionsInOrganizationById200Response) Get() *ListCollectionsInOrganizationById200Response {
	return v.value
}

func (v *NullableListCollectionsInOrganizationById200Response) Set(val *ListCollectionsInOrganizationById200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListCollectionsInOrganizationById200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListCollectionsInOrganizationById200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCollectionsInOrganizationById200Response(val *ListCollectionsInOrganizationById200Response) *NullableListCollectionsInOrganizationById200Response {
	return &NullableListCollectionsInOrganizationById200Response{value: val, isSet: true}
}

func (v NullableListCollectionsInOrganizationById200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCollectionsInOrganizationById200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}