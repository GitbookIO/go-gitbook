# RevisionTypeUpdate

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
**Type** | **string** | Revision created when updating a change request with changes from primary. | 

## Methods

### NewRevisionTypeUpdate

`func NewRevisionTypeUpdate(object string, id string, parents []string, pages []RevisionPage, files []RevisionFile, urls RevisionBaseUrls, type_ string, ) *RevisionTypeUpdate`

NewRevisionTypeUpdate instantiates a new RevisionTypeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionTypeUpdateWithDefaults

`func NewRevisionTypeUpdateWithDefaults() *RevisionTypeUpdate`

NewRevisionTypeUpdateWithDefaults instantiates a new RevisionTypeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *RevisionTypeUpdate) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RevisionTypeUpdate) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RevisionTypeUpdate) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *RevisionTypeUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionTypeUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionTypeUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetParents

`func (o *RevisionTypeUpdate) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *RevisionTypeUpdate) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *RevisionTypeUpdate) SetParents(v []string)`

SetParents sets Parents field to given value.


### GetPages

`func (o *RevisionTypeUpdate) GetPages() []RevisionPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RevisionTypeUpdate) GetPagesOk() (*[]RevisionPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RevisionTypeUpdate) SetPages(v []RevisionPage)`

SetPages sets Pages field to given value.


### GetFiles

`func (o *RevisionTypeUpdate) GetFiles() []RevisionFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RevisionTypeUpdate) GetFilesOk() (*[]RevisionFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RevisionTypeUpdate) SetFiles(v []RevisionFile)`

SetFiles sets Files field to given value.


### GetGit

`func (o *RevisionTypeUpdate) GetGit() RevisionBaseGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *RevisionTypeUpdate) GetGitOk() (*RevisionBaseGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *RevisionTypeUpdate) SetGit(v RevisionBaseGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *RevisionTypeUpdate) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetUrls

`func (o *RevisionTypeUpdate) GetUrls() RevisionBaseUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *RevisionTypeUpdate) GetUrlsOk() (*RevisionBaseUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *RevisionTypeUpdate) SetUrls(v RevisionBaseUrls)`

SetUrls sets Urls field to given value.


### GetType

`func (o *RevisionTypeUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevisionTypeUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevisionTypeUpdate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


