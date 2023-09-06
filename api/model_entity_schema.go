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

// checks if the EntitySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitySchema{}

// EntitySchema struct for EntitySchema
type EntitySchema struct {
	// Type of an entity
	Type  string               `json:"type"`
	Title EntityRawSchemaTitle `json:"title"`
	// Ordered list of all properties stored in entities.
	Properties []EntityPropertySchema `json:"properties"`
	// Count of entities created in this schema.
	Entities    float32               `json:"entities"`
	Integration *Integration          `json:"integration,omitempty"`
	Urls        EntitySchemaAllOfUrls `json:"urls"`
}

// NewEntitySchema instantiates a new EntitySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySchema(type_ string, title EntityRawSchemaTitle, properties []EntityPropertySchema, entities float32, urls EntitySchemaAllOfUrls) *EntitySchema {
	this := EntitySchema{}
	this.Type = type_
	this.Title = title
	this.Properties = properties
	this.Entities = entities
	this.Urls = urls
	return &this
}

// NewEntitySchemaWithDefaults instantiates a new EntitySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySchemaWithDefaults() *EntitySchema {
	this := EntitySchema{}
	return &this
}

// GetType returns the Type field value
func (o *EntitySchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntitySchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntitySchema) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *EntitySchema) GetTitle() EntityRawSchemaTitle {
	if o == nil {
		var ret EntityRawSchemaTitle
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EntitySchema) GetTitleOk() (*EntityRawSchemaTitle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *EntitySchema) SetTitle(v EntityRawSchemaTitle) {
	o.Title = v
}

// GetProperties returns the Properties field value
func (o *EntitySchema) GetProperties() []EntityPropertySchema {
	if o == nil {
		var ret []EntityPropertySchema
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *EntitySchema) GetPropertiesOk() ([]EntityPropertySchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *EntitySchema) SetProperties(v []EntityPropertySchema) {
	o.Properties = v
}

// GetEntities returns the Entities field value
func (o *EntitySchema) GetEntities() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *EntitySchema) GetEntitiesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entities, true
}

// SetEntities sets field value
func (o *EntitySchema) SetEntities(v float32) {
	o.Entities = v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *EntitySchema) GetIntegration() Integration {
	if o == nil || IsNil(o.Integration) {
		var ret Integration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySchema) GetIntegrationOk() (*Integration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *EntitySchema) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given Integration and assigns it to the Integration field.
func (o *EntitySchema) SetIntegration(v Integration) {
	o.Integration = &v
}

// GetUrls returns the Urls field value
func (o *EntitySchema) GetUrls() EntitySchemaAllOfUrls {
	if o == nil {
		var ret EntitySchemaAllOfUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *EntitySchema) GetUrlsOk() (*EntitySchemaAllOfUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *EntitySchema) SetUrls(v EntitySchemaAllOfUrls) {
	o.Urls = v
}

func (o EntitySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["properties"] = o.Properties
	toSerialize["entities"] = o.Entities
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableEntitySchema struct {
	value *EntitySchema
	isSet bool
}

func (v NullableEntitySchema) Get() *EntitySchema {
	return v.value
}

func (v *NullableEntitySchema) Set(val *EntitySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySchema(val *EntitySchema) *NullableEntitySchema {
	return &NullableEntitySchema{value: val, isSet: true}
}

func (v NullableEntitySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}