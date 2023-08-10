# UserBackOfficeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RiskEvaluation** | [**UserRiskEvaluation**](UserRiskEvaluation.md) |  | 
**AuthProviders** | [**[]FirebaseUserInfo**](FirebaseUserInfo.md) |  | 
**CreatedAt** | **string** |  | 
**LastSignInAt** | **string** |  | 
**Disabled** | **bool** |  | 

## Methods

### NewUserBackOfficeInfo

`func NewUserBackOfficeInfo(id string, riskEvaluation UserRiskEvaluation, authProviders []FirebaseUserInfo, createdAt string, lastSignInAt string, disabled bool, ) *UserBackOfficeInfo`

NewUserBackOfficeInfo instantiates a new UserBackOfficeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBackOfficeInfoWithDefaults

`func NewUserBackOfficeInfoWithDefaults() *UserBackOfficeInfo`

NewUserBackOfficeInfoWithDefaults instantiates a new UserBackOfficeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBackOfficeInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBackOfficeInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBackOfficeInfo) SetId(v string)`

SetId sets Id field to given value.


### GetRiskEvaluation

`func (o *UserBackOfficeInfo) GetRiskEvaluation() UserRiskEvaluation`

GetRiskEvaluation returns the RiskEvaluation field if non-nil, zero value otherwise.

### GetRiskEvaluationOk

`func (o *UserBackOfficeInfo) GetRiskEvaluationOk() (*UserRiskEvaluation, bool)`

GetRiskEvaluationOk returns a tuple with the RiskEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEvaluation

`func (o *UserBackOfficeInfo) SetRiskEvaluation(v UserRiskEvaluation)`

SetRiskEvaluation sets RiskEvaluation field to given value.


### GetAuthProviders

`func (o *UserBackOfficeInfo) GetAuthProviders() []FirebaseUserInfo`

GetAuthProviders returns the AuthProviders field if non-nil, zero value otherwise.

### GetAuthProvidersOk

`func (o *UserBackOfficeInfo) GetAuthProvidersOk() (*[]FirebaseUserInfo, bool)`

GetAuthProvidersOk returns a tuple with the AuthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviders

`func (o *UserBackOfficeInfo) SetAuthProviders(v []FirebaseUserInfo)`

SetAuthProviders sets AuthProviders field to given value.


### GetCreatedAt

`func (o *UserBackOfficeInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserBackOfficeInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserBackOfficeInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastSignInAt

`func (o *UserBackOfficeInfo) GetLastSignInAt() string`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *UserBackOfficeInfo) GetLastSignInAtOk() (*string, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *UserBackOfficeInfo) SetLastSignInAt(v string)`

SetLastSignInAt sets LastSignInAt field to given value.


### GetDisabled

`func (o *UserBackOfficeInfo) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserBackOfficeInfo) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserBackOfficeInfo) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


