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

// checks if the SubscriptionChannelOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionChannelOneOf2{}

// SubscriptionChannelOneOf2 struct for SubscriptionChannelOneOf2
type SubscriptionChannelOneOf2 struct {
	Channel       string  `json:"channel"`
	Space         string  `json:"space"`
	ChangeRequest *string `json:"changeRequest,omitempty"`
	Comment       string  `json:"comment"`
	CommentReply  string  `json:"commentReply"`
}

// NewSubscriptionChannelOneOf2 instantiates a new SubscriptionChannelOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionChannelOneOf2(channel string, space string, comment string, commentReply string) *SubscriptionChannelOneOf2 {
	this := SubscriptionChannelOneOf2{}
	this.Channel = channel
	this.Space = space
	this.Comment = comment
	this.CommentReply = commentReply
	return &this
}

// NewSubscriptionChannelOneOf2WithDefaults instantiates a new SubscriptionChannelOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionChannelOneOf2WithDefaults() *SubscriptionChannelOneOf2 {
	this := SubscriptionChannelOneOf2{}
	return &this
}

// GetChannel returns the Channel field value
func (o *SubscriptionChannelOneOf2) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SubscriptionChannelOneOf2) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *SubscriptionChannelOneOf2) SetChannel(v string) {
	o.Channel = v
}

// GetSpace returns the Space field value
func (o *SubscriptionChannelOneOf2) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *SubscriptionChannelOneOf2) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *SubscriptionChannelOneOf2) SetSpace(v string) {
	o.Space = v
}

// GetChangeRequest returns the ChangeRequest field value if set, zero value otherwise.
func (o *SubscriptionChannelOneOf2) GetChangeRequest() string {
	if o == nil || IsNil(o.ChangeRequest) {
		var ret string
		return ret
	}
	return *o.ChangeRequest
}

// GetChangeRequestOk returns a tuple with the ChangeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionChannelOneOf2) GetChangeRequestOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeRequest) {
		return nil, false
	}
	return o.ChangeRequest, true
}

// HasChangeRequest returns a boolean if a field has been set.
func (o *SubscriptionChannelOneOf2) HasChangeRequest() bool {
	if o != nil && !IsNil(o.ChangeRequest) {
		return true
	}

	return false
}

// SetChangeRequest gets a reference to the given string and assigns it to the ChangeRequest field.
func (o *SubscriptionChannelOneOf2) SetChangeRequest(v string) {
	o.ChangeRequest = &v
}

// GetComment returns the Comment field value
func (o *SubscriptionChannelOneOf2) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *SubscriptionChannelOneOf2) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *SubscriptionChannelOneOf2) SetComment(v string) {
	o.Comment = v
}

// GetCommentReply returns the CommentReply field value
func (o *SubscriptionChannelOneOf2) GetCommentReply() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommentReply
}

// GetCommentReplyOk returns a tuple with the CommentReply field value
// and a boolean to check if the value has been set.
func (o *SubscriptionChannelOneOf2) GetCommentReplyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentReply, true
}

// SetCommentReply sets field value
func (o *SubscriptionChannelOneOf2) SetCommentReply(v string) {
	o.CommentReply = v
}

func (o SubscriptionChannelOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionChannelOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["space"] = o.Space
	if !IsNil(o.ChangeRequest) {
		toSerialize["changeRequest"] = o.ChangeRequest
	}
	toSerialize["comment"] = o.Comment
	toSerialize["commentReply"] = o.CommentReply
	return toSerialize, nil
}

type NullableSubscriptionChannelOneOf2 struct {
	value *SubscriptionChannelOneOf2
	isSet bool
}

func (v NullableSubscriptionChannelOneOf2) Get() *SubscriptionChannelOneOf2 {
	return v.value
}

func (v *NullableSubscriptionChannelOneOf2) Set(val *SubscriptionChannelOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionChannelOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionChannelOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionChannelOneOf2(val *SubscriptionChannelOneOf2) *NullableSubscriptionChannelOneOf2 {
	return &NullableSubscriptionChannelOneOf2{value: val, isSet: true}
}

func (v NullableSubscriptionChannelOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionChannelOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
