# IntegrationEnvironmentApiTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | **string** | API authentication token representing the integration. | 
**Installation** | Pointer to **string** | API authentication token representing the current installation. | [optional] 

## Methods

### NewIntegrationEnvironmentApiTokens

`func NewIntegrationEnvironmentApiTokens(integration string, ) *IntegrationEnvironmentApiTokens`

NewIntegrationEnvironmentApiTokens instantiates a new IntegrationEnvironmentApiTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEnvironmentApiTokensWithDefaults

`func NewIntegrationEnvironmentApiTokensWithDefaults() *IntegrationEnvironmentApiTokens`

NewIntegrationEnvironmentApiTokensWithDefaults instantiates a new IntegrationEnvironmentApiTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *IntegrationEnvironmentApiTokens) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationEnvironmentApiTokens) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationEnvironmentApiTokens) SetIntegration(v string)`

SetIntegration sets Integration field to given value.


### GetInstallation

`func (o *IntegrationEnvironmentApiTokens) GetInstallation() string`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *IntegrationEnvironmentApiTokens) GetInstallationOk() (*string, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *IntegrationEnvironmentApiTokens) SetInstallation(v string)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *IntegrationEnvironmentApiTokens) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


