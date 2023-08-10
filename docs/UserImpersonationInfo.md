# UserImpersonationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthURL** | **string** |  | 
**ImpersonatorId** | **string** |  | 

## Methods

### NewUserImpersonationInfo

`func NewUserImpersonationInfo(authURL string, impersonatorId string, ) *UserImpersonationInfo`

NewUserImpersonationInfo instantiates a new UserImpersonationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImpersonationInfoWithDefaults

`func NewUserImpersonationInfoWithDefaults() *UserImpersonationInfo`

NewUserImpersonationInfoWithDefaults instantiates a new UserImpersonationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthURL

`func (o *UserImpersonationInfo) GetAuthURL() string`

GetAuthURL returns the AuthURL field if non-nil, zero value otherwise.

### GetAuthURLOk

`func (o *UserImpersonationInfo) GetAuthURLOk() (*string, bool)`

GetAuthURLOk returns a tuple with the AuthURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthURL

`func (o *UserImpersonationInfo) SetAuthURL(v string)`

SetAuthURL sets AuthURL field to given value.


### GetImpersonatorId

`func (o *UserImpersonationInfo) GetImpersonatorId() string`

GetImpersonatorId returns the ImpersonatorId field if non-nil, zero value otherwise.

### GetImpersonatorIdOk

`func (o *UserImpersonationInfo) GetImpersonatorIdOk() (*string, bool)`

GetImpersonatorIdOk returns a tuple with the ImpersonatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonatorId

`func (o *UserImpersonationInfo) SetImpersonatorId(v string)`

SetImpersonatorId sets ImpersonatorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


