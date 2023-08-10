# FetchEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | Pointer to **string** | The space ID, if requests are specific to a single space | [optional] 
**InstallationId** | Pointer to **string** | The installation ID, if requests are specific to a single installation | [optional] 
**Auth** | Pointer to [**FetchEventAllOfAuth**](FetchEventAllOfAuth.md) |  | [optional] 
**Type** | **string** |  | 
**Request** | [**FetchRequest**](FetchRequest.md) |  | 

## Methods

### NewFetchEventAllOf

`func NewFetchEventAllOf(type_ string, request FetchRequest, ) *FetchEventAllOf`

NewFetchEventAllOf instantiates a new FetchEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchEventAllOfWithDefaults

`func NewFetchEventAllOfWithDefaults() *FetchEventAllOf`

NewFetchEventAllOfWithDefaults instantiates a new FetchEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *FetchEventAllOf) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *FetchEventAllOf) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *FetchEventAllOf) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *FetchEventAllOf) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetInstallationId

`func (o *FetchEventAllOf) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *FetchEventAllOf) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *FetchEventAllOf) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *FetchEventAllOf) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetAuth

`func (o *FetchEventAllOf) GetAuth() FetchEventAllOfAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *FetchEventAllOf) GetAuthOk() (*FetchEventAllOfAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *FetchEventAllOf) SetAuth(v FetchEventAllOfAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *FetchEventAllOf) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetType

`func (o *FetchEventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FetchEventAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FetchEventAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetRequest

`func (o *FetchEventAllOf) GetRequest() FetchRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *FetchEventAllOf) GetRequestOk() (*FetchRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *FetchEventAllOf) SetRequest(v FetchRequest)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


