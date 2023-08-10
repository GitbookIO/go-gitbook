# RequestSpaceTrackPageViewVisitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousId** | **string** | GitBook&#39;s unique identifier of the visitor. | 
**Cookies** | **map[string]string** | The visitors cookies. | 
**UserAgent** | **string** | User-agent of the visitor. | 

## Methods

### NewRequestSpaceTrackPageViewVisitor

`func NewRequestSpaceTrackPageViewVisitor(anonymousId string, cookies map[string]string, userAgent string, ) *RequestSpaceTrackPageViewVisitor`

NewRequestSpaceTrackPageViewVisitor instantiates a new RequestSpaceTrackPageViewVisitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSpaceTrackPageViewVisitorWithDefaults

`func NewRequestSpaceTrackPageViewVisitorWithDefaults() *RequestSpaceTrackPageViewVisitor`

NewRequestSpaceTrackPageViewVisitorWithDefaults instantiates a new RequestSpaceTrackPageViewVisitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousId

`func (o *RequestSpaceTrackPageViewVisitor) GetAnonymousId() string`

GetAnonymousId returns the AnonymousId field if non-nil, zero value otherwise.

### GetAnonymousIdOk

`func (o *RequestSpaceTrackPageViewVisitor) GetAnonymousIdOk() (*string, bool)`

GetAnonymousIdOk returns a tuple with the AnonymousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousId

`func (o *RequestSpaceTrackPageViewVisitor) SetAnonymousId(v string)`

SetAnonymousId sets AnonymousId field to given value.


### GetCookies

`func (o *RequestSpaceTrackPageViewVisitor) GetCookies() map[string]string`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *RequestSpaceTrackPageViewVisitor) GetCookiesOk() (*map[string]string, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *RequestSpaceTrackPageViewVisitor) SetCookies(v map[string]string)`

SetCookies sets Cookies field to given value.


### GetUserAgent

`func (o *RequestSpaceTrackPageViewVisitor) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *RequestSpaceTrackPageViewVisitor) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *RequestSpaceTrackPageViewVisitor) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


