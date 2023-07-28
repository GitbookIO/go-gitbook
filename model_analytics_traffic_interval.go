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

// AnalyticsTrafficInterval the model 'AnalyticsTrafficInterval'
type AnalyticsTrafficInterval string

// List of AnalyticsTrafficInterval
const (
	ANALYTICSTRAFFICINTERVAL_DAILY   AnalyticsTrafficInterval = "daily"
	ANALYTICSTRAFFICINTERVAL_WEEKLY  AnalyticsTrafficInterval = "weekly"
	ANALYTICSTRAFFICINTERVAL_MONTHLY AnalyticsTrafficInterval = "monthly"
)

// All allowed values of AnalyticsTrafficInterval enum
var AllowedAnalyticsTrafficIntervalEnumValues = []AnalyticsTrafficInterval{
	"daily",
	"weekly",
	"monthly",
}

func (v *AnalyticsTrafficInterval) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnalyticsTrafficInterval(value)
	for _, existing := range AllowedAnalyticsTrafficIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnalyticsTrafficInterval", value)
}

// NewAnalyticsTrafficIntervalFromValue returns a pointer to a valid AnalyticsTrafficInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnalyticsTrafficIntervalFromValue(v string) (*AnalyticsTrafficInterval, error) {
	ev := AnalyticsTrafficInterval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnalyticsTrafficInterval: valid values are %v", v, AllowedAnalyticsTrafficIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnalyticsTrafficInterval) IsValid() bool {
	for _, existing := range AllowedAnalyticsTrafficIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalyticsTrafficInterval value
func (v AnalyticsTrafficInterval) Ptr() *AnalyticsTrafficInterval {
	return &v
}

type NullableAnalyticsTrafficInterval struct {
	value *AnalyticsTrafficInterval
	isSet bool
}

func (v NullableAnalyticsTrafficInterval) Get() *AnalyticsTrafficInterval {
	return v.value
}

func (v *NullableAnalyticsTrafficInterval) Set(val *AnalyticsTrafficInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsTrafficInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsTrafficInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsTrafficInterval(val *AnalyticsTrafficInterval) *NullableAnalyticsTrafficInterval {
	return &NullableAnalyticsTrafficInterval{value: val, isSet: true}
}

func (v NullableAnalyticsTrafficInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsTrafficInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
