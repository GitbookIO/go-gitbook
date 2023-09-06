# StartRecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Optional title describing the recording | [optional] 
**Context** | [**RecordingContext**](RecordingContext.md) |  | 
**ExternalId** | Pointer to **string** | ID in the original source of the recording. | [optional] 
**ExternalURL** | Pointer to **string** | URL of the original source of the recording. | [optional] 

## Methods

### NewStartRecordingRequest

`func NewStartRecordingRequest(context RecordingContext, ) *StartRecordingRequest`

NewStartRecordingRequest instantiates a new StartRecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartRecordingRequestWithDefaults

`func NewStartRecordingRequestWithDefaults() *StartRecordingRequest`

NewStartRecordingRequestWithDefaults instantiates a new StartRecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StartRecordingRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StartRecordingRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StartRecordingRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StartRecordingRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContext

`func (o *StartRecordingRequest) GetContext() RecordingContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StartRecordingRequest) GetContextOk() (*RecordingContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StartRecordingRequest) SetContext(v RecordingContext)`

SetContext sets Context field to given value.


### GetExternalId

`func (o *StartRecordingRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StartRecordingRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StartRecordingRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *StartRecordingRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalURL

`func (o *StartRecordingRequest) GetExternalURL() string`

GetExternalURL returns the ExternalURL field if non-nil, zero value otherwise.

### GetExternalURLOk

`func (o *StartRecordingRequest) GetExternalURLOk() (*string, bool)`

GetExternalURLOk returns a tuple with the ExternalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalURL

`func (o *StartRecordingRequest) SetExternalURL(v string)`

SetExternalURL sets ExternalURL field to given value.

### HasExternalURL

`func (o *StartRecordingRequest) HasExternalURL() bool`

HasExternalURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


