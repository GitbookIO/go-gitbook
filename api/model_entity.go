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

// checks if the Entity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Entity{}

// Entity struct for Entity
type Entity struct {
	// Unique ID of the entity in the context of the integration's entity type
	EntityId string `json:"entityId"`
	// Map of values stored as properties on the entity
	Properties map[string]UpsertEntityPropertiesValue `json:"properties"`
	// Unique ID for the entity in GitBook
	Id string `json:"id"`
	// Type of an entity
	Type string          `json:"type"`
	Urls EntityAllOfUrls `json:"urls"`
}

// NewEntity instantiates a new Entity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntity(entityId string, properties map[string]UpsertEntityPropertiesValue, id string, type_ string, urls EntityAllOfUrls) *Entity {
	this := Entity{}
	this.EntityId = entityId
	this.Properties = properties
	this.Id = id
	this.Type = type_
	this.Urls = urls
	return &this
}

// NewEntityWithDefaults instantiates a new Entity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWithDefaults() *Entity {
	this := Entity{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *Entity) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *Entity) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *Entity) SetEntityId(v string) {
	o.EntityId = v
}

// GetProperties returns the Properties field value
func (o *Entity) GetProperties() map[string]UpsertEntityPropertiesValue {
	if o == nil {
		var ret map[string]UpsertEntityPropertiesValue
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *Entity) GetPropertiesOk() (*map[string]UpsertEntityPropertiesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *Entity) SetProperties(v map[string]UpsertEntityPropertiesValue) {
	o.Properties = v
}

// GetId returns the Id field value
func (o *Entity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Entity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Entity) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Entity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Entity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Entity) SetType(v string) {
	o.Type = v
}

// GetUrls returns the Urls field value
func (o *Entity) GetUrls() EntityAllOfUrls {
	if o == nil {
		var ret EntityAllOfUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Entity) GetUrlsOk() (*EntityAllOfUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Entity) SetUrls(v EntityAllOfUrls) {
	o.Urls = v
}

func (o Entity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Entity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityId"] = o.EntityId
	toSerialize["properties"] = o.Properties
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableEntity struct {
	value *Entity
	isSet bool
}

func (v NullableEntity) Get() *Entity {
	return v.value
}

func (v *NullableEntity) Set(val *Entity) {
	v.value = val
	v.isSet = true
}

func (v NullableEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntity(val *Entity) *NullableEntity {
	return &NullableEntity{value: val, isSet: true}
}

func (v NullableEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
