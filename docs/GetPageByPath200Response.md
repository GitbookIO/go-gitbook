# GetPageByPath200Response

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

### NewGetPageByPath200Response

`func NewGetPageByPath200Response(id string, title string, markdown string, document JSONDocumentDocument, kind string, type_ string, slug string, path string, pages []RevisionPageDocumentAllOfPagesInner, ) *GetPageByPath200Response`

NewGetPageByPath200Response instantiates a new GetPageByPath200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPageByPath200ResponseWithDefaults

`func NewGetPageByPath200ResponseWithDefaults() *GetPageByPath200Response`

NewGetPageByPath200ResponseWithDefaults instantiates a new GetPageByPath200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPageByPath200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPageByPath200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPageByPath200Response) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *GetPageByPath200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetPageByPath200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetPageByPath200Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMarkdown

`func (o *GetPageByPath200Response) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *GetPageByPath200Response) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *GetPageByPath200Response) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetDocument

`func (o *GetPageByPath200Response) GetDocument() JSONDocumentDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *GetPageByPath200Response) GetDocumentOk() (*JSONDocumentDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *GetPageByPath200Response) SetDocument(v JSONDocumentDocument)`

SetDocument sets Document field to given value.


### GetKind

`func (o *GetPageByPath200Response) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetPageByPath200Response) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetPageByPath200Response) SetKind(v string)`

SetKind sets Kind field to given value.


### GetType

`func (o *GetPageByPath200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPageByPath200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPageByPath200Response) SetType(v string)`

SetType sets Type field to given value.


### GetSlug

`func (o *GetPageByPath200Response) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GetPageByPath200Response) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GetPageByPath200Response) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPath

`func (o *GetPageByPath200Response) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetPageByPath200Response) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetPageByPath200Response) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *GetPageByPath200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPageByPath200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPageByPath200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPageByPath200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPages

`func (o *GetPageByPath200Response) GetPages() []RevisionPageDocumentAllOfPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetPageByPath200Response) GetPagesOk() (*[]RevisionPageDocumentAllOfPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetPageByPath200Response) SetPages(v []RevisionPageDocumentAllOfPagesInner)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


