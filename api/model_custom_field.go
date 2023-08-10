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

// checks if the CustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomField{}

// CustomField Custom field
type CustomField struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Type        CustomFieldType `json:"type"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   string          `json:"updatedAt"`
	Urls        CustomFieldUrls `json:"urls"`
}

// NewCustomField instantiates a new CustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomField(id string, name string, title string, description string, type_ CustomFieldType, createdAt string, updatedAt string, urls CustomFieldUrls) *CustomField {
	this := CustomField{}
	this.Id = id
	this.Name = name
	this.Title = title
	this.Description = description
	this.Type = type_
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Urls = urls
	return &this
}

// NewCustomFieldWithDefaults instantiates a new CustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldWithDefaults() *CustomField {
	this := CustomField{}
	return &this
}

// GetId returns the Id field value
func (o *CustomField) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomField) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CustomField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomField) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *CustomField) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CustomField) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *CustomField) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CustomField) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *CustomField) GetType() CustomFieldType {
	if o == nil {
		var ret CustomFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetTypeOk() (*CustomFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomField) SetType(v CustomFieldType) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomField) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomField) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomField) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomField) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUrls returns the Urls field value
func (o *CustomField) GetUrls() CustomFieldUrls {
	if o == nil {
		var ret CustomFieldUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetUrlsOk() (*CustomFieldUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *CustomField) SetUrls(v CustomFieldUrls) {
	o.Urls = v
}

func (o CustomField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

type NullableCustomField struct {
	value *CustomField
	isSet bool
}

func (v NullableCustomField) Get() *CustomField {
	return v.value
}

func (v *NullableCustomField) Set(val *CustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomField(val *CustomField) *NullableCustomField {
	return &NullableCustomField{value: val, isSet: true}
}

func (v NullableCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
