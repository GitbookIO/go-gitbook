# ChangeRequestReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;change-request-review\&quot; | 
**Id** | **string** | Unique identifier for the review. | 
**Revision** | **string** | The revision this review was made against. | 
**Reviewer** | [**User**](User.md) |  | 
**RequestedBy** | Pointer to [**User**](User.md) |  | [optional] 
**Status** | [**ChangeRequestReviewStatus**](ChangeRequestReviewStatus.md) |  | 
**Comment** | Pointer to [**Comment**](Comment.md) |  | [optional] 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewChangeRequestReview

`func NewChangeRequestReview(object string, id string, revision string, reviewer User, status ChangeRequestReviewStatus, createdAt string, updatedAt string, ) *ChangeRequestReview`

NewChangeRequestReview instantiates a new ChangeRequestReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestReviewWithDefaults

`func NewChangeRequestReviewWithDefaults() *ChangeRequestReview`

NewChangeRequestReviewWithDefaults instantiates a new ChangeRequestReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ChangeRequestReview) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChangeRequestReview) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChangeRequestReview) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *ChangeRequestReview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRequestReview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRequestReview) SetId(v string)`

SetId sets Id field to given value.


### GetRevision

`func (o *ChangeRequestReview) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ChangeRequestReview) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ChangeRequestReview) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetReviewer

`func (o *ChangeRequestReview) GetReviewer() User`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *ChangeRequestReview) GetReviewerOk() (*User, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *ChangeRequestReview) SetReviewer(v User)`

SetReviewer sets Reviewer field to given value.


### GetRequestedBy

`func (o *ChangeRequestReview) GetRequestedBy() User`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *ChangeRequestReview) GetRequestedByOk() (*User, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *ChangeRequestReview) SetRequestedBy(v User)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *ChangeRequestReview) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetStatus

`func (o *ChangeRequestReview) GetStatus() ChangeRequestReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeRequestReview) GetStatusOk() (*ChangeRequestReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeRequestReview) SetStatus(v ChangeRequestReviewStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *ChangeRequestReview) GetComment() Comment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ChangeRequestReview) GetCommentOk() (*Comment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ChangeRequestReview) SetComment(v Comment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ChangeRequestReview) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChangeRequestReview) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestReview) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestReview) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ChangeRequestReview) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChangeRequestReview) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChangeRequestReview) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


