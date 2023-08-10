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
	"fmt"
)

// IntegrationInstallationStatus the model 'IntegrationInstallationStatus'
type IntegrationInstallationStatus string

// List of IntegrationInstallationStatus
const (
	INTEGRATIONINSTALLATIONSTATUS_ACTIVE  IntegrationInstallationStatus = "active"
	INTEGRATIONINSTALLATIONSTATUS_PENDING IntegrationInstallationStatus = "pending"
	INTEGRATIONINSTALLATIONSTATUS_PAUSED  IntegrationInstallationStatus = "paused"
)

// All allowed values of IntegrationInstallationStatus enum
var AllowedIntegrationInstallationStatusEnumValues = []IntegrationInstallationStatus{
	"active",
	"pending",
	"paused",
}

func (v *IntegrationInstallationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationInstallationStatus(value)
	for _, existing := range AllowedIntegrationInstallationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationInstallationStatus", value)
}

// NewIntegrationInstallationStatusFromValue returns a pointer to a valid IntegrationInstallationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationInstallationStatusFromValue(v string) (*IntegrationInstallationStatus, error) {
	ev := IntegrationInstallationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationInstallationStatus: valid values are %v", v, AllowedIntegrationInstallationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationInstallationStatus) IsValid() bool {
	for _, existing := range AllowedIntegrationInstallationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationInstallationStatus value
func (v IntegrationInstallationStatus) Ptr() *IntegrationInstallationStatus {
	return &v
}

type NullableIntegrationInstallationStatus struct {
	value *IntegrationInstallationStatus
	isSet bool
}

func (v NullableIntegrationInstallationStatus) Get() *IntegrationInstallationStatus {
	return v.value
}

func (v *NullableIntegrationInstallationStatus) Set(val *IntegrationInstallationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationInstallationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationInstallationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationInstallationStatus(val *IntegrationInstallationStatus) *NullableIntegrationInstallationStatus {
	return &NullableIntegrationInstallationStatus{value: val, isSet: true}
}

func (v NullableIntegrationInstallationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationInstallationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
