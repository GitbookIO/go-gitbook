# SpaceIntegrationBlocksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique named identifier for the integration | 
**Blocks** | [**[]IntegrationBlock**](IntegrationBlock.md) |  | 

## Methods

### NewSpaceIntegrationBlocksInner

`func NewSpaceIntegrationBlocksInner(name string, blocks []IntegrationBlock, ) *SpaceIntegrationBlocksInner`

NewSpaceIntegrationBlocksInner instantiates a new SpaceIntegrationBlocksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceIntegrationBlocksInnerWithDefaults

`func NewSpaceIntegrationBlocksInnerWithDefaults() *SpaceIntegrationBlocksInner`

NewSpaceIntegrationBlocksInnerWithDefaults instantiates a new SpaceIntegrationBlocksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpaceIntegrationBlocksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceIntegrationBlocksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceIntegrationBlocksInner) SetName(v string)`

SetName sets Name field to given value.


### GetBlocks

`func (o *SpaceIntegrationBlocksInner) GetBlocks() []IntegrationBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *SpaceIntegrationBlocksInner) GetBlocksOk() (*[]IntegrationBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *SpaceIntegrationBlocksInner) SetBlocks(v []IntegrationBlock)`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


