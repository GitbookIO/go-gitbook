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

// checks if the EntitySchemaAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitySchemaAllOf{}

// EntitySchemaAllOf struct for EntitySchemaAllOf
type EntitySchemaAllOf struct {
	// Count of entities created in this schema.
	Entities    float32               `json:"entities"`
	Integration *Integration          `json:"integration,omitempty"`
	Urls        EntitySchemaAllOfUrls `json:"urls"`
}

// NewEntitySchemaAllOf instantiates a new EntitySchemaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySchemaAllOf(entities float32, urls EntitySchemaAllOfUrls) *EntitySchemaAllOf {
	this := EntitySchemaAllOf{}
	this.Entities = entities
	this.Urls = urls
	return &this
}

// NewEntitySchemaAllOfWithDefaults instantiates a new EntitySchemaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySchemaAllOfWithDefaults() *EntitySchemaAllOf {
	this := EntitySchemaAllOf{}
	return &this
}

// GetEntities returns the Entities field value
func (o *EntitySchemaAllOf) GetEntities() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOf) GetEntitiesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *EntitySchemaAllOf) SetEntities(v float32) {
	o.Entities = v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *EntitySchemaAllOf) GetIntegration() Integration {
	if o == nil || IsNil(o.Integration) {
		var ret Integration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOf) GetIntegrationOk() (*Integration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *EntitySchemaAllOf) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given Integration and assigns it to the Integration field.
func (o *EntitySchemaAllOf) SetIntegration(v Integration) {
	o.Integration = &v
}

// GetUrls returns the Urls field value
func (o *EntitySchemaAllOf) GetUrls() EntitySchemaAllOfUrls {
	if o == nil {
		var ret EntitySchemaAllOfUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *EntitySchemaAllOf) GetUrlsOk() (*EntitySchemaAllOfUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *EntitySchemaAllOf) SetUrls(v EntitySchemaAllOfUrls) {
	o.Urls = v
}

func (o EntitySchemaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitySchemaAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entities"] = o.Entities
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableEntitySchemaAllOf struct {
	value *EntitySchemaAllOf
	isSet bool
}

func (v NullableEntitySchemaAllOf) Get() *EntitySchemaAllOf {
	return v.value
}

func (v *NullableEntitySchemaAllOf) Set(val *EntitySchemaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySchemaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySchemaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySchemaAllOf(val *EntitySchemaAllOf) *NullableEntitySchemaAllOf {
	return &NullableEntitySchemaAllOf{value: val, isSet: true}
}

func (v NullableEntitySchemaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySchemaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}