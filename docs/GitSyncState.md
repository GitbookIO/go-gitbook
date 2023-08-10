# GitSyncState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallationProvider** | Pointer to **string** | The provider of the Git Sync installation. | [optional] 
**Operation** | Pointer to [**GitSyncOperation**](GitSyncOperation.md) |  | [optional] 
**Url** | Pointer to **string** | The URL to the repository tree, used when rendering public content. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGitSyncState

`func NewGitSyncState() *GitSyncState`

NewGitSyncState instantiates a new GitSyncState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitSyncStateWithDefaults

`func NewGitSyncStateWithDefaults() *GitSyncState`

NewGitSyncStateWithDefaults instantiates a new GitSyncState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallationProvider

`func (o *GitSyncState) GetInstallationProvider() string`

GetInstallationProvider returns the InstallationProvider field if non-nil, zero value otherwise.

### GetInstallationProviderOk

`func (o *GitSyncState) GetInstallationProviderOk() (*string, bool)`

GetInstallationProviderOk returns a tuple with the InstallationProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationProvider

`func (o *GitSyncState) SetInstallationProvider(v string)`

SetInstallationProvider sets InstallationProvider field to given value.

### HasInstallationProvider

`func (o *GitSyncState) HasInstallationProvider() bool`

HasInstallationProvider returns a boolean if a field has been set.

### GetOperation

`func (o *GitSyncState) GetOperation() GitSyncOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GitSyncState) GetOperationOk() (*GitSyncOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GitSyncState) SetOperation(v GitSyncOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GitSyncState) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetUrl

`func (o *GitSyncState) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitSyncState) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitSyncState) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GitSyncState) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GitSyncState) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GitSyncState) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GitSyncState) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GitSyncState) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


