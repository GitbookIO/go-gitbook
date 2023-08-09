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
)

// checks if the UpdateCommentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCommentSchema{}

// UpdateCommentSchema struct for UpdateCommentSchema
type UpdateCommentSchema struct {
	// Whether the comment is resolved or not.
	Resolved *bool     `json:"resolved,omitempty"`
	Body     *Document `json:"body,omitempty"`
	// Reactions to add to the comment.
	AddedReactions []string `json:"addedReactions,omitempty"`
	// Reactions to remove from the comment.
	RemovedReactions []string `json:"removedReactions,omitempty"`
}

// NewUpdateCommentSchema instantiates a new UpdateCommentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCommentSchema() *UpdateCommentSchema {
	this := UpdateCommentSchema{}
	return &this
}

// NewUpdateCommentSchemaWithDefaults instantiates a new UpdateCommentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCommentSchemaWithDefaults() *UpdateCommentSchema {
	this := UpdateCommentSchema{}
	return &this
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *UpdateCommentSchema) GetResolved() bool {
	if o == nil || IsNil(o.Resolved) {
		var ret bool
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCommentSchema) GetResolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *UpdateCommentSchema) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given bool and assigns it to the Resolved field.
func (o *UpdateCommentSchema) SetResolved(v bool) {
	o.Resolved = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UpdateCommentSchema) GetBody() Document {
	if o == nil || IsNil(o.Body) {
		var ret Document
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCommentSchema) GetBodyOk() (*Document, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UpdateCommentSchema) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given Document and assigns it to the Body field.
func (o *UpdateCommentSchema) SetBody(v Document) {
	o.Body = &v
}

// GetAddedReactions returns the AddedReactions field value if set, zero value otherwise.
func (o *UpdateCommentSchema) GetAddedReactions() []string {
	if o == nil || IsNil(o.AddedReactions) {
		var ret []string
		return ret
	}
	return o.AddedReactions
}

// GetAddedReactionsOk returns a tuple with the AddedReactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCommentSchema) GetAddedReactionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddedReactions) {
		return nil, false
	}
	return o.AddedReactions, true
}

// HasAddedReactions returns a boolean if a field has been set.
func (o *UpdateCommentSchema) HasAddedReactions() bool {
	if o != nil && !IsNil(o.AddedReactions) {
		return true
	}

	return false
}

// SetAddedReactions gets a reference to the given []string and assigns it to the AddedReactions field.
func (o *UpdateCommentSchema) SetAddedReactions(v []string) {
	o.AddedReactions = v
}

// GetRemovedReactions returns the RemovedReactions field value if set, zero value otherwise.
func (o *UpdateCommentSchema) GetRemovedReactions() []string {
	if o == nil || IsNil(o.RemovedReactions) {
		var ret []string
		return ret
	}
	return o.RemovedReactions
}

// GetRemovedReactionsOk returns a tuple with the RemovedReactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCommentSchema) GetRemovedReactionsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemovedReactions) {
		return nil, false
	}
	return o.RemovedReactions, true
}

// HasRemovedReactions returns a boolean if a field has been set.
func (o *UpdateCommentSchema) HasRemovedReactions() bool {
	if o != nil && !IsNil(o.RemovedReactions) {
		return true
	}

	return false
}

// SetRemovedReactions gets a reference to the given []string and assigns it to the RemovedReactions field.
func (o *UpdateCommentSchema) SetRemovedReactions(v []string) {
	o.RemovedReactions = v
}

func (o UpdateCommentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCommentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resolved) {
		toSerialize["resolved"] = o.Resolved
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.AddedReactions) {
		toSerialize["addedReactions"] = o.AddedReactions
	}
	if !IsNil(o.RemovedReactions) {
		toSerialize["removedReactions"] = o.RemovedReactions
	}
	return toSerialize, nil
}

type NullableUpdateCommentSchema struct {
	value *UpdateCommentSchema
	isSet bool
}

func (v NullableUpdateCommentSchema) Get() *UpdateCommentSchema {
	return v.value
}

func (v *NullableUpdateCommentSchema) Set(val *UpdateCommentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCommentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCommentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCommentSchema(val *UpdateCommentSchema) *NullableUpdateCommentSchema {
	return &NullableUpdateCommentSchema{value: val, isSet: true}
}

func (v NullableUpdateCommentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCommentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}