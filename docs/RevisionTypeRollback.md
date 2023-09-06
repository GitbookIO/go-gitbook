# RevisionTypeRollback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;revision\&quot; | 
**Id** | **string** | Unique identifier for the revision | 
**Parents** | **[]string** | IDs of the parent revisions | 
**Pages** | [**[]RevisionPage**](RevisionPage.md) |  | 
**Files** | [**[]RevisionFile**](RevisionFile.md) |  | 
**Git** | Pointer to [**RevisionBaseGit**](RevisionBaseGit.md) |  | [optional] 
**Urls** | [**RevisionBaseUrls**](RevisionBaseUrls.md) |  | 
**Type** | **string** | Revision created as a rollback of a previous revision. | 

## Methods

### NewRevisionTypeRollback

`func NewRevisionTypeRollback(object string, id string, parents []string, pages []RevisionPage, files []RevisionFile, urls RevisionBaseUrls, type_ string, ) *RevisionTypeRollback`

NewRevisionTypeRollback instantiates a new RevisionTypeRollback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionTypeRollbackWithDefaults

`func NewRevisionTypeRollbackWithDefaults() *RevisionTypeRollback`

NewRevisionTypeRollbackWithDefaults instantiates a new RevisionTypeRollback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *RevisionTypeRollback) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RevisionTypeRollback) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RevisionTypeRollback) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *RevisionTypeRollback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionTypeRollback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionTypeRollback) SetId(v string)`

SetId sets Id field to given value.


### GetParents

`func (o *RevisionTypeRollback) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *RevisionTypeRollback) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *RevisionTypeRollback) SetParents(v []string)`

SetParents sets Parents field to given value.


### GetPages

`func (o *RevisionTypeRollback) GetPages() []RevisionPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RevisionTypeRollback) GetPagesOk() (*[]RevisionPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RevisionTypeRollback) SetPages(v []RevisionPage)`

SetPages sets Pages field to given value.


### GetFiles

`func (o *RevisionTypeRollback) GetFiles() []RevisionFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RevisionTypeRollback) GetFilesOk() (*[]RevisionFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RevisionTypeRollback) SetFiles(v []RevisionFile)`

SetFiles sets Files field to given value.


### GetGit

`func (o *RevisionTypeRollback) GetGit() RevisionBaseGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *RevisionTypeRollback) GetGitOk() (*RevisionBaseGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *RevisionTypeRollback) SetGit(v RevisionBaseGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *RevisionTypeRollback) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetUrls

`func (o *RevisionTypeRollback) GetUrls() RevisionBaseUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *RevisionTypeRollback) GetUrlsOk() (*RevisionBaseUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *RevisionTypeRollback) SetUrls(v RevisionBaseUrls)`

SetUrls sets Urls field to given value.


### GetType

`func (o *RevisionTypeRollback) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevisionTypeRollback) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevisionTypeRollback) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


