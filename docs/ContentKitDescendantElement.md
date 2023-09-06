# ContentKitDescendantElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Style** | Pointer to **string** |  | [optional] 
**OnPress** | [**ContentKitAction**](ContentKitAction.md) |  | 
**Icon** | Pointer to [**ContentKitCardIcon**](ContentKitCardIcon.md) |  | [optional] 
**TrailingIcon** | Pointer to [**ContentKitIcon**](ContentKitIcon.md) |  | [optional] 
**Label** | **string** | Text label displayed next to the input. | 
**Tooltip** | Pointer to **string** |  | [optional] 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**State** | **string** | State binding. The value of the input will be stored as a property in the state named after this ID. | 
**InitialValue** | Pointer to **bool** | Value to initialize the switch with. | [optional] 
**Placeholder** | Pointer to **string** | Text that appears in the form control when it has no value set | [optional] 
**Multiline** | Pointer to **bool** |  | [optional] 
**Align** | Pointer to **string** |  | [optional] [default to "start"]
**Children** | [**ContentKitLinkChildren**](ContentKitLinkChildren.md) |  | 
**Grow** | Pointer to **float32** | specifies how much of the remaining space in the container should be assigned to the element | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**AspectRatio** | **float32** |  | 
**Source** | [**ContentKitImageSource**](ContentKitImageSource.md) |  | 
**Buttons** | Pointer to [**[]ContentKitButton**](ContentKitButton.md) | Buttons displayed in the top right corner of the card. | [optional] 
**Data** | Pointer to [**map[string]ContentKitWebFrameDataValue**](ContentKitWebFrameDataValue.md) | Data to communicated to the webframe&#39;s content. Each state update will cause the webframe to receive a message. | [optional] 
**Content** | [**ContentKitMarkdownContent**](ContentKitMarkdownContent.md) |  | 
**Syntax** | Pointer to **string** | Syntax to use for highlighting (ex: javascript, python) | [optional] 
**LineNumbers** | Pointer to [**ContentKitCodeBlockLineNumbers**](ContentKitCodeBlockLineNumbers.md) |  | [optional] 
**OnContentChange** | Pointer to [**ContentKitAction**](ContentKitAction.md) |  | [optional] 
**Header** | Pointer to [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) | Header displayed before the code lines | [optional] 
**Footer** | Pointer to [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) | Footer displayed after the code lines | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Hint** | Pointer to [**ContentKitInputHint**](ContentKitInputHint.md) |  | [optional] 
**Element** | [**ContentKitInputElement**](ContentKitInputElement.md) |  | 
**OnValueChange** | Pointer to [**ContentKitAction**](ContentKitAction.md) |  | [optional] 
**Multiple** | Pointer to **bool** | Should the select accept the selection of multiple options. If true, the state will be an array. | [optional] 
**AcceptInput** | Pointer to **bool** | Should the filter input be allowed to be selected as an option. | [optional] 
**Options** | [**ContentKitSelectOptions**](ContentKitSelectOptions.md) |  | 
**Value** | [**ContentKitRadioValue**](ContentKitRadioValue.md) |  | 
**Target** | [**ContentKitLinkTarget**](ContentKitLinkTarget.md) |  | 

## Methods

### NewContentKitDescendantElement

`func NewContentKitDescendantElement(type_ string, onPress ContentKitAction, label string, state string, children ContentKitLinkChildren, aspectRatio float32, source ContentKitImageSource, content ContentKitMarkdownContent, element ContentKitInputElement, options ContentKitSelectOptions, value ContentKitRadioValue, target ContentKitLinkTarget, ) *ContentKitDescendantElement`

NewContentKitDescendantElement instantiates a new ContentKitDescendantElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitDescendantElementWithDefaults

`func NewContentKitDescendantElementWithDefaults() *ContentKitDescendantElement`

