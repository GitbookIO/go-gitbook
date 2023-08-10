# RequestReviewersForChangeRequest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]ChangeRequestRequestedReviewer**](ChangeRequestRequestedReviewer.md) | The user requests that were sent. | [optional] 

## Methods

### NewRequestReviewersForChangeRequest200Response

`func NewRequestReviewersForChangeRequest200Response() *RequestReviewersForChangeRequest200Response`

NewRequestReviewersForChangeRequest200Response instantiates a new RequestReviewersForChangeRequest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestReviewersForChangeRequest200ResponseWithDefaults

`func NewRequestReviewersForChangeRequest200ResponseWithDefaults() *RequestReviewersForChangeRequest200Response`

NewRequestReviewersForChangeRequest200ResponseWithDefaults instantiates a new RequestReviewersForChangeRequest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *RequestReviewersForChangeRequest200Response) GetUsers() []ChangeRequestRequestedReviewer`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RequestReviewersForChangeRequest200Response) GetUsersOk() (*[]ChangeRequestRequestedReviewer, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RequestReviewersForChangeRequest200Response) SetUsers(v []ChangeRequestRequestedReviewer)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RequestReviewersForChangeRequest200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


