# RequestImportGitRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL of the Git repository to import. It can contain basic auth credentials. | 
**Ref** | **string** | Git ref to import in the format \&quot;refs/heads/main\&quot; | 
**RepoCacheID** | Pointer to **string** | Unique identifier to use to cache the Git repository across multiple operations. | [optional] 
**RepoTreeURL** | Pointer to **string** | URL to use as a prefix for external file references. | [optional] 
**RepoCommitURL** | Pointer to **string** | URL to use as a prefix for the commit URL. | [optional] 
**RepoProjectDirectory** | Pointer to **string** | Path to a root directory for the project in the repository. | [optional] 
**Force** | Pointer to **bool** |  | [optional] 
**Standalone** | Pointer to **bool** | If true, the import will generate a revision without updating the space primary content. | [optional] 
**GitInfo** | Pointer to [**RequestUpdateSpaceGitInfo**](RequestUpdateSpaceGitInfo.md) |  | [optional] 

## Methods

### NewRequestImportGitRepository

`func NewRequestImportGitRepository(url string, ref string, ) *RequestImportGitRepository`

NewRequestImportGitRepository instantiates a new RequestImportGitRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestImportGitRepositoryWithDefaults

`func NewRequestImportGitRepositoryWithDefaults() *RequestImportGitRepository`

NewRequestImportGitRepositoryWithDefaults instantiates a new RequestImportGitRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RequestImportGitRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestImportGitRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestImportGitRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRef

`func (o *RequestImportGitRepository) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RequestImportGitRepository) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RequestImportGitRepository) SetRef(v string)`

SetRef sets Ref field to given value.


### GetRepoCacheID

`func (o *RequestImportGitRepository) GetRepoCacheID() string`

GetRepoCacheID returns the RepoCacheID field if non-nil, zero value otherwise.

### GetRepoCacheIDOk

`func (o *RequestImportGitRepository) GetRepoCacheIDOk() (*string, bool)`

GetRepoCacheIDOk returns a tuple with the RepoCacheID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCacheID

`func (o *RequestImportGitRepository) SetRepoCacheID(v string)`

SetRepoCacheID sets RepoCacheID field to given value.

### HasRepoCacheID

`func (o *RequestImportGitRepository) HasRepoCacheID() bool`

HasRepoCacheID returns a boolean if a field has been set.

### GetRepoTreeURL

`func (o *RequestImportGitRepository) GetRepoTreeURL() string`

GetRepoTreeURL returns the RepoTreeURL field if non-nil, zero value otherwise.

### GetRepoTreeURLOk

`func (o *RequestImportGitRepository) GetRepoTreeURLOk() (*string, bool)`

GetRepoTreeURLOk returns a tuple with the RepoTreeURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoTreeURL

`func (o *RequestImportGitRepository) SetRepoTreeURL(v string)`

SetRepoTreeURL sets RepoTreeURL field to given value.

### HasRepoTreeURL

`func (o *RequestImportGitRepository) HasRepoTreeURL() bool`

HasRepoTreeURL returns a boolean if a field has been set.

### GetRepoCommitURL

`func (o *RequestImportGitRepository) GetRepoCommitURL() string`

GetRepoCommitURL returns the RepoCommitURL field if non-nil, zero value otherwise.

### GetRepoCommitURLOk

`func (o *RequestImportGitRepository) GetRepoCommitURLOk() (*string, bool)`

GetRepoCommitURLOk returns a tuple with the RepoCommitURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCommitURL

`func (o *RequestImportGitRepository) SetRepoCommitURL(v string)`

SetRepoCommitURL sets RepoCommitURL field to given value.

### HasRepoCommitURL

`func (o *RequestImportGitRepository) HasRepoCommitURL() bool`

HasRepoCommitURL returns a boolean if a field has been set.

### GetRepoProjectDirectory

`func (o *RequestImportGitRepository) GetRepoProjectDirectory() string`

GetRepoProjectDirectory returns the RepoProjectDirectory field if non-nil, zero value otherwise.

### GetRepoProjectDirectoryOk

`func (o *RequestImportGitRepository) GetRepoProjectDirectoryOk() (*string, bool)`

GetRepoProjectDirectoryOk returns a tuple with the RepoProjectDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoProjectDirectory

`func (o *RequestImportGitRepository) SetRepoProjectDirectory(v string)`

SetRepoProjectDirectory sets RepoProjectDirectory field to given value.

### HasRepoProjectDirectory

`func (o *RequestImportGitRepository) HasRepoProjectDirectory() bool`

HasRepoProjectDirectory returns a boolean if a field has been set.

### GetForce

`func (o *RequestImportGitRepository) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *RequestImportGitRepository) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *RequestImportGitRepository) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *RequestImportGitRepository) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetStandalone

`func (o *RequestImportGitRepository) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *RequestImportGitRepository) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *RequestImportGitRepository) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *RequestImportGitRepository) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetGitInfo

`func (o *RequestImportGitRepository) GetGitInfo() RequestUpdateSpaceGitInfo`

GetGitInfo returns the GitInfo field if non-nil, zero value otherwise.

### GetGitInfoOk

`func (o *RequestImportGitRepository) GetGitInfoOk() (*RequestUpdateSpaceGitInfo, bool)`

GetGitInfoOk returns a tuple with the GitInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitInfo

`func (o *RequestImportGitRepository) SetGitInfo(v RequestUpdateSpaceGitInfo)`

SetGitInfo sets GitInfo field to given value.

### HasGitInfo

`func (o *RequestImportGitRepository) HasGitInfo() bool`

HasGitInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