NewContentKitDescendantElementWithDefaults instantiates a new ContentKitDescendantElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitDescendantElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitDescendantElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitDescendantElement) SetType(v string)`

SetType sets Type field to given value.


### GetStyle

`func (o *ContentKitDescendantElement) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ContentKitDescendantElement) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ContentKitDescendantElement) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ContentKitDescendantElement) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetOnPress

`func (o *ContentKitDescendantElement) GetOnPress() ContentKitAction`

GetOnPress returns the OnPress field if non-nil, zero value otherwise.

### GetOnPressOk

`func (o *ContentKitDescendantElement) GetOnPressOk() (*ContentKitAction, bool)`

GetOnPressOk returns a tuple with the OnPress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPress

`func (o *ContentKitDescendantElement) SetOnPress(v ContentKitAction)`

SetOnPress sets OnPress field to given value.


### GetIcon

`func (o *ContentKitDescendantElement) GetIcon() ContentKitCardIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ContentKitDescendantElement) GetIconOk() (*ContentKitCardIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ContentKitDescendantElement) SetIcon(v ContentKitCardIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ContentKitDescendantElement) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetTrailingIcon

`func (o *ContentKitDescendantElement) GetTrailingIcon() ContentKitIcon`

GetTrailingIcon returns the TrailingIcon field if non-nil, zero value otherwise.

### GetTrailingIconOk

`func (o *ContentKitDescendantElement) GetTrailingIconOk() (*ContentKitIcon, bool)`

GetTrailingIconOk returns a tuple with the TrailingIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingIcon

`func (o *ContentKitDescendantElement) SetTrailingIcon(v ContentKitIcon)`

SetTrailingIcon sets TrailingIcon field to given value.

### HasTrailingIcon

`func (o *ContentKitDescendantElement) HasTrailingIcon() bool`

HasTrailingIcon returns a boolean if a field has been set.

### GetLabel

`func (o *ContentKitDescendantElement) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentKitDescendantElement) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentKitDescendantElement) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTooltip

`func (o *ContentKitDescendantElement) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ContentKitDescendantElement) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ContentKitDescendantElement) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *ContentKitDescendantElement) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetConfirm

`func (o *ContentKitDescendantElement) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitDescendantElement) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitDescendantElement) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitDescendantElement) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.

### GetDisabled

`func (o *ContentKitDescendantElement) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ContentKitDescendantElement) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ContentKitDescendantElement) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ContentKitDescendantElement) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetState

`func (o *ContentKitDescendantElement) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitDescendantElement) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitDescendantElement) SetState(v string)`

SetState sets State field to given value.


### GetInitialValue

`func (o *ContentKitDescendantElement) GetInitialValue() bool`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *ContentKitDescendantElement) GetInitialValueOk() (*bool, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *ContentKitDescendantElement) SetInitialValue(v bool)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *ContentKitDescendantElement) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ContentKitDescendantElement) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ContentKitDescendantElement) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ContentKitDescendantElement) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ContentKitDescendantElement) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMultiline

`func (o *ContentKitDescendantElement) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *ContentKitDescendantElement) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *ContentKitDescendantElement) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *ContentKitDescendantElement) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetAlign

`func (o *ContentKitDescendantElement) GetAlign() string`

GetAlign returns the Align field if non-nil, zero value otherwise.

### GetAlignOk

`func (o *ContentKitDescendantElement) GetAlignOk() (*string, bool)`

GetAlignOk returns a tuple with the Align field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlign

`func (o *ContentKitDescendantElement) SetAlign(v string)`

SetAlign sets Align field to given value.

### HasAlign

`func (o *ContentKitDescendantElement) HasAlign() bool`

