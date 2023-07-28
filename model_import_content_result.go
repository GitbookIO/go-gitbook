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

// checks if the ImportContentResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportContentResult{}

// ImportContentResult struct for ImportContentResult
type ImportContentResult struct {
	// ID of the newly created revision.
	Revision string `json:"revision"`
	// How many resources were imported
	ImportedResources float32 `json:"importedResources"`
	// How many resources were processed
	TotalResources float32 `json:"totalResources"`
}

// NewImportContentResult instantiates a new ImportContentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportContentResult(revision string, importedResources float32, totalResources float32) *ImportContentResult {
	this := ImportContentResult{}
	this.Revision = revision
	this.ImportedResources = importedResources
	this.TotalResources = totalResources
	return &this
}

// NewImportContentResultWithDefaults instantiates a new ImportContentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportContentResultWithDefaults() *ImportContentResult {
	this := ImportContentResult{}
	return &this
}

// GetRevision returns the Revision field value
func (o *ImportContentResult) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *ImportContentResult) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *ImportContentResult) SetRevision(v string) {
	o.Revision = v
}

// GetImportedResources returns the ImportedResources field value
func (o *ImportContentResult) GetImportedResources() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ImportedResources
}

// GetImportedResourcesOk returns a tuple with the ImportedResources field value
// and a boolean to check if the value has been set.
func (o *ImportContentResult) GetImportedResourcesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportedResources, true
}

// SetImportedResources sets field value
func (o *ImportContentResult) SetImportedResources(v float32) {
	o.ImportedResources = v
}

// GetTotalResources returns the TotalResources field value
func (o *ImportContentResult) GetTotalResources() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalResources
}

// GetTotalResourcesOk returns a tuple with the TotalResources field value
// and a boolean to check if the value has been set.
func (o *ImportContentResult) GetTotalResourcesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalResources, true
}

// SetTotalResources sets field value
func (o *ImportContentResult) SetTotalResources(v float32) {
	o.TotalResources = v
}

func (o ImportContentResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportContentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revision"] = o.Revision
	toSerialize["importedResources"] = o.ImportedResources
	toSerialize["totalResources"] = o.TotalResources
	return toSerialize, nil
}

type NullableImportContentResult struct {
	value *ImportContentResult
	isSet bool
}

func (v NullableImportContentResult) Get() *ImportContentResult {
	return v.value
}

func (v *NullableImportContentResult) Set(val *ImportContentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableImportContentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableImportContentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportContentResult(val *ImportContentResult) *NullableImportContentResult {
	return &NullableImportContentResult{value: val, isSet: true}
}

func (v NullableImportContentResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportContentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
