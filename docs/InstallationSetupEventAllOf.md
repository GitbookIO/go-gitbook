# InstallationSetupEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Previous** | Pointer to [**InstallationSetupEventAllOfPrevious**](InstallationSetupEventAllOfPrevious.md) |  | [optional] 

## Methods

### NewInstallationSetupEventAllOf

`func NewInstallationSetupEventAllOf(type_ string, status IntegrationInstallationStatus, ) *InstallationSetupEventAllOf`

NewInstallationSetupEventAllOf instantiates a new InstallationSetupEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationSetupEventAllOfWithDefaults

`func NewInstallationSetupEventAllOfWithDefaults() *InstallationSetupEventAllOf`

NewInstallationSetupEventAllOfWithDefaults instantiates a new InstallationSetupEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InstallationSetupEventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstallationSetupEventAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstallationSetupEventAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *InstallationSetupEventAllOf) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstallationSetupEventAllOf) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstallationSetupEventAllOf) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetPrevious

`func (o *InstallationSetupEventAllOf) GetPrevious() InstallationSetupEventAllOfPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *InstallationSetupEventAllOf) GetPreviousOk() (*InstallationSetupEventAllOfPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *InstallationSetupEventAllOf) SetPrevious(v InstallationSetupEventAllOfPrevious)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *InstallationSetupEventAllOf) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


