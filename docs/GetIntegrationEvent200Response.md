# GetIntegrationEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**IntegrationEvent**](IntegrationEvent.md) |  | 
**Trace** | Pointer to [**IntegrationEventTrace**](IntegrationEventTrace.md) |  | [optional] 

## Methods

### NewGetIntegrationEvent200Response

`func NewGetIntegrationEvent200Response(event IntegrationEvent, ) *GetIntegrationEvent200Response`

NewGetIntegrationEvent200Response instantiates a new GetIntegrationEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrationEvent200ResponseWithDefaults

`func NewGetIntegrationEvent200ResponseWithDefaults() *GetIntegrationEvent200Response`

NewGetIntegrationEvent200ResponseWithDefaults instantiates a new GetIntegrationEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *GetIntegrationEvent200Response) GetEvent() IntegrationEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GetIntegrationEvent200Response) GetEventOk() (*IntegrationEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GetIntegrationEvent200Response) SetEvent(v IntegrationEvent)`

SetEvent sets Event field to given value.


### GetTrace

`func (o *GetIntegrationEvent200Response) GetTrace() IntegrationEventTrace`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *GetIntegrationEvent200Response) GetTraceOk() (*IntegrationEventTrace, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *GetIntegrationEvent200Response) SetTrace(v IntegrationEventTrace)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *GetIntegrationEvent200Response) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