HasAlign returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitDescendantElement) GetChildren() ContentKitLinkChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitDescendantElement) GetChildrenOk() (*ContentKitLinkChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitDescendantElement) SetChildren(v ContentKitLinkChildren)`

SetChildren sets Children field to given value.


### GetGrow

`func (o *ContentKitDescendantElement) GetGrow() float32`

GetGrow returns the Grow field if non-nil, zero value otherwise.

### GetGrowOk

`func (o *ContentKitDescendantElement) GetGrowOk() (*float32, bool)`

GetGrowOk returns a tuple with the Grow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrow

`func (o *ContentKitDescendantElement) SetGrow(v float32)`

SetGrow sets Grow field to given value.

### HasGrow

`func (o *ContentKitDescendantElement) HasGrow() bool`

HasGrow returns a boolean if a field has been set.

### GetSize

`func (o *ContentKitDescendantElement) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentKitDescendantElement) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentKitDescendantElement) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentKitDescendantElement) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetAspectRatio

`func (o *ContentKitDescendantElement) GetAspectRatio() float32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *ContentKitDescendantElement) GetAspectRatioOk() (*float32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *ContentKitDescendantElement) SetAspectRatio(v float32)`

SetAspectRatio sets AspectRatio field to given value.


### GetSource

`func (o *ContentKitDescendantElement) GetSource() ContentKitImageSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContentKitDescendantElement) GetSourceOk() (*ContentKitImageSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContentKitDescendantElement) SetSource(v ContentKitImageSource)`

SetSource sets Source field to given value.


### GetButtons

`func (o *ContentKitDescendantElement) GetButtons() []ContentKitButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ContentKitDescendantElement) GetButtonsOk() (*[]ContentKitButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ContentKitDescendantElement) SetButtons(v []ContentKitButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ContentKitDescendantElement) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetData

`func (o *ContentKitDescendantElement) GetData() map[string]ContentKitWebFrameDataValue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentKitDescendantElement) GetDataOk() (*map[string]ContentKitWebFrameDataValue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentKitDescendantElement) SetData(v map[string]ContentKitWebFrameDataValue)`

SetData sets Data field to given value.

### HasData

`func (o *ContentKitDescendantElement) HasData() bool`

HasData returns a boolean if a field has been set.

### GetContent

`func (o *ContentKitDescendantElement) GetContent() ContentKitMarkdownContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentKitDescendantElement) GetContentOk() (*ContentKitMarkdownContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentKitDescendantElement) SetContent(v ContentKitMarkdownContent)`

SetContent sets Content field to given value.


### GetSyntax

`func (o *ContentKitDescendantElement) GetSyntax() string`

GetSyntax returns the Syntax field if non-nil, zero value otherwise.

### GetSyntaxOk

`func (o *ContentKitDescendantElement) GetSyntaxOk() (*string, bool)`

GetSyntaxOk returns a tuple with the Syntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntax

`func (o *ContentKitDescendantElement) SetSyntax(v string)`

SetSyntax sets Syntax field to given value.

### HasSyntax

`func (o *ContentKitDescendantElement) HasSyntax() bool`

HasSyntax returns a boolean if a field has been set.

### GetLineNumbers

`func (o *ContentKitDescendantElement) GetLineNumbers() ContentKitCodeBlockLineNumbers`

GetLineNumbers returns the LineNumbers field if non-nil, zero value otherwise.

### GetLineNumbersOk

`func (o *ContentKitDescendantElement) GetLineNumbersOk() (*ContentKitCodeBlockLineNumbers, bool)`

GetLineNumbersOk returns a tuple with the LineNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumbers

`func (o *ContentKitDescendantElement) SetLineNumbers(v ContentKitCodeBlockLineNumbers)`

SetLineNumbers sets LineNumbers field to given value.

### HasLineNumbers

`func (o *ContentKitDescendantElement) HasLineNumbers() bool`

HasLineNumbers returns a boolean if a field has been set.

### GetOnContentChange

`func (o *ContentKitDescendantElement) GetOnContentChange() ContentKitAction`

GetOnContentChange returns the OnContentChange field if non-nil, zero value otherwise.

### GetOnContentChangeOk

`func (o *ContentKitDescendantElement) GetOnContentChangeOk() (*ContentKitAction, bool)`

GetOnContentChangeOk returns a tuple with the OnContentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnContentChange

`func (o *ContentKitDescendantElement) SetOnContentChange(v ContentKitAction)`

SetOnContentChange sets OnContentChange field to given value.

### HasOnContentChange

`func (o *ContentKitDescendantElement) HasOnContentChange() bool`

HasOnContentChange returns a boolean if a field has been set.

### GetHeader

`func (o *ContentKitDescendantElement) GetHeader() []ContentKitDescendantElement`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *ContentKitDescendantElement) GetHeaderOk() (*[]ContentKitDescendantElement, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *ContentKitDescendantElement) SetHeader(v []ContentKitDescendantElement)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *ContentKitDescendantElement) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetFooter

`func (o *ContentKitDescendantElement) GetFooter() []ContentKitDescendantElement`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *ContentKitDescendantElement) GetFooterOk() (*[]ContentKitDescendantElement, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *ContentKitDescendantElement) SetFooter(v []ContentKitDescendantElement)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *ContentKitDescendantElement) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetTitle

`func (o *ContentKitDescendantElement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContentKitDescendantElement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContentKitDescendantElement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContentKitDescendantElement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHint

`func (o *ContentKitDescendantElement) GetHint() ContentKitInputHint`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ContentKitDescendantElement) GetHintOk() (*ContentKitInputHint, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ContentKitDescendantElement) SetHint(v ContentKitInputHint)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ContentKitDescendantElement) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetElement

