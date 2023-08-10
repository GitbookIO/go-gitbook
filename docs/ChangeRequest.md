# ChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;change-request\&quot; | 
**Id** | **string** | Unique identifier for the change request | 
**Number** | **float32** | Incremental identifier of the change request | 
**Status** | [**ChangeRequestStatus**](ChangeRequestStatus.md) |  | 
**Subject** | **string** | Subject of the change request | 
**CreatedBy** | [**User**](User.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Urls** | [**ChangeRequestUrls**](ChangeRequestUrls.md) |  | 

## Methods

### NewChangeRequest

`func NewChangeRequest(object string, id string, number float32, status ChangeRequestStatus, subject string, createdBy User, createdAt string, updatedAt string, urls ChangeRequestUrls, ) *ChangeRequest`

NewChangeRequest instantiates a new ChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestWithDefaults

`func NewChangeRequestWithDefaults() *ChangeRequest`

NewChangeRequestWithDefaults instantiates a new ChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ChangeRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChangeRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChangeRequest) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *ChangeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *ChangeRequest) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ChangeRequest) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ChangeRequest) SetNumber(v float32)`

SetNumber sets Number field to given value.


### GetStatus

`func (o *ChangeRequest) GetStatus() ChangeRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeRequest) GetStatusOk() (*ChangeRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeRequest) SetStatus(v ChangeRequestStatus)`

SetStatus sets Status field to given value.


### GetSubject

`func (o *ChangeRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ChangeRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ChangeRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCreatedBy

`func (o *ChangeRequest) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChangeRequest) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChangeRequest) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *ChangeRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ChangeRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChangeRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChangeRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrls

`func (o *ChangeRequest) GetUrls() ChangeRequestUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ChangeRequest) GetUrlsOk() (*ChangeRequestUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ChangeRequest) SetUrls(v ChangeRequestUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


