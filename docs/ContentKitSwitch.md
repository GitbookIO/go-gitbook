# ContentKitSwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**State** | **string** | State binding. The value of the input will be stored as a property in the state named after this ID. | 
**InitialValue** | Pointer to **bool** | Value to initialize the switch with. | [optional] 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 

## Methods

### NewContentKitSwitch

`func NewContentKitSwitch(type_ string, state string, ) *ContentKitSwitch`

NewContentKitSwitch instantiates a new ContentKitSwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitSwitchWithDefaults

`func NewContentKitSwitchWithDefaults() *ContentKitSwitch`

NewContentKitSwitchWithDefaults instantiates a new ContentKitSwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitSwitch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitSwitch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitSwitch) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *ContentKitSwitch) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitSwitch) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitSwitch) SetState(v string)`

SetState sets State field to given value.


### GetInitialValue

`func (o *ContentKitSwitch) GetInitialValue() bool`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *ContentKitSwitch) GetInitialValueOk() (*bool, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *ContentKitSwitch) SetInitialValue(v bool)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *ContentKitSwitch) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetConfirm

`func (o *ContentKitSwitch) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitSwitch) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitSwitch) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitSwitch) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