`func (o *ContentKitDescendantElement) GetElement() ContentKitInputElement`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ContentKitDescendantElement) GetElementOk() (*ContentKitInputElement, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ContentKitDescendantElement) SetElement(v ContentKitInputElement)`

SetElement sets Element field to given value.


### GetOnValueChange

`func (o *ContentKitDescendantElement) GetOnValueChange() ContentKitAction`

GetOnValueChange returns the OnValueChange field if non-nil, zero value otherwise.

### GetOnValueChangeOk

`func (o *ContentKitDescendantElement) GetOnValueChangeOk() (*ContentKitAction, bool)`

GetOnValueChangeOk returns a tuple with the OnValueChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnValueChange

`func (o *ContentKitDescendantElement) SetOnValueChange(v ContentKitAction)`

SetOnValueChange sets OnValueChange field to given value.

### HasOnValueChange

`func (o *ContentKitDescendantElement) HasOnValueChange() bool`

HasOnValueChange returns a boolean if a field has been set.

### GetMultiple

`func (o *ContentKitDescendantElement) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *ContentKitDescendantElement) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *ContentKitDescendantElement) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *ContentKitDescendantElement) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetAcceptInput

`func (o *ContentKitDescendantElement) GetAcceptInput() bool`

GetAcceptInput returns the AcceptInput field if non-nil, zero value otherwise.

### GetAcceptInputOk

`func (o *ContentKitDescendantElement) GetAcceptInputOk() (*bool, bool)`

GetAcceptInputOk returns a tuple with the AcceptInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptInput

`func (o *ContentKitDescendantElement) SetAcceptInput(v bool)`

SetAcceptInput sets AcceptInput field to given value.

### HasAcceptInput

`func (o *ContentKitDescendantElement) HasAcceptInput() bool`

HasAcceptInput returns a boolean if a field has been set.

### GetOptions

`func (o *ContentKitDescendantElement) GetOptions() ContentKitSelectOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ContentKitDescendantElement) GetOptionsOk() (*ContentKitSelectOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ContentKitDescendantElement) SetOptions(v ContentKitSelectOptions)`

SetOptions sets Options field to given value.


### GetValue

`func (o *ContentKitDescendantElement) GetValue() ContentKitRadioValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentKitDescendantElement) GetValueOk() (*ContentKitRadioValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentKitDescendantElement) SetValue(v ContentKitRadioValue)`

SetValue sets Value field to given value.


### GetTarget

`func (o *ContentKitDescendantElement) GetTarget() ContentKitLinkTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContentKitDescendantElement) GetTargetOk() (*ContentKitLinkTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContentKitDescendantElement) SetTarget(v ContentKitLinkTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


