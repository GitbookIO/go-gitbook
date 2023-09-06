# ContentKitLinkTargetUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the URL along with the port number if required. | 
**Pathname** | **string** | Path of the URL prefixed with a &#x60;/&#x60;. | 
**Query** | Pointer to [**map[string]ContentKitWebFrameDataValue**](ContentKitWebFrameDataValue.md) |  | [optional] 

## Methods

### NewContentKitLinkTargetUrl

`func NewContentKitLinkTargetUrl(host string, pathname string, ) *ContentKitLinkTargetUrl`

NewContentKitLinkTargetUrl instantiates a new ContentKitLinkTargetUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitLinkTargetUrlWithDefaults

`func NewContentKitLinkTargetUrlWithDefaults() *ContentKitLinkTargetUrl`

NewContentKitLinkTargetUrlWithDefaults instantiates a new ContentKitLinkTargetUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ContentKitLinkTargetUrl) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ContentKitLinkTargetUrl) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ContentKitLinkTargetUrl) SetHost(v string)`

SetHost sets Host field to given value.


### GetPathname

`func (o *ContentKitLinkTargetUrl) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *ContentKitLinkTargetUrl) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *ContentKitLinkTargetUrl) SetPathname(v string)`

SetPathname sets Pathname field to given value.


### GetQuery

`func (o *ContentKitLinkTargetUrl) GetQuery() map[string]ContentKitWebFrameDataValue`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ContentKitLinkTargetUrl) GetQueryOk() (*map[string]ContentKitWebFrameDataValue, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ContentKitLinkTargetUrl) SetQuery(v map[string]ContentKitWebFrameDataValue)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ContentKitLinkTargetUrl) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


