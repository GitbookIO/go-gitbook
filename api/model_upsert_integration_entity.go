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

// checks if the UpsertIntegrationEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertIntegrationEntity{}

// UpsertIntegrationEntity Entity to create or update in an integration's installation.
type UpsertIntegrationEntity struct {
	// Unique ID of the entity in the context of the integration's installation
	EntityId string `json:"entityId"`
	// Title of the entity
	Title string `json:"title"`
	// Longer text description of the entity
	Description string `json:"description"`
	// URL to open the entity
	Target string `json:"target"`
	// Metadata stored by the integration on the entity
	Metadata map[string]UpsertIntegrationEntityMetadataValue `json:"metadata"`
}

// NewUpsertIntegrationEntity instantiates a new UpsertIntegrationEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertIntegrationEntity(entityId string, title string, description string, target string, metadata map[string]UpsertIntegrationEntityMetadataValue) *UpsertIntegrationEntity {
	this := UpsertIntegrationEntity{}
	this.EntityId = entityId
	this.Title = title
	this.Description = description
	this.Target = target
	this.Metadata = metadata
	return &this
}

// NewUpsertIntegrationEntityWithDefaults instantiates a new UpsertIntegrationEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertIntegrationEntityWithDefaults() *UpsertIntegrationEntity {
	this := UpsertIntegrationEntity{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *UpsertIntegrationEntity) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *UpsertIntegrationEntity) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *UpsertIntegrationEntity) SetEntityId(v string) {
	o.EntityId = v
}

// GetTitle returns the Title field value
func (o *UpsertIntegrationEntity) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UpsertIntegrationEntity) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UpsertIntegrationEntity) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *UpsertIntegrationEntity) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpsertIntegrationEntity) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpsertIntegrationEntity) SetDescription(v string) {
	o.Description = v
}

// GetTarget returns the Target field value
func (o *UpsertIntegrationEntity) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *UpsertIntegrationEntity) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *UpsertIntegrationEntity) SetTarget(v string) {
	o.Target = v
}

// GetMetadata returns the Metadata field value
func (o *UpsertIntegrationEntity) GetMetadata() map[string]UpsertIntegrationEntityMetadataValue {
	if o == nil {
		var ret map[string]UpsertIntegrationEntityMetadataValue
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *UpsertIntegrationEntity) GetMetadataOk() (*map[string]UpsertIntegrationEntityMetadataValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *UpsertIntegrationEntity) SetMetadata(v map[string]UpsertIntegrationEntityMetadataValue) {
	o.Metadata = v
}

func (o UpsertIntegrationEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertIntegrationEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityId"] = o.EntityId
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["target"] = o.Target
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

type NullableUpsertIntegrationEntity struct {
	value *UpsertIntegrationEntity
	isSet bool
}

func (v NullableUpsertIntegrationEntity) Get() *UpsertIntegrationEntity {
	return v.value
}

func (v *NullableUpsertIntegrationEntity) Set(val *UpsertIntegrationEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertIntegrationEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertIntegrationEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertIntegrationEntity(val *UpsertIntegrationEntity) *NullableUpsertIntegrationEntity {
	return &NullableUpsertIntegrationEntity{value: val, isSet: true}
}

func (v NullableUpsertIntegrationEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertIntegrationEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}