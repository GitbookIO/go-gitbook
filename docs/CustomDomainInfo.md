# CustomDomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**Status** | [**CloudflareHostnameStatus**](CloudflareHostnameStatus.md) |  | 
**CreatedAt** | **string** |  | 
**Ssl** | Pointer to [**CloudflareHostnameTLSInfo**](CloudflareHostnameTLSInfo.md) |  | [optional] 
**VerificationErrors** | **[]string** |  | 

## Methods

### NewCustomDomainInfo

`func NewCustomDomainInfo(hostname string, status CloudflareHostnameStatus, createdAt string, verificationErrors []string, ) *CustomDomainInfo`

NewCustomDomainInfo instantiates a new CustomDomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainInfoWithDefaults

`func NewCustomDomainInfoWithDefaults() *CustomDomainInfo`

NewCustomDomainInfoWithDefaults instantiates a new CustomDomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *CustomDomainInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CustomDomainInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CustomDomainInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *CustomDomainInfo) GetStatus() CloudflareHostnameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDomainInfo) GetStatusOk() (*CloudflareHostnameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDomainInfo) SetStatus(v CloudflareHostnameStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *CustomDomainInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomDomainInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomDomainInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetSsl

`func (o *CustomDomainInfo) GetSsl() CloudflareHostnameTLSInfo`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *CustomDomainInfo) GetSslOk() (*CloudflareHostnameTLSInfo, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *CustomDomainInfo) SetSsl(v CloudflareHostnameTLSInfo)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *CustomDomainInfo) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetVerificationErrors

`func (o *CustomDomainInfo) GetVerificationErrors() []string`

GetVerificationErrors returns the VerificationErrors field if non-nil, zero value otherwise.

### GetVerificationErrorsOk

`func (o *CustomDomainInfo) GetVerificationErrorsOk() (*[]string, bool)`

GetVerificationErrorsOk returns a tuple with the VerificationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationErrors

`func (o *CustomDomainInfo) SetVerificationErrors(v []string)`

SetVerificationErrors sets VerificationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


