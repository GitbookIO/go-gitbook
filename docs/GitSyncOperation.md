# GitSyncOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**GitSyncOperationState**](GitSyncOperationState.md) |  | 
**StartedAt** | **string** |  | 
**CompletedAt** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** | Error details, defined if state is &#x60;failure&#x60;. | [optional] 

## Methods

### NewGitSyncOperation

`func NewGitSyncOperation(state GitSyncOperationState, startedAt string, ) *GitSyncOperation`

NewGitSyncOperation instantiates a new GitSyncOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitSyncOperationWithDefaults

`func NewGitSyncOperationWithDefaults() *GitSyncOperation`

NewGitSyncOperationWithDefaults instantiates a new GitSyncOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GitSyncOperation) GetState() GitSyncOperationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GitSyncOperation) GetStateOk() (*GitSyncOperationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GitSyncOperation) SetState(v GitSyncOperationState)`

SetState sets State field to given value.


### GetStartedAt

`func (o *GitSyncOperation) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GitSyncOperation) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GitSyncOperation) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetCompletedAt

`func (o *GitSyncOperation) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GitSyncOperation) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GitSyncOperation) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GitSyncOperation) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetError

`func (o *GitSyncOperation) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GitSyncOperation) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GitSyncOperation) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GitSyncOperation) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


