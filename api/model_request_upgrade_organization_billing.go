/*
GitBook API

The GitBook API

API version: 0.0.1-beta
Contact: support@gitbook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the RequestUpgradeOrganizationBilling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestUpgradeOrganizationBilling{}

// RequestUpgradeOrganizationBilling struct for RequestUpgradeOrganizationBilling
type RequestUpgradeOrganizationBilling struct {
	Product  BillingProduct  `json:"product"`
	Interval BillingInterval `json:"interval"`
	// Reason that triggered the billing upgrade
	Reason *string `json:"reason,omitempty"`
	// Mode to use for the upgrade (default value is `commit`): - `auto`: user is redirect to checkout if possible, other a preview of the auto-upgrade is returned. - `commit`: a checkout session is returned or an auto-upgrade is done - `preview`: a preview invoice is always returned
	Mode *string `json:"mode,omitempty"`
}

// NewRequestUpgradeOrganizationBilling instantiates a new RequestUpgradeOrganizationBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestUpgradeOrganizationBilling(product BillingProduct, interval BillingInterval) *RequestUpgradeOrganizationBilling {
	this := RequestUpgradeOrganizationBilling{}
	this.Product = product
	this.Interval = interval
	return &this
}

// NewRequestUpgradeOrganizationBillingWithDefaults instantiates a new RequestUpgradeOrganizationBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestUpgradeOrganizationBillingWithDefaults() *RequestUpgradeOrganizationBilling {
	this := RequestUpgradeOrganizationBilling{}
	return &this
}

// GetProduct returns the Product field value
func (o *RequestUpgradeOrganizationBilling) GetProduct() BillingProduct {
	if o == nil {
		var ret BillingProduct
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *RequestUpgradeOrganizationBilling) GetProductOk() (*BillingProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *RequestUpgradeOrganizationBilling) SetProduct(v BillingProduct) {
	o.Product = v
}

// GetInterval returns the Interval field value
func (o *RequestUpgradeOrganizationBilling) GetInterval() BillingInterval {
	if o == nil {
		var ret BillingInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *RequestUpgradeOrganizationBilling) GetIntervalOk() (*BillingInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *RequestUpgradeOrganizationBilling) SetInterval(v BillingInterval) {
	o.Interval = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RequestUpgradeOrganizationBilling) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUpgradeOrganizationBilling) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RequestUpgradeOrganizationBilling) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RequestUpgradeOrganizationBilling) SetReason(v string) {
	o.Reason = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RequestUpgradeOrganizationBilling) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestUpgradeOrganizationBilling) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RequestUpgradeOrganizationBilling) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RequestUpgradeOrganizationBilling) SetMode(v string) {
	o.Mode = &v
}

func (o RequestUpgradeOrganizationBilling) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestUpgradeOrganizationBilling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["interval"] = o.Interval
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableRequestUpgradeOrganizationBilling struct {
	value *RequestUpgradeOrganizationBilling
	isSet bool
}

func (v NullableRequestUpgradeOrganizationBilling) Get() *RequestUpgradeOrganizationBilling {
	return v.value
}

func (v *NullableRequestUpgradeOrganizationBilling) Set(val *RequestUpgradeOrganizationBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestUpgradeOrganizationBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestUpgradeOrganizationBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestUpgradeOrganizationBilling(val *RequestUpgradeOrganizationBilling) *NullableRequestUpgradeOrganizationBilling {
	return &NullableRequestUpgradeOrganizationBilling{value: val, isSet: true}
}

func (v NullableRequestUpgradeOrganizationBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestUpgradeOrganizationBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}