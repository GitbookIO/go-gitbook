# IntegrationEventLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The message of the log entry. | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** | The level of the log entry. | [optional] 

## Methods

### NewIntegrationEventLog

`func NewIntegrationEventLog() *IntegrationEventLog`

NewIntegrationEventLog instantiates a new IntegrationEventLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEventLogWithDefaults

`func NewIntegrationEventLogWithDefaults() *IntegrationEventLog`

NewIntegrationEventLogWithDefaults instantiates a new IntegrationEventLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *IntegrationEventLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IntegrationEventLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IntegrationEventLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IntegrationEventLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *IntegrationEventLog) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IntegrationEventLog) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IntegrationEventLog) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *IntegrationEventLog) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLevel

`func (o *IntegrationEventLog) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *IntegrationEventLog) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *IntegrationEventLog) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *IntegrationEventLog) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


