# IntegrationUrls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **string** | URL of the icon associated to the integration | [optional] 
**App** | **string** | URL of the integration in the application | 
**Assets** | **string** | URL of the integration&#39;s assets. | 
**PublicEndpoint** | **string** | Public HTTP endpoint for the integration | 

## Methods

### NewIntegrationUrls

`func NewIntegrationUrls(app string, assets string, publicEndpoint string, ) *IntegrationUrls`

NewIntegrationUrls instantiates a new IntegrationUrls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationUrlsWithDefaults

`func NewIntegrationUrlsWithDefaults() *IntegrationUrls`

NewIntegrationUrlsWithDefaults instantiates a new IntegrationUrls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *IntegrationUrls) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IntegrationUrls) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IntegrationUrls) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IntegrationUrls) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetApp

`func (o *IntegrationUrls) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *IntegrationUrls) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *IntegrationUrls) SetApp(v string)`

SetApp sets App field to given value.


### GetAssets

`func (o *IntegrationUrls) GetAssets() string`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *IntegrationUrls) GetAssetsOk() (*string, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *IntegrationUrls) SetAssets(v string)`

SetAssets sets Assets field to given value.


### GetPublicEndpoint

`func (o *IntegrationUrls) GetPublicEndpoint() string`

GetPublicEndpoint returns the PublicEndpoint field if non-nil, zero value otherwise.

### GetPublicEndpointOk

`func (o *IntegrationUrls) GetPublicEndpointOk() (*string, bool)`

GetPublicEndpointOk returns a tuple with the PublicEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpoint

`func (o *IntegrationUrls) SetPublicEndpoint(v string)`

SetPublicEndpoint sets PublicEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


