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

// checks if the CommentAllOfTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentAllOfTarget{}

// CommentAllOfTarget Information about the target of the comment.
type CommentAllOfTarget struct {
	Node *CommentAllOfTargetNode `json:"node,omitempty"`
	// The change request containing this comment, if the comment was made inside a change request.
	ChangeRequest *string `json:"changeRequest,omitempty"`
	// The review containing this comment, if this comment was made as part of a review.
	Review *string `json:"review,omitempty"`
	// Information about the page, if this comment refers to a specific page.
	Page *RevisionPageBase `json:"page,omitempty"`
	// The space containing this comment.
	Space string `json:"space"`
	// The revision in which the target can be found in.
	Revision string `json:"revision"`
}

// NewCommentAllOfTarget instantiates a new CommentAllOfTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentAllOfTarget(space string, revision string) *CommentAllOfTarget {
	this := CommentAllOfTarget{}
	this.Space = space
	this.Revision = revision
	return &this
}

// NewCommentAllOfTargetWithDefaults instantiates a new CommentAllOfTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentAllOfTargetWithDefaults() *CommentAllOfTarget {
	this := CommentAllOfTarget{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *CommentAllOfTarget) GetNode() CommentAllOfTargetNode {
	if o == nil || IsNil(o.Node) {
		var ret CommentAllOfTargetNode
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentAllOfTarget) GetNodeOk() (*CommentAllOfTargetNode, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *CommentAllOfTarget) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given CommentAllOfTargetNode and assigns it to the Node field.
func (o *CommentAllOfTarget) SetNode(v CommentAllOfTargetNode) {
	o.Node = &v
}

// GetChangeRequest returns the ChangeRequest field value if set, zero value otherwise.
func (o *CommentAllOfTarget) GetChangeRequest() string {
	if o == nil || IsNil(o.ChangeRequest) {
		var ret string
		return ret
	}
	return *o.ChangeRequest
}

// GetChangeRequestOk returns a tuple with the ChangeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentAllOfTarget) GetChangeRequestOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeRequest) {
		return nil, false
	}
	return o.ChangeRequest, true
}

// HasChangeRequest returns a boolean if a field has been set.
func (o *CommentAllOfTarget) HasChangeRequest() bool {
	if o != nil && !IsNil(o.ChangeRequest) {
		return true
	}

	return false
}

// SetChangeRequest gets a reference to the given string and assigns it to the ChangeRequest field.
func (o *CommentAllOfTarget) SetChangeRequest(v string) {
	o.ChangeRequest = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *CommentAllOfTarget) GetReview() string {
	if o == nil || IsNil(o.Review) {
		var ret string
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentAllOfTarget) GetReviewOk() (*string, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *CommentAllOfTarget) HasReview() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given string and assigns it to the Review field.
func (o *CommentAllOfTarget) SetReview(v string) {
	o.Review = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CommentAllOfTarget) GetPage() RevisionPageBase {
	if o == nil || IsNil(o.Page) {
		var ret RevisionPageBase
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentAllOfTarget) GetPageOk() (*RevisionPageBase, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CommentAllOfTarget) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given RevisionPageBase and assigns it to the Page field.
func (o *CommentAllOfTarget) SetPage(v RevisionPageBase) {
	o.Page = &v
}

// GetSpace returns the Space field value
func (o *CommentAllOfTarget) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *CommentAllOfTarget) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *CommentAllOfTarget) SetSpace(v string) {
	o.Space = v
}

// GetRevision returns the Revision field value
func (o *CommentAllOfTarget) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *CommentAllOfTarget) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *CommentAllOfTarget) SetRevision(v string) {
	o.Revision = v
}

func (o CommentAllOfTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentAllOfTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.ChangeRequest) {
		toSerialize["changeRequest"] = o.ChangeRequest
	}
	if !IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	toSerialize["space"] = o.Space
	toSerialize["revision"] = o.Revision
	return toSerialize, nil
}

type NullableCommentAllOfTarget struct {
	value *CommentAllOfTarget
	isSet bool
}

func (v NullableCommentAllOfTarget) Get() *CommentAllOfTarget {
	return v.value
}

func (v *NullableCommentAllOfTarget) Set(val *CommentAllOfTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentAllOfTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentAllOfTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentAllOfTarget(val *CommentAllOfTarget) *NullableCommentAllOfTarget {
	return &NullableCommentAllOfTarget{value: val, isSet: true}
}

func (v NullableCommentAllOfTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentAllOfTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
