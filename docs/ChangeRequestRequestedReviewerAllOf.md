# ChangeRequestRequestedReviewerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;change-request-requested-reviewer\&quot; | 
**Revision** | **string** | The revision of the content when the request was made. | 
**RequestedBy** | [**User**](User.md) |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewChangeRequestRequestedReviewerAllOf

`func NewChangeRequestRequestedReviewerAllOf(object string, revision string, requestedBy User, createdAt string, ) *ChangeRequestRequestedReviewerAllOf`

NewChangeRequestRequestedReviewerAllOf instantiates a new ChangeRequestRequestedReviewerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestRequestedReviewerAllOfWithDefaults

`func NewChangeRequestRequestedReviewerAllOfWithDefaults() *ChangeRequestRequestedReviewerAllOf`

NewChangeRequestRequestedReviewerAllOfWithDefaults instantiates a new ChangeRequestRequestedReviewerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ChangeRequestRequestedReviewerAllOf) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChangeRequestRequestedReviewerAllOf) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChangeRequestRequestedReviewerAllOf) SetObject(v string)`

SetObject sets Object field to given value.


### GetRevision

`func (o *ChangeRequestRequestedReviewerAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ChangeRequestRequestedReviewerAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ChangeRequestRequestedReviewerAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetRequestedBy

`func (o *ChangeRequestRequestedReviewerAllOf) GetRequestedBy() User`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *ChangeRequestRequestedReviewerAllOf) GetRequestedByOk() (*User, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *ChangeRequestRequestedReviewerAllOf) SetRequestedBy(v User)`

SetRequestedBy sets RequestedBy field to given value.


### GetCreatedAt

`func (o *ChangeRequestRequestedReviewerAllOf) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestRequestedReviewerAllOf) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestRequestedReviewerAllOf) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


