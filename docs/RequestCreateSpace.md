# RequestCreateSpace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Emoji** | Pointer to **string** | Unicode codepoint or character of the emoji | [optional] 
**Private** | Pointer to **bool** | Private spaces are no longer supported by GitBook. | [optional] 
**Type** | Pointer to [**SpaceType**](SpaceType.md) |  | [optional] 
**Parent** | Pointer to **string** | ID of a parent collection | [optional] 

## Methods

### NewRequestCreateSpace

`func NewRequestCreateSpace() *RequestCreateSpace`

NewRequestCreateSpace instantiates a new RequestCreateSpace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateSpaceWithDefaults

`func NewRequestCreateSpaceWithDefaults() *RequestCreateSpace`

NewRequestCreateSpaceWithDefaults instantiates a new RequestCreateSpace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RequestCreateSpace) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestCreateSpace) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestCreateSpace) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RequestCreateSpace) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEmoji

`func (o *RequestCreateSpace) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *RequestCreateSpace) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *RequestCreateSpace) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *RequestCreateSpace) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetPrivate

`func (o *RequestCreateSpace) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *RequestCreateSpace) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *RequestCreateSpace) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *RequestCreateSpace) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetType

`func (o *RequestCreateSpace) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestCreateSpace) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestCreateSpace) SetType(v SpaceType)`

SetType sets Type field to given value.

### HasType

`func (o *RequestCreateSpace) HasType() bool`

HasType returns a boolean if a field has been set.

### GetParent

`func (o *RequestCreateSpace) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RequestCreateSpace) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RequestCreateSpace) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RequestCreateSpace) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


