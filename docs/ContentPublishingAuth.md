# ContentPublishingAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;publishing-auth\&quot; | 
**Id** | **string** | Unique identifier for the content. | 
**PrivateKey** | **string** | Private key used to sign JWT tokens. | 
**FallbackURL** | Pointer to **string** | URL to redirect to when the visitor auth secret is invalid. | [optional] 

## Methods

### NewContentPublishingAuth

`func NewContentPublishingAuth(object string, id string, privateKey string, ) *ContentPublishingAuth`

NewContentPublishingAuth instantiates a new ContentPublishingAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPublishingAuthWithDefaults

`func NewContentPublishingAuthWithDefaults() *ContentPublishingAuth`

NewContentPublishingAuthWithDefaults instantiates a new ContentPublishingAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ContentPublishingAuth) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ContentPublishingAuth) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ContentPublishingAuth) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *ContentPublishingAuth) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentPublishingAuth) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentPublishingAuth) SetId(v string)`

SetId sets Id field to given value.


### GetPrivateKey

`func (o *ContentPublishingAuth) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ContentPublishingAuth) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ContentPublishingAuth) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetFallbackURL

`func (o *ContentPublishingAuth) GetFallbackURL() string`

GetFallbackURL returns the FallbackURL field if non-nil, zero value otherwise.

### GetFallbackURLOk

`func (o *ContentPublishingAuth) GetFallbackURLOk() (*string, bool)`

GetFallbackURLOk returns a tuple with the FallbackURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackURL

`func (o *ContentPublishingAuth) SetFallbackURL(v string)`

SetFallbackURL sets FallbackURL field to given value.

### HasFallbackURL

`func (o *ContentPublishingAuth) HasFallbackURL() bool`

HasFallbackURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


