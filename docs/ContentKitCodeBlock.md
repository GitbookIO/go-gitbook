# ContentKitCodeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Content** | [**ContentKitCodeBlockContent**](ContentKitCodeBlockContent.md) |  | 
**Syntax** | Pointer to **string** | Syntax to use for highlighting (ex: javascript, python) | [optional] 
**LineNumbers** | Pointer to [**ContentKitCodeBlockLineNumbers**](ContentKitCodeBlockLineNumbers.md) |  | [optional] 
**Buttons** | Pointer to [**[]ContentKitButton**](ContentKitButton.md) | Controls button shown as an overlay in a corner of the code block. | [optional] 
**State** | Pointer to **string** | State binding when editable. The value of the input will be stored as a property in the state named after this ID. | [optional] 
**OnContentChange** | Pointer to [**ContentKitAction**](ContentKitAction.md) |  | [optional] 
**Header** | Pointer to [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) | Header displayed before the code lines | [optional] 
**Footer** | Pointer to [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) | Footer displayed after the code lines | [optional] 

## Methods

### NewContentKitCodeBlock

`func NewContentKitCodeBlock(type_ string, content ContentKitCodeBlockContent, ) *ContentKitCodeBlock`

NewContentKitCodeBlock instantiates a new ContentKitCodeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitCodeBlockWithDefaults

`func NewContentKitCodeBlockWithDefaults() *ContentKitCodeBlock`

NewContentKitCodeBlockWithDefaults instantiates a new ContentKitCodeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitCodeBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitCodeBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitCodeBlock) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *ContentKitCodeBlock) GetContent() ContentKitCodeBlockContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentKitCodeBlock) GetContentOk() (*ContentKitCodeBlockContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentKitCodeBlock) SetContent(v ContentKitCodeBlockContent)`

SetContent sets Content field to given value.


### GetSyntax

`func (o *ContentKitCodeBlock) GetSyntax() string`

GetSyntax returns the Syntax field if non-nil, zero value otherwise.

### GetSyntaxOk

`func (o *ContentKitCodeBlock) GetSyntaxOk() (*string, bool)`

GetSyntaxOk returns a tuple with the Syntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntax

`func (o *ContentKitCodeBlock) SetSyntax(v string)`

SetSyntax sets Syntax field to given value.

### HasSyntax

`func (o *ContentKitCodeBlock) HasSyntax() bool`

HasSyntax returns a boolean if a field has been set.

### GetLineNumbers

`func (o *ContentKitCodeBlock) GetLineNumbers() ContentKitCodeBlockLineNumbers`

GetLineNumbers returns the LineNumbers field if non-nil, zero value otherwise.

### GetLineNumbersOk

`func (o *ContentKitCodeBlock) GetLineNumbersOk() (*ContentKitCodeBlockLineNumbers, bool)`

GetLineNumbersOk returns a tuple with the LineNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumbers

`func (o *ContentKitCodeBlock) SetLineNumbers(v ContentKitCodeBlockLineNumbers)`

SetLineNumbers sets LineNumbers field to given value.

### HasLineNumbers

`func (o *ContentKitCodeBlock) HasLineNumbers() bool`

HasLineNumbers returns a boolean if a field has been set.

### GetButtons

`func (o *ContentKitCodeBlock) GetButtons() []ContentKitButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ContentKitCodeBlock) GetButtonsOk() (*[]ContentKitButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ContentKitCodeBlock) SetButtons(v []ContentKitButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ContentKitCodeBlock) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetState

`func (o *ContentKitCodeBlock) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitCodeBlock) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitCodeBlock) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ContentKitCodeBlock) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOnContentChange

`func (o *ContentKitCodeBlock) GetOnContentChange() ContentKitAction`

GetOnContentChange returns the OnContentChange field if non-nil, zero value otherwise.

### GetOnContentChangeOk

`func (o *ContentKitCodeBlock) GetOnContentChangeOk() (*ContentKitAction, bool)`

GetOnContentChangeOk returns a tuple with the OnContentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnContentChange

`func (o *ContentKitCodeBlock) SetOnContentChange(v ContentKitAction)`

SetOnContentChange sets OnContentChange field to given value.

### HasOnContentChange

`func (o *ContentKitCodeBlock) HasOnContentChange() bool`

HasOnContentChange returns a boolean if a field has been set.

### GetHeader

`func (o *ContentKitCodeBlock) GetHeader() []ContentKitDescendantElement`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ContentKitCodeBlock) GetHeaderOk() (*[]ContentKitDescendantElement, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ContentKitCodeBlock) SetHeader(v []ContentKitDescendantElement)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ContentKitCodeBlock) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetFooter

`func (o *ContentKitCodeBlock) GetFooter() []ContentKitDescendantElement`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *ContentKitCodeBlock) GetFooterOk() (*[]ContentKitDescendantElement, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *ContentKitCodeBlock) SetFooter(v []ContentKitDescendantElement)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *ContentKitCodeBlock) HasFooter() bool`

HasFooter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


