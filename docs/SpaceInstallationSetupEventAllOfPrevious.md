# SpaceInstallationSetupEventAllOfPrevious

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Configuration** | Pointer to **map[string]interface{}** | The previous configuration of the Space installation. | [optional] 

## Methods

### NewSpaceInstallationSetupEventAllOfPrevious

`func NewSpaceInstallationSetupEventAllOfPrevious(status IntegrationInstallationStatus, ) *SpaceInstallationSetupEventAllOfPrevious`

NewSpaceInstallationSetupEventAllOfPrevious instantiates a new SpaceInstallationSetupEventAllOfPrevious object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceInstallationSetupEventAllOfPreviousWithDefaults

`func NewSpaceInstallationSetupEventAllOfPreviousWithDefaults() *SpaceInstallationSetupEventAllOfPrevious`

NewSpaceInstallationSetupEventAllOfPreviousWithDefaults instantiates a new SpaceInstallationSetupEventAllOfPrevious object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SpaceInstallationSetupEventAllOfPrevious) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpaceInstallationSetupEventAllOfPrevious) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpaceInstallationSetupEventAllOfPrevious) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetConfiguration

`func (o *SpaceInstallationSetupEventAllOfPrevious) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SpaceInstallationSetupEventAllOfPrevious) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SpaceInstallationSetupEventAllOfPrevious) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SpaceInstallationSetupEventAllOfPrevious) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


