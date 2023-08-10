# EmojiReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | **string** | The Emoji of the reaction | 
**Count** | **float32** | The number of users who reacted with this emoji | 
**Users** | [**[]EmojiReactionUsersInner**](EmojiReactionUsersInner.md) | The users who reacted with this emoji | 

## Methods

### NewEmojiReaction

`func NewEmojiReaction(emoji string, count float32, users []EmojiReactionUsersInner, ) *EmojiReaction`

NewEmojiReaction instantiates a new EmojiReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmojiReactionWithDefaults

`func NewEmojiReactionWithDefaults() *EmojiReaction`

NewEmojiReactionWithDefaults instantiates a new EmojiReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *EmojiReaction) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *EmojiReaction) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *EmojiReaction) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.


### GetCount

`func (o *EmojiReaction) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EmojiReaction) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EmojiReaction) SetCount(v float32)`

SetCount sets Count field to given value.


### GetUsers

`func (o *EmojiReaction) GetUsers() []EmojiReactionUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EmojiReaction) GetUsersOk() (*[]EmojiReactionUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EmojiReaction) SetUsers(v []EmojiReactionUsersInner)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


