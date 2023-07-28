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

// OrganizationUseCase the model 'OrganizationUseCase'
type OrganizationUseCase string

// List of OrganizationUseCase
const (
	ORGANIZATIONUSECASE_PRODUCT_DOCS        OrganizationUseCase = "productDocs"
	ORGANIZATIONUSECASE_TEAM_KNOWLEDGE_BASE OrganizationUseCase = "teamKnowledgeBase"
	ORGANIZATIONUSECASE_DESIGN_SYSTEM       OrganizationUseCase = "designSystem"
	ORGANIZATIONUSECASE_OPEN_SOURCE_DOCS    OrganizationUseCase = "openSourceDocs"
	ORGANIZATIONUSECASE_NOTES               OrganizationUseCase = "notes"
	ORGANIZATIONUSECASE_OTHER               OrganizationUseCase = "other"
)

// All allowed values of OrganizationUseCase enum
var AllowedOrganizationUseCaseEnumValues = []OrganizationUseCase{
	"productDocs",
	"teamKnowledgeBase",
	"designSystem",
	"openSourceDocs",
	"notes",
	"other",
}

func (v *OrganizationUseCase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationUseCase(value)
	for _, existing := range AllowedOrganizationUseCaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationUseCase", value)
}

// NewOrganizationUseCaseFromValue returns a pointer to a valid OrganizationUseCase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationUseCaseFromValue(v string) (*OrganizationUseCase, error) {
	ev := OrganizationUseCase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationUseCase: valid values are %v", v, AllowedOrganizationUseCaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationUseCase) IsValid() bool {
	for _, existing := range AllowedOrganizationUseCaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationUseCase value
func (v OrganizationUseCase) Ptr() *OrganizationUseCase {
	return &v
}

type NullableOrganizationUseCase struct {
	value *OrganizationUseCase
	isSet bool
}

func (v NullableOrganizationUseCase) Get() *OrganizationUseCase {
	return v.value
}

func (v *NullableOrganizationUseCase) Set(val *OrganizationUseCase) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUseCase) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUseCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUseCase(val *OrganizationUseCase) *NullableOrganizationUseCase {
	return &NullableOrganizationUseCase{value: val, isSet: true}
}

func (v NullableOrganizationUseCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUseCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}