# CommentAllOf

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

## Methods

### NewCommentAllOf

`func NewCommentAllOf(object string, id string, postedBy User, postedAt string, reactions []EmojiReaction, replies float32, body Document, target CommentAllOfTarget, urls CommentAllOfUrls, ) *CommentAllOf`

NewCommentAllOf instantiates a new CommentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentAllOfWithDefaults

`func NewCommentAllOfWithDefaults() *CommentAllOf`

NewCommentAllOfWithDefaults instantiates a new CommentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CommentAllOf) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CommentAllOf) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CommentAllOf) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *CommentAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetPostedBy

`func (o *CommentAllOf) GetPostedBy() User`

GetPostedBy returns the PostedBy field if non-nil, zero value otherwise.

### GetPostedByOk

`func (o *CommentAllOf) GetPostedByOk() (*User, bool)`

GetPostedByOk returns a tuple with the PostedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedBy

`func (o *CommentAllOf) SetPostedBy(v User)`

SetPostedBy sets PostedBy field to given value.


### GetPostedAt

`func (o *CommentAllOf) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *CommentAllOf) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *CommentAllOf) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.


### GetEditedAt

`func (o *CommentAllOf) GetEditedAt() string`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *CommentAllOf) GetEditedAtOk() (*string, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *CommentAllOf) SetEditedAt(v string)`

SetEditedAt sets EditedAt field to given value.

### HasEditedAt

`func (o *CommentAllOf) HasEditedAt() bool`

HasEditedAt returns a boolean if a field has been set.

### GetReactions

`func (o *CommentAllOf) GetReactions() []EmojiReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CommentAllOf) GetReactionsOk() (*[]EmojiReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CommentAllOf) SetReactions(v []EmojiReaction)`

SetReactions sets Reactions field to given value.


### GetReplies

`func (o *CommentAllOf) GetReplies() float32`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *CommentAllOf) GetRepliesOk() (*float32, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *CommentAllOf) SetReplies(v float32)`

SetReplies sets Replies field to given value.


### GetBody

`func (o *CommentAllOf) GetBody() Document`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CommentAllOf) GetBodyOk() (*Document, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CommentAllOf) SetBody(v Document)`

SetBody sets Body field to given value.


### GetTarget

`func (o *CommentAllOf) GetTarget() CommentAllOfTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CommentAllOf) GetTargetOk() (*CommentAllOfTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CommentAllOf) SetTarget(v CommentAllOfTarget)`

SetTarget sets Target field to given value.


### GetUrls

`func (o *CommentAllOf) GetUrls() CommentAllOfUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CommentAllOf) GetUrlsOk() (*CommentAllOfUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CommentAllOf) SetUrls(v CommentAllOfUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


