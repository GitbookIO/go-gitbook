# CommentAllOfTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | Pointer to [**CommentAllOfTargetNode**](CommentAllOfTargetNode.md) |  | [optional] 
**ChangeRequest** | Pointer to **string** | The change request containing this comment, if the comment was made inside a change request. | [optional] 
**Review** | Pointer to **string** | The review containing this comment, if this comment was made as part of a review. | [optional] 
**Page** | Pointer to [**RevisionPageBase**](RevisionPageBase.md) | Information about the page, if this comment refers to a specific page. | [optional] 
**Space** | **string** | The space containing this comment. | 
**Revision** | **string** | The revision in which the target can be found in. | 

## Methods

### NewCommentAllOfTarget

`func NewCommentAllOfTarget(space string, revision string, ) *CommentAllOfTarget`

NewCommentAllOfTarget instantiates a new CommentAllOfTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentAllOfTargetWithDefaults

`func NewCommentAllOfTargetWithDefaults() *CommentAllOfTarget`

NewCommentAllOfTargetWithDefaults instantiates a new CommentAllOfTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *CommentAllOfTarget) GetNode() CommentAllOfTargetNode`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *CommentAllOfTarget) GetNodeOk() (*CommentAllOfTargetNode, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *CommentAllOfTarget) SetNode(v CommentAllOfTargetNode)`

SetNode sets Node field to given value.

### HasNode

`func (o *CommentAllOfTarget) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetChangeRequest

`func (o *CommentAllOfTarget) GetChangeRequest() string`

GetChangeRequest returns the ChangeRequest field if non-nil, zero value otherwise.

### GetChangeRequestOk

`func (o *CommentAllOfTarget) GetChangeRequestOk() (*string, bool)`

GetChangeRequestOk returns a tuple with the ChangeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequest

`func (o *CommentAllOfTarget) SetChangeRequest(v string)`

SetChangeRequest sets ChangeRequest field to given value.

### HasChangeRequest

`func (o *CommentAllOfTarget) HasChangeRequest() bool`

HasChangeRequest returns a boolean if a field has been set.

### GetReview

`func (o *CommentAllOfTarget) GetReview() string`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *CommentAllOfTarget) GetReviewOk() (*string, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *CommentAllOfTarget) SetReview(v string)`

SetReview sets Review field to given value.

### HasReview

`func (o *CommentAllOfTarget) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetPage

`func (o *CommentAllOfTarget) GetPage() RevisionPageBase`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CommentAllOfTarget) GetPageOk() (*RevisionPageBase, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CommentAllOfTarget) SetPage(v RevisionPageBase)`

SetPage sets Page field to given value.

### HasPage

`func (o *CommentAllOfTarget) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSpace

`func (o *CommentAllOfTarget) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *CommentAllOfTarget) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *CommentAllOfTarget) SetSpace(v string)`

SetSpace sets Space field to given value.


### GetRevision

`func (o *CommentAllOfTarget) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CommentAllOfTarget) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CommentAllOfTarget) SetRevision(v string)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


