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

// checks if the SyncIntegrationEntitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncIntegrationEntitiesRequest{}

// SyncIntegrationEntitiesRequest struct for SyncIntegrationEntitiesRequest
type SyncIntegrationEntitiesRequest struct {
	// Array of entities to synchronize. Entities will be created and updated, missing entities will be deleted.
	Entities []UpsertIntegrationEntity `json:"entities"`
}

// NewSyncIntegrationEntitiesRequest instantiates a new SyncIntegrationEntitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncIntegrationEntitiesRequest(entities []UpsertIntegrationEntity) *SyncIntegrationEntitiesRequest {
	this := SyncIntegrationEntitiesRequest{}
	this.Entities = entities
	return &this
}

// NewSyncIntegrationEntitiesRequestWithDefaults instantiates a new SyncIntegrationEntitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncIntegrationEntitiesRequestWithDefaults() *SyncIntegrationEntitiesRequest {
	this := SyncIntegrationEntitiesRequest{}
	return &this
}

// GetEntities returns the Entities field value
func (o *SyncIntegrationEntitiesRequest) GetEntities() []UpsertIntegrationEntity {
	if o == nil {
		var ret []UpsertIntegrationEntity
		return ret
	}

	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
func (o *SyncIntegrationEntitiesRequest) GetEntitiesOk() ([]UpsertIntegrationEntity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entities, true
}

// SetEntities sets field value
func (o *SyncIntegrationEntitiesRequest) SetEntities(v []UpsertIntegrationEntity) {
	o.Entities = v
}

func (o SyncIntegrationEntitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncIntegrationEntitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entities"] = o.Entities
	return toSerialize, nil
}

type NullableSyncIntegrationEntitiesRequest struct {
	value *SyncIntegrationEntitiesRequest
	isSet bool
}

func (v NullableSyncIntegrationEntitiesRequest) Get() *SyncIntegrationEntitiesRequest {
	return v.value
}

func (v *NullableSyncIntegrationEntitiesRequest) Set(val *SyncIntegrationEntitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncIntegrationEntitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncIntegrationEntitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncIntegrationEntitiesRequest(val *SyncIntegrationEntitiesRequest) *NullableSyncIntegrationEntitiesRequest {
	return &NullableSyncIntegrationEntitiesRequest{value: val, isSet: true}
}

func (v NullableSyncIntegrationEntitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncIntegrationEntitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
