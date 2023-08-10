# CommentReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;comment-reply\&quot; | 
**Id** | **string** | Unique identifier for the reply. | 
**PostedBy** | [**User**](User.md) |  | 
**PostedAt** | **string** |  | 
**EditedAt** | Pointer to **string** |  | [optional] 
**Reactions** | [**[]EmojiReaction**](EmojiReaction.md) |  | 
**Body** | [**Document**](Document.md) |  | 
**Urls** | [**CommentReplyUrls**](CommentReplyUrls.md) |  | 

## Methods

### NewCommentReply

`func NewCommentReply(object string, id string, postedBy User, postedAt string, reactions []EmojiReaction, body Document, urls CommentReplyUrls, ) *CommentReply`

NewCommentReply instantiates a new CommentReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentReplyWithDefaults

`func NewCommentReplyWithDefaults() *CommentReply`

NewCommentReplyWithDefaults instantiates a new CommentReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CommentReply) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CommentReply) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CommentReply) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *CommentReply) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommentReply) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommentReply) SetId(v string)`

SetId sets Id field to given value.


### GetPostedBy

`func (o *CommentReply) GetPostedBy() User`

GetPostedBy returns the PostedBy field if non-nil, zero value otherwise.

### GetPostedByOk

`func (o *CommentReply) GetPostedByOk() (*User, bool)`

GetPostedByOk returns a tuple with the PostedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedBy

`func (o *CommentReply) SetPostedBy(v User)`

SetPostedBy sets PostedBy field to given value.


### GetPostedAt

`func (o *CommentReply) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *CommentReply) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *CommentReply) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.


### GetEditedAt

`func (o *CommentReply) GetEditedAt() string`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *CommentReply) GetEditedAtOk() (*string, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *CommentReply) SetEditedAt(v string)`

SetEditedAt sets EditedAt field to given value.

### HasEditedAt

`func (o *CommentReply) HasEditedAt() bool`

HasEditedAt returns a boolean if a field has been set.

### GetReactions

`func (o *CommentReply) GetReactions() []EmojiReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CommentReply) GetReactionsOk() (*[]EmojiReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CommentReply) SetReactions(v []EmojiReaction)`

SetReactions sets Reactions field to given value.


### GetBody

`func (o *CommentReply) GetBody() Document`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CommentReply) GetBodyOk() (*Document, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CommentReply) SetBody(v Document)`

SetBody sets Body field to given value.


### GetUrls

`func (o *CommentReply) GetUrls() CommentReplyUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CommentReply) GetUrlsOk() (*CommentReplyUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CommentReply) SetUrls(v CommentReplyUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


