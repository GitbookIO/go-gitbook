# UserAPIToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The API token ID. | 
**Label** | **string** | The API token name. | 
**CreatedAt** | **string** |  | 

## Methods

### NewUserAPIToken

`func NewUserAPIToken(id string, label string, createdAt string, ) *UserAPIToken`

NewUserAPIToken instantiates a new UserAPIToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAPITokenWithDefaults

`func NewUserAPITokenWithDefaults() *UserAPIToken`

NewUserAPITokenWithDefaults instantiates a new UserAPIToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAPIToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAPIToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAPIToken) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *UserAPIToken) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UserAPIToken) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UserAPIToken) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedAt

`func (o *UserAPIToken) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserAPIToken) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserAPIToken) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


