# RequestSpaceTrackPageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageId** | **string** | Unique identifier of the page. | 
**Visitor** | [**RequestSpaceTrackPageViewVisitor**](RequestSpaceTrackPageViewVisitor.md) |  | 
**Url** | **string** | The GitBook content&#39;s URL visited (including URL params). | 
**Referrer** | **string** | The URL of referrer that linked to the page. | 

## Methods

### NewRequestSpaceTrackPageView

`func NewRequestSpaceTrackPageView(pageId string, visitor RequestSpaceTrackPageViewVisitor, url string, referrer string, ) *RequestSpaceTrackPageView`

NewRequestSpaceTrackPageView instantiates a new RequestSpaceTrackPageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSpaceTrackPageViewWithDefaults

`func NewRequestSpaceTrackPageViewWithDefaults() *RequestSpaceTrackPageView`

NewRequestSpaceTrackPageViewWithDefaults instantiates a new RequestSpaceTrackPageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageId

`func (o *RequestSpaceTrackPageView) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *RequestSpaceTrackPageView) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *RequestSpaceTrackPageView) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetVisitor

`func (o *RequestSpaceTrackPageView) GetVisitor() RequestSpaceTrackPageViewVisitor`

GetVisitor returns the Visitor field if non-nil, zero value otherwise.

### GetVisitorOk

`func (o *RequestSpaceTrackPageView) GetVisitorOk() (*RequestSpaceTrackPageViewVisitor, bool)`

GetVisitorOk returns a tuple with the Visitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitor

`func (o *RequestSpaceTrackPageView) SetVisitor(v RequestSpaceTrackPageViewVisitor)`

SetVisitor sets Visitor field to given value.


### GetUrl

`func (o *RequestSpaceTrackPageView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestSpaceTrackPageView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestSpaceTrackPageView) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReferrer

`func (o *RequestSpaceTrackPageView) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *RequestSpaceTrackPageView) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *RequestSpaceTrackPageView) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


