# ContentKitRadio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**State** | **string** | State binding. The value of the input will be stored as a property in the state named after this ID. | 
**Value** | [**ContentKitRadioValue**](ContentKitRadioValue.md) |  | 
**Confirm** | Pointer to [**ContentKitConfirm**](ContentKitConfirm.md) |  | [optional] 

## Methods

### NewContentKitRadio

`func NewContentKitRadio(type_ string, state string, value ContentKitRadioValue, ) *ContentKitRadio`

NewContentKitRadio instantiates a new ContentKitRadio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitRadioWithDefaults

`func NewContentKitRadioWithDefaults() *ContentKitRadio`

NewContentKitRadioWithDefaults instantiates a new ContentKitRadio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitRadio) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitRadio) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitRadio) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *ContentKitRadio) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitRadio) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitRadio) SetState(v string)`

SetState sets State field to given value.


### GetValue

`func (o *ContentKitRadio) GetValue() ContentKitRadioValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentKitRadio) GetValueOk() (*ContentKitRadioValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentKitRadio) SetValue(v ContentKitRadioValue)`

SetValue sets Value field to given value.


### GetConfirm

`func (o *ContentKitRadio) GetConfirm() ContentKitConfirm`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *ContentKitRadio) GetConfirmOk() (*ContentKitConfirm, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *ContentKitRadio) SetConfirm(v ContentKitConfirm)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *ContentKitRadio) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


