# IntegrationEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | Pointer to **string** | Authentication token to use with the HTTP API. Depending on the context, the token might be representing the installation or the integration. | [optional] 
**Integration** | [**Integration**](Integration.md) |  | 
**Installation** | Pointer to [**IntegrationInstallation**](IntegrationInstallation.md) |  | [optional] 
**SpaceInstallation** | Pointer to [**IntegrationSpaceInstallation**](IntegrationSpaceInstallation.md) |  | [optional] 
**Secrets** | **map[string]string** | Secrets stored on the integration and passed at runtime. | 
**ApiEndpoint** | **string** | URL of the HTTP API | 
**ApiTokens** | [**IntegrationEnvironmentApiTokens**](IntegrationEnvironmentApiTokens.md) |  | 

## Methods

### NewIntegrationEnvironment

`func NewIntegrationEnvironment(integration Integration, secrets map[string]string, apiEndpoint string, apiTokens IntegrationEnvironmentApiTokens, ) *IntegrationEnvironment`

NewIntegrationEnvironment instantiates a new IntegrationEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEnvironmentWithDefaults

`func NewIntegrationEnvironmentWithDefaults() *IntegrationEnvironment`

NewIntegrationEnvironmentWithDefaults instantiates a new IntegrationEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *IntegrationEnvironment) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *IntegrationEnvironment) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *IntegrationEnvironment) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *IntegrationEnvironment) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetIntegration

`func (o *IntegrationEnvironment) GetIntegration() Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationEnvironment) GetIntegrationOk() (*Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationEnvironment) SetIntegration(v Integration)`

SetIntegration sets Integration field to given value.


### GetInstallation

`func (o *IntegrationEnvironment) GetInstallation() IntegrationInstallation`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *IntegrationEnvironment) GetInstallationOk() (*IntegrationInstallation, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *IntegrationEnvironment) SetInstallation(v IntegrationInstallation)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *IntegrationEnvironment) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetSpaceInstallation

`func (o *IntegrationEnvironment) GetSpaceInstallation() IntegrationSpaceInstallation`

GetSpaceInstallation returns the SpaceInstallation field if non-nil, zero value otherwise.

### GetSpaceInstallationOk

`func (o *IntegrationEnvironment) GetSpaceInstallationOk() (*IntegrationSpaceInstallation, bool)`

GetSpaceInstallationOk returns a tuple with the SpaceInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceInstallation

`func (o *IntegrationEnvironment) SetSpaceInstallation(v IntegrationSpaceInstallation)`

SetSpaceInstallation sets SpaceInstallation field to given value.

### HasSpaceInstallation

`func (o *IntegrationEnvironment) HasSpaceInstallation() bool`

HasSpaceInstallation returns a boolean if a field has been set.

### GetSecrets

`func (o *IntegrationEnvironment) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *IntegrationEnvironment) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *IntegrationEnvironment) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.


### GetApiEndpoint

`func (o *IntegrationEnvironment) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *IntegrationEnvironment) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *IntegrationEnvironment) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.


### GetApiTokens

`func (o *IntegrationEnvironment) GetApiTokens() IntegrationEnvironmentApiTokens`

GetApiTokens returns the ApiTokens field if non-nil, zero value otherwise.

### GetApiTokensOk

`func (o *IntegrationEnvironment) GetApiTokensOk() (*IntegrationEnvironmentApiTokens, bool)`

GetApiTokensOk returns a tuple with the ApiTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokens

`func (o *IntegrationEnvironment) SetApiTokens(v IntegrationEnvironmentApiTokens)`

SetApiTokens sets ApiTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


