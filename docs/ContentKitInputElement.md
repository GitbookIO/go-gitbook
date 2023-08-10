# ContentKitInputElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**State** | **string** | State binding. The value of the input will be stored as a property in the state named after this ID. | 
**InitialValue** | Pointer to **bool** | Value to initialize the switch with. | [optional] 
**Placeholder** | Pointer to **string** | Text that appears in the form control when it has no value set | [optional] 
**Multiline** | Pointer to **bool** |  | [optional] 
**Multiple** | Pointer to **bool** | Should the select accept the selection of multiple options. If true, the state will be an array. | [optional] 
**Options** | [**ContentKitSelectOptions**](ContentKitSelectOptions.md) |  | 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 
**Value** | [**ContentKitCheckboxValue**](ContentKitCheckboxValue.md) |  | 
**Style** | Pointer to **string** |  | [optional] 
**OnPress** | [**ContentKitAction**](ContentKitAction.md) |  | 
**Icon** | Pointer to [**ContentKitIcon**](ContentKitIcon.md) |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Tooltip** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewContentKitInputElement

`func NewContentKitInputElement(type_ string, state string, options ContentKitSelectOptions, value ContentKitCheckboxValue, onPress ContentKitAction, ) *ContentKitInputElement`

NewContentKitInputElement instantiates a new ContentKitInputElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitInputElementWithDefaults

`func NewContentKitInputElementWithDefaults() *ContentKitInputElement`

NewContentKitInputElementWithDefaults instantiates a new ContentKitInputElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitInputElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitInputElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitInputElement) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *ContentKitInputElement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitInputElement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitInputElement) SetState(v string)`

SetState sets State field to given value.


### GetInitialValue

`func (o *ContentKitInputElement) GetInitialValue() bool`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *ContentKitInputElement) GetInitialValueOk() (*bool, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *ContentKitInputElement) SetInitialValue(v bool)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *ContentKitInputElement) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ContentKitInputElement) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ContentKitInputElement) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ContentKitInputElement) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ContentKitInputElement) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMultiline

`func (o *ContentKitInputElement) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *ContentKitInputElement) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *ContentKitInputElement) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *ContentKitInputElement) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetMultiple

`func (o *ContentKitInputElement) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *ContentKitInputElement) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *ContentKitInputElement) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *ContentKitInputElement) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetOptions

`func (o *ContentKitInputElement) GetOptions() ContentKitSelectOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ContentKitInputElement) GetOptionsOk() (*ContentKitSelectOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ContentKitInputElement) SetOptions(v ContentKitSelectOptions)`

SetOptions sets Options field to given value.


### GetConfirm

`func (o *ContentKitInputElement) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitInputElement) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitInputElement) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitInputElement) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.

### GetValue

`func (o *ContentKitInputElement) GetValue() ContentKitCheckboxValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentKitInputElement) GetValueOk() (*ContentKitCheckboxValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentKitInputElement) SetValue(v ContentKitCheckboxValue)`

SetValue sets Value field to given value.


### GetStyle

`func (o *ContentKitInputElement) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ContentKitInputElement) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ContentKitInputElement) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ContentKitInputElement) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetOnPress

`func (o *ContentKitInputElement) GetOnPress() ContentKitAction`

GetOnPress returns the OnPress field if non-nil, zero value otherwise.

### GetOnPressOk

`func (o *ContentKitInputElement) GetOnPressOk() (*ContentKitAction, bool)`

GetOnPressOk returns a tuple with the OnPress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPress

`func (o *ContentKitInputElement) SetOnPress(v ContentKitAction)`

SetOnPress sets OnPress field to given value.


### GetIcon

`func (o *ContentKitInputElement) GetIcon() ContentKitIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ContentKitInputElement) GetIconOk() (*ContentKitIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ContentKitInputElement) SetIcon(v ContentKitIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ContentKitInputElement) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLabel

`func (o *ContentKitInputElement) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentKitInputElement) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentKitInputElement) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ContentKitInputElement) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTooltip

`func (o *ContentKitInputElement) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ContentKitInputElement) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ContentKitInputElement) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *ContentKitInputElement) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetDisabled

`func (o *ContentKitInputElement) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ContentKitInputElement) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ContentKitInputElement) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ContentKitInputElement) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


