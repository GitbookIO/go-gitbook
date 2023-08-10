# SyncDirectorySyncGroupsToTeamsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToSync** | [**[]SyncDirectorySyncGroupsToTeamsRequestToSyncInner**](SyncDirectorySyncGroupsToTeamsRequestToSyncInner.md) | A list of groups to teams pairs to sync. The group_id is required and the team_id is optional. If the latter is omitted, a new team will be created from the group&#39;s data.  | 

## Methods

### NewSyncDirectorySyncGroupsToTeamsRequest

`func NewSyncDirectorySyncGroupsToTeamsRequest(toSync []SyncDirectorySyncGroupsToTeamsRequestToSyncInner, ) *SyncDirectorySyncGroupsToTeamsRequest`

NewSyncDirectorySyncGroupsToTeamsRequest instantiates a new SyncDirectorySyncGroupsToTeamsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncDirectorySyncGroupsToTeamsRequestWithDefaults

`func NewSyncDirectorySyncGroupsToTeamsRequestWithDefaults() *SyncDirectorySyncGroupsToTeamsRequest`

NewSyncDirectorySyncGroupsToTeamsRequestWithDefaults instantiates a new SyncDirectorySyncGroupsToTeamsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToSync

`func (o *SyncDirectorySyncGroupsToTeamsRequest) GetToSync() []SyncDirectorySyncGroupsToTeamsRequestToSyncInner`

GetToSync returns the ToSync field if non-nil, zero value otherwise.

### GetToSyncOk

`func (o *SyncDirectorySyncGroupsToTeamsRequest) GetToSyncOk() (*[]SyncDirectorySyncGroupsToTeamsRequestToSyncInner, bool)`

GetToSyncOk returns a tuple with the ToSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSync

`func (o *SyncDirectorySyncGroupsToTeamsRequest) SetToSync(v []SyncDirectorySyncGroupsToTeamsRequestToSyncInner)`

SetToSync sets ToSync field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


