// Copyright 2023 GitBook, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
	"fmt"
)

// CloudflareHostnameStatus The Cloudflare Hostname status
type CloudflareHostnameStatus string

// List of CloudflareHostnameStatus
const (
	CLOUDFLAREHOSTNAMESTATUS_PENDING CloudflareHostnameStatus = "pending"
	CLOUDFLAREHOSTNAMESTATUS_ACTIVE  CloudflareHostnameStatus = "active"
	CLOUDFLAREHOSTNAMESTATUS_BLOCKED CloudflareHostnameStatus = "blocked"
	CLOUDFLAREHOSTNAMESTATUS_MOVED   CloudflareHostnameStatus = "moved"
	CLOUDFLAREHOSTNAMESTATUS_DELETED CloudflareHostnameStatus = "deleted"
)

// All allowed values of CloudflareHostnameStatus enum
var AllowedCloudflareHostnameStatusEnumValues = []CloudflareHostnameStatus{
	"pending",
	"active",
	"blocked",
	"moved",
	"deleted",
}

func (v *CloudflareHostnameStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudflareHostnameStatus(value)
	for _, existing := range AllowedCloudflareHostnameStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudflareHostnameStatus", value)
}

// NewCloudflareHostnameStatusFromValue returns a pointer to a valid CloudflareHostnameStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudflareHostnameStatusFromValue(v string) (*CloudflareHostnameStatus, error) {
	ev := CloudflareHostnameStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudflareHostnameStatus: valid values are %v", v, AllowedCloudflareHostnameStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudflareHostnameStatus) IsValid() bool {
	for _, existing := range AllowedCloudflareHostnameStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudflareHostnameStatus value
func (v CloudflareHostnameStatus) Ptr() *CloudflareHostnameStatus {
	return &v
}

type NullableCloudflareHostnameStatus struct {
	value *CloudflareHostnameStatus
	isSet bool
}

func (v NullableCloudflareHostnameStatus) Get() *CloudflareHostnameStatus {
	return v.value
}

func (v *NullableCloudflareHostnameStatus) Set(val *CloudflareHostnameStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudflareHostnameStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudflareHostnameStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudflareHostnameStatus(val *CloudflareHostnameStatus) *NullableCloudflareHostnameStatus {
	return &NullableCloudflareHostnameStatus{value: val, isSet: true}
}

func (v NullableCloudflareHostnameStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudflareHostnameStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
