# RevisionPageDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the page in the revision | 
**Title** | **string** | Title of the page | 
**Markdown** | **string** | Content of the document formatted as markdown | 
**Document** | [**JSONDocumentDocument**](JSONDocumentDocument.md) |  | 
**Kind** | **string** |  | 
**Type** | **string** |  | 
**Slug** | **string** | Page&#39;s slug in its direct parent | 
**Path** | **string** | Complete path to access the page in the revision. | 
**Description** | Pointer to **string** |  | [optional] 
**Pages** | [**[]RevisionPageDocumentAllOfPagesInner**](RevisionPageDocumentAllOfPagesInner.md) |  | 

## Methods

### NewRevisionPageDocument

`func NewRevisionPageDocument(id string, title string, markdown string, document JSONDocumentDocument, kind string, type_ string, slug string, path string, pages []RevisionPageDocumentAllOfPagesInner, ) *RevisionPageDocument`

NewRevisionPageDocument instantiates a new RevisionPageDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionPageDocumentWithDefaults

`func NewRevisionPageDocumentWithDefaults() *RevisionPageDocument`

NewRevisionPageDocumentWithDefaults instantiates a new RevisionPageDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RevisionPageDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionPageDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionPageDocument) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *RevisionPageDocument) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RevisionPageDocument) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RevisionPageDocument) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMarkdown

`func (o *RevisionPageDocument) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *RevisionPageDocument) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *RevisionPageDocument) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetDocument

`func (o *RevisionPageDocument) GetDocument() JSONDocumentDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *RevisionPageDocument) GetDocumentOk() (*JSONDocumentDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *RevisionPageDocument) SetDocument(v JSONDocumentDocument)`

SetDocument sets Document field to given value.


### GetKind

`func (o *RevisionPageDocument) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevisionPageDocument) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevisionPageDocument) SetKind(v string)`

SetKind sets Kind field to given value.


### GetType

`func (o *RevisionPageDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevisionPageDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevisionPageDocument) SetType(v string)`

SetType sets Type field to given value.


### GetSlug

`func (o *RevisionPageDocument) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RevisionPageDocument) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RevisionPageDocument) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPath

`func (o *RevisionPageDocument) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RevisionPageDocument) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RevisionPageDocument) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *RevisionPageDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RevisionPageDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RevisionPageDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RevisionPageDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPages

`func (o *RevisionPageDocument) GetPages() []RevisionPageDocumentAllOfPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RevisionPageDocument) GetPagesOk() (*[]RevisionPageDocumentAllOfPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RevisionPageDocument) SetPages(v []RevisionPageDocumentAllOfPagesInner)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


