# InstallationSetupEventAllOfPrevious

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Configuration** | Pointer to **map[string]interface{}** | The previous configuration of the installation at the account level. | [optional] 

## Methods

### NewInstallationSetupEventAllOfPrevious

`func NewInstallationSetupEventAllOfPrevious(status IntegrationInstallationStatus, ) *InstallationSetupEventAllOfPrevious`

NewInstallationSetupEventAllOfPrevious instantiates a new InstallationSetupEventAllOfPrevious object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationSetupEventAllOfPreviousWithDefaults

`func NewInstallationSetupEventAllOfPreviousWithDefaults() *InstallationSetupEventAllOfPrevious`

NewInstallationSetupEventAllOfPreviousWithDefaults instantiates a new InstallationSetupEventAllOfPrevious object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InstallationSetupEventAllOfPrevious) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstallationSetupEventAllOfPrevious) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstallationSetupEventAllOfPrevious) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetConfiguration

`func (o *InstallationSetupEventAllOfPrevious) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InstallationSetupEventAllOfPrevious) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InstallationSetupEventAllOfPrevious) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *InstallationSetupEventAllOfPrevious) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


