# RevisionPageDocumentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Type** | **string** |  | 
**Slug** | **string** | Page&#39;s slug in its direct parent | 
**Path** | **string** | Complete path to access the page in the revision. | 
**Description** | Pointer to **string** |  | [optional] 
**Pages** | [**[]RevisionPageDocumentAllOfPagesInner**](RevisionPageDocumentAllOfPagesInner.md) |  | 

## Methods

### NewRevisionPageDocumentAllOf

`func NewRevisionPageDocumentAllOf(kind string, type_ string, slug string, path string, pages []RevisionPageDocumentAllOfPagesInner, ) *RevisionPageDocumentAllOf`

NewRevisionPageDocumentAllOf instantiates a new RevisionPageDocumentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionPageDocumentAllOfWithDefaults

`func NewRevisionPageDocumentAllOfWithDefaults() *RevisionPageDocumentAllOf`

NewRevisionPageDocumentAllOfWithDefaults instantiates a new RevisionPageDocumentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RevisionPageDocumentAllOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevisionPageDocumentAllOf) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevisionPageDocumentAllOf) SetKind(v string)`

SetKind sets Kind field to given value.


### GetType

`func (o *RevisionPageDocumentAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevisionPageDocumentAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevisionPageDocumentAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetSlug

`func (o *RevisionPageDocumentAllOf) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RevisionPageDocumentAllOf) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RevisionPageDocumentAllOf) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPath

`func (o *RevisionPageDocumentAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RevisionPageDocumentAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RevisionPageDocumentAllOf) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *RevisionPageDocumentAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RevisionPageDocumentAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RevisionPageDocumentAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RevisionPageDocumentAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPages

`func (o *RevisionPageDocumentAllOf) GetPages() []RevisionPageDocumentAllOfPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RevisionPageDocumentAllOf) GetPagesOk() (*[]RevisionPageDocumentAllOfPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RevisionPageDocumentAllOf) SetPages(v []RevisionPageDocumentAllOfPagesInner)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


