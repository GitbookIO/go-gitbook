# TriggerContentIndexingContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | **string** | The unique identifier of the Space to index. | 
**Force** | **bool** | Whether to force a complete re-indexing of the Space. | 

## Methods

### NewTriggerContentIndexingContext

`func NewTriggerContentIndexingContext(spaceId string, force bool, ) *TriggerContentIndexingContext`

NewTriggerContentIndexingContext instantiates a new TriggerContentIndexingContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerContentIndexingContextWithDefaults

`func NewTriggerContentIndexingContextWithDefaults() *TriggerContentIndexingContext`

NewTriggerContentIndexingContextWithDefaults instantiates a new TriggerContentIndexingContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *TriggerContentIndexingContext) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *TriggerContentIndexingContext) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *TriggerContentIndexingContext) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetForce

`func (o *TriggerContentIndexingContext) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *TriggerContentIndexingContext) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *TriggerContentIndexingContext) SetForce(v bool)`

SetForce sets Force field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


