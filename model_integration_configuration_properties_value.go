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

// checks if the IntegrationConfigurationPropertiesValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationConfigurationPropertiesValue{}

// IntegrationConfigurationPropertiesValue struct for IntegrationConfigurationPropertiesValue
type IntegrationConfigurationPropertiesValue struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        string  `json:"type"`
	Default     *bool   `json:"default,omitempty"`
	CallbackUrl string  `json:"callback_url"`
	ButtonText  string  `json:"button_text"`
}

// NewIntegrationConfigurationPropertiesValue instantiates a new IntegrationConfigurationPropertiesValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationConfigurationPropertiesValue(type_ string, callbackUrl string, buttonText string) *IntegrationConfigurationPropertiesValue {
	this := IntegrationConfigurationPropertiesValue{}
	this.Type = type_
	this.CallbackUrl = callbackUrl
	this.ButtonText = buttonText
	return &this
}

// NewIntegrationConfigurationPropertiesValueWithDefaults instantiates a new IntegrationConfigurationPropertiesValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationConfigurationPropertiesValueWithDefaults() *IntegrationConfigurationPropertiesValue {
	this := IntegrationConfigurationPropertiesValue{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IntegrationConfigurationPropertiesValue) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValue) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IntegrationConfigurationPropertiesValue) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IntegrationConfigurationPropertiesValue) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationConfigurationPropertiesValue) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValue) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationConfigurationPropertiesValue) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationConfigurationPropertiesValue) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *IntegrationConfigurationPropertiesValue) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValue) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationConfigurationPropertiesValue) SetType(v string) {
	o.Type = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *IntegrationConfigurationPropertiesValue) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValue) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *IntegrationConfigurationPropertiesValue) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *IntegrationConfigurationPropertiesValue) SetDefault(v bool) {
	o.Default = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *IntegrationConfigurationPropertiesValue) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValue) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *IntegrationConfigurationPropertiesValue) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetButtonText returns the ButtonText field value
func (o *IntegrationConfigurationPropertiesValue) GetButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ButtonText
}

// GetButtonTextOk returns a tuple with the ButtonText field value
// and a boolean to check if the value has been set.
func (o *IntegrationConfigurationPropertiesValue) GetButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ButtonText, true
}

// SetButtonText sets field value
func (o *IntegrationConfigurationPropertiesValue) SetButtonText(v string) {
	o.ButtonText = v
}

func (o IntegrationConfigurationPropertiesValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationConfigurationPropertiesValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	toSerialize["callback_url"] = o.CallbackUrl
	toSerialize["button_text"] = o.ButtonText
	return toSerialize, nil
}

type NullableIntegrationConfigurationPropertiesValue struct {
	value *IntegrationConfigurationPropertiesValue
	isSet bool
}

func (v NullableIntegrationConfigurationPropertiesValue) Get() *IntegrationConfigurationPropertiesValue {
	return v.value
}

func (v *NullableIntegrationConfigurationPropertiesValue) Set(val *IntegrationConfigurationPropertiesValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationConfigurationPropertiesValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationConfigurationPropertiesValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationConfigurationPropertiesValue(val *IntegrationConfigurationPropertiesValue) *NullableIntegrationConfigurationPropertiesValue {
	return &NullableIntegrationConfigurationPropertiesValue{value: val, isSet: true}
}

func (v NullableIntegrationConfigurationPropertiesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationConfigurationPropertiesValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
