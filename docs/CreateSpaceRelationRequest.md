# CreateSpaceRelationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | ID of the other space | 
**Type** | [**SpaceRelationType**](SpaceRelationType.md) |  | 

## Methods

### NewCreateSpaceRelationRequest

`func NewCreateSpaceRelationRequest(target string, type_ SpaceRelationType, ) *CreateSpaceRelationRequest`

NewCreateSpaceRelationRequest instantiates a new CreateSpaceRelationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSpaceRelationRequestWithDefaults

`func NewCreateSpaceRelationRequestWithDefaults() *CreateSpaceRelationRequest`

NewCreateSpaceRelationRequestWithDefaults instantiates a new CreateSpaceRelationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *CreateSpaceRelationRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateSpaceRelationRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateSpaceRelationRequest) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetType

`func (o *CreateSpaceRelationRequest) GetType() SpaceRelationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSpaceRelationRequest) GetTypeOk() (*SpaceRelationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSpaceRelationRequest) SetType(v SpaceRelationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


