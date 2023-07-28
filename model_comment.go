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

// checks if the Comment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Comment{}

// Comment struct for Comment
type Comment struct {
	// Type of Object, always equals to \"comment\"
	Object string `json:"object"`
	// Unique identifier for the comment.
	Id        string          `json:"id"`
	PostedBy  User            `json:"postedBy"`
	PostedAt  string          `json:"postedAt"`
	EditedAt  *string         `json:"editedAt,omitempty"`
	Reactions []EmojiReaction `json:"reactions"`
	// The number of replies to this comment.
	Replies float32            `json:"replies"`
	Body    Document           `json:"body"`
	Target  CommentAllOfTarget `json:"target"`
	Urls    CommentAllOfUrls   `json:"urls"`
	// Status of the comment.
	Status     string `json:"status"`
	ResolvedAt string `json:"resolvedAt"`
	ResolvedBy User   `json:"resolvedBy"`
}

// NewComment instantiates a new Comment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment(object string, id string, postedBy User, postedAt string, reactions []EmojiReaction, replies float32, body Document, target CommentAllOfTarget, urls CommentAllOfUrls, status string, resolvedAt string, resolvedBy User) *Comment {
	this := Comment{}
	this.Object = object
	this.Id = id
	this.PostedBy = postedBy
	this.PostedAt = postedAt
	this.Reactions = reactions
	this.Replies = replies
	this.Body = body
	this.Target = target
	this.Urls = urls
	this.Status = status
	this.ResolvedAt = resolvedAt
	this.ResolvedBy = resolvedBy
	return &this
}

// NewCommentWithDefaults instantiates a new Comment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentWithDefaults() *Comment {
	this := Comment{}
	return &this
}

// GetObject returns the Object field value
func (o *Comment) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Comment) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Comment) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Comment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Comment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Comment) SetId(v string) {
	o.Id = v
}

// GetPostedBy returns the PostedBy field value
func (o *Comment) GetPostedBy() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.PostedBy
}

// GetPostedByOk returns a tuple with the PostedBy field value
// and a boolean to check if the value has been set.
func (o *Comment) GetPostedByOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostedBy, true
}

// SetPostedBy sets field value
func (o *Comment) SetPostedBy(v User) {
	o.PostedBy = v
}

// GetPostedAt returns the PostedAt field value
func (o *Comment) GetPostedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostedAt
}

// GetPostedAtOk returns a tuple with the PostedAt field value
// and a boolean to check if the value has been set.
func (o *Comment) GetPostedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostedAt, true
}

// SetPostedAt sets field value
func (o *Comment) SetPostedAt(v string) {
	o.PostedAt = v
}

// GetEditedAt returns the EditedAt field value if set, zero value otherwise.
func (o *Comment) GetEditedAt() string {
	if o == nil || IsNil(o.EditedAt) {
		var ret string
		return ret
	}
	return *o.EditedAt
}

// GetEditedAtOk returns a tuple with the EditedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetEditedAtOk() (*string, bool) {
	if o == nil || IsNil(o.EditedAt) {
		return nil, false
	}
	return o.EditedAt, true
}

// HasEditedAt returns a boolean if a field has been set.
func (o *Comment) HasEditedAt() bool {
	if o != nil && !IsNil(o.EditedAt) {
		return true
	}

	return false
}

// SetEditedAt gets a reference to the given string and assigns it to the EditedAt field.
func (o *Comment) SetEditedAt(v string) {
	o.EditedAt = &v
}

// GetReactions returns the Reactions field value
func (o *Comment) GetReactions() []EmojiReaction {
	if o == nil {
		var ret []EmojiReaction
		return ret
	}

	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value
// and a boolean to check if the value has been set.
func (o *Comment) GetReactionsOk() ([]EmojiReaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reactions, true
}

// SetReactions sets field value
func (o *Comment) SetReactions(v []EmojiReaction) {
	o.Reactions = v
}

// GetReplies returns the Replies field value
func (o *Comment) GetReplies() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Replies
}

// GetRepliesOk returns a tuple with the Replies field value
// and a boolean to check if the value has been set.
func (o *Comment) GetRepliesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replies, true
}

// SetReplies sets field value
func (o *Comment) SetReplies(v float32) {
	o.Replies = v
}

// GetBody returns the Body field value
func (o *Comment) GetBody() Document {
	if o == nil {
		var ret Document
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *Comment) GetBodyOk() (*Document, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *Comment) SetBody(v Document) {
	o.Body = v
}

// GetTarget returns the Target field value
func (o *Comment) GetTarget() CommentAllOfTarget {
	if o == nil {
		var ret CommentAllOfTarget
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Comment) GetTargetOk() (*CommentAllOfTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Comment) SetTarget(v CommentAllOfTarget) {
	o.Target = v
}

// GetUrls returns the Urls field value
func (o *Comment) GetUrls() CommentAllOfUrls {
	if o == nil {
		var ret CommentAllOfUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Comment) GetUrlsOk() (*CommentAllOfUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Comment) SetUrls(v CommentAllOfUrls) {
	o.Urls = v
}

// GetStatus returns the Status field value
func (o *Comment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Comment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Comment) SetStatus(v string) {
	o.Status = v
}

// GetResolvedAt returns the ResolvedAt field value
func (o *Comment) GetResolvedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResolvedAt
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value
// and a boolean to check if the value has been set.
func (o *Comment) GetResolvedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResolvedAt, true
}

// SetResolvedAt sets field value
func (o *Comment) SetResolvedAt(v string) {
	o.ResolvedAt = v
}

// GetResolvedBy returns the ResolvedBy field value
func (o *Comment) GetResolvedBy() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.ResolvedBy
}

// GetResolvedByOk returns a tuple with the ResolvedBy field value
// and a boolean to check if the value has been set.
func (o *Comment) GetResolvedByOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResolvedBy, true
}

// SetResolvedBy sets field value
func (o *Comment) SetResolvedBy(v User) {
	o.ResolvedBy = v
}

func (o Comment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Comment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["postedBy"] = o.PostedBy
	toSerialize["postedAt"] = o.PostedAt
	if !IsNil(o.EditedAt) {
		toSerialize["editedAt"] = o.EditedAt
	}
	toSerialize["reactions"] = o.Reactions
	toSerialize["replies"] = o.Replies
	toSerialize["body"] = o.Body
	toSerialize["target"] = o.Target
	toSerialize["urls"] = o.Urls
	toSerialize["status"] = o.Status
	toSerialize["resolvedAt"] = o.ResolvedAt
	toSerialize["resolvedBy"] = o.ResolvedBy
	return toSerialize, nil
}

type NullableComment struct {
	value *Comment
	isSet bool
}

func (v NullableComment) Get() *Comment {
	return v.value
}

func (v *NullableComment) Set(val *Comment) {
	v.value = val
	v.isSet = true
}

func (v NullableComment) IsSet() bool {
	return v.isSet
}

func (v *NullableComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment(val *Comment) *NullableComment {
	return &NullableComment{value: val, isSet: true}
}

func (v NullableComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}