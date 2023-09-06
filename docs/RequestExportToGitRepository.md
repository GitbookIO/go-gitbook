# RequestExportToGitRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL of the Git repository to export to. It can contain basic auth credentials. | 
**Ref** | **string** | Git ref to push the commit to in the format \&quot;refs/heads/main\&quot; | 
**CommitMessage** | **string** | Message for the commit generated by the export | 
**RepoCacheID** | Pointer to **string** | Unique identifier to use to cache the Git repository across multiple operations. | [optional] 
**RepoTreeURL** | Pointer to **string** | URL to use as a prefix for external file references. | [optional] 
**RepoCommitURL** | Pointer to **string** | URL to use as a prefix for the commit URL. | [optional] 
**RepoProjectDirectory** | Pointer to **string** | Path to a root directory for the project in the repository. | [optional] 
**Force** | Pointer to **bool** |  | [optional] 
**GitInfo** | Pointer to [**RequestUpdateSpaceGitInfo**](RequestUpdateSpaceGitInfo.md) |  | [optional] 

## Methods

### NewRequestExportToGitRepository

`func NewRequestExportToGitRepository(url string, ref string, commitMessage string, ) *RequestExportToGitRepository`

NewRequestExportToGitRepository instantiates a new RequestExportToGitRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestExportToGitRepositoryWithDefaults

`func NewRequestExportToGitRepositoryWithDefaults() *RequestExportToGitRepository`

NewRequestExportToGitRepositoryWithDefaults instantiates a new RequestExportToGitRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RequestExportToGitRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestExportToGitRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestExportToGitRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRef

`func (o *RequestExportToGitRepository) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RequestExportToGitRepository) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RequestExportToGitRepository) SetRef(v string)`

SetRef sets Ref field to given value.


### GetCommitMessage

`func (o *RequestExportToGitRepository) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *RequestExportToGitRepository) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *RequestExportToGitRepository) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.


### GetRepoCacheID

`func (o *RequestExportToGitRepository) GetRepoCacheID() string`

GetRepoCacheID returns the RepoCacheID field if non-nil, zero value otherwise.

### GetRepoCacheIDOk

`func (o *RequestExportToGitRepository) GetRepoCacheIDOk() (*string, bool)`

GetRepoCacheIDOk returns a tuple with the RepoCacheID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCacheID

`func (o *RequestExportToGitRepository) SetRepoCacheID(v string)`

SetRepoCacheID sets RepoCacheID field to given value.

### HasRepoCacheID

`func (o *RequestExportToGitRepository) HasRepoCacheID() bool`

HasRepoCacheID returns a boolean if a field has been set.

### GetRepoTreeURL

`func (o *RequestExportToGitRepository) GetRepoTreeURL() string`

GetRepoTreeURL returns the RepoTreeURL field if non-nil, zero value otherwise.

### GetRepoTreeURLOk

`func (o *RequestExportToGitRepository) GetRepoTreeURLOk() (*string, bool)`

GetRepoTreeURLOk returns a tuple with the RepoTreeURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoTreeURL

`func (o *RequestExportToGitRepository) SetRepoTreeURL(v string)`

SetRepoTreeURL sets RepoTreeURL field to given value.

### HasRepoTreeURL

`func (o *RequestExportToGitRepository) HasRepoTreeURL() bool`

HasRepoTreeURL returns a boolean if a field has been set.

### GetRepoCommitURL

`func (o *RequestExportToGitRepository) GetRepoCommitURL() string`

GetRepoCommitURL returns the RepoCommitURL field if non-nil, zero value otherwise.

### GetRepoCommitURLOk

`func (o *RequestExportToGitRepository) GetRepoCommitURLOk() (*string, bool)`

GetRepoCommitURLOk returns a tuple with the RepoCommitURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCommitURL

`func (o *RequestExportToGitRepository) SetRepoCommitURL(v string)`

SetRepoCommitURL sets RepoCommitURL field to given value.

### HasRepoCommitURL

`func (o *RequestExportToGitRepository) HasRepoCommitURL() bool`

HasRepoCommitURL returns a boolean if a field has been set.

### GetRepoProjectDirectory

`func (o *RequestExportToGitRepository) GetRepoProjectDirectory() string`

GetRepoProjectDirectory returns the RepoProjectDirectory field if non-nil, zero value otherwise.

### GetRepoProjectDirectoryOk

`func (o *RequestExportToGitRepository) GetRepoProjectDirectoryOk() (*string, bool)`

GetRepoProjectDirectoryOk returns a tuple with the RepoProjectDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoProjectDirectory

`func (o *RequestExportToGitRepository) SetRepoProjectDirectory(v string)`

SetRepoProjectDirectory sets RepoProjectDirectory field to given value.

### HasRepoProjectDirectory

`func (o *RequestExportToGitRepository) HasRepoProjectDirectory() bool`

HasRepoProjectDirectory returns a boolean if a field has been set.

### GetForce

`func (o *RequestExportToGitRepository) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *RequestExportToGitRepository) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *RequestExportToGitRepository) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *RequestExportToGitRepository) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetGitInfo

`func (o *RequestExportToGitRepository) GetGitInfo() RequestUpdateSpaceGitInfo`

GetGitInfo returns the GitInfo field if non-nil, zero value otherwise.

### GetGitInfoOk

`func (o *RequestExportToGitRepository) GetGitInfoOk() (*RequestUpdateSpaceGitInfo, bool)`

GetGitInfoOk returns a tuple with the GitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitInfo

`func (o *RequestExportToGitRepository) SetGitInfo(v RequestUpdateSpaceGitInfo)`

SetGitInfo sets GitInfo field to given value.

### HasGitInfo

`func (o *RequestExportToGitRepository) HasGitInfo() bool`

HasGitInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

