# SyncDirectorySyncGroupsToTeamsRequestToSyncInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The unique identifier of the group for the given identity provider. | 
**TeamId** | Pointer to **string** | The unique identifier of the team for the given identity provider. | [optional] 

## Methods

### NewSyncDirectorySyncGroupsToTeamsRequestToSyncInner

`func NewSyncDirectorySyncGroupsToTeamsRequestToSyncInner(groupId string, ) *SyncDirectorySyncGroupsToTeamsRequestToSyncInner`

NewSyncDirectorySyncGroupsToTeamsRequestToSyncInner instantiates a new SyncDirectorySyncGroupsToTeamsRequestToSyncInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncDirectorySyncGroupsToTeamsRequestToSyncInnerWithDefaults

`func NewSyncDirectorySyncGroupsToTeamsRequestToSyncInnerWithDefaults() *SyncDirectorySyncGroupsToTeamsRequestToSyncInner`

NewSyncDirectorySyncGroupsToTeamsRequestToSyncInnerWithDefaults instantiates a new SyncDirectorySyncGroupsToTeamsRequestToSyncInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetTeamId

`func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *SyncDirectorySyncGroupsToTeamsRequestToSyncInner) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


