# SpaceViewEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**PageId** | Pointer to **string** | Unique identifier of the visited page. | [optional] 
**Visitor** | [**SpaceViewEventAllOfVisitor**](SpaceViewEventAllOfVisitor.md) |  | 
**Url** | **string** | The GitBook content&#39;s URL visited (including URL params). | 
**Referrer** | **string** | The URL of referrer that linked to the page. | 

## Methods

### NewSpaceViewEventAllOf

`func NewSpaceViewEventAllOf(type_ string, visitor SpaceViewEventAllOfVisitor, url string, referrer string, ) *SpaceViewEventAllOf`

NewSpaceViewEventAllOf instantiates a new SpaceViewEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceViewEventAllOfWithDefaults

`func NewSpaceViewEventAllOfWithDefaults() *SpaceViewEventAllOf`

NewSpaceViewEventAllOfWithDefaults instantiates a new SpaceViewEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SpaceViewEventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceViewEventAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceViewEventAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetPageId

`func (o *SpaceViewEventAllOf) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *SpaceViewEventAllOf) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *SpaceViewEventAllOf) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *SpaceViewEventAllOf) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetVisitor

`func (o *SpaceViewEventAllOf) GetVisitor() SpaceViewEventAllOfVisitor`

GetVisitor returns the Visitor field if non-nil, zero value otherwise.

### GetVisitorOk

`func (o *SpaceViewEventAllOf) GetVisitorOk() (*SpaceViewEventAllOfVisitor, bool)`

GetVisitorOk returns a tuple with the Visitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitor

`func (o *SpaceViewEventAllOf) SetVisitor(v SpaceViewEventAllOfVisitor)`

SetVisitor sets Visitor field to given value.


### GetUrl

`func (o *SpaceViewEventAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SpaceViewEventAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SpaceViewEventAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReferrer

`func (o *SpaceViewEventAllOf) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *SpaceViewEventAllOf) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *SpaceViewEventAllOf) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


