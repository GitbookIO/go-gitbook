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

// SpaceType - `application`: A fully-packaged application, like a mobile application, desktop application, or a CLI-type tool. - `website`: A single web page or a collection of related web pages under a single domain. - `service`: An independently-deployable software unit that is usually is operated by a person or a team. - `library`: A reusable collection of objects, functions, and methods. A library is typically used by other spaces. - `cloud-resource`: An entity or service provided by a cloud vendor, with consumer-managed configuration and monitoring. - `database`: A databse entity, persisting data. A database is typically used by services or applications. - `other`: A space that doesn’t match the service, library, or application type.
type SpaceType string

// List of SpaceType
const (
	SPACETYPE_APPLICATION    SpaceType = "application"
	SPACETYPE_WEBSITE        SpaceType = "website"
	SPACETYPE_SERVICE        SpaceType = "service"
	SPACETYPE_LIBRARY        SpaceType = "library"
	SPACETYPE_CLOUD_RESOURCE SpaceType = "cloud-resource"
	SPACETYPE_DATABASE       SpaceType = "database"
	SPACETYPE_OTHER          SpaceType = "other"
)

// All allowed values of SpaceType enum
var AllowedSpaceTypeEnumValues = []SpaceType{
	"application",
	"website",
	"service",
	"library",
	"cloud-resource",
	"database",
	"other",
}

func (v *SpaceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SpaceType(value)
	for _, existing := range AllowedSpaceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SpaceType", value)
}

// NewSpaceTypeFromValue returns a pointer to a valid SpaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSpaceTypeFromValue(v string) (*SpaceType, error) {
	ev := SpaceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SpaceType: valid values are %v", v, AllowedSpaceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SpaceType) IsValid() bool {
	for _, existing := range AllowedSpaceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpaceType value
func (v SpaceType) Ptr() *SpaceType {
	return &v
}

type NullableSpaceType struct {
	value *SpaceType
	isSet bool
}

func (v NullableSpaceType) Get() *SpaceType {
	return v.value
}

func (v *NullableSpaceType) Set(val *SpaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceType(val *SpaceType) *NullableSpaceType {
	return &NullableSpaceType{value: val, isSet: true}
}

func (v NullableSpaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}