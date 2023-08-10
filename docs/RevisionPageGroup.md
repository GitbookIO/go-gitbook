# RevisionPageGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the page in the revision | 
**Title** | **string** | Title of the page | 
**Kind** | **string** |  | 
**Type** | **string** |  | 
**Slug** | **string** | Page&#39;s slug in its direct parent | 
**Path** | **string** | Complete path to access the page in the revision. | 
**Pages** | [**[]RevisionPageDocumentAllOfPagesInner**](RevisionPageDocumentAllOfPagesInner.md) |  | 

## Methods

### NewRevisionPageGroup

`func NewRevisionPageGroup(id string, title string, kind string, type_ string, slug string, path string, pages []RevisionPageDocumentAllOfPagesInner, ) *RevisionPageGroup`

NewRevisionPageGroup instantiates a new RevisionPageGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionPageGroupWithDefaults

`func NewRevisionPageGroupWithDefaults() *RevisionPageGroup`

NewRevisionPageGroupWithDefaults instantiates a new RevisionPageGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RevisionPageGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionPageGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionPageGroup) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *RevisionPageGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RevisionPageGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RevisionPageGroup) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetKind

`func (o *RevisionPageGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevisionPageGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevisionPageGroup) SetKind(v string)`

SetKind sets Kind field to given value.


### GetType

`func (o *RevisionPageGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevisionPageGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevisionPageGroup) SetType(v string)`

SetType sets Type field to given value.


### GetSlug

`func (o *RevisionPageGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RevisionPageGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RevisionPageGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPath

`func (o *RevisionPageGroup) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RevisionPageGroup) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RevisionPageGroup) SetPath(v string)`

SetPath sets Path field to given value.


### GetPages

`func (o *RevisionPageGroup) GetPages() []RevisionPageDocumentAllOfPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RevisionPageGroup) GetPagesOk() (*[]RevisionPageDocumentAllOfPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RevisionPageGroup) SetPages(v []RevisionPageDocumentAllOfPagesInner)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


