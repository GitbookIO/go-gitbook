# EntityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the entity in GitBook | 
**Type** | **string** | Type of an entity | 
**Urls** | [**EntityAllOfUrls**](EntityAllOfUrls.md) |  | 

## Methods

### NewEntityAllOf

`func NewEntityAllOf(id string, type_ string, urls EntityAllOfUrls, ) *EntityAllOf`

NewEntityAllOf instantiates a new EntityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityAllOfWithDefaults

`func NewEntityAllOfWithDefaults() *EntityAllOf`

NewEntityAllOfWithDefaults instantiates a new EntityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *EntityAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetUrls

`func (o *EntityAllOf) GetUrls() EntityAllOfUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *EntityAllOf) GetUrlsOk() (*EntityAllOfUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *EntityAllOf) SetUrls(v EntityAllOfUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


