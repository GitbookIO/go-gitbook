# SubmitChangeRequestReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ChangeRequestReviewStatus**](ChangeRequestReviewStatus.md) |  | 
**Comment** | Pointer to [**Document**](Document.md) |  | [optional] 

## Methods

### NewSubmitChangeRequestReviewRequest

`func NewSubmitChangeRequestReviewRequest(status ChangeRequestReviewStatus, ) *SubmitChangeRequestReviewRequest`

NewSubmitChangeRequestReviewRequest instantiates a new SubmitChangeRequestReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitChangeRequestReviewRequestWithDefaults

`func NewSubmitChangeRequestReviewRequestWithDefaults() *SubmitChangeRequestReviewRequest`

NewSubmitChangeRequestReviewRequestWithDefaults instantiates a new SubmitChangeRequestReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubmitChangeRequestReviewRequest) GetStatus() ChangeRequestReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitChangeRequestReviewRequest) GetStatusOk() (*ChangeRequestReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitChangeRequestReviewRequest) SetStatus(v ChangeRequestReviewStatus)`

SetStatus sets Status field to given value.


### GetComment

`func (o *SubmitChangeRequestReviewRequest) GetComment() Document`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SubmitChangeRequestReviewRequest) GetCommentOk() (*Document, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SubmitChangeRequestReviewRequest) SetComment(v Document)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SubmitChangeRequestReviewRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


