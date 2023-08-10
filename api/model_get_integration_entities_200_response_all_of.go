/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the GetIntegrationEntities200ResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIntegrationEntities200ResponseAllOf{}

// GetIntegrationEntities200ResponseAllOf struct for GetIntegrationEntities200ResponseAllOf
type GetIntegrationEntities200ResponseAllOf struct {
	Items []IntegrationEntity `json:"items"`
}

// NewGetIntegrationEntities200ResponseAllOf instantiates a new GetIntegrationEntities200ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIntegrationEntities200ResponseAllOf(items []IntegrationEntity) *GetIntegrationEntities200ResponseAllOf {
	this := GetIntegrationEntities200ResponseAllOf{}
	this.Items = items
	return &this
}

// NewGetIntegrationEntities200ResponseAllOfWithDefaults instantiates a new GetIntegrationEntities200ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIntegrationEntities200ResponseAllOfWithDefaults() *GetIntegrationEntities200ResponseAllOf {
	this := GetIntegrationEntities200ResponseAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *GetIntegrationEntities200ResponseAllOf) GetItems() []IntegrationEntity {
	if o == nil {
		var ret []IntegrationEntity
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetIntegrationEntities200ResponseAllOf) GetItemsOk() ([]IntegrationEntity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GetIntegrationEntities200ResponseAllOf) SetItems(v []IntegrationEntity) {
	o.Items = v
}

func (o GetIntegrationEntities200ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIntegrationEntities200ResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableGetIntegrationEntities200ResponseAllOf struct {
	value *GetIntegrationEntities200ResponseAllOf
	isSet bool
}

func (v NullableGetIntegrationEntities200ResponseAllOf) Get() *GetIntegrationEntities200ResponseAllOf {
	return v.value
}

func (v *NullableGetIntegrationEntities200ResponseAllOf) Set(val *GetIntegrationEntities200ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIntegrationEntities200ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIntegrationEntities200ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIntegrationEntities200ResponseAllOf(val *GetIntegrationEntities200ResponseAllOf) *NullableGetIntegrationEntities200ResponseAllOf {
	return &NullableGetIntegrationEntities200ResponseAllOf{value: val, isSet: true}
}

func (v NullableGetIntegrationEntities200ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIntegrationEntities200ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
