# ContentKitRootElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Children** | [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) |  | 
**Controls** | Pointer to [**[]ContentKitBlockControlsInner**](ContentKitBlockControlsInner.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Subtitle** | Pointer to [**[]ContentKitInlineElement**](ContentKitInlineElement.md) |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**ReturnValue** | Pointer to **map[string]interface{}** | Data passed back to the parent view when the modal is closed. These data are accessible in the \&quot;@ui.modal.close\&quot; | [optional] 
**Submit** | Pointer to [**ContentKitButton**](ContentKitButton.md) |  | [optional] 

## Methods

### NewContentKitRootElement

`func NewContentKitRootElement(type_ string, children []ContentKitDescendantElement, ) *ContentKitRootElement`

NewContentKitRootElement instantiates a new ContentKitRootElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitRootElementWithDefaults

`func NewContentKitRootElementWithDefaults() *ContentKitRootElement`

NewContentKitRootElementWithDefaults instantiates a new ContentKitRootElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitRootElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitRootElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitRootElement) SetType(v string)`

SetType sets Type field to given value.


### GetChildren

`func (o *ContentKitRootElement) GetChildren() []ContentKitDescendantElement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitRootElement) GetChildrenOk() (*[]ContentKitDescendantElement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitRootElement) SetChildren(v []ContentKitDescendantElement)`

SetChildren sets Children field to given value.


### GetControls

`func (o *ContentKitRootElement) GetControls() []ContentKitBlockControlsInner`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *ContentKitRootElement) GetControlsOk() (*[]ContentKitBlockControlsInner, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *ContentKitRootElement) SetControls(v []ContentKitBlockControlsInner)`

SetControls sets Controls field to given value.

### HasControls

`func (o *ContentKitRootElement) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetTitle

`func (o *ContentKitRootElement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentKitRootElement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentKitRootElement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentKitRootElement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSubtitle

`func (o *ContentKitRootElement) GetSubtitle() []ContentKitInlineElement`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ContentKitRootElement) GetSubtitleOk() (*[]ContentKitInlineElement, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ContentKitRootElement) SetSubtitle(v []ContentKitInlineElement)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ContentKitRootElement) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetSize

`func (o *ContentKitRootElement) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentKitRootElement) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentKitRootElement) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentKitRootElement) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetReturnValue

`func (o *ContentKitRootElement) GetReturnValue() map[string]interface{}`

GetReturnValue returns the ReturnValue field if non-nil, zero value otherwise.

### GetReturnValueOk

`func (o *ContentKitRootElement) GetReturnValueOk() (*map[string]interface{}, bool)`

GetReturnValueOk returns a tuple with the ReturnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnValue

`func (o *ContentKitRootElement) SetReturnValue(v map[string]interface{})`

SetReturnValue sets ReturnValue field to given value.

### HasReturnValue

`func (o *ContentKitRootElement) HasReturnValue() bool`

HasReturnValue returns a boolean if a field has been set.

### GetSubmit

`func (o *ContentKitRootElement) GetSubmit() ContentKitButton`

GetSubmit returns the Submit field if non-nil, zero value otherwise.

### GetSubmitOk

`func (o *ContentKitRootElement) GetSubmitOk() (*ContentKitButton, bool)`

GetSubmitOk returns a tuple with the Submit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmit

`func (o *ContentKitRootElement) SetSubmit(v ContentKitButton)`

SetSubmit sets Submit field to given value.

### HasSubmit

`func (o *ContentKitRootElement) HasSubmit() bool`

HasSubmit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


