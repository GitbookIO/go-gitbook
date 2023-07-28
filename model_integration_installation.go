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

// checks if the IntegrationInstallation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationInstallation{}

// IntegrationInstallation Installation of an integration on an account
type IntegrationInstallation struct {
	Id             string                                `json:"id"`
	Status         IntegrationInstallationStatus         `json:"status"`
	SpaceSelection IntegrationInstallationSpaceSelection `json:"space_selection"`
	// Configuration of the integration at the account level
	Configuration map[string]interface{}      `json:"configuration"`
	Urls          IntegrationInstallationUrls `json:"urls"`
	// External IDs assigned by the integration.
	ExternalIds []string                      `json:"externalIds"`
	Target      IntegrationInstallationTarget `json:"target"`
}

// NewIntegrationInstallation instantiates a new IntegrationInstallation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationInstallation(id string, status IntegrationInstallationStatus, spaceSelection IntegrationInstallationSpaceSelection, configuration map[string]interface{}, urls IntegrationInstallationUrls, externalIds []string, target IntegrationInstallationTarget) *IntegrationInstallation {
	this := IntegrationInstallation{}
	this.Id = id
	this.Status = status
	this.SpaceSelection = spaceSelection
	this.Configuration = configuration
	this.Urls = urls
	this.ExternalIds = externalIds
	this.Target = target
	return &this
}

// NewIntegrationInstallationWithDefaults instantiates a new IntegrationInstallation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationInstallationWithDefaults() *IntegrationInstallation {
	this := IntegrationInstallation{}
	return &this
}

// GetId returns the Id field value
func (o *IntegrationInstallation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntegrationInstallation) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *IntegrationInstallation) GetStatus() IntegrationInstallationStatus {
	if o == nil {
		var ret IntegrationInstallationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallation) GetStatusOk() (*IntegrationInstallationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IntegrationInstallation) SetStatus(v IntegrationInstallationStatus) {
	o.Status = v
}

// GetSpaceSelection returns the SpaceSelection field value
func (o *IntegrationInstallation) GetSpaceSelection() IntegrationInstallationSpaceSelection {
	if o == nil {
		var ret IntegrationInstallationSpaceSelection
		return ret
	}

	return o.SpaceSelection
}

// GetSpaceSelectionOk returns a tuple with the SpaceSelection field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallation) GetSpaceSelectionOk() (*IntegrationInstallationSpaceSelection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceSelection, true
}

// SetSpaceSelection sets field value
func (o *IntegrationInstallation) SetSpaceSelection(v IntegrationInstallationSpaceSelection) {
	o.SpaceSelection = v
}

// GetConfiguration returns the Configuration field value
func (o *IntegrationInstallation) GetConfiguration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallation) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// SetConfiguration sets field value
func (o *IntegrationInstallation) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetUrls returns the Urls field value
func (o *IntegrationInstallation) GetUrls() IntegrationInstallationUrls {
	if o == nil {
		var ret IntegrationInstallationUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallation) GetUrlsOk() (*IntegrationInstallationUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *IntegrationInstallation) SetUrls(v IntegrationInstallationUrls) {
	o.Urls = v
}

// GetExternalIds returns the ExternalIds field value
func (o *IntegrationInstallation) GetExternalIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallation) GetExternalIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalIds, true
}

// SetExternalIds sets field value
func (o *IntegrationInstallation) SetExternalIds(v []string) {
	o.ExternalIds = v
}

// GetTarget returns the Target field value
func (o *IntegrationInstallation) GetTarget() IntegrationInstallationTarget {
	if o == nil {
		var ret IntegrationInstallationTarget
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *IntegrationInstallation) GetTargetOk() (*IntegrationInstallationTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *IntegrationInstallation) SetTarget(v IntegrationInstallationTarget) {
	o.Target = v
}

func (o IntegrationInstallation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationInstallation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["space_selection"] = o.SpaceSelection
	toSerialize["configuration"] = o.Configuration
	toSerialize["urls"] = o.Urls
	toSerialize["externalIds"] = o.ExternalIds
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

type NullableIntegrationInstallation struct {
	value *IntegrationInstallation
	isSet bool
}

func (v NullableIntegrationInstallation) Get() *IntegrationInstallation {
	return v.value
}

func (v *NullableIntegrationInstallation) Set(val *IntegrationInstallation) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationInstallation) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationInstallation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationInstallation(val *IntegrationInstallation) *NullableIntegrationInstallation {
	return &NullableIntegrationInstallation{value: val, isSet: true}
}

func (v NullableIntegrationInstallation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationInstallation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}