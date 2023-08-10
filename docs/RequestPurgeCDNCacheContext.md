# RequestPurgeCDNCacheContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PurgeCDNCacheContextType**](PurgeCDNCacheContextType.md) |  | 
**Values** | **[]string** | The list of tags or hosts to purge | 

## Methods

### NewRequestPurgeCDNCacheContext

`func NewRequestPurgeCDNCacheContext(type_ PurgeCDNCacheContextType, values []string, ) *RequestPurgeCDNCacheContext`

NewRequestPurgeCDNCacheContext instantiates a new RequestPurgeCDNCacheContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPurgeCDNCacheContextWithDefaults

`func NewRequestPurgeCDNCacheContextWithDefaults() *RequestPurgeCDNCacheContext`

NewRequestPurgeCDNCacheContextWithDefaults instantiates a new RequestPurgeCDNCacheContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestPurgeCDNCacheContext) GetType() PurgeCDNCacheContextType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestPurgeCDNCacheContext) GetTypeOk() (*PurgeCDNCacheContextType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestPurgeCDNCacheContext) SetType(v PurgeCDNCacheContextType)`

SetType sets Type field to given value.


### GetValues

`func (o *RequestPurgeCDNCacheContext) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RequestPurgeCDNCacheContext) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RequestPurgeCDNCacheContext) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


