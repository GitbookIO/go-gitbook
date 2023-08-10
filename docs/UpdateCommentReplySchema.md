# UpdateCommentReplySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**Document**](Document.md) |  | [optional] 
**AddedReactions** | Pointer to **[]string** | Reactions to add to the comment. | [optional] 
**RemovedReactions** | Pointer to **[]string** | Reactions to remove from the comment. | [optional] 

## Methods

### NewUpdateCommentReplySchema

`func NewUpdateCommentReplySchema() *UpdateCommentReplySchema`

NewUpdateCommentReplySchema instantiates a new UpdateCommentReplySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCommentReplySchemaWithDefaults

`func NewUpdateCommentReplySchemaWithDefaults() *UpdateCommentReplySchema`

NewUpdateCommentReplySchemaWithDefaults instantiates a new UpdateCommentReplySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *UpdateCommentReplySchema) GetBody() Document`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateCommentReplySchema) GetBodyOk() (*Document, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateCommentReplySchema) SetBody(v Document)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateCommentReplySchema) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAddedReactions

`func (o *UpdateCommentReplySchema) GetAddedReactions() []string`

GetAddedReactions returns the AddedReactions field if non-nil, zero value otherwise.

### GetAddedReactionsOk

`func (o *UpdateCommentReplySchema) GetAddedReactionsOk() (*[]string, bool)`

GetAddedReactionsOk returns a tuple with the AddedReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedReactions

`func (o *UpdateCommentReplySchema) SetAddedReactions(v []string)`

SetAddedReactions sets AddedReactions field to given value.

### HasAddedReactions

`func (o *UpdateCommentReplySchema) HasAddedReactions() bool`

HasAddedReactions returns a boolean if a field has been set.

### GetRemovedReactions

`func (o *UpdateCommentReplySchema) GetRemovedReactions() []string`

GetRemovedReactions returns the RemovedReactions field if non-nil, zero value otherwise.

### GetRemovedReactionsOk

`func (o *UpdateCommentReplySchema) GetRemovedReactionsOk() (*[]string, bool)`

GetRemovedReactionsOk returns a tuple with the RemovedReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedReactions

`func (o *UpdateCommentReplySchema) SetRemovedReactions(v []string)`

SetRemovedReactions sets RemovedReactions field to given value.

### HasRemovedReactions

`func (o *UpdateCommentReplySchema) HasRemovedReactions() bool`

HasRemovedReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


