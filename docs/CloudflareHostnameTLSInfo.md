# CloudflareHostnameTLSInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**CloudflareHostnameTLSStatus**](CloudflareHostnameTLSStatus.md) |  | 
**Method** | [**CloudflareHostnameTLSValidationMethod**](CloudflareHostnameTLSValidationMethod.md) |  | 
**CertificateAuthority** | Pointer to **string** |  | [optional] 
**Certificates** | [**[]CloudflareHostnameTLSCertificate**](CloudflareHostnameTLSCertificate.md) |  | 
**ValidationErrors** | [**[]CloudflareHostnameTLSValidationError**](CloudflareHostnameTLSValidationError.md) |  | 

## Methods

### NewCloudflareHostnameTLSInfo

`func NewCloudflareHostnameTLSInfo(status CloudflareHostnameTLSStatus, method CloudflareHostnameTLSValidationMethod, certificates []CloudflareHostnameTLSCertificate, validationErrors []CloudflareHostnameTLSValidationError, ) *CloudflareHostnameTLSInfo`

NewCloudflareHostnameTLSInfo instantiates a new CloudflareHostnameTLSInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudflareHostnameTLSInfoWithDefaults

`func NewCloudflareHostnameTLSInfoWithDefaults() *CloudflareHostnameTLSInfo`

NewCloudflareHostnameTLSInfoWithDefaults instantiates a new CloudflareHostnameTLSInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudflareHostnameTLSInfo) GetStatus() CloudflareHostnameTLSStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudflareHostnameTLSInfo) GetStatusOk() (*CloudflareHostnameTLSStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudflareHostnameTLSInfo) SetStatus(v CloudflareHostnameTLSStatus)`

SetStatus sets Status field to given value.


### GetMethod

`func (o *CloudflareHostnameTLSInfo) GetMethod() CloudflareHostnameTLSValidationMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CloudflareHostnameTLSInfo) GetMethodOk() (*CloudflareHostnameTLSValidationMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CloudflareHostnameTLSInfo) SetMethod(v CloudflareHostnameTLSValidationMethod)`

SetMethod sets Method field to given value.


### GetCertificateAuthority

`func (o *CloudflareHostnameTLSInfo) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *CloudflareHostnameTLSInfo) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *CloudflareHostnameTLSInfo) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *CloudflareHostnameTLSInfo) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetCertificates

`func (o *CloudflareHostnameTLSInfo) GetCertificates() []CloudflareHostnameTLSCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *CloudflareHostnameTLSInfo) GetCertificatesOk() (*[]CloudflareHostnameTLSCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *CloudflareHostnameTLSInfo) SetCertificates(v []CloudflareHostnameTLSCertificate)`

SetCertificates sets Certificates field to given value.


### GetValidationErrors

`func (o *CloudflareHostnameTLSInfo) GetValidationErrors() []CloudflareHostnameTLSValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CloudflareHostnameTLSInfo) GetValidationErrorsOk() (*[]CloudflareHostnameTLSValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CloudflareHostnameTLSInfo) SetValidationErrors(v []CloudflareHostnameTLSValidationError)`

SetValidationErrors sets ValidationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


