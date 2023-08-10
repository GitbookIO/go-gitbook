# ContentKitBlockControlsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to [**ContentKitIcon**](ContentKitIcon.md) |  | [optional] 
**Label** | **string** |  | 
**OnPress** | [**ContentKitAction**](ContentKitAction.md) |  | 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 

## Methods

### NewContentKitBlockControlsInner

`func NewContentKitBlockControlsInner(label string, onPress ContentKitAction, ) *ContentKitBlockControlsInner`

NewContentKitBlockControlsInner instantiates a new ContentKitBlockControlsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitBlockControlsInnerWithDefaults

`func NewContentKitBlockControlsInnerWithDefaults() *ContentKitBlockControlsInner`

NewContentKitBlockControlsInnerWithDefaults instantiates a new ContentKitBlockControlsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *ContentKitBlockControlsInner) GetIcon() ContentKitIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ContentKitBlockControlsInner) GetIconOk() (*ContentKitIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ContentKitBlockControlsInner) SetIcon(v ContentKitIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ContentKitBlockControlsInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLabel

`func (o *ContentKitBlockControlsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentKitBlockControlsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentKitBlockControlsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOnPress

`func (o *ContentKitBlockControlsInner) GetOnPress() ContentKitAction`

GetOnPress returns the OnPress field if non-nil, zero value otherwise.

### GetOnPressOk

`func (o *ContentKitBlockControlsInner) GetOnPressOk() (*ContentKitAction, bool)`

GetOnPressOk returns a tuple with the OnPress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPress

`func (o *ContentKitBlockControlsInner) SetOnPress(v ContentKitAction)`

SetOnPress sets OnPress field to given value.


### GetConfirm

`func (o *ContentKitBlockControlsInner) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitBlockControlsInner) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitBlockControlsInner) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitBlockControlsInner) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


