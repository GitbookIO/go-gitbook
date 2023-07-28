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

// checks if the ListOrganizationCustomFields200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationCustomFields200ResponseAllOf{}

// ListOrganizationCustomFields200ResponseAllOf struct for ListOrganizationCustomFields200ResponseAllOf
type ListOrganizationCustomFields200ResponseAllOf struct {
	Items []CustomField `json:"items"`
}

// NewListOrganizationCustomFields200ResponseAllOf instantiates a new ListOrganizationCustomFields200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationCustomFields200ResponseAllOf(items []CustomField) *ListOrganizationCustomFields200ResponseAllOf {
	this := ListOrganizationCustomFields200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewListOrganizationCustomFields200ResponseAllOfWithDefaults instantiates a new ListOrganizationCustomFields200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationCustomFields200ResponseAllOfWithDefaults() *ListOrganizationCustomFields200ResponseAllOf {
	this := ListOrganizationCustomFields200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *ListOrganizationCustomFields200ResponseAllOf) GetItems() []CustomField {
	if o == nil {
		var ret []CustomField
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationCustomFields200ResponseAllOf) GetItemsOk() ([]CustomField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListOrganizationCustomFields200ResponseAllOf) SetItems(v []CustomField) {
	o.Items = v
}

func (o ListOrganizationCustomFields200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationCustomFields200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableListOrganizationCustomFields200ResponseAllOf struct {
	value *ListOrganizationCustomFields200ResponseAllOf
	isSet bool
}

func (v NullableListOrganizationCustomFields200ResponseAllOf) Get() *ListOrganizationCustomFields200ResponseAllOf {
	return v.value
}

func (v *NullableListOrganizationCustomFields200ResponseAllOf) Set(val *ListOrganizationCustomFields200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationCustomFields200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationCustomFields200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationCustomFields200ResponseAllOf(val *ListOrganizationCustomFields200ResponseAllOf) *NullableListOrganizationCustomFields200ResponseAllOf {
	return &NullableListOrganizationCustomFields200ResponseAllOf{value: val, isSet: true}
}

func (v NullableListOrganizationCustomFields200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationCustomFields200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}