# UserImpersonation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RiskEvaluation** | [**UserRiskEvaluation**](UserRiskEvaluation.md) |  | 
**AuthProviders** | [**[]FirebaseUserInfo**](FirebaseUserInfo.md) |  | 
**CreatedAt** | **string** |  | 
**LastSignInAt** | **string** |  | 
**Disabled** | **bool** |  | 
**Impersonation** | [**UserImpersonationInfo**](UserImpersonationInfo.md) |  | 

## Methods

### NewUserImpersonation

`func NewUserImpersonation(id string, riskEvaluation UserRiskEvaluation, authProviders []FirebaseUserInfo, createdAt string, lastSignInAt string, disabled bool, impersonation UserImpersonationInfo, ) *UserImpersonation`

NewUserImpersonation instantiates a new UserImpersonation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImpersonationWithDefaults

`func NewUserImpersonationWithDefaults() *UserImpersonation`

NewUserImpersonationWithDefaults instantiates a new UserImpersonation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserImpersonation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserImpersonation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserImpersonation) SetId(v string)`

SetId sets Id field to given value.


### GetRiskEvaluation

`func (o *UserImpersonation) GetRiskEvaluation() UserRiskEvaluation`

GetRiskEvaluation returns the RiskEvaluation field if non-nil, zero value otherwise.

### GetRiskEvaluationOk

`func (o *UserImpersonation) GetRiskEvaluationOk() (*UserRiskEvaluation, bool)`

GetRiskEvaluationOk returns a tuple with the RiskEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskEvaluation

`func (o *UserImpersonation) SetRiskEvaluation(v UserRiskEvaluation)`

SetRiskEvaluation sets RiskEvaluation field to given value.


### GetAuthProviders

`func (o *UserImpersonation) GetAuthProviders() []FirebaseUserInfo`

GetAuthProviders returns the AuthProviders field if non-nil, zero value otherwise.

### GetAuthProvidersOk

`func (o *UserImpersonation) GetAuthProvidersOk() (*[]FirebaseUserInfo, bool)`

GetAuthProvidersOk returns a tuple with the AuthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviders

`func (o *UserImpersonation) SetAuthProviders(v []FirebaseUserInfo)`

SetAuthProviders sets AuthProviders field to given value.


### GetCreatedAt

`func (o *UserImpersonation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserImpersonation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserImpersonation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastSignInAt

`func (o *UserImpersonation) GetLastSignInAt() string`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *UserImpersonation) GetLastSignInAtOk() (*string, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *UserImpersonation) SetLastSignInAt(v string)`

SetLastSignInAt sets LastSignInAt field to given value.


### GetDisabled

`func (o *UserImpersonation) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserImpersonation) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserImpersonation) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetImpersonation

`func (o *UserImpersonation) GetImpersonation() UserImpersonationInfo`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *UserImpersonation) GetImpersonationOk() (*UserImpersonationInfo, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *UserImpersonation) SetImpersonation(v UserImpersonationInfo)`

SetImpersonation sets Impersonation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


