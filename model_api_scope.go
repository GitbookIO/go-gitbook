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

// APIScope struct for APIScope
type APIScope struct {
	APIIntegrationScope *APIIntegrationScope
	IntegrationScope    *IntegrationScope
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *APIScope) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into APIIntegrationScope
	err = json.Unmarshal(data, &dst.APIIntegrationScope)
	if err == nil {
		jsonAPIIntegrationScope, _ := json.Marshal(dst.APIIntegrationScope)
		if string(jsonAPIIntegrationScope) == "{}" { // empty struct
			dst.APIIntegrationScope = nil
		} else {
			return nil // data stored in dst.APIIntegrationScope, return on the first match
		}
	} else {
		dst.APIIntegrationScope = nil
	}

	// try to unmarshal JSON data into IntegrationScope
	err = json.Unmarshal(data, &dst.IntegrationScope)
	if err == nil {
		jsonIntegrationScope, _ := json.Marshal(dst.IntegrationScope)
		if string(jsonIntegrationScope) == "{}" { // empty struct
			dst.IntegrationScope = nil
		} else {
			return nil // data stored in dst.IntegrationScope, return on the first match
		}
	} else {
		dst.IntegrationScope = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(APIScope)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *APIScope) MarshalJSON() ([]byte, error) {
	if src.APIIntegrationScope != nil {
		return json.Marshal(&src.APIIntegrationScope)
	}

	if src.IntegrationScope != nil {
		return json.Marshal(&src.IntegrationScope)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAPIScope struct {
	value *APIScope
	isSet bool
}

func (v NullableAPIScope) Get() *APIScope {
	return v.value
}

func (v *NullableAPIScope) Set(val *APIScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIScope(val *APIScope) *NullableAPIScope {
	return &NullableAPIScope{value: val, isSet: true}
}

func (v NullableAPIScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
