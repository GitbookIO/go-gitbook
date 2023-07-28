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

// checks if the CloudflareHostnameTLSInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudflareHostnameTLSInfo{}

// CloudflareHostnameTLSInfo The Cloudflare Hostname TLS information
type CloudflareHostnameTLSInfo struct {
	Status               CloudflareHostnameTLSStatus            `json:"status"`
	Method               CloudflareHostnameTLSValidationMethod  `json:"method"`
	CertificateAuthority *string                                `json:"certificateAuthority,omitempty"`
	Certificates         []CloudflareHostnameTLSCertificate     `json:"certificates"`
	ValidationErrors     []CloudflareHostnameTLSValidationError `json:"validationErrors"`
}

// NewCloudflareHostnameTLSInfo instantiates a new CloudflareHostnameTLSInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudflareHostnameTLSInfo(status CloudflareHostnameTLSStatus, method CloudflareHostnameTLSValidationMethod, certificates []CloudflareHostnameTLSCertificate, validationErrors []CloudflareHostnameTLSValidationError) *CloudflareHostnameTLSInfo {
	this := CloudflareHostnameTLSInfo{}
	this.Status = status
	this.Method = method
	this.Certificates = certificates
	this.ValidationErrors = validationErrors
	return &this
}

// NewCloudflareHostnameTLSInfoWithDefaults instantiates a new CloudflareHostnameTLSInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudflareHostnameTLSInfoWithDefaults() *CloudflareHostnameTLSInfo {
	this := CloudflareHostnameTLSInfo{}
	return &this
}

// GetStatus returns the Status field value
func (o *CloudflareHostnameTLSInfo) GetStatus() CloudflareHostnameTLSStatus {
	if o == nil {
		var ret CloudflareHostnameTLSStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSInfo) GetStatusOk() (*CloudflareHostnameTLSStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CloudflareHostnameTLSInfo) SetStatus(v CloudflareHostnameTLSStatus) {
	o.Status = v
}

// GetMethod returns the Method field value
func (o *CloudflareHostnameTLSInfo) GetMethod() CloudflareHostnameTLSValidationMethod {
	if o == nil {
		var ret CloudflareHostnameTLSValidationMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSInfo) GetMethodOk() (*CloudflareHostnameTLSValidationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *CloudflareHostnameTLSInfo) SetMethod(v CloudflareHostnameTLSValidationMethod) {
	o.Method = v
}

// GetCertificateAuthority returns the CertificateAuthority field value if set, zero value otherwise.
func (o *CloudflareHostnameTLSInfo) GetCertificateAuthority() string {
	if o == nil || IsNil(o.CertificateAuthority) {
		var ret string
		return ret
	}
	return *o.CertificateAuthority
}

// GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSInfo) GetCertificateAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateAuthority) {
		return nil, false
	}
	return o.CertificateAuthority, true
}

// HasCertificateAuthority returns a boolean if a field has been set.
func (o *CloudflareHostnameTLSInfo) HasCertificateAuthority() bool {
	if o != nil && !IsNil(o.CertificateAuthority) {
		return true
	}

	return false
}

// SetCertificateAuthority gets a reference to the given string and assigns it to the CertificateAuthority field.
func (o *CloudflareHostnameTLSInfo) SetCertificateAuthority(v string) {
	o.CertificateAuthority = &v
}

// GetCertificates returns the Certificates field value
func (o *CloudflareHostnameTLSInfo) GetCertificates() []CloudflareHostnameTLSCertificate {
	if o == nil {
		var ret []CloudflareHostnameTLSCertificate
		return ret
	}

	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSInfo) GetCertificatesOk() ([]CloudflareHostnameTLSCertificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificates, true
}

// SetCertificates sets field value
func (o *CloudflareHostnameTLSInfo) SetCertificates(v []CloudflareHostnameTLSCertificate) {
	o.Certificates = v
}

// GetValidationErrors returns the ValidationErrors field value
func (o *CloudflareHostnameTLSInfo) GetValidationErrors() []CloudflareHostnameTLSValidationError {
	if o == nil {
		var ret []CloudflareHostnameTLSValidationError
		return ret
	}

	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value
// and a boolean to check if the value has been set.
func (o *CloudflareHostnameTLSInfo) GetValidationErrorsOk() ([]CloudflareHostnameTLSValidationError, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// SetValidationErrors sets field value
func (o *CloudflareHostnameTLSInfo) SetValidationErrors(v []CloudflareHostnameTLSValidationError) {
	o.ValidationErrors = v
}

func (o CloudflareHostnameTLSInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudflareHostnameTLSInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["method"] = o.Method
	if !IsNil(o.CertificateAuthority) {
		toSerialize["certificateAuthority"] = o.CertificateAuthority
	}
	toSerialize["certificates"] = o.Certificates
	toSerialize["validationErrors"] = o.ValidationErrors
	return toSerialize, nil
}

type NullableCloudflareHostnameTLSInfo struct {
	value *CloudflareHostnameTLSInfo
	isSet bool
}

func (v NullableCloudflareHostnameTLSInfo) Get() *CloudflareHostnameTLSInfo {
	return v.value
}

func (v *NullableCloudflareHostnameTLSInfo) Set(val *CloudflareHostnameTLSInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudflareHostnameTLSInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudflareHostnameTLSInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudflareHostnameTLSInfo(val *CloudflareHostnameTLSInfo) *NullableCloudflareHostnameTLSInfo {
	return &NullableCloudflareHostnameTLSInfo{value: val, isSet: true}
}

func (v NullableCloudflareHostnameTLSInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudflareHostnameTLSInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
