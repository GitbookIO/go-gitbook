# IntegrationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the event. | 
**IntegrationId** | **string** | Unique ID of the integration. | 
**InstallationId** | Pointer to **string** | Unique ID of the integration installation. | [optional] 
**CreatedAt** | **string** |  | 
**Payload** | [**Event**](Event.md) |  | 
**Status** | **string** | Status of the event. | 

## Methods

### NewIntegrationEvent

`func NewIntegrationEvent(id string, integrationId string, createdAt string, payload Event, status string, ) *IntegrationEvent`

NewIntegrationEvent instantiates a new IntegrationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEventWithDefaults

`func NewIntegrationEventWithDefaults() *IntegrationEvent`

NewIntegrationEventWithDefaults instantiates a new IntegrationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationEvent) SetId(v string)`

SetId sets Id field to given value.


### GetIntegrationId

`func (o *IntegrationEvent) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationEvent) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *IntegrationEvent) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetInstallationId

`func (o *IntegrationEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *IntegrationEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *IntegrationEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *IntegrationEvent) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IntegrationEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPayload

`func (o *IntegrationEvent) GetPayload() Event`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IntegrationEvent) GetPayloadOk() (*Event, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IntegrationEvent) SetPayload(v Event)`

SetPayload sets Payload field to given value.


### GetStatus

`func (o *IntegrationEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationEvent) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


