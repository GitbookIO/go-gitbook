# Revision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;revision\&quot; | 
**Id** | **string** | Unique identifier for the revision | 
**Parents** | **[]string** | IDs of the parent revisions | 
**Pages** | [**[]RevisionPage**](RevisionPage.md) |  | 
**Files** | [**[]RevisionFile**](RevisionFile.md) |  | 
**Git** | Pointer to [**RevisionGit**](RevisionGit.md) |  | [optional] 
**Urls** | [**RevisionUrls**](RevisionUrls.md) |  | 

## Methods

### NewRevision

`func NewRevision(object string, id string, parents []string, pages []RevisionPage, files []RevisionFile, urls RevisionUrls, ) *Revision`

NewRevision instantiates a new Revision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionWithDefaults

`func NewRevisionWithDefaults() *Revision`

NewRevisionWithDefaults instantiates a new Revision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Revision) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Revision) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Revision) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Revision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Revision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Revision) SetId(v string)`

SetId sets Id field to given value.


### GetParents

`func (o *Revision) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Revision) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Revision) SetParents(v []string)`

SetParents sets Parents field to given value.


### GetPages

`func (o *Revision) GetPages() []RevisionPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *Revision) GetPagesOk() (*[]RevisionPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *Revision) SetPages(v []RevisionPage)`

SetPages sets Pages field to given value.


### GetFiles

`func (o *Revision) GetFiles() []RevisionFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Revision) GetFilesOk() (*[]RevisionFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Revision) SetFiles(v []RevisionFile)`

SetFiles sets Files field to given value.


### GetGit

`func (o *Revision) GetGit() RevisionGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *Revision) GetGitOk() (*RevisionGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *Revision) SetGit(v RevisionGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *Revision) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetUrls

`func (o *Revision) GetUrls() RevisionUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Revision) GetUrlsOk() (*RevisionUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Revision) SetUrls(v RevisionUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


