# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Markdown** | **string** | Content of the document formatted as markdown | 
**Document** | [**JSONDocumentDocument**](JSONDocumentDocument.md) |  | 

## Methods

### NewDocument

`func NewDocument(markdown string, document JSONDocumentDocument, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkdown

`func (o *Document) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *Document) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *Document) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetDocument

`func (o *Document) GetDocument() JSONDocumentDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *Document) GetDocumentOk() (*JSONDocumentDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *Document) SetDocument(v JSONDocumentDocument)`

SetDocument sets Document field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


