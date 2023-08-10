# SpaceGitSyncCompletedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 
**State** | **string** |  | 
**RevisionId** | **string** | Unique identifier of the new content revision | 
**CommitId** | **string** | Unique identifier for the commit (sha) | 

## Methods

### NewSpaceGitSyncCompletedEvent

`func NewSpaceGitSyncCompletedEvent(eventId string, type_ string, installationId string, spaceId string, state string, revisionId string, commitId string, ) *SpaceGitSyncCompletedEvent`

NewSpaceGitSyncCompletedEvent instantiates a new SpaceGitSyncCompletedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceGitSyncCompletedEventWithDefaults

`func NewSpaceGitSyncCompletedEventWithDefaults() *SpaceGitSyncCompletedEvent`

NewSpaceGitSyncCompletedEventWithDefaults instantiates a new SpaceGitSyncCompletedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SpaceGitSyncCompletedEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpaceGitSyncCompletedEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SpaceGitSyncCompletedEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SpaceGitSyncCompletedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceGitSyncCompletedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceGitSyncCompletedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *SpaceGitSyncCompletedEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SpaceGitSyncCompletedEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SpaceGitSyncCompletedEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *SpaceGitSyncCompletedEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SpaceGitSyncCompletedEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SpaceGitSyncCompletedEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetState

`func (o *SpaceGitSyncCompletedEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SpaceGitSyncCompletedEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SpaceGitSyncCompletedEvent) SetState(v string)`

SetState sets State field to given value.


### GetRevisionId

`func (o *SpaceGitSyncCompletedEvent) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *SpaceGitSyncCompletedEvent) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *SpaceGitSyncCompletedEvent) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetCommitId

`func (o *SpaceGitSyncCompletedEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *SpaceGitSyncCompletedEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *SpaceGitSyncCompletedEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


