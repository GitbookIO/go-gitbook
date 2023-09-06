# IntegrationSpaceInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | **string** | Unique name identifier of the integration | 
**Installation** | **string** | ID of the integration installation | 
**Space** | **string** | ID of the space the integration is installed on. | 
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Configuration** | **map[string]interface{}** | Configuration of the integration for this space | 
**ExternalIds** | **[]string** | External IDs assigned by the integration. | 
**Urls** | [**IntegrationSpaceInstallationUrls**](IntegrationSpaceInstallationUrls.md) |  | 

## Methods

### NewIntegrationSpaceInstallation

`func NewIntegrationSpaceInstallation(integration string, installation string, space string, status IntegrationInstallationStatus, configuration map[string]interface{}, externalIds []string, urls IntegrationSpaceInstallationUrls, ) *IntegrationSpaceInstallation`

NewIntegrationSpaceInstallation instantiates a new IntegrationSpaceInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSpaceInstallationWithDefaults

`func NewIntegrationSpaceInstallationWithDefaults() *IntegrationSpaceInstallation`

NewIntegrationSpaceInstallationWithDefaults instantiates a new IntegrationSpaceInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *IntegrationSpaceInstallation) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationSpaceInstallation) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationSpaceInstallation) SetIntegration(v string)`

SetIntegration sets Integration field to given value.


### GetInstallation

`func (o *IntegrationSpaceInstallation) GetInstallation() string`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *IntegrationSpaceInstallation) GetInstallationOk() (*string, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *IntegrationSpaceInstallation) SetInstallation(v string)`

SetInstallation sets Installation field to given value.


### GetSpace

`func (o *IntegrationSpaceInstallation) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IntegrationSpaceInstallation) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IntegrationSpaceInstallation) SetSpace(v string)`

SetSpace sets Space field to given value.


### GetStatus

`func (o *IntegrationSpaceInstallation) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationSpaceInstallation) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationSpaceInstallation) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetConfiguration

`func (o *IntegrationSpaceInstallation) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationSpaceInstallation) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationSpaceInstallation) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.


### GetExternalIds

`func (o *IntegrationSpaceInstallation) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *IntegrationSpaceInstallation) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *IntegrationSpaceInstallation) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.


### GetUrls

`func (o *IntegrationSpaceInstallation) GetUrls() IntegrationSpaceInstallationUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *IntegrationSpaceInstallation) GetUrlsOk() (*IntegrationSpaceInstallationUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *IntegrationSpaceInstallation) SetUrls(v IntegrationSpaceInstallationUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


