# SMTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Target host | 
**Type** | **string** |  | 
**Probeid** | Pointer to **int32** | Probe identifier | [optional] 
**ProbeFilters** | Pointer to **int32** | Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are &#39;EU&#39;, &#39;NA&#39;, &#39;APAC&#39; and &#39;LATAM&#39;. | [optional] 
**Ipv6** | Pointer to **bool** | Use ipv6 instead of ipv4 | [optional] 
**ResponsetimeThreshold** | Pointer to **int32** | Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.) | [optional] [default to 30000]
**Port** | Pointer to **int32** | (smtp specific) Target port | [optional] [default to 25]
**Auth** | Pointer to **string** | (smtp specific) Username and password for target HTTP authentication. | [optional] 
**Stringtoexpect** | Pointer to **string** | (smtp specific) String to expect in response | [optional] 
**Encryption** | Pointer to **bool** | (smtp specific) Connection encryption | [optional] [default to false]

## Methods

### NewSMTP

`func NewSMTP(host string, type_ string, ) *SMTP`

NewSMTP instantiates a new SMTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPWithDefaults

`func NewSMTPWithDefaults() *SMTP`

NewSMTPWithDefaults instantiates a new SMTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SMTP) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTP) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTP) SetHost(v string)`

SetHost sets Host field to given value.


### GetType

`func (o *SMTP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SMTP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SMTP) SetType(v string)`

SetType sets Type field to given value.


### GetProbeid

`func (o *SMTP) GetProbeid() int32`

GetProbeid returns the Probeid field if non-nil, zero value otherwise.

### GetProbeidOk

`func (o *SMTP) GetProbeidOk() (*int32, bool)`

GetProbeidOk returns a tuple with the Probeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeid

`func (o *SMTP) SetProbeid(v int32)`

SetProbeid sets Probeid field to given value.

### HasProbeid

`func (o *SMTP) HasProbeid() bool`

HasProbeid returns a boolean if a field has been set.

### GetProbeFilters

`func (o *SMTP) GetProbeFilters() int32`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *SMTP) GetProbeFiltersOk() (*int32, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *SMTP) SetProbeFilters(v int32)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *SMTP) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetIpv6

`func (o *SMTP) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *SMTP) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *SMTP) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *SMTP) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *SMTP) GetResponsetimeThreshold() int32`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *SMTP) GetResponsetimeThresholdOk() (*int32, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *SMTP) SetResponsetimeThreshold(v int32)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *SMTP) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetPort

`func (o *SMTP) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTP) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTP) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTP) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuth

`func (o *SMTP) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *SMTP) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *SMTP) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *SMTP) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *SMTP) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *SMTP) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *SMTP) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *SMTP) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.

### GetEncryption

`func (o *SMTP) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *SMTP) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *SMTP) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *SMTP) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


