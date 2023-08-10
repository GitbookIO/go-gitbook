# ContentKitTextInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**State** | **string** | State binding. The value of the input will be stored as a property in the state named after this ID. | 
**InitialValue** | Pointer to **string** | Text value to initialize the input with. | [optional] 
**Placeholder** | Pointer to **string** | Text that appears in the form control when it has no value set | [optional] 
**Multiline** | Pointer to **bool** |  | [optional] 

## Methods

### NewContentKitTextInput

`func NewContentKitTextInput(type_ string, state string, ) *ContentKitTextInput`

NewContentKitTextInput instantiates a new ContentKitTextInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitTextInputWithDefaults

`func NewContentKitTextInputWithDefaults() *ContentKitTextInput`

NewContentKitTextInputWithDefaults instantiates a new ContentKitTextInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitTextInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitTextInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitTextInput) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *ContentKitTextInput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitTextInput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitTextInput) SetState(v string)`

SetState sets State field to given value.


### GetInitialValue

`func (o *ContentKitTextInput) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *ContentKitTextInput) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *ContentKitTextInput) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *ContentKitTextInput) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ContentKitTextInput) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ContentKitTextInput) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ContentKitTextInput) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ContentKitTextInput) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMultiline

`func (o *ContentKitTextInput) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *ContentKitTextInput) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *ContentKitTextInput) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *ContentKitTextInput) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


