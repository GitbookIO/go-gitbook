/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the UserRiskEvaluation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskEvaluation{}

// UserRiskEvaluation The GitBook User risk evaluation.
type UserRiskEvaluation struct {
	// True if the user was originally considered as risky
	WasRisky bool `json:"wasRisky"`
	// True if the user is currently considered as risky
	IsRisky bool `json:"isRisky"`
	// True if the user went through the verification process
	IsVerified bool `json:"isVerified"`
	// Risk score of the user
	RiskScore float32 `json:"riskScore"`
	// Number of verification steps completed by the user
	CompletedSteps float32 `json:"completedSteps"`
	// Total number of verification steps expected
	ExpectedSteps float32 `json:"expectedSteps"`
	// User completed the Google Account verification step
	GoogleLogin bool `json:"googleLogin"`
	// User completed the GitHub Account verification step
	GithubLogin bool `json:"githubLogin"`
	// User completed the Email verification step
	EmailVerified       bool    `json:"emailVerified"`
	ActiveDaysRemaining float32 `json:"activeDaysRemaining"`
}

// NewUserRiskEvaluation instantiates a new UserRiskEvaluation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskEvaluation(wasRisky bool, isRisky bool, isVerified bool, riskScore float32, completedSteps float32, expectedSteps float32, googleLogin bool, githubLogin bool, emailVerified bool, activeDaysRemaining float32) *UserRiskEvaluation {
	this := UserRiskEvaluation{}
	this.WasRisky = wasRisky
	this.IsRisky = isRisky
	this.IsVerified = isVerified
	this.RiskScore = riskScore
	this.CompletedSteps = completedSteps
	this.ExpectedSteps = expectedSteps
	this.GoogleLogin = googleLogin
	this.GithubLogin = githubLogin
	this.EmailVerified = emailVerified
	this.ActiveDaysRemaining = activeDaysRemaining
	return &this
}

// NewUserRiskEvaluationWithDefaults instantiates a new UserRiskEvaluation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskEvaluationWithDefaults() *UserRiskEvaluation {
	this := UserRiskEvaluation{}
	return &this
}

// GetWasRisky returns the WasRisky field value
func (o *UserRiskEvaluation) GetWasRisky() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WasRisky
}

// GetWasRiskyOk returns a tuple with the WasRisky field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetWasRiskyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WasRisky, true
}

// SetWasRisky sets field value
func (o *UserRiskEvaluation) SetWasRisky(v bool) {
	o.WasRisky = v
}

// GetIsRisky returns the IsRisky field value
func (o *UserRiskEvaluation) GetIsRisky() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRisky
}

// GetIsRiskyOk returns a tuple with the IsRisky field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetIsRiskyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRisky, true
}

// SetIsRisky sets field value
func (o *UserRiskEvaluation) SetIsRisky(v bool) {
	o.IsRisky = v
}

// GetIsVerified returns the IsVerified field value
func (o *UserRiskEvaluation) GetIsVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsVerified
}

// GetIsVerifiedOk returns a tuple with the IsVerified field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetIsVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVerified, true
}

// SetIsVerified sets field value
func (o *UserRiskEvaluation) SetIsVerified(v bool) {
	o.IsVerified = v
}

// GetRiskScore returns the RiskScore field value
func (o *UserRiskEvaluation) GetRiskScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetRiskScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskScore, true
}

// SetRiskScore sets field value
func (o *UserRiskEvaluation) SetRiskScore(v float32) {
	o.RiskScore = v
}

// GetCompletedSteps returns the CompletedSteps field value
func (o *UserRiskEvaluation) GetCompletedSteps() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CompletedSteps
}

// GetCompletedStepsOk returns a tuple with the CompletedSteps field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetCompletedStepsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedSteps, true
}

// SetCompletedSteps sets field value
func (o *UserRiskEvaluation) SetCompletedSteps(v float32) {
	o.CompletedSteps = v
}

// GetExpectedSteps returns the ExpectedSteps field value
func (o *UserRiskEvaluation) GetExpectedSteps() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpectedSteps
}

// GetExpectedStepsOk returns a tuple with the ExpectedSteps field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetExpectedStepsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpectedSteps, true
}

// SetExpectedSteps sets field value
func (o *UserRiskEvaluation) SetExpectedSteps(v float32) {
	o.ExpectedSteps = v
}

// GetGoogleLogin returns the GoogleLogin field value
func (o *UserRiskEvaluation) GetGoogleLogin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GoogleLogin
}

// GetGoogleLoginOk returns a tuple with the GoogleLogin field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetGoogleLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleLogin, true
}

// SetGoogleLogin sets field value
func (o *UserRiskEvaluation) SetGoogleLogin(v bool) {
	o.GoogleLogin = v
}

// GetGithubLogin returns the GithubLogin field value
func (o *UserRiskEvaluation) GetGithubLogin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GithubLogin
}

// GetGithubLoginOk returns a tuple with the GithubLogin field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetGithubLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GithubLogin, true
}

// SetGithubLogin sets field value
func (o *UserRiskEvaluation) SetGithubLogin(v bool) {
	o.GithubLogin = v
}

// GetEmailVerified returns the EmailVerified field value
func (o *UserRiskEvaluation) GetEmailVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailVerified, true
}

// SetEmailVerified sets field value
func (o *UserRiskEvaluation) SetEmailVerified(v bool) {
	o.EmailVerified = v
}

// GetActiveDaysRemaining returns the ActiveDaysRemaining field value
func (o *UserRiskEvaluation) GetActiveDaysRemaining() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ActiveDaysRemaining
}

// GetActiveDaysRemainingOk returns a tuple with the ActiveDaysRemaining field value
// and a boolean to check if the value has been set.
func (o *UserRiskEvaluation) GetActiveDaysRemainingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveDaysRemaining, true
}

// SetActiveDaysRemaining sets field value
func (o *UserRiskEvaluation) SetActiveDaysRemaining(v float32) {
	o.ActiveDaysRemaining = v
}

func (o UserRiskEvaluation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskEvaluation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wasRisky"] = o.WasRisky
	toSerialize["isRisky"] = o.IsRisky
	toSerialize["isVerified"] = o.IsVerified
	toSerialize["riskScore"] = o.RiskScore
	toSerialize["completedSteps"] = o.CompletedSteps
	toSerialize["expectedSteps"] = o.ExpectedSteps
	toSerialize["googleLogin"] = o.GoogleLogin
	toSerialize["githubLogin"] = o.GithubLogin
	toSerialize["emailVerified"] = o.EmailVerified
	toSerialize["activeDaysRemaining"] = o.ActiveDaysRemaining
	return toSerialize, nil
}

type NullableUserRiskEvaluation struct {
	value *UserRiskEvaluation
	isSet bool
}

func (v NullableUserRiskEvaluation) Get() *UserRiskEvaluation {
	return v.value
}

func (v *NullableUserRiskEvaluation) Set(val *UserRiskEvaluation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskEvaluation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskEvaluation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskEvaluation(val *UserRiskEvaluation) *NullableUserRiskEvaluation {
	return &NullableUserRiskEvaluation{value: val, isSet: true}
}

func (v NullableUserRiskEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskEvaluation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
