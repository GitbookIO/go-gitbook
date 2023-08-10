# SearchAIQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**PreviousQueries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchAIQuery

`func NewSearchAIQuery(query string, ) *SearchAIQuery`

NewSearchAIQuery instantiates a new SearchAIQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAIQueryWithDefaults

`func NewSearchAIQueryWithDefaults() *SearchAIQuery`

NewSearchAIQueryWithDefaults instantiates a new SearchAIQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchAIQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchAIQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchAIQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetPreviousQueries

`func (o *SearchAIQuery) GetPreviousQueries() []string`

GetPreviousQueries returns the PreviousQueries field if non-nil, zero value otherwise.

### GetPreviousQueriesOk

`func (o *SearchAIQuery) GetPreviousQueriesOk() (*[]string, bool)`

GetPreviousQueriesOk returns a tuple with the PreviousQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousQueries

`func (o *SearchAIQuery) SetPreviousQueries(v []string)`

SetPreviousQueries sets PreviousQueries field to given value.

### HasPreviousQueries

`func (o *SearchAIQuery) HasPreviousQueries() bool`

HasPreviousQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


