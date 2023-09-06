# ContentKitSelectOptionsOneOfUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the URL along with the port number if required. | 
**Pathname** | **string** | Path of the URL prefixed with a &#x60;/&#x60;. | 
**Query** | Pointer to [**map[string]ContentKitWebFrameDataValue**](ContentKitWebFrameDataValue.md) |  | [optional] 

## Methods

### NewContentKitSelectOptionsOneOfUrl

`func NewContentKitSelectOptionsOneOfUrl(host string, pathname string, ) *ContentKitSelectOptionsOneOfUrl`

NewContentKitSelectOptionsOneOfUrl instantiates a new ContentKitSelectOptionsOneOfUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitSelectOptionsOneOfUrlWithDefaults

`func NewContentKitSelectOptionsOneOfUrlWithDefaults() *ContentKitSelectOptionsOneOfUrl`

NewContentKitSelectOptionsOneOfUrlWithDefaults instantiates a new ContentKitSelectOptionsOneOfUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ContentKitSelectOptionsOneOfUrl) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ContentKitSelectOptionsOneOfUrl) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ContentKitSelectOptionsOneOfUrl) SetHost(v string)`

SetHost sets Host field to given value.


### GetPathname

`func (o *ContentKitSelectOptionsOneOfUrl) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *ContentKitSelectOptionsOneOfUrl) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *ContentKitSelectOptionsOneOfUrl) SetPathname(v string)`

SetPathname sets Pathname field to given value.


### GetQuery

`func (o *ContentKitSelectOptionsOneOfUrl) GetQuery() map[string]ContentKitWebFrameDataValue`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ContentKitSelectOptionsOneOfUrl) GetQueryOk() (*map[string]ContentKitWebFrameDataValue, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ContentKitSelectOptionsOneOfUrl) SetQuery(v map[string]ContentKitWebFrameDataValue)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ContentKitSelectOptionsOneOfUrl) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


