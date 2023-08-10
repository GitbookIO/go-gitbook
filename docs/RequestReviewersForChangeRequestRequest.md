# RequestReviewersForChangeRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]string** | An array of user ids that will be requested. | 
**Subject** | Pointer to **string** | Optionally, update the subject of the change request when requesting reviewers. | [optional] 

## Methods

### NewRequestReviewersForChangeRequestRequest

`func NewRequestReviewersForChangeRequestRequest(users []string, ) *RequestReviewersForChangeRequestRequest`

NewRequestReviewersForChangeRequestRequest instantiates a new RequestReviewersForChangeRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestReviewersForChangeRequestRequestWithDefaults

`func NewRequestReviewersForChangeRequestRequestWithDefaults() *RequestReviewersForChangeRequestRequest`

NewRequestReviewersForChangeRequestRequestWithDefaults instantiates a new RequestReviewersForChangeRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *RequestReviewersForChangeRequestRequest) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RequestReviewersForChangeRequestRequest) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RequestReviewersForChangeRequestRequest) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetSubject

`func (o *RequestReviewersForChangeRequestRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestReviewersForChangeRequestRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestReviewersForChangeRequestRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RequestReviewersForChangeRequestRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


