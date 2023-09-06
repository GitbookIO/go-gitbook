# RevisionBase

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

## Methods

### NewRevisionBase

`func NewRevisionBase(object string, id string, parents []string, pages []RevisionPage, files []RevisionFile, urls RevisionBaseUrls, ) *RevisionBase`

NewRevisionBase instantiates a new RevisionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionBaseWithDefaults

`func NewRevisionBaseWithDefaults() *RevisionBase`

NewRevisionBaseWithDefaults instantiates a new RevisionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *RevisionBase) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RevisionBase) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RevisionBase) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *RevisionBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionBase) SetId(v string)`

SetId sets Id field to given value.


### GetParents

`func (o *RevisionBase) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *RevisionBase) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *RevisionBase) SetParents(v []string)`

SetParents sets Parents field to given value.


### GetPages

`func (o *RevisionBase) GetPages() []RevisionPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RevisionBase) GetPagesOk() (*[]RevisionPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RevisionBase) SetPages(v []RevisionPage)`

SetPages sets Pages field to given value.


### GetFiles

`func (o *RevisionBase) GetFiles() []RevisionFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RevisionBase) GetFilesOk() (*[]RevisionFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RevisionBase) SetFiles(v []RevisionFile)`

SetFiles sets Files field to given value.


### GetGit

`func (o *RevisionBase) GetGit() RevisionBaseGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *RevisionBase) GetGitOk() (*RevisionBaseGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *RevisionBase) SetGit(v RevisionBaseGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *RevisionBase) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetUrls

`func (o *RevisionBase) GetUrls() RevisionBaseUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *RevisionBase) GetUrlsOk() (*RevisionBaseUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *RevisionBase) SetUrls(v RevisionBaseUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


