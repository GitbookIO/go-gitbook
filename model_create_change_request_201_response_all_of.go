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

// checks if the CreateChangeRequest201ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChangeRequest201ResponseAllOf{}

// CreateChangeRequest201ResponseAllOf struct for CreateChangeRequest201ResponseAllOf
type CreateChangeRequest201ResponseAllOf struct {
	// Deprecated
	ChangeRequest string `json:"changeRequest"`
}

// NewCreateChangeRequest201ResponseAllOf instantiates a new CreateChangeRequest201ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChangeRequest201ResponseAllOf(changeRequest string) *CreateChangeRequest201ResponseAllOf {
	this := CreateChangeRequest201ResponseAllOf{}
	this.ChangeRequest = changeRequest
	return &this
}

// NewCreateChangeRequest201ResponseAllOfWithDefaults instantiates a new CreateChangeRequest201ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChangeRequest201ResponseAllOfWithDefaults() *CreateChangeRequest201ResponseAllOf {
	this := CreateChangeRequest201ResponseAllOf{}
	return &this
}

// GetChangeRequest returns the ChangeRequest field value
// Deprecated
func (o *CreateChangeRequest201ResponseAllOf) GetChangeRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangeRequest
}

// GetChangeRequestOk returns a tuple with the ChangeRequest field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateChangeRequest201ResponseAllOf) GetChangeRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequest, true
}

// SetChangeRequest sets field value
// Deprecated
func (o *CreateChangeRequest201ResponseAllOf) SetChangeRequest(v string) {
	o.ChangeRequest = v
}

func (o CreateChangeRequest201ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChangeRequest201ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["changeRequest"] = o.ChangeRequest
	return toSerialize, nil
}

type NullableCreateChangeRequest201ResponseAllOf struct {
	value *CreateChangeRequest201ResponseAllOf
	isSet bool
}

func (v NullableCreateChangeRequest201ResponseAllOf) Get() *CreateChangeRequest201ResponseAllOf {
	return v.value
}

func (v *NullableCreateChangeRequest201ResponseAllOf) Set(val *CreateChangeRequest201ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChangeRequest201ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChangeRequest201ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChangeRequest201ResponseAllOf(val *CreateChangeRequest201ResponseAllOf) *NullableCreateChangeRequest201ResponseAllOf {
	return &NullableCreateChangeRequest201ResponseAllOf{value: val, isSet: true}
}

func (v NullableCreateChangeRequest201ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChangeRequest201ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
