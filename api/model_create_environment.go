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

// checks if the CreateEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEnvironment{}

// CreateEnvironment struct for CreateEnvironment
type CreateEnvironment struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	// Whether or not an environment is considered to have elevated responsibilities over other environments. Useful for distinguishing a production environment from a staging environment. Multiple primary environments are allowed. Your organization must have at least one primary environment.
	Primary *bool `json:"primary,omitempty"`
}

// NewCreateEnvironment instantiates a new CreateEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEnvironment(name string, title string) *CreateEnvironment {
	this := CreateEnvironment{}
	this.Name = name
	this.Title = title
	return &this
}

// NewCreateEnvironmentWithDefaults instantiates a new CreateEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEnvironmentWithDefaults() *CreateEnvironment {
	this := CreateEnvironment{}
	return &this
}

// GetName returns the Name field value
func (o *CreateEnvironment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateEnvironment) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *CreateEnvironment) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironment) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateEnvironment) SetTitle(v string) {
	o.Title = v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *CreateEnvironment) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironment) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *CreateEnvironment) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *CreateEnvironment) SetPrimary(v bool) {
	o.Primary = &v
}

func (o CreateEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	return toSerialize, nil
}

type NullableCreateEnvironment struct {
	value *CreateEnvironment
	isSet bool
}

func (v NullableCreateEnvironment) Get() *CreateEnvironment {
	return v.value
}

func (v *NullableCreateEnvironment) Set(val *CreateEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEnvironment(val *CreateEnvironment) *NullableCreateEnvironment {
	return &NullableCreateEnvironment{value: val, isSet: true}
}

func (v NullableCreateEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}