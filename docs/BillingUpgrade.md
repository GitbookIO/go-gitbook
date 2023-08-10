# BillingUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** |  | 
**SessionId** | **string** | Stripe payment session ID | 
**Invoice** | [**BillingInvoicePreview**](BillingInvoicePreview.md) |  | 

## Methods

### NewBillingUpgrade

`func NewBillingUpgrade(result string, sessionId string, invoice BillingInvoicePreview, ) *BillingUpgrade`

NewBillingUpgrade instantiates a new BillingUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingUpgradeWithDefaults

`func NewBillingUpgradeWithDefaults() *BillingUpgrade`

NewBillingUpgradeWithDefaults instantiates a new BillingUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *BillingUpgrade) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BillingUpgrade) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BillingUpgrade) SetResult(v string)`

SetResult sets Result field to given value.


### GetSessionId

`func (o *BillingUpgrade) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *BillingUpgrade) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *BillingUpgrade) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetInvoice

`func (o *BillingUpgrade) GetInvoice() BillingInvoicePreview`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *BillingUpgrade) GetInvoiceOk() (*BillingInvoicePreview, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *BillingUpgrade) SetInvoice(v BillingInvoicePreview)`

SetInvoice sets Invoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


