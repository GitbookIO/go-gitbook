# RequestUpgradeOrganizationBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**BillingProduct**](BillingProduct.md) |  | 
**Interval** | [**BillingInterval**](BillingInterval.md) |  | 
**Reason** | Pointer to **string** | Reason that triggered the billing upgrade | [optional] 
**Mode** | Pointer to **string** | Mode to use for the upgrade (default value is &#x60;commit&#x60;): - &#x60;auto&#x60;: user is redirect to checkout if possible, other a preview of the auto-upgrade is returned. - &#x60;commit&#x60;: a checkout session is returned or an auto-upgrade is done - &#x60;preview&#x60;: a preview invoice is always returned  | [optional] 

## Methods

### NewRequestUpgradeOrganizationBilling

`func NewRequestUpgradeOrganizationBilling(product BillingProduct, interval BillingInterval, ) *RequestUpgradeOrganizationBilling`

NewRequestUpgradeOrganizationBilling instantiates a new RequestUpgradeOrganizationBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUpgradeOrganizationBillingWithDefaults

`func NewRequestUpgradeOrganizationBillingWithDefaults() *RequestUpgradeOrganizationBilling`

NewRequestUpgradeOrganizationBillingWithDefaults instantiates a new RequestUpgradeOrganizationBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *RequestUpgradeOrganizationBilling) GetProduct() BillingProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *RequestUpgradeOrganizationBilling) GetProductOk() (*BillingProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *RequestUpgradeOrganizationBilling) SetProduct(v BillingProduct)`

SetProduct sets Product field to given value.


### GetInterval

`func (o *RequestUpgradeOrganizationBilling) GetInterval() BillingInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RequestUpgradeOrganizationBilling) GetIntervalOk() (*BillingInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RequestUpgradeOrganizationBilling) SetInterval(v BillingInterval)`

SetInterval sets Interval field to given value.


### GetReason

`func (o *RequestUpgradeOrganizationBilling) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RequestUpgradeOrganizationBilling) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RequestUpgradeOrganizationBilling) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RequestUpgradeOrganizationBilling) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMode

`func (o *RequestUpgradeOrganizationBilling) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RequestUpgradeOrganizationBilling) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RequestUpgradeOrganizationBilling) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RequestUpgradeOrganizationBilling) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


