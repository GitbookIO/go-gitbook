# RequestUpdateIntegrationInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIds** | Pointer to **[]string** | External IDs assigned by the integration. | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | Configuration of the integration at the account level | [optional] 
**SpaceSelection** | Pointer to [**IntegrationInstallationSpaceSelection**](IntegrationInstallationSpaceSelection.md) |  | [optional] 

## Methods

### NewRequestUpdateIntegrationInstallation

`func NewRequestUpdateIntegrationInstallation() *RequestUpdateIntegrationInstallation`

NewRequestUpdateIntegrationInstallation instantiates a new RequestUpdateIntegrationInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUpdateIntegrationInstallationWithDefaults

`func NewRequestUpdateIntegrationInstallationWithDefaults() *RequestUpdateIntegrationInstallation`

NewRequestUpdateIntegrationInstallationWithDefaults instantiates a new RequestUpdateIntegrationInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIds

`func (o *RequestUpdateIntegrationInstallation) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *RequestUpdateIntegrationInstallation) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *RequestUpdateIntegrationInstallation) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *RequestUpdateIntegrationInstallation) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetConfiguration

`func (o *RequestUpdateIntegrationInstallation) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RequestUpdateIntegrationInstallation) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RequestUpdateIntegrationInstallation) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RequestUpdateIntegrationInstallation) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetSpaceSelection

`func (o *RequestUpdateIntegrationInstallation) GetSpaceSelection() IntegrationInstallationSpaceSelection`

GetSpaceSelection returns the SpaceSelection field if non-nil, zero value otherwise.

### GetSpaceSelectionOk

`func (o *RequestUpdateIntegrationInstallation) GetSpaceSelectionOk() (*IntegrationInstallationSpaceSelection, bool)`

GetSpaceSelectionOk returns a tuple with the SpaceSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceSelection

`func (o *RequestUpdateIntegrationInstallation) SetSpaceSelection(v IntegrationInstallationSpaceSelection)`

SetSpaceSelection sets SpaceSelection field to given value.

### HasSpaceSelection

`func (o *RequestUpdateIntegrationInstallation) HasSpaceSelection() bool`

HasSpaceSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


