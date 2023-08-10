# PostCommentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | Pointer to **string** | The node to which the comment is posted, if any. | [optional] 
**Page** | Pointer to **string** | The page to which the comment is posted, if any. | [optional] 
**Body** | [**Document**](Document.md) |  | 

## Methods

### NewPostCommentSchema

`func NewPostCommentSchema(body Document, ) *PostCommentSchema`

NewPostCommentSchema instantiates a new PostCommentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCommentSchemaWithDefaults

`func NewPostCommentSchemaWithDefaults() *PostCommentSchema`

NewPostCommentSchemaWithDefaults instantiates a new PostCommentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *PostCommentSchema) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *PostCommentSchema) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *PostCommentSchema) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *PostCommentSchema) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPage

`func (o *PostCommentSchema) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PostCommentSchema) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PostCommentSchema) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *PostCommentSchema) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetBody

`func (o *PostCommentSchema) GetBody() Document`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PostCommentSchema) GetBodyOk() (*Document, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PostCommentSchema) SetBody(v Document)`

SetBody sets Body field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


