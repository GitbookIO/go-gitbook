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

// checks if the UpdateCollectionByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCollectionByIdRequest{}

// UpdateCollectionByIdRequest struct for UpdateCollectionByIdRequest
type UpdateCollectionByIdRequest struct {
	Visibility *ContentVisibility `json:"visibility,omitempty"`
}

// NewUpdateCollectionByIdRequest instantiates a new UpdateCollectionByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCollectionByIdRequest() *UpdateCollectionByIdRequest {
	this := UpdateCollectionByIdRequest{}
	return &this
}

// NewUpdateCollectionByIdRequestWithDefaults instantiates a new UpdateCollectionByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCollectionByIdRequestWithDefaults() *UpdateCollectionByIdRequest {
	this := UpdateCollectionByIdRequest{}
	return &this
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *UpdateCollectionByIdRequest) GetVisibility() ContentVisibility {
	if o == nil || IsNil(o.Visibility) {
		var ret ContentVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCollectionByIdRequest) GetVisibilityOk() (*ContentVisibility, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *UpdateCollectionByIdRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given ContentVisibility and assigns it to the Visibility field.
func (o *UpdateCollectionByIdRequest) SetVisibility(v ContentVisibility) {
	o.Visibility = &v
}

func (o UpdateCollectionByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCollectionByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableUpdateCollectionByIdRequest struct {
	value *UpdateCollectionByIdRequest
	isSet bool
}

func (v NullableUpdateCollectionByIdRequest) Get() *UpdateCollectionByIdRequest {
	return v.value
}

func (v *NullableUpdateCollectionByIdRequest) Set(val *UpdateCollectionByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCollectionByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCollectionByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCollectionByIdRequest(val *UpdateCollectionByIdRequest) *NullableUpdateCollectionByIdRequest {
	return &NullableUpdateCollectionByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateCollectionByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCollectionByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
