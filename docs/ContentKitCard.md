# ContentKitCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to [**ContentKitCardHint**](ContentKitCardHint.md) |  | [optional] 
**Icon** | Pointer to [**ContentKitCardIcon**](ContentKitCardIcon.md) |  | [optional] 
**OnPress** | Pointer to [**ContentKitAction**](ContentKitAction.md) |  | [optional] 
**Children** | Pointer to [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) |  | [optional] 
**Buttons** | Pointer to [**[]ContentKitButton**](ContentKitButton.md) | Buttons displayed in the top right corner of the card. | [optional] 

## Methods

### NewContentKitCard

`func NewContentKitCard(type_ string, ) *ContentKitCard`

NewContentKitCard instantiates a new ContentKitCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitCardWithDefaults

`func NewContentKitCardWithDefaults() *ContentKitCard`

NewContentKitCardWithDefaults instantiates a new ContentKitCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitCard) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *ContentKitCard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentKitCard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentKitCard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentKitCard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHint

`func (o *ContentKitCard) GetHint() ContentKitCardHint`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ContentKitCard) GetHintOk() (*ContentKitCardHint, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ContentKitCard) SetHint(v ContentKitCardHint)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ContentKitCard) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetIcon

`func (o *ContentKitCard) GetIcon() ContentKitCardIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ContentKitCard) GetIconOk() (*ContentKitCardIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ContentKitCard) SetIcon(v ContentKitCardIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ContentKitCard) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetOnPress

`func (o *ContentKitCard) GetOnPress() ContentKitAction`

GetOnPress returns the OnPress field if non-nil, zero value otherwise.

### GetOnPressOk

`func (o *ContentKitCard) GetOnPressOk() (*ContentKitAction, bool)`

GetOnPressOk returns a tuple with the OnPress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPress

`func (o *ContentKitCard) SetOnPress(v ContentKitAction)`

SetOnPress sets OnPress field to given value.

### HasOnPress

`func (o *ContentKitCard) HasOnPress() bool`

HasOnPress returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitCard) GetChildren() []ContentKitDescendantElement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitCard) GetChildrenOk() (*[]ContentKitDescendantElement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitCard) SetChildren(v []ContentKitDescendantElement)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ContentKitCard) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetButtons

`func (o *ContentKitCard) GetButtons() []ContentKitButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ContentKitCard) GetButtonsOk() (*[]ContentKitButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ContentKitCard) SetButtons(v []ContentKitButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ContentKitCard) HasButtons() bool`

HasButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


