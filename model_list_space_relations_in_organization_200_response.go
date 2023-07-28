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

// checks if the ListSpaceRelationsInOrganization200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSpaceRelationsInOrganization200Response{}

// ListSpaceRelationsInOrganization200Response struct for ListSpaceRelationsInOrganization200Response
type ListSpaceRelationsInOrganization200Response struct {
	Next *ListNext `json:"next,omitempty"`
	// Total count of objects in the list
	Count *float32            `json:"count,omitempty"`
	Items []SpaceRelationPair `json:"items"`
}

// NewListSpaceRelationsInOrganization200Response instantiates a new ListSpaceRelationsInOrganization200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSpaceRelationsInOrganization200Response(items []SpaceRelationPair) *ListSpaceRelationsInOrganization200Response {
	this := ListSpaceRelationsInOrganization200Response{}
	this.Items = items
	return &this
}

// NewListSpaceRelationsInOrganization200ResponseWithDefaults instantiates a new ListSpaceRelationsInOrganization200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSpaceRelationsInOrganization200ResponseWithDefaults() *ListSpaceRelationsInOrganization200Response {
	this := ListSpaceRelationsInOrganization200Response{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListSpaceRelationsInOrganization200Response) GetNext() ListNext {
	if o == nil || IsNil(o.Next) {
		var ret ListNext
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSpaceRelationsInOrganization200Response) GetNextOk() (*ListNext, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListSpaceRelationsInOrganization200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ListNext and assigns it to the Next field.
func (o *ListSpaceRelationsInOrganization200Response) SetNext(v ListNext) {
	o.Next = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListSpaceRelationsInOrganization200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSpaceRelationsInOrganization200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListSpaceRelationsInOrganization200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ListSpaceRelationsInOrganization200Response) SetCount(v float32) {
	o.Count = &v
}

// GetItems returns the Items field value
func (o *ListSpaceRelationsInOrganization200Response) GetItems() []SpaceRelationPair {
	if o == nil {
		var ret []SpaceRelationPair
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListSpaceRelationsInOrganization200Response) GetItemsOk() ([]SpaceRelationPair, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListSpaceRelationsInOrganization200Response) SetItems(v []SpaceRelationPair) {
	o.Items = v
}

func (o ListSpaceRelationsInOrganization200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSpaceRelationsInOrganization200Response) ToMap() (map[string]interface{}, error) {
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

type NullableListSpaceRelationsInOrganization200Response struct {
	value *ListSpaceRelationsInOrganization200Response
	isSet bool
}

func (v NullableListSpaceRelationsInOrganization200Response) Get() *ListSpaceRelationsInOrganization200Response {
	return v.value
}

func (v *NullableListSpaceRelationsInOrganization200Response) Set(val *ListSpaceRelationsInOrganization200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSpaceRelationsInOrganization200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSpaceRelationsInOrganization200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSpaceRelationsInOrganization200Response(val *ListSpaceRelationsInOrganization200Response) *NullableListSpaceRelationsInOrganization200Response {
	return &NullableListSpaceRelationsInOrganization200Response{value: val, isSet: true}
}

func (v NullableListSpaceRelationsInOrganization200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSpaceRelationsInOrganization200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
