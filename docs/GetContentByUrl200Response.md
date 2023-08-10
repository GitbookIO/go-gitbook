# GetContentByUrl200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | [**Collection**](Collection.md) |  | 
**Space** | [**Space**](Space.md) |  | 
**ChangeRequest** | Pointer to [**ChangeRequest**](ChangeRequest.md) |  | [optional] 
**Page** | Pointer to [**GetPageByPath200Response**](GetPageByPath200Response.md) |  | [optional] 

## Methods

### NewGetContentByUrl200Response

`func NewGetContentByUrl200Response(collection Collection, space Space, ) *GetContentByUrl200Response`

NewGetContentByUrl200Response instantiates a new GetContentByUrl200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContentByUrl200ResponseWithDefaults

`func NewGetContentByUrl200ResponseWithDefaults() *GetContentByUrl200Response`

NewGetContentByUrl200ResponseWithDefaults instantiates a new GetContentByUrl200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *GetContentByUrl200Response) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *GetContentByUrl200Response) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *GetContentByUrl200Response) SetCollection(v Collection)`

SetCollection sets Collection field to given value.


### GetSpace

`func (o *GetContentByUrl200Response) GetSpace() Space`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *GetContentByUrl200Response) GetSpaceOk() (*Space, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *GetContentByUrl200Response) SetSpace(v Space)`

SetSpace sets Space field to given value.


### GetChangeRequest

`func (o *GetContentByUrl200Response) GetChangeRequest() ChangeRequest`

GetChangeRequest returns the ChangeRequest field if non-nil, zero value otherwise.

### GetChangeRequestOk

`func (o *GetContentByUrl200Response) GetChangeRequestOk() (*ChangeRequest, bool)`

GetChangeRequestOk returns a tuple with the ChangeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequest

`func (o *GetContentByUrl200Response) SetChangeRequest(v ChangeRequest)`

SetChangeRequest sets ChangeRequest field to given value.

### HasChangeRequest

`func (o *GetContentByUrl200Response) HasChangeRequest() bool`

HasChangeRequest returns a boolean if a field has been set.

### GetPage

`func (o *GetContentByUrl200Response) GetPage() GetPageByPath200Response`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetContentByUrl200Response) GetPageOk() (*GetPageByPath200Response, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetContentByUrl200Response) SetPage(v GetPageByPath200Response)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetContentByUrl200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


