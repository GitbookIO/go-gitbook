# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;comment\&quot; | 
**Id** | **string** | Unique identifier for the comment. | 
**PostedBy** | [**User**](User.md) |  | 
**PostedAt** | **string** |  | 
**EditedAt** | Pointer to **string** |  | [optional] 
**Reactions** | [**[]EmojiReaction**](EmojiReaction.md) |  | 
**Replies** | **float32** | The number of replies to this comment. | 
**Body** | [**Document**](Document.md) |  | 
**Target** | [**CommentAllOfTarget**](CommentAllOfTarget.md) |  | 
**Urls** | [**CommentAllOfUrls**](CommentAllOfUrls.md) |  | 
**Status** | **string** | Status of the comment. | 
**ResolvedAt** | **string** |  | 
**ResolvedBy** | [**User**](User.md) |  | 

## Methods

### NewComment

`func NewComment(object string, id string, postedBy User, postedAt string, reactions []EmojiReaction, replies float32, body Document, target CommentAllOfTarget, urls CommentAllOfUrls, status string, resolvedAt string, resolvedBy User, ) *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Comment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Comment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Comment) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.


### GetPostedBy

`func (o *Comment) GetPostedBy() User`

GetPostedBy returns the PostedBy field if non-nil, zero value otherwise.

### GetPostedByOk

`func (o *Comment) GetPostedByOk() (*User, bool)`

GetPostedByOk returns a tuple with the PostedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedBy

`func (o *Comment) SetPostedBy(v User)`

SetPostedBy sets PostedBy field to given value.


### GetPostedAt

`func (o *Comment) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *Comment) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *Comment) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.


### GetEditedAt

`func (o *Comment) GetEditedAt() string`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *Comment) GetEditedAtOk() (*string, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *Comment) SetEditedAt(v string)`

SetEditedAt sets EditedAt field to given value.

### HasEditedAt

`func (o *Comment) HasEditedAt() bool`

HasEditedAt returns a boolean if a field has been set.

### GetReactions

`func (o *Comment) GetReactions() []EmojiReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Comment) GetReactionsOk() (*[]EmojiReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Comment) SetReactions(v []EmojiReaction)`

SetReactions sets Reactions field to given value.


### GetReplies

`func (o *Comment) GetReplies() float32`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *Comment) GetRepliesOk() (*float32, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *Comment) SetReplies(v float32)`

SetReplies sets Replies field to given value.


### GetBody

`func (o *Comment) GetBody() Document`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Comment) GetBodyOk() (*Document, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Comment) SetBody(v Document)`

SetBody sets Body field to given value.


### GetTarget

`func (o *Comment) GetTarget() CommentAllOfTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Comment) GetTargetOk() (*CommentAllOfTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Comment) SetTarget(v CommentAllOfTarget)`

SetTarget sets Target field to given value.


### GetUrls

`func (o *Comment) GetUrls() CommentAllOfUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Comment) GetUrlsOk() (*CommentAllOfUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Comment) SetUrls(v CommentAllOfUrls)`

SetUrls sets Urls field to given value.


### GetStatus

`func (o *Comment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Comment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Comment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResolvedAt

`func (o *Comment) GetResolvedAt() string`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Comment) GetResolvedAtOk() (*string, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Comment) SetResolvedAt(v string)`

SetResolvedAt sets ResolvedAt field to given value.


### GetResolvedBy

`func (o *Comment) GetResolvedBy() User`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *Comment) GetResolvedByOk() (*User, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *Comment) SetResolvedBy(v User)`

SetResolvedBy sets ResolvedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


