# IntegrationBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID in the integration for the block. It also represents the UI component used. | 
**Title** | **string** | Short descriptive title for the block. | 
**Description** | Pointer to **string** | Long descriptive text for the block. | [optional] 
**Icon** | Pointer to **string** | URL of the icon to represent this block. | [optional] 
**UrlUnfurl** | Pointer to **[]string** | URLs patterns to convert as this block. | [optional] 
**Markdown** | Pointer to [**IntegrationBlockMarkdown**](IntegrationBlockMarkdown.md) |  | [optional] 

## Methods

### NewIntegrationBlock

`func NewIntegrationBlock(id string, title string, ) *IntegrationBlock`

NewIntegrationBlock instantiates a new IntegrationBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationBlockWithDefaults

`func NewIntegrationBlockWithDefaults() *IntegrationBlock`

NewIntegrationBlockWithDefaults instantiates a new IntegrationBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationBlock) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *IntegrationBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IntegrationBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IntegrationBlock) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *IntegrationBlock) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationBlock) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationBlock) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationBlock) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *IntegrationBlock) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IntegrationBlock) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IntegrationBlock) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IntegrationBlock) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetUrlUnfurl

`func (o *IntegrationBlock) GetUrlUnfurl() []string`

GetUrlUnfurl returns the UrlUnfurl field if non-nil, zero value otherwise.

### GetUrlUnfurlOk

`func (o *IntegrationBlock) GetUrlUnfurlOk() (*[]string, bool)`

GetUrlUnfurlOk returns a tuple with the UrlUnfurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlUnfurl

`func (o *IntegrationBlock) SetUrlUnfurl(v []string)`

SetUrlUnfurl sets UrlUnfurl field to given value.

### HasUrlUnfurl

`func (o *IntegrationBlock) HasUrlUnfurl() bool`

HasUrlUnfurl returns a boolean if a field has been set.

### GetMarkdown

`func (o *IntegrationBlock) GetMarkdown() IntegrationBlockMarkdown`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *IntegrationBlock) GetMarkdownOk() (*IntegrationBlockMarkdown, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *IntegrationBlock) SetMarkdown(v IntegrationBlockMarkdown)`

SetMarkdown sets Markdown field to given value.

### HasMarkdown

`func (o *IntegrationBlock) HasMarkdown() bool`

HasMarkdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


