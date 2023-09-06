# RevisionTypeMergeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Revision created when merging a change request with primary. | 
**MergedFrom** | [**ChangeRequest**](ChangeRequest.md) |  | 

## Methods

### NewRevisionTypeMergeAllOf

`func NewRevisionTypeMergeAllOf(type_ string, mergedFrom ChangeRequest, ) *RevisionTypeMergeAllOf`

NewRevisionTypeMergeAllOf instantiates a new RevisionTypeMergeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionTypeMergeAllOfWithDefaults

`func NewRevisionTypeMergeAllOfWithDefaults() *RevisionTypeMergeAllOf`

NewRevisionTypeMergeAllOfWithDefaults instantiates a new RevisionTypeMergeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RevisionTypeMergeAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RevisionTypeMergeAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RevisionTypeMergeAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetMergedFrom

`func (o *RevisionTypeMergeAllOf) GetMergedFrom() ChangeRequest`

GetMergedFrom returns the MergedFrom field if non-nil, zero value otherwise.

### GetMergedFromOk

`func (o *RevisionTypeMergeAllOf) GetMergedFromOk() (*ChangeRequest, bool)`

GetMergedFromOk returns a tuple with the MergedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedFrom

`func (o *RevisionTypeMergeAllOf) SetMergedFrom(v ChangeRequest)`

SetMergedFrom sets MergedFrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


