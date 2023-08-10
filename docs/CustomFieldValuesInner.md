# CustomFieldValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomField** | [**CustomField**](CustomField.md) |  | 
**Value** | Pointer to [**CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewCustomFieldValuesInner

`func NewCustomFieldValuesInner(customField CustomField, ) *CustomFieldValuesInner`

NewCustomFieldValuesInner instantiates a new CustomFieldValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldValuesInnerWithDefaults

`func NewCustomFieldValuesInnerWithDefaults() *CustomFieldValuesInner`

NewCustomFieldValuesInnerWithDefaults instantiates a new CustomFieldValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomField

`func (o *CustomFieldValuesInner) GetCustomField() CustomField`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *CustomFieldValuesInner) GetCustomFieldOk() (*CustomField, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *CustomFieldValuesInner) SetCustomField(v CustomField)`

SetCustomField sets CustomField field to given value.


### GetValue

`func (o *CustomFieldValuesInner) GetValue() CustomFieldValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldValuesInner) GetValueOk() (*CustomFieldValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldValuesInner) SetValue(v CustomFieldValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldValuesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


