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

// IntegrationConfiguration - struct for IntegrationConfiguration
type IntegrationConfiguration struct {
	IntegrationConfigurationComponent *IntegrationConfigurationComponent
	IntegrationConfigurationSchema    *IntegrationConfigurationSchema
}

// IntegrationConfigurationComponentAsIntegrationConfiguration is a convenience function that returns IntegrationConfigurationComponent wrapped in IntegrationConfiguration
func IntegrationConfigurationComponentAsIntegrationConfiguration(v *IntegrationConfigurationComponent) IntegrationConfiguration {
	return IntegrationConfiguration{
		IntegrationConfigurationComponent: v,
	}
}

// IntegrationConfigurationSchemaAsIntegrationConfiguration is a convenience function that returns IntegrationConfigurationSchema wrapped in IntegrationConfiguration
func IntegrationConfigurationSchemaAsIntegrationConfiguration(v *IntegrationConfigurationSchema) IntegrationConfiguration {
	return IntegrationConfiguration{
		IntegrationConfigurationSchema: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IntegrationConfiguration) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IntegrationConfigurationComponent
	err = newStrictDecoder(data).Decode(&dst.IntegrationConfigurationComponent)
	if err == nil {
		jsonIntegrationConfigurationComponent, _ := json.Marshal(dst.IntegrationConfigurationComponent)
		if string(jsonIntegrationConfigurationComponent) == "{}" { // empty struct
			dst.IntegrationConfigurationComponent = nil
		} else {
			match++
		}
	} else {
		dst.IntegrationConfigurationComponent = nil
	}

	// try to unmarshal data into IntegrationConfigurationSchema
	err = newStrictDecoder(data).Decode(&dst.IntegrationConfigurationSchema)
	if err == nil {
		jsonIntegrationConfigurationSchema, _ := json.Marshal(dst.IntegrationConfigurationSchema)
		if string(jsonIntegrationConfigurationSchema) == "{}" { // empty struct
			dst.IntegrationConfigurationSchema = nil
		} else {
			match++
		}
	} else {
		dst.IntegrationConfigurationSchema = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IntegrationConfigurationComponent = nil
		dst.IntegrationConfigurationSchema = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IntegrationConfiguration)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IntegrationConfiguration)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IntegrationConfiguration) MarshalJSON() ([]byte, error) {
	if src.IntegrationConfigurationComponent != nil {
		return json.Marshal(&src.IntegrationConfigurationComponent)
	}

	if src.IntegrationConfigurationSchema != nil {
		return json.Marshal(&src.IntegrationConfigurationSchema)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IntegrationConfiguration) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.IntegrationConfigurationComponent != nil {
		return obj.IntegrationConfigurationComponent
	}

	if obj.IntegrationConfigurationSchema != nil {
		return obj.IntegrationConfigurationSchema
	}

	// all schemas are nil
	return nil
}

type NullableIntegrationConfiguration struct {
	value *IntegrationConfiguration
	isSet bool
}

func (v NullableIntegrationConfiguration) Get() *IntegrationConfiguration {
	return v.value
}

func (v *NullableIntegrationConfiguration) Set(val *IntegrationConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfiguration(val *IntegrationConfiguration) *NullableIntegrationConfiguration {
	return &NullableIntegrationConfiguration{value: val, isSet: true}
}

func (v NullableIntegrationConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
