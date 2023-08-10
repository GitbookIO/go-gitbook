# ContentKitModal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to [**[]ContentKitInlineElement**](ContentKitInlineElement.md) |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**ReturnValue** | Pointer to **map[string]interface{}** | Data passed back to the parent view when the modal is closed. These data are accessible in the \&quot;@ui.modal.close\&quot; | [optional] 
**Children** | [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) |  | 
**Submit** | Pointer to [**ContentKitButton**](ContentKitButton.md) |  | [optional] 

## Methods

### NewContentKitModal

`func NewContentKitModal(type_ string, children []ContentKitDescendantElement, ) *ContentKitModal`

NewContentKitModal instantiates a new ContentKitModal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitModalWithDefaults

`func NewContentKitModalWithDefaults() *ContentKitModal`

NewContentKitModalWithDefaults instantiates a new ContentKitModal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitModal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitModal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitModal) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *ContentKitModal) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentKitModal) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentKitModal) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentKitModal) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSubtitle

`func (o *ContentKitModal) GetSubtitle() []ContentKitInlineElement`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ContentKitModal) GetSubtitleOk() (*[]ContentKitInlineElement, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ContentKitModal) SetSubtitle(v []ContentKitInlineElement)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ContentKitModal) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetSize

`func (o *ContentKitModal) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentKitModal) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentKitModal) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentKitModal) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetReturnValue

`func (o *ContentKitModal) GetReturnValue() map[string]interface{}`

GetReturnValue returns the ReturnValue field if non-nil, zero value otherwise.

### GetReturnValueOk

`func (o *ContentKitModal) GetReturnValueOk() (*map[string]interface{}, bool)`

GetReturnValueOk returns a tuple with the ReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValue

`func (o *ContentKitModal) SetReturnValue(v map[string]interface{})`

SetReturnValue sets ReturnValue field to given value.

### HasReturnValue

`func (o *ContentKitModal) HasReturnValue() bool`

HasReturnValue returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitModal) GetChildren() []ContentKitDescendantElement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitModal) GetChildrenOk() (*[]ContentKitDescendantElement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitModal) SetChildren(v []ContentKitDescendantElement)`

SetChildren sets Children field to given value.


### GetSubmit

`func (o *ContentKitModal) GetSubmit() ContentKitButton`

GetSubmit returns the Submit field if non-nil, zero value otherwise.

### GetSubmitOk

`func (o *ContentKitModal) GetSubmitOk() (*ContentKitButton, bool)`

GetSubmitOk returns a tuple with the Submit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmit

`func (o *ContentKitModal) SetSubmit(v ContentKitButton)`

SetSubmit sets Submit field to given value.

### HasSubmit

`func (o *ContentKitModal) HasSubmit() bool`

HasSubmit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


