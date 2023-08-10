# SpaceRelationPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**SpaceRelationType**](SpaceRelationType.md) |  | 
**Target** | [**SpaceRelationEdge**](SpaceRelationEdge.md) |  | 
**CreatedAt** | **string** |  | 
**Urls** | [**SpaceRelationUrls**](SpaceRelationUrls.md) |  | 
**Source** | [**SpaceRelationEdge**](SpaceRelationEdge.md) |  | 

## Methods

### NewSpaceRelationPair

`func NewSpaceRelationPair(id string, type_ SpaceRelationType, target SpaceRelationEdge, createdAt string, urls SpaceRelationUrls, source SpaceRelationEdge, ) *SpaceRelationPair`

NewSpaceRelationPair instantiates a new SpaceRelationPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceRelationPairWithDefaults

`func NewSpaceRelationPairWithDefaults() *SpaceRelationPair`

NewSpaceRelationPairWithDefaults instantiates a new SpaceRelationPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpaceRelationPair) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpaceRelationPair) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpaceRelationPair) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SpaceRelationPair) GetType() SpaceRelationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceRelationPair) GetTypeOk() (*SpaceRelationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceRelationPair) SetType(v SpaceRelationType)`

SetType sets Type field to given value.


### GetTarget

`func (o *SpaceRelationPair) GetTarget() SpaceRelationEdge`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SpaceRelationPair) GetTargetOk() (*SpaceRelationEdge, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SpaceRelationPair) SetTarget(v SpaceRelationEdge)`

SetTarget sets Target field to given value.


### GetCreatedAt

`func (o *SpaceRelationPair) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpaceRelationPair) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpaceRelationPair) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUrls

`func (o *SpaceRelationPair) GetUrls() SpaceRelationUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SpaceRelationPair) GetUrlsOk() (*SpaceRelationUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SpaceRelationPair) SetUrls(v SpaceRelationUrls)`

SetUrls sets Urls field to given value.


### GetSource

`func (o *SpaceRelationPair) GetSource() SpaceRelationEdge`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SpaceRelationPair) GetSourceOk() (*SpaceRelationEdge, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SpaceRelationPair) SetSource(v SpaceRelationEdge)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


