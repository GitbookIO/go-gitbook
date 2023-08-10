# BillingInvoicePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the invoice | 
**AmountDueToday** | **float32** | Amount that will be immediately charged. | 
**CustomerBalance** | **float32** | Current balance, if any, being stored on the customer. If positive, the customer has credit to apply to their next invoice. | 
**RemainingCustomerBalance** | **float32** | Current balance after potential upgrade. | 
**Lines** | [**[]BillingInvoicePreviewLinesInner**](BillingInvoicePreviewLinesInner.md) | Details of the change happening on the subscription. | 

## Methods

### NewBillingInvoicePreview

`func NewBillingInvoicePreview(amount float32, amountDueToday float32, customerBalance float32, remainingCustomerBalance float32, lines []BillingInvoicePreviewLinesInner, ) *BillingInvoicePreview`

NewBillingInvoicePreview instantiates a new BillingInvoicePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoicePreviewWithDefaults

`func NewBillingInvoicePreviewWithDefaults() *BillingInvoicePreview`

NewBillingInvoicePreviewWithDefaults instantiates a new BillingInvoicePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingInvoicePreview) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingInvoicePreview) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingInvoicePreview) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAmountDueToday

`func (o *BillingInvoicePreview) GetAmountDueToday() float32`

GetAmountDueToday returns the AmountDueToday field if non-nil, zero value otherwise.

### GetAmountDueTodayOk

`func (o *BillingInvoicePreview) GetAmountDueTodayOk() (*float32, bool)`

GetAmountDueTodayOk returns a tuple with the AmountDueToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueToday

`func (o *BillingInvoicePreview) SetAmountDueToday(v float32)`

SetAmountDueToday sets AmountDueToday field to given value.


### GetCustomerBalance

`func (o *BillingInvoicePreview) GetCustomerBalance() float32`

GetCustomerBalance returns the CustomerBalance field if non-nil, zero value otherwise.

### GetCustomerBalanceOk

`func (o *BillingInvoicePreview) GetCustomerBalanceOk() (*float32, bool)`

GetCustomerBalanceOk returns a tuple with the CustomerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBalance

`func (o *BillingInvoicePreview) SetCustomerBalance(v float32)`

SetCustomerBalance sets CustomerBalance field to given value.


### GetRemainingCustomerBalance

`func (o *BillingInvoicePreview) GetRemainingCustomerBalance() float32`

GetRemainingCustomerBalance returns the RemainingCustomerBalance field if non-nil, zero value otherwise.

### GetRemainingCustomerBalanceOk

`func (o *BillingInvoicePreview) GetRemainingCustomerBalanceOk() (*float32, bool)`

GetRemainingCustomerBalanceOk returns a tuple with the RemainingCustomerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCustomerBalance

`func (o *BillingInvoicePreview) SetRemainingCustomerBalance(v float32)`

SetRemainingCustomerBalance sets RemainingCustomerBalance field to given value.


### GetLines

`func (o *BillingInvoicePreview) GetLines() []BillingInvoicePreviewLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *BillingInvoicePreview) GetLinesOk() (*[]BillingInvoicePreviewLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *BillingInvoicePreview) SetLines(v []BillingInvoicePreviewLinesInner)`

SetLines sets Lines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


