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

// GitSyncOperationState * `running`: The operation is still running * `failure`: The operation failed * `success`: The operation was successful
type GitSyncOperationState string

// List of GitSyncOperationState
const (
	GITSYNCOPERATIONSTATE_RUNNING GitSyncOperationState = "running"
	GITSYNCOPERATIONSTATE_FAILURE GitSyncOperationState = "failure"
	GITSYNCOPERATIONSTATE_SUCCESS GitSyncOperationState = "success"
)

// All allowed values of GitSyncOperationState enum
var AllowedGitSyncOperationStateEnumValues = []GitSyncOperationState{
	"running",
	"failure",
	"success",
}

func (v *GitSyncOperationState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GitSyncOperationState(value)
	for _, existing := range AllowedGitSyncOperationStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GitSyncOperationState", value)
}

// NewGitSyncOperationStateFromValue returns a pointer to a valid GitSyncOperationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGitSyncOperationStateFromValue(v string) (*GitSyncOperationState, error) {
	ev := GitSyncOperationState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GitSyncOperationState: valid values are %v", v, AllowedGitSyncOperationStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GitSyncOperationState) IsValid() bool {
	for _, existing := range AllowedGitSyncOperationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GitSyncOperationState value
func (v GitSyncOperationState) Ptr() *GitSyncOperationState {
	return &v
}

type NullableGitSyncOperationState struct {
	value *GitSyncOperationState
	isSet bool
}

func (v NullableGitSyncOperationState) Get() *GitSyncOperationState {
	return v.value
}

func (v *NullableGitSyncOperationState) Set(val *GitSyncOperationState) {
	v.value = val
	v.isSet = true
}

func (v NullableGitSyncOperationState) IsSet() bool {
	return v.isSet
}

func (v *NullableGitSyncOperationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitSyncOperationState(val *GitSyncOperationState) *NullableGitSyncOperationState {
	return &NullableGitSyncOperationState{value: val, isSet: true}
}

func (v NullableGitSyncOperationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitSyncOperationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}