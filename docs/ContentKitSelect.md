# ContentKitSelect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**State** | **string** | State binding. The value of the input will be stored as a property in the state named after this ID. | 
**InitialValue** | Pointer to [**ContentKitSelectInitialValue**](ContentKitSelectInitialValue.md) |  | [optional] 
**Placeholder** | Pointer to **string** | Text that appears in the form control when it has no value set | [optional] 
**Multiple** | Pointer to **bool** | Should the select accept the selection of multiple options. If true, the state will be an array. | [optional] 
**Options** | [**ContentKitSelectOptions**](ContentKitSelectOptions.md) |  | 

## Methods

### NewContentKitSelect

`func NewContentKitSelect(type_ string, state string, options ContentKitSelectOptions, ) *ContentKitSelect`

NewContentKitSelect instantiates a new ContentKitSelect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitSelectWithDefaults

`func NewContentKitSelectWithDefaults() *ContentKitSelect`

NewContentKitSelectWithDefaults instantiates a new ContentKitSelect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitSelect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitSelect) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitSelect) SetType(v string)`

SetType sets Type field to given value.


### GetState

`func (o *ContentKitSelect) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContentKitSelect) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContentKitSelect) SetState(v string)`

SetState sets State field to given value.


### GetInitialValue

`func (o *ContentKitSelect) GetInitialValue() ContentKitSelectInitialValue`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *ContentKitSelect) GetInitialValueOk() (*ContentKitSelectInitialValue, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *ContentKitSelect) SetInitialValue(v ContentKitSelectInitialValue)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *ContentKitSelect) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetPlaceholder

`func (o *ContentKitSelect) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *ContentKitSelect) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *ContentKitSelect) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *ContentKitSelect) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetMultiple

`func (o *ContentKitSelect) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *ContentKitSelect) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *ContentKitSelect) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *ContentKitSelect) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetOptions

`func (o *ContentKitSelect) GetOptions() ContentKitSelectOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ContentKitSelect) GetOptionsOk() (*ContentKitSelectOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ContentKitSelect) SetOptions(v ContentKitSelectOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


