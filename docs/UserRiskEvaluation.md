# UserRiskEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WasRisky** | **bool** | True if the user was originally considered as risky | 
**IsRisky** | **bool** | True if the user is currently considered as risky | 
**IsVerified** | **bool** | True if the user went through the verification process | 
**RiskScore** | **float32** | Risk score of the user | 
**CompletedSteps** | **float32** | Number of verification steps completed by the user | 
**ExpectedSteps** | **float32** | Total number of verification steps expected | 
**GoogleLogin** | **bool** | User completed the Google Account verification step | 
**GithubLogin** | **bool** | User completed the GitHub Account verification step | 
**EmailVerified** | **bool** | User completed the Email verification step | 
**ActiveDaysRemaining** | **float32** |  | 

## Methods

### NewUserRiskEvaluation

`func NewUserRiskEvaluation(wasRisky bool, isRisky bool, isVerified bool, riskScore float32, completedSteps float32, expectedSteps float32, googleLogin bool, githubLogin bool, emailVerified bool, activeDaysRemaining float32, ) *UserRiskEvaluation`

NewUserRiskEvaluation instantiates a new UserRiskEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRiskEvaluationWithDefaults

`func NewUserRiskEvaluationWithDefaults() *UserRiskEvaluation`

NewUserRiskEvaluationWithDefaults instantiates a new UserRiskEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWasRisky

`func (o *UserRiskEvaluation) GetWasRisky() bool`

GetWasRisky returns the WasRisky field if non-nil, zero value otherwise.

### GetWasRiskyOk

`func (o *UserRiskEvaluation) GetWasRiskyOk() (*bool, bool)`

GetWasRiskyOk returns a tuple with the WasRisky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasRisky

`func (o *UserRiskEvaluation) SetWasRisky(v bool)`

SetWasRisky sets WasRisky field to given value.


### GetIsRisky

`func (o *UserRiskEvaluation) GetIsRisky() bool`

GetIsRisky returns the IsRisky field if non-nil, zero value otherwise.

### GetIsRiskyOk

`func (o *UserRiskEvaluation) GetIsRiskyOk() (*bool, bool)`

GetIsRiskyOk returns a tuple with the IsRisky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRisky

`func (o *UserRiskEvaluation) SetIsRisky(v bool)`

SetIsRisky sets IsRisky field to given value.


### GetIsVerified

`func (o *UserRiskEvaluation) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *UserRiskEvaluation) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *UserRiskEvaluation) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.


### GetRiskScore

`func (o *UserRiskEvaluation) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *UserRiskEvaluation) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *UserRiskEvaluation) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.


### GetCompletedSteps

`func (o *UserRiskEvaluation) GetCompletedSteps() float32`

GetCompletedSteps returns the CompletedSteps field if non-nil, zero value otherwise.

### GetCompletedStepsOk

`func (o *UserRiskEvaluation) GetCompletedStepsOk() (*float32, bool)`

GetCompletedStepsOk returns a tuple with the CompletedSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedSteps

`func (o *UserRiskEvaluation) SetCompletedSteps(v float32)`

SetCompletedSteps sets CompletedSteps field to given value.


### GetExpectedSteps

`func (o *UserRiskEvaluation) GetExpectedSteps() float32`

GetExpectedSteps returns the ExpectedSteps field if non-nil, zero value otherwise.

### GetExpectedStepsOk

`func (o *UserRiskEvaluation) GetExpectedStepsOk() (*float32, bool)`

GetExpectedStepsOk returns a tuple with the ExpectedSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSteps

`func (o *UserRiskEvaluation) SetExpectedSteps(v float32)`

SetExpectedSteps sets ExpectedSteps field to given value.


### GetGoogleLogin

`func (o *UserRiskEvaluation) GetGoogleLogin() bool`

GetGoogleLogin returns the GoogleLogin field if non-nil, zero value otherwise.

### GetGoogleLoginOk

`func (o *UserRiskEvaluation) GetGoogleLoginOk() (*bool, bool)`

GetGoogleLoginOk returns a tuple with the GoogleLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleLogin

`func (o *UserRiskEvaluation) SetGoogleLogin(v bool)`

SetGoogleLogin sets GoogleLogin field to given value.


### GetGithubLogin

`func (o *UserRiskEvaluation) GetGithubLogin() bool`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *UserRiskEvaluation) GetGithubLoginOk() (*bool, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *UserRiskEvaluation) SetGithubLogin(v bool)`

SetGithubLogin sets GithubLogin field to given value.


### GetEmailVerified

`func (o *UserRiskEvaluation) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserRiskEvaluation) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserRiskEvaluation) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetActiveDaysRemaining

`func (o *UserRiskEvaluation) GetActiveDaysRemaining() float32`

GetActiveDaysRemaining returns the ActiveDaysRemaining field if non-nil, zero value otherwise.

### GetActiveDaysRemainingOk

`func (o *UserRiskEvaluation) GetActiveDaysRemainingOk() (*float32, bool)`

GetActiveDaysRemainingOk returns a tuple with the ActiveDaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDaysRemaining

`func (o *UserRiskEvaluation) SetActiveDaysRemaining(v float32)`

SetActiveDaysRemaining sets ActiveDaysRemaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


