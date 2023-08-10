# ContentKitBlockControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to [**ContentKitIcon**](ContentKitIcon.md) |  | [optional] 
**Label** | **string** |  | 
**OnPress** | [**ContentKitAction**](ContentKitAction.md) |  | 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 

## Methods

### NewContentKitBlockControl

`func NewContentKitBlockControl(label string, onPress ContentKitAction, ) *ContentKitBlockControl`

NewContentKitBlockControl instantiates a new ContentKitBlockControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitBlockControlWithDefaults

`func NewContentKitBlockControlWithDefaults() *ContentKitBlockControl`

NewContentKitBlockControlWithDefaults instantiates a new ContentKitBlockControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *ContentKitBlockControl) GetIcon() ContentKitIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ContentKitBlockControl) GetIconOk() (*ContentKitIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ContentKitBlockControl) SetIcon(v ContentKitIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ContentKitBlockControl) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLabel

`func (o *ContentKitBlockControl) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentKitBlockControl) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentKitBlockControl) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOnPress

`func (o *ContentKitBlockControl) GetOnPress() ContentKitAction`

GetOnPress returns the OnPress field if non-nil, zero value otherwise.

### GetOnPressOk

`func (o *ContentKitBlockControl) GetOnPressOk() (*ContentKitAction, bool)`

GetOnPressOk returns a tuple with the OnPress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPress

`func (o *ContentKitBlockControl) SetOnPress(v ContentKitAction)`

SetOnPress sets OnPress field to given value.


### GetConfirm

`func (o *ContentKitBlockControl) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitBlockControl) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitBlockControl) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitBlockControl) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


