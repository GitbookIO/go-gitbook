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
	"fmt"
)

// IntegrationInstallationTarget - struct for IntegrationInstallationTarget
type IntegrationInstallationTarget struct {
	OrganizationTarget *OrganizationTarget
	UserTarget         *UserTarget
}

// OrganizationTargetAsIntegrationInstallationTarget is a convenience function that returns OrganizationTarget wrapped in IntegrationInstallationTarget
func OrganizationTargetAsIntegrationInstallationTarget(v *OrganizationTarget) IntegrationInstallationTarget {
	return IntegrationInstallationTarget{
		OrganizationTarget: v,
	}
}

// UserTargetAsIntegrationInstallationTarget is a convenience function that returns UserTarget wrapped in IntegrationInstallationTarget
func UserTargetAsIntegrationInstallationTarget(v *UserTarget) IntegrationInstallationTarget {
	return IntegrationInstallationTarget{
		UserTarget: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IntegrationInstallationTarget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OrganizationTarget
	err = newStrictDecoder(data).Decode(&dst.OrganizationTarget)
	if err == nil {
		jsonOrganizationTarget, _ := json.Marshal(dst.OrganizationTarget)
		if string(jsonOrganizationTarget) == "{}" { // empty struct
			dst.OrganizationTarget = nil
		} else {
			match++
		}
	} else {
		dst.OrganizationTarget = nil
	}

	// try to unmarshal data into UserTarget
	err = newStrictDecoder(data).Decode(&dst.UserTarget)
	if err == nil {
		jsonUserTarget, _ := json.Marshal(dst.UserTarget)
		if string(jsonUserTarget) == "{}" { // empty struct
			dst.UserTarget = nil
		} else {
			match++
		}
	} else {
		dst.UserTarget = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OrganizationTarget = nil
		dst.UserTarget = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IntegrationInstallationTarget)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IntegrationInstallationTarget)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IntegrationInstallationTarget) MarshalJSON() ([]byte, error) {
	if src.OrganizationTarget != nil {
		return json.Marshal(&src.OrganizationTarget)
	}

	if src.UserTarget != nil {
		return json.Marshal(&src.UserTarget)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IntegrationInstallationTarget) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OrganizationTarget != nil {
		return obj.OrganizationTarget
	}

	if obj.UserTarget != nil {
		return obj.UserTarget
	}

	// all schemas are nil
	return nil
}

type NullableIntegrationInstallationTarget struct {
	value *IntegrationInstallationTarget
	isSet bool
}

func (v NullableIntegrationInstallationTarget) Get() *IntegrationInstallationTarget {
	return v.value
}

func (v *NullableIntegrationInstallationTarget) Set(val *IntegrationInstallationTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationInstallationTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationInstallationTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationInstallationTarget(val *IntegrationInstallationTarget) *NullableIntegrationInstallationTarget {
	return &NullableIntegrationInstallationTarget{value: val, isSet: true}
}

func (v NullableIntegrationInstallationTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationInstallationTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
