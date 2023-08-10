# SyncDirectorySyncGroupsToTeams200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Synced** | [**[]OrganizationDirectorySyncGroupTeamStatus**](OrganizationDirectorySyncGroupTeamStatus.md) | The list of synced pairs, mapped to the original pairs given to the sync parameters. Either the group_id and team_id will be defined, or the error will be defined to describe why it failed. Use the success parameter to know if the pair sync was sucessful or not.  | 

## Methods

### NewSyncDirectorySyncGroupsToTeams200Response

`func NewSyncDirectorySyncGroupsToTeams200Response(synced []OrganizationDirectorySyncGroupTeamStatus, ) *SyncDirectorySyncGroupsToTeams200Response`

NewSyncDirectorySyncGroupsToTeams200Response instantiates a new SyncDirectorySyncGroupsToTeams200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncDirectorySyncGroupsToTeams200ResponseWithDefaults

`func NewSyncDirectorySyncGroupsToTeams200ResponseWithDefaults() *SyncDirectorySyncGroupsToTeams200Response`

NewSyncDirectorySyncGroupsToTeams200ResponseWithDefaults instantiates a new SyncDirectorySyncGroupsToTeams200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSynced

`func (o *SyncDirectorySyncGroupsToTeams200Response) GetSynced() []OrganizationDirectorySyncGroupTeamStatus`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *SyncDirectorySyncGroupsToTeams200Response) GetSyncedOk() (*[]OrganizationDirectorySyncGroupTeamStatus, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *SyncDirectorySyncGroupsToTeams200Response) SetSynced(v []OrganizationDirectorySyncGroupTeamStatus)`

SetSynced sets Synced field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


