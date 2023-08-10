# ContentKitButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Style** | Pointer to **string** |  | [optional] 
**OnPress** | [**ContentKitAction**](ContentKitAction.md) |  | 
**Icon** | Pointer to [**ContentKitIcon**](ContentKitIcon.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Tooltip** | Pointer to **string** |  | [optional] 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewContentKitButton

`func NewContentKitButton(type_ string, onPress ContentKitAction, ) *ContentKitButton`

NewContentKitButton instantiates a new ContentKitButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitButtonWithDefaults

`func NewContentKitButtonWithDefaults() *ContentKitButton`

NewContentKitButtonWithDefaults instantiates a new ContentKitButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitButton) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitButton) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitButton) SetType(v string)`

SetType sets Type field to given value.


### GetStyle

`func (o *ContentKitButton) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ContentKitButton) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ContentKitButton) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ContentKitButton) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetOnPress

`func (o *ContentKitButton) GetOnPress() ContentKitAction`

GetOnPress returns the OnPress field if non-nil, zero value otherwise.

### GetOnPressOk

`func (o *ContentKitButton) GetOnPressOk() (*ContentKitAction, bool)`

GetOnPressOk returns a tuple with the OnPress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPress

`func (o *ContentKitButton) SetOnPress(v ContentKitAction)`

SetOnPress sets OnPress field to given value.


### GetIcon

`func (o *ContentKitButton) GetIcon() ContentKitIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ContentKitButton) GetIconOk() (*ContentKitIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ContentKitButton) SetIcon(v ContentKitIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ContentKitButton) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLabel

`func (o *ContentKitButton) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentKitButton) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentKitButton) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ContentKitButton) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTooltip

`func (o *ContentKitButton) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ContentKitButton) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ContentKitButton) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *ContentKitButton) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetConfirm

`func (o *ContentKitButton) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitButton) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitButton) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitButton) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.

### GetDisabled

`func (o *ContentKitButton) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ContentKitButton) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ContentKitButton) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ContentKitButton) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


