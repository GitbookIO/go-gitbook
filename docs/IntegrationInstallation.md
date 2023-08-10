# IntegrationInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**SpaceSelection** | [**IntegrationInstallationSpaceSelection**](IntegrationInstallationSpaceSelection.md) |  | 
**Configuration** | **map[string]interface{}** | Configuration of the integration at the account level | 
**Urls** | [**IntegrationInstallationUrls**](IntegrationInstallationUrls.md) |  | 
**ExternalIds** | **[]string** | External IDs assigned by the integration. | 
**Target** | [**IntegrationInstallationTarget**](IntegrationInstallationTarget.md) |  | 

## Methods

### NewIntegrationInstallation

`func NewIntegrationInstallation(id string, status IntegrationInstallationStatus, spaceSelection IntegrationInstallationSpaceSelection, configuration map[string]interface{}, urls IntegrationInstallationUrls, externalIds []string, target IntegrationInstallationTarget, ) *IntegrationInstallation`

NewIntegrationInstallation instantiates a new IntegrationInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationInstallationWithDefaults

`func NewIntegrationInstallationWithDefaults() *IntegrationInstallation`

NewIntegrationInstallationWithDefaults instantiates a new IntegrationInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationInstallation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationInstallation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationInstallation) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *IntegrationInstallation) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationInstallation) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationInstallation) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetSpaceSelection

`func (o *IntegrationInstallation) GetSpaceSelection() IntegrationInstallationSpaceSelection`

GetSpaceSelection returns the SpaceSelection field if non-nil, zero value otherwise.

### GetSpaceSelectionOk

`func (o *IntegrationInstallation) GetSpaceSelectionOk() (*IntegrationInstallationSpaceSelection, bool)`

GetSpaceSelectionOk returns a tuple with the SpaceSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceSelection

`func (o *IntegrationInstallation) SetSpaceSelection(v IntegrationInstallationSpaceSelection)`

SetSpaceSelection sets SpaceSelection field to given value.


### GetConfiguration

`func (o *IntegrationInstallation) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationInstallation) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationInstallation) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.


### GetUrls

`func (o *IntegrationInstallation) GetUrls() IntegrationInstallationUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *IntegrationInstallation) GetUrlsOk() (*IntegrationInstallationUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *IntegrationInstallation) SetUrls(v IntegrationInstallationUrls)`

SetUrls sets Urls field to given value.


### GetExternalIds

`func (o *IntegrationInstallation) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *IntegrationInstallation) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *IntegrationInstallation) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.


### GetTarget

`func (o *IntegrationInstallation) GetTarget() IntegrationInstallationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IntegrationInstallation) GetTargetOk() (*IntegrationInstallationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IntegrationInstallation) SetTarget(v IntegrationInstallationTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


