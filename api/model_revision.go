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

// Revision - struct for Revision
type Revision struct {
	RevisionTypeEdits    *RevisionTypeEdits
	RevisionTypeMerge    *RevisionTypeMerge
	RevisionTypeRollback *RevisionTypeRollback
	RevisionTypeUpdate   *RevisionTypeUpdate
}

// RevisionTypeEditsAsRevision is a convenience function that returns RevisionTypeEdits wrapped in Revision
func RevisionTypeEditsAsRevision(v *RevisionTypeEdits) Revision {
	return Revision{
		RevisionTypeEdits: v,
	}
}

// RevisionTypeMergeAsRevision is a convenience function that returns RevisionTypeMerge wrapped in Revision
func RevisionTypeMergeAsRevision(v *RevisionTypeMerge) Revision {
	return Revision{
		RevisionTypeMerge: v,
	}
}

// RevisionTypeRollbackAsRevision is a convenience function that returns RevisionTypeRollback wrapped in Revision
func RevisionTypeRollbackAsRevision(v *RevisionTypeRollback) Revision {
	return Revision{
		RevisionTypeRollback: v,
	}
}

// RevisionTypeUpdateAsRevision is a convenience function that returns RevisionTypeUpdate wrapped in Revision
func RevisionTypeUpdateAsRevision(v *RevisionTypeUpdate) Revision {
	return Revision{
		RevisionTypeUpdate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Revision) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RevisionTypeEdits
	err = newStrictDecoder(data).Decode(&dst.RevisionTypeEdits)
	if err == nil {
		jsonRevisionTypeEdits, _ := json.Marshal(dst.RevisionTypeEdits)
		if string(jsonRevisionTypeEdits) == "{}" { // empty struct
			dst.RevisionTypeEdits = nil
		} else {
			match++
		}
	} else {
		dst.RevisionTypeEdits = nil
	}

	// try to unmarshal data into RevisionTypeMerge
	err = newStrictDecoder(data).Decode(&dst.RevisionTypeMerge)
	if err == nil {
		jsonRevisionTypeMerge, _ := json.Marshal(dst.RevisionTypeMerge)
		if string(jsonRevisionTypeMerge) == "{}" { // empty struct
			dst.RevisionTypeMerge = nil
		} else {
			match++
		}
	} else {
		dst.RevisionTypeMerge = nil
	}

	// try to unmarshal data into RevisionTypeRollback
	err = newStrictDecoder(data).Decode(&dst.RevisionTypeRollback)
	if err == nil {
		jsonRevisionTypeRollback, _ := json.Marshal(dst.RevisionTypeRollback)
		if string(jsonRevisionTypeRollback) == "{}" { // empty struct
			dst.RevisionTypeRollback = nil
		} else {
			match++
		}
	} else {
		dst.RevisionTypeRollback = nil
	}

	// try to unmarshal data into RevisionTypeUpdate
	err = newStrictDecoder(data).Decode(&dst.RevisionTypeUpdate)
	if err == nil {
		jsonRevisionTypeUpdate, _ := json.Marshal(dst.RevisionTypeUpdate)
		if string(jsonRevisionTypeUpdate) == "{}" { // empty struct
			dst.RevisionTypeUpdate = nil
		} else {
			match++
		}
	} else {
		dst.RevisionTypeUpdate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RevisionTypeEdits = nil
		dst.RevisionTypeMerge = nil
		dst.RevisionTypeRollback = nil
		dst.RevisionTypeUpdate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Revision)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Revision)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Revision) MarshalJSON() ([]byte, error) {
	if src.RevisionTypeEdits != nil {
		return json.Marshal(&src.RevisionTypeEdits)
	}

	if src.RevisionTypeMerge != nil {
		return json.Marshal(&src.RevisionTypeMerge)
	}

	if src.RevisionTypeRollback != nil {
		return json.Marshal(&src.RevisionTypeRollback)
	}

	if src.RevisionTypeUpdate != nil {
		return json.Marshal(&src.RevisionTypeUpdate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Revision) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RevisionTypeEdits != nil {
		return obj.RevisionTypeEdits
	}

	if obj.RevisionTypeMerge != nil {
		return obj.RevisionTypeMerge
	}

	if obj.RevisionTypeRollback != nil {
		return obj.RevisionTypeRollback
	}

	if obj.RevisionTypeUpdate != nil {
		return obj.RevisionTypeUpdate
	}

	// all schemas are nil
	return nil
}

type NullableRevision struct {
	value *Revision
	isSet bool
}

func (v NullableRevision) Get() *Revision {
	return v.value
}

func (v *NullableRevision) Set(val *Revision) {
	v.value = val
	v.isSet = true
}

func (v NullableRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevision(val *Revision) *NullableRevision {
	return &NullableRevision{value: val, isSet: true}
}

func (v NullableRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
