# SpaceInstallationSetupEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Previous** | Pointer to [**SpaceInstallationSetupEventAllOfPrevious**](SpaceInstallationSetupEventAllOfPrevious.md) |  | [optional] 

## Methods

### NewSpaceInstallationSetupEventAllOf

`func NewSpaceInstallationSetupEventAllOf(type_ string, status IntegrationInstallationStatus, ) *SpaceInstallationSetupEventAllOf`

NewSpaceInstallationSetupEventAllOf instantiates a new SpaceInstallationSetupEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceInstallationSetupEventAllOfWithDefaults

`func NewSpaceInstallationSetupEventAllOfWithDefaults() *SpaceInstallationSetupEventAllOf`

NewSpaceInstallationSetupEventAllOfWithDefaults instantiates a new SpaceInstallationSetupEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SpaceInstallationSetupEventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceInstallationSetupEventAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceInstallationSetupEventAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *SpaceInstallationSetupEventAllOf) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpaceInstallationSetupEventAllOf) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpaceInstallationSetupEventAllOf) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetPrevious

`func (o *SpaceInstallationSetupEventAllOf) GetPrevious() SpaceInstallationSetupEventAllOfPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *SpaceInstallationSetupEventAllOf) GetPreviousOk() (*SpaceInstallationSetupEventAllOfPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *SpaceInstallationSetupEventAllOf) SetPrevious(v SpaceInstallationSetupEventAllOfPrevious)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *SpaceInstallationSetupEventAllOf) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


