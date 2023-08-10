# UpdateCommentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolved** | Pointer to **bool** | Whether the comment is resolved or not. | [optional] 
**Body** | Pointer to [**Document**](Document.md) |  | [optional] 
**AddedReactions** | Pointer to **[]string** | Reactions to add to the comment. | [optional] 
**RemovedReactions** | Pointer to **[]string** | Reactions to remove from the comment. | [optional] 

## Methods

### NewUpdateCommentSchema

`func NewUpdateCommentSchema() *UpdateCommentSchema`

NewUpdateCommentSchema instantiates a new UpdateCommentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCommentSchemaWithDefaults

`func NewUpdateCommentSchemaWithDefaults() *UpdateCommentSchema`

NewUpdateCommentSchemaWithDefaults instantiates a new UpdateCommentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolved

`func (o *UpdateCommentSchema) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *UpdateCommentSchema) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *UpdateCommentSchema) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *UpdateCommentSchema) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetBody

`func (o *UpdateCommentSchema) GetBody() Document`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateCommentSchema) GetBodyOk() (*Document, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateCommentSchema) SetBody(v Document)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateCommentSchema) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAddedReactions

`func (o *UpdateCommentSchema) GetAddedReactions() []string`

GetAddedReactions returns the AddedReactions field if non-nil, zero value otherwise.

### GetAddedReactionsOk

`func (o *UpdateCommentSchema) GetAddedReactionsOk() (*[]string, bool)`

GetAddedReactionsOk returns a tuple with the AddedReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedReactions

`func (o *UpdateCommentSchema) SetAddedReactions(v []string)`

SetAddedReactions sets AddedReactions field to given value.

### HasAddedReactions

`func (o *UpdateCommentSchema) HasAddedReactions() bool`

HasAddedReactions returns a boolean if a field has been set.

### GetRemovedReactions

`func (o *UpdateCommentSchema) GetRemovedReactions() []string`

GetRemovedReactions returns the RemovedReactions field if non-nil, zero value otherwise.

### GetRemovedReactionsOk

`func (o *UpdateCommentSchema) GetRemovedReactionsOk() (*[]string, bool)`

GetRemovedReactionsOk returns a tuple with the RemovedReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedReactions

`func (o *UpdateCommentSchema) SetRemovedReactions(v []string)`

SetRemovedReactions sets RemovedReactions field to given value.

### HasRemovedReactions

`func (o *UpdateCommentSchema) HasRemovedReactions() bool`

HasRemovedReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


