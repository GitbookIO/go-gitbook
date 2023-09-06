# ContentKitSelectOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | **string** |  | 
**Icon** | Pointer to [**ContentKitCardIcon**](ContentKitCardIcon.md) |  | [optional] 

## Methods

### NewContentKitSelectOption

`func NewContentKitSelectOption(id string, label string, ) *ContentKitSelectOption`

NewContentKitSelectOption instantiates a new ContentKitSelectOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitSelectOptionWithDefaults

`func NewContentKitSelectOptionWithDefaults() *ContentKitSelectOption`

NewContentKitSelectOptionWithDefaults instantiates a new ContentKitSelectOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContentKitSelectOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentKitSelectOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentKitSelectOption) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ContentKitSelectOption) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContentKitSelectOption) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContentKitSelectOption) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetIcon

`func (o *ContentKitSelectOption) GetIcon() ContentKitCardIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ContentKitSelectOption) GetIconOk() (*ContentKitCardIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ContentKitSelectOption) SetIcon(v ContentKitCardIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ContentKitSelectOption) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


