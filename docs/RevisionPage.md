# RevisionPage

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
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewRevisionPage

`func NewRevisionPage(id string, title string, markdown string, document JSONDocumentDocument, kind string, type_ string, slug string, path string, pages []RevisionPageDocumentAllOfPagesInner, ) *RevisionPage`

NewRevisionPage instantiates a new RevisionPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionPageWithDefaults

`func NewRevisionPageWithDefaults() *RevisionPage`

NewRevisionPageWithDefaults instantiates a new RevisionPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RevisionPage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RevisionPage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RevisionPage) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *RevisionPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RevisionPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RevisionPage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMarkdown

`func (o *RevisionPage) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *RevisionPage) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *RevisionPage) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetDocument

`func (o *RevisionPage) GetDocument() JSONDocumentDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *RevisionPage) GetDocumentOk() (*JSONDocumentDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *RevisionPage) SetDocument(v JSONDocumentDocument)`

SetDocument sets Document field to given value.


### GetKind

`func (o *RevisionPage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevisionPage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevisionPage) SetKind(v string)`

SetKind sets Kind field to given value.


### GetType

`func (o *RevisionPage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevisionPage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevisionPage) SetType(v string)`

SetType sets Type field to given value.


### GetSlug

`func (o *RevisionPage) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RevisionPage) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RevisionPage) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPath

`func (o *RevisionPage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RevisionPage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RevisionPage) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *RevisionPage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RevisionPage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RevisionPage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RevisionPage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPages

`func (o *RevisionPage) GetPages() []RevisionPageDocumentAllOfPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RevisionPage) GetPagesOk() (*[]RevisionPageDocumentAllOfPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RevisionPage) SetPages(v []RevisionPageDocumentAllOfPagesInner)`

SetPages sets Pages field to given value.


### GetHref

`func (o *RevisionPage) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RevisionPage) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RevisionPage) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RevisionPage) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


