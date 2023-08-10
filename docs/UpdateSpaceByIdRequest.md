# UpdateSpaceByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visibility** | Pointer to [**ContentVisibility**](ContentVisibility.md) |  | [optional] 
**Type** | Pointer to [**SpaceType**](SpaceType.md) |  | [optional] 

## Methods

### NewUpdateSpaceByIdRequest

`func NewUpdateSpaceByIdRequest() *UpdateSpaceByIdRequest`

NewUpdateSpaceByIdRequest instantiates a new UpdateSpaceByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpaceByIdRequestWithDefaults

`func NewUpdateSpaceByIdRequestWithDefaults() *UpdateSpaceByIdRequest`

NewUpdateSpaceByIdRequestWithDefaults instantiates a new UpdateSpaceByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibility

`func (o *UpdateSpaceByIdRequest) GetVisibility() ContentVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateSpaceByIdRequest) GetVisibilityOk() (*ContentVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateSpaceByIdRequest) SetVisibility(v ContentVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateSpaceByIdRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetType

`func (o *UpdateSpaceByIdRequest) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSpaceByIdRequest) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSpaceByIdRequest) SetType(v SpaceType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateSpaceByIdRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


