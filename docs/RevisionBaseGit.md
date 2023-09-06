# RevisionBaseGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | **string** | SHA of the Git commit. | 
**Message** | **string** | Git commit message. | 
**CreatedByGitBook** | **bool** | Whether not this commit was created by GitBook, while exporting the revision. | 
**Url** | Pointer to **string** | URL of the Git commit. | [optional] 

## Methods

### NewRevisionBaseGit

`func NewRevisionBaseGit(oid string, message string, createdByGitBook bool, ) *RevisionBaseGit`

NewRevisionBaseGit instantiates a new RevisionBaseGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionBaseGitWithDefaults

`func NewRevisionBaseGitWithDefaults() *RevisionBaseGit`

NewRevisionBaseGitWithDefaults instantiates a new RevisionBaseGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *RevisionBaseGit) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *RevisionBaseGit) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *RevisionBaseGit) SetOid(v string)`

SetOid sets Oid field to given value.


### GetMessage

`func (o *RevisionBaseGit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RevisionBaseGit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RevisionBaseGit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCreatedByGitBook

`func (o *RevisionBaseGit) GetCreatedByGitBook() bool`

GetCreatedByGitBook returns the CreatedByGitBook field if non-nil, zero value otherwise.

### GetCreatedByGitBookOk

`func (o *RevisionBaseGit) GetCreatedByGitBookOk() (*bool, bool)`

GetCreatedByGitBookOk returns a tuple with the CreatedByGitBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByGitBook

`func (o *RevisionBaseGit) SetCreatedByGitBook(v bool)`

SetCreatedByGitBook sets CreatedByGitBook field to given value.


### GetUrl

`func (o *RevisionBaseGit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RevisionBaseGit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RevisionBaseGit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RevisionBaseGit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


