# CreateEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | **string** |  | 
**Primary** | Pointer to **bool** | Whether or not an environment is considered to have elevated responsibilities over other environments. Useful for distinguishing a production environment from a staging environment. Multiple primary environments are allowed. Your organization must have at least one primary environment.  | [optional] 

## Methods

### NewCreateEnvironment

`func NewCreateEnvironment(name string, title string, ) *CreateEnvironment`

NewCreateEnvironment instantiates a new CreateEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentWithDefaults

`func NewCreateEnvironmentWithDefaults() *CreateEnvironment`

NewCreateEnvironmentWithDefaults instantiates a new CreateEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateEnvironment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateEnvironment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateEnvironment) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrimary

`func (o *CreateEnvironment) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *CreateEnvironment) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *CreateEnvironment) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *CreateEnvironment) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


