# ContentKitInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Label** | **string** | Text label displayed next to the input. | 
**Hint** | Pointer to [**ContentKitInputHint**](ContentKitInputHint.md) |  | [optional] 
**Element** | [**ContentKitInputElement**](ContentKitInputElement.md) |  | 

## Methods

### NewContentKitInput

`func NewContentKitInput(type_ string, label string, element ContentKitInputElement, ) *ContentKitInput`

NewContentKitInput instantiates a new ContentKitInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitInputWithDefaults

`func NewContentKitInputWithDefaults() *ContentKitInput`

NewContentKitInputWithDefaults instantiates a new ContentKitInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitInput) SetType(v string)`

SetType sets Type field to given value.


### GetLabel

`func (o *ContentKitInput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentKitInput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentKitInput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetHint

`func (o *ContentKitInput) GetHint() ContentKitInputHint`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ContentKitInput) GetHintOk() (*ContentKitInputHint, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ContentKitInput) SetHint(v ContentKitInputHint)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ContentKitInput) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetElement

`func (o *ContentKitInput) GetElement() ContentKitInputElement`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ContentKitInput) GetElementOk() (*ContentKitInputElement, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ContentKitInput) SetElement(v ContentKitInputElement)`

SetElement sets Element field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


