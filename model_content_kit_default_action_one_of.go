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

// checks if the ContentKitDefaultActionOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitDefaultActionOneOf{}

// ContentKitDefaultActionOneOf Action to open an overlay modal defined by \"componentId\".
type ContentKitDefaultActionOneOf struct {
	Action      string                 `json:"action"`
	ComponentId string                 `json:"componentId"`
	Props       map[string]interface{} `json:"props"`
}

// NewContentKitDefaultActionOneOf instantiates a new ContentKitDefaultActionOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitDefaultActionOneOf(action string, componentId string, props map[string]interface{}) *ContentKitDefaultActionOneOf {
	this := ContentKitDefaultActionOneOf{}
	this.Action = action
	this.ComponentId = componentId
	this.Props = props
	return &this
}

// NewContentKitDefaultActionOneOfWithDefaults instantiates a new ContentKitDefaultActionOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitDefaultActionOneOfWithDefaults() *ContentKitDefaultActionOneOf {
	this := ContentKitDefaultActionOneOf{}
	return &this
}

// GetAction returns the Action field value
func (o *ContentKitDefaultActionOneOf) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContentKitDefaultActionOneOf) SetAction(v string) {
	o.Action = v
}

// GetComponentId returns the ComponentId field value
func (o *ContentKitDefaultActionOneOf) GetComponentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf) GetComponentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentId, true
}

// SetComponentId sets field value
func (o *ContentKitDefaultActionOneOf) SetComponentId(v string) {
	o.ComponentId = v
}

// GetProps returns the Props field value
func (o *ContentKitDefaultActionOneOf) GetProps() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Props
}

// GetPropsOk returns a tuple with the Props field value
// and a boolean to check if the value has been set.
func (o *ContentKitDefaultActionOneOf) GetPropsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Props, true
}

// SetProps sets field value
func (o *ContentKitDefaultActionOneOf) SetProps(v map[string]interface{}) {
	o.Props = v
}

func (o ContentKitDefaultActionOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitDefaultActionOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["componentId"] = o.ComponentId
	toSerialize["props"] = o.Props
	return toSerialize, nil
}

type NullableContentKitDefaultActionOneOf struct {
	value *ContentKitDefaultActionOneOf
	isSet bool
}

func (v NullableContentKitDefaultActionOneOf) Get() *ContentKitDefaultActionOneOf {
	return v.value
}

func (v *NullableContentKitDefaultActionOneOf) Set(val *ContentKitDefaultActionOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitDefaultActionOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitDefaultActionOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitDefaultActionOneOf(val *ContentKitDefaultActionOneOf) *NullableContentKitDefaultActionOneOf {
	return &NullableContentKitDefaultActionOneOf{value: val, isSet: true}
}

func (v NullableContentKitDefaultActionOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitDefaultActionOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
