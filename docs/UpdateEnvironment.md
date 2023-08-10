# UpdateEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to **bool** | Whether or not an environment is considered to have elevated responsibilities over other environments. Useful for distinguishing a production environment from a staging environment. Multiple primary environments are allowed. Your organization must have at least one primary environment.  | [optional] 

## Methods

### NewUpdateEnvironment

`func NewUpdateEnvironment() *UpdateEnvironment`

NewUpdateEnvironment instantiates a new UpdateEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnvironmentWithDefaults

`func NewUpdateEnvironmentWithDefaults() *UpdateEnvironment`

NewUpdateEnvironmentWithDefaults instantiates a new UpdateEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateEnvironment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateEnvironment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateEnvironment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateEnvironment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPrimary

`func (o *UpdateEnvironment) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *UpdateEnvironment) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *UpdateEnvironment) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *UpdateEnvironment) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


