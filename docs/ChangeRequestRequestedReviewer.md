# ChangeRequestRequestedReviewer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;change-request-requested-reviewer\&quot; | 
**Revision** | **string** | The revision of the content when the request was made. | 
**RequestedBy** | [**User**](User.md) |  | 
**CreatedAt** | **string** |  | 
**Kind** | **string** |  | 
**User** | [**User**](User.md) |  | 
**Team** | [**Team**](Team.md) |  | 

## Methods

### NewChangeRequestRequestedReviewer

`func NewChangeRequestRequestedReviewer(object string, revision string, requestedBy User, createdAt string, kind string, user User, team Team, ) *ChangeRequestRequestedReviewer`

NewChangeRequestRequestedReviewer instantiates a new ChangeRequestRequestedReviewer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestRequestedReviewerWithDefaults

`func NewChangeRequestRequestedReviewerWithDefaults() *ChangeRequestRequestedReviewer`

NewChangeRequestRequestedReviewerWithDefaults instantiates a new ChangeRequestRequestedReviewer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ChangeRequestRequestedReviewer) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChangeRequestRequestedReviewer) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChangeRequestRequestedReviewer) SetObject(v string)`

SetObject sets Object field to given value.


### GetRevision

`func (o *ChangeRequestRequestedReviewer) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ChangeRequestRequestedReviewer) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ChangeRequestRequestedReviewer) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetRequestedBy

`func (o *ChangeRequestRequestedReviewer) GetRequestedBy() User`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *ChangeRequestRequestedReviewer) GetRequestedByOk() (*User, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *ChangeRequestRequestedReviewer) SetRequestedBy(v User)`

SetRequestedBy sets RequestedBy field to given value.


### GetCreatedAt

`func (o *ChangeRequestRequestedReviewer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestRequestedReviewer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestRequestedReviewer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetKind

`func (o *ChangeRequestRequestedReviewer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ChangeRequestRequestedReviewer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ChangeRequestRequestedReviewer) SetKind(v string)`

SetKind sets Kind field to given value.


### GetUser

`func (o *ChangeRequestRequestedReviewer) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChangeRequestRequestedReviewer) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChangeRequestRequestedReviewer) SetUser(v User)`

SetUser sets User field to given value.


### GetTeam

`func (o *ChangeRequestRequestedReviewer) GetTeam() Team`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ChangeRequestRequestedReviewer) GetTeamOk() (*Team, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ChangeRequestRequestedReviewer) SetTeam(v Team)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


