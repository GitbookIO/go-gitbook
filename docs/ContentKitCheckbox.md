# ContentKitCheckbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**State** | **string** | State binding. The value of the input will be stored as a property in the state named after this ID. | 
**Value** | [**ContentKitCheckboxValue**](ContentKitCheckboxValue.md) |  | 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 

## Methods

### NewContentKitCheckbox

`func NewContentKitCheckbox(type_ string, state string, value ContentKitCheckboxValue, ) *ContentKitCheckbox`

NewContentKitCheckbox instantiates a new ContentKitCheckbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitCheckboxWithDefaults

`func NewContentKitCheckboxWithDefaults() *ContentKitCheckbox`

NewContentKitCheckboxWithDefaults instantiates a new ContentKitCheckbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitCheckbox) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitCheckbox) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitCheckbox) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *ContentKitCheckbox) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitCheckbox) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitCheckbox) SetState(v string)`

SetState sets State field to given value.


### GetValue

`func (o *ContentKitCheckbox) GetValue() ContentKitCheckboxValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentKitCheckbox) GetValueOk() (*ContentKitCheckboxValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentKitCheckbox) SetValue(v ContentKitCheckboxValue)`

SetValue sets Value field to given value.


### GetConfirm

`func (o *ContentKitCheckbox) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitCheckbox) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitCheckbox) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitCheckbox) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


