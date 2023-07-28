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

// checks if the ContentKitBox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitBox{}

// ContentKitBox struct for ContentKitBox
type ContentKitBox struct {
	Type string `json:"type"`
	// specifies how much of the remaining space in the container should be assigned to the element
	Grow     *float32                      `json:"grow,omitempty"`
	Children []ContentKitDescendantElement `json:"children"`
}

// NewContentKitBox instantiates a new ContentKitBox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitBox(type_ string, children []ContentKitDescendantElement) *ContentKitBox {
	this := ContentKitBox{}
	this.Type = type_
	this.Children = children
	return &this
}

// NewContentKitBoxWithDefaults instantiates a new ContentKitBox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitBoxWithDefaults() *ContentKitBox {
	this := ContentKitBox{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitBox) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitBox) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitBox) SetType(v string) {
	o.Type = v
}

// GetGrow returns the Grow field value if set, zero value otherwise.
func (o *ContentKitBox) GetGrow() float32 {
	if o == nil || IsNil(o.Grow) {
		var ret float32
		return ret
	}
	return *o.Grow
}

// GetGrowOk returns a tuple with the Grow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitBox) GetGrowOk() (*float32, bool) {
	if o == nil || IsNil(o.Grow) {
		return nil, false
	}
	return o.Grow, true
}

// HasGrow returns a boolean if a field has been set.
func (o *ContentKitBox) HasGrow() bool {
	if o != nil && !IsNil(o.Grow) {
		return true
	}

	return false
}

// SetGrow gets a reference to the given float32 and assigns it to the Grow field.
func (o *ContentKitBox) SetGrow(v float32) {
	o.Grow = &v
}

// GetChildren returns the Children field value
func (o *ContentKitBox) GetChildren() []ContentKitDescendantElement {
	if o == nil {
		var ret []ContentKitDescendantElement
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *ContentKitBox) GetChildrenOk() ([]ContentKitDescendantElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *ContentKitBox) SetChildren(v []ContentKitDescendantElement) {
	o.Children = v
}

func (o ContentKitBox) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitBox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Grow) {
		toSerialize["grow"] = o.Grow
	}
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

type NullableContentKitBox struct {
	value *ContentKitBox
	isSet bool
}

func (v NullableContentKitBox) Get() *ContentKitBox {
	return v.value
}

func (v *NullableContentKitBox) Set(val *ContentKitBox) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitBox) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitBox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitBox(val *ContentKitBox) *NullableContentKitBox {
	return &NullableContentKitBox{value: val, isSet: true}
}

func (v NullableContentKitBox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitBox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
