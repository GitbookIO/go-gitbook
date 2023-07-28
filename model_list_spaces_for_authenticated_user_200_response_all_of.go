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

// checks if the ListSpacesForAuthenticatedUser200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSpacesForAuthenticatedUser200ResponseAllOf{}

// ListSpacesForAuthenticatedUser200ResponseAllOf struct for ListSpacesForAuthenticatedUser200ResponseAllOf
type ListSpacesForAuthenticatedUser200ResponseAllOf struct {
	Items []Space `json:"items"`
}

// NewListSpacesForAuthenticatedUser200ResponseAllOf instantiates a new ListSpacesForAuthenticatedUser200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSpacesForAuthenticatedUser200ResponseAllOf(items []Space) *ListSpacesForAuthenticatedUser200ResponseAllOf {
	this := ListSpacesForAuthenticatedUser200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewListSpacesForAuthenticatedUser200ResponseAllOfWithDefaults instantiates a new ListSpacesForAuthenticatedUser200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSpacesForAuthenticatedUser200ResponseAllOfWithDefaults() *ListSpacesForAuthenticatedUser200ResponseAllOf {
	this := ListSpacesForAuthenticatedUser200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *ListSpacesForAuthenticatedUser200ResponseAllOf) GetItems() []Space {
	if o == nil {
		var ret []Space
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListSpacesForAuthenticatedUser200ResponseAllOf) GetItemsOk() ([]Space, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListSpacesForAuthenticatedUser200ResponseAllOf) SetItems(v []Space) {
	o.Items = v
}

func (o ListSpacesForAuthenticatedUser200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSpacesForAuthenticatedUser200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListSpacesForAuthenticatedUser200ResponseAllOf struct {
	value *ListSpacesForAuthenticatedUser200ResponseAllOf
	isSet bool
}

func (v NullableListSpacesForAuthenticatedUser200ResponseAllOf) Get() *ListSpacesForAuthenticatedUser200ResponseAllOf {
	return v.value
}

func (v *NullableListSpacesForAuthenticatedUser200ResponseAllOf) Set(val *ListSpacesForAuthenticatedUser200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListSpacesForAuthenticatedUser200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListSpacesForAuthenticatedUser200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSpacesForAuthenticatedUser200ResponseAllOf(val *ListSpacesForAuthenticatedUser200ResponseAllOf) *NullableListSpacesForAuthenticatedUser200ResponseAllOf {
	return &NullableListSpacesForAuthenticatedUser200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListSpacesForAuthenticatedUser200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSpacesForAuthenticatedUser200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}