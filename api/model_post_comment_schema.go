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
)

// checks if the PostCommentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCommentSchema{}

// PostCommentSchema struct for PostCommentSchema
type PostCommentSchema struct {
	// The node to which the comment is posted, if any.
	Node *string `json:"node,omitempty"`
	// The page to which the comment is posted, if any.
	Page *string  `json:"page,omitempty"`
	Body Document `json:"body"`
}

// NewPostCommentSchema instantiates a new PostCommentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCommentSchema(body Document) *PostCommentSchema {
	this := PostCommentSchema{}
	this.Body = body
	return &this
}

// NewPostCommentSchemaWithDefaults instantiates a new PostCommentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCommentSchemaWithDefaults() *PostCommentSchema {
	this := PostCommentSchema{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *PostCommentSchema) GetNode() string {
	if o == nil || IsNil(o.Node) {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCommentSchema) GetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *PostCommentSchema) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *PostCommentSchema) SetNode(v string) {
	o.Node = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PostCommentSchema) GetPage() string {
	if o == nil || IsNil(o.Page) {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCommentSchema) GetPageOk() (*string, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PostCommentSchema) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *PostCommentSchema) SetPage(v string) {
	o.Page = &v
}

// GetBody returns the Body field value
func (o *PostCommentSchema) GetBody() Document {
	if o == nil {
		var ret Document
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *PostCommentSchema) GetBodyOk() (*Document, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *PostCommentSchema) SetBody(v Document) {
	o.Body = v
}

func (o PostCommentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCommentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

type NullablePostCommentSchema struct {
	value *PostCommentSchema
	isSet bool
}

func (v NullablePostCommentSchema) Get() *PostCommentSchema {
	return v.value
}

func (v *NullablePostCommentSchema) Set(val *PostCommentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCommentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCommentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCommentSchema(val *PostCommentSchema) *NullablePostCommentSchema {
	return &NullablePostCommentSchema{value: val, isSet: true}
}

func (v NullablePostCommentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCommentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
