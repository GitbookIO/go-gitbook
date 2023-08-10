# SpaceRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**SpaceRelationType**](SpaceRelationType.md) |  | 
**Target** | [**SpaceRelationTarget**](SpaceRelationTarget.md) |  | 
**CreatedAt** | **string** |  | 
**Urls** | [**SpaceRelationUrls**](SpaceRelationUrls.md) |  | 

## Methods

### NewSpaceRelation

`func NewSpaceRelation(id string, type_ SpaceRelationType, target SpaceRelationTarget, createdAt string, urls SpaceRelationUrls, ) *SpaceRelation`

NewSpaceRelation instantiates a new SpaceRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceRelationWithDefaults

`func NewSpaceRelationWithDefaults() *SpaceRelation`

NewSpaceRelationWithDefaults instantiates a new SpaceRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpaceRelation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpaceRelation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpaceRelation) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SpaceRelation) GetType() SpaceRelationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceRelation) GetTypeOk() (*SpaceRelationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceRelation) SetType(v SpaceRelationType)`

SetType sets Type field to given value.


### GetTarget

`func (o *SpaceRelation) GetTarget() SpaceRelationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SpaceRelation) GetTargetOk() (*SpaceRelationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SpaceRelation) SetTarget(v SpaceRelationTarget)`

SetTarget sets Target field to given value.


### GetCreatedAt

`func (o *SpaceRelation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpaceRelation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpaceRelation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUrls

`func (o *SpaceRelation) GetUrls() SpaceRelationUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SpaceRelation) GetUrlsOk() (*SpaceRelationUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SpaceRelation) SetUrls(v SpaceRelationUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


