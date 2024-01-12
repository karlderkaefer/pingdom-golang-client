# SingleGetQueryParametersParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Target host | 
**Type** | **string** |  | 
**Probeid** | Pointer to **int32** | Probe identifier | [optional] 
**ProbeFilters** | Pointer to **int32** | Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are &#39;EU&#39;, &#39;NA&#39;, &#39;APAC&#39; and &#39;LATAM&#39;. | [optional] 
**Ipv6** | Pointer to **bool** | Use ipv6 instead of ipv4 | [optional] 
**ResponsetimeThreshold** | Pointer to **int32** | Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.) | [optional] [default to 30000]
**Url** | **string** | (httpcustom specific) Target path to XML file on server | 
**Encryption** | Pointer to **bool** | (imap specific) Connection encryption | [optional] [default to false]
**Port** | **int32** | (imap specific) Target port | [default to 143]
**Auth** | Pointer to **string** | (smtp specific) Username and password for target HTTP authentication. | [optional] 
**Shouldcontain** | Pointer to **string** | (http specific) Target site should contain this string | [optional] 
**Shouldnotcontain** | Pointer to **string** | (http specific) Target site should NOT contain this string | [optional] 
**Postdata** | Pointer to **string** | (http specific) Data that should be posted to the web page, for example submission data for a sign-up or login form. The data needs to be formatted in the same way as a web browser would send it to the web server | [optional] 
**RequestheaderX** | Pointer to **string** | (http specific) Custom HTTP header name. Replace {X} with a number unique for each header argument. | [optional] 
**VerifyCertificate** | Pointer to **bool** | (http specific) Treat target site as down if an invalid/unverifiable certificate is found. | [optional] [default to true]
**SslDownDaysBefore** | Pointer to **int32** | (http specific) Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if &#x60;verify_certificate&#x60; is set to &#x60;false&#x60;. | [optional] [default to 0]
**Additionalurls** | Pointer to **string** | (httpcustom specific) ;-separated list of addidional URLs with hostname included. | [optional] 
**Stringtosend** | Pointer to **string** | (udp specific) String to send | [optional] 
**Stringtoexpect** | Pointer to **string** | (imap specific) String to expect in response | [optional] 
**Expectedip** | **string** | (dns specific) Expected ip | 
**Nameserver** | **string** | (dns specific) Nameserver | 

## Methods

### NewSingleGetQueryParametersParameter

`func NewSingleGetQueryParametersParameter(host string, type_ string, url string, port int32, expectedip string, nameserver string, ) *SingleGetQueryParametersParameter`

NewSingleGetQueryParametersParameter instantiates a new SingleGetQueryParametersParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleGetQueryParametersParameterWithDefaults

`func NewSingleGetQueryParametersParameterWithDefaults() *SingleGetQueryParametersParameter`

NewSingleGetQueryParametersParameterWithDefaults instantiates a new SingleGetQueryParametersParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SingleGetQueryParametersParameter) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SingleGetQueryParametersParameter) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SingleGetQueryParametersParameter) SetHost(v string)`

SetHost sets Host field to given value.


### GetType

`func (o *SingleGetQueryParametersParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SingleGetQueryParametersParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SingleGetQueryParametersParameter) SetType(v string)`

SetType sets Type field to given value.


### GetProbeid

`func (o *SingleGetQueryParametersParameter) GetProbeid() int32`

GetProbeid returns the Probeid field if non-nil, zero value otherwise.

### GetProbeidOk

`func (o *SingleGetQueryParametersParameter) GetProbeidOk() (*int32, bool)`

GetProbeidOk returns a tuple with the Probeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeid

`func (o *SingleGetQueryParametersParameter) SetProbeid(v int32)`

SetProbeid sets Probeid field to given value.

### HasProbeid

`func (o *SingleGetQueryParametersParameter) HasProbeid() bool`

HasProbeid returns a boolean if a field has been set.

### GetProbeFilters

`func (o *SingleGetQueryParametersParameter) GetProbeFilters() int32`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *SingleGetQueryParametersParameter) GetProbeFiltersOk() (*int32, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *SingleGetQueryParametersParameter) SetProbeFilters(v int32)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *SingleGetQueryParametersParameter) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetIpv6

`func (o *SingleGetQueryParametersParameter) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *SingleGetQueryParametersParameter) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *SingleGetQueryParametersParameter) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *SingleGetQueryParametersParameter) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *SingleGetQueryParametersParameter) GetResponsetimeThreshold() int32`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *SingleGetQueryParametersParameter) GetResponsetimeThresholdOk() (*int32, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *SingleGetQueryParametersParameter) SetResponsetimeThreshold(v int32)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *SingleGetQueryParametersParameter) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetUrl

`func (o *SingleGetQueryParametersParameter) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SingleGetQueryParametersParameter) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SingleGetQueryParametersParameter) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEncryption

`func (o *SingleGetQueryParametersParameter) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *SingleGetQueryParametersParameter) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *SingleGetQueryParametersParameter) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *SingleGetQueryParametersParameter) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *SingleGetQueryParametersParameter) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SingleGetQueryParametersParameter) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SingleGetQueryParametersParameter) SetPort(v int32)`

SetPort sets Port field to given value.


### GetAuth

`func (o *SingleGetQueryParametersParameter) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *SingleGetQueryParametersParameter) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *SingleGetQueryParametersParameter) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *SingleGetQueryParametersParameter) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetShouldcontain

`func (o *SingleGetQueryParametersParameter) GetShouldcontain() string`

GetShouldcontain returns the Shouldcontain field if non-nil, zero value otherwise.

### GetShouldcontainOk

`func (o *SingleGetQueryParametersParameter) GetShouldcontainOk() (*string, bool)`

GetShouldcontainOk returns a tuple with the Shouldcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldcontain

`func (o *SingleGetQueryParametersParameter) SetShouldcontain(v string)`

SetShouldcontain sets Shouldcontain field to given value.

### HasShouldcontain

`func (o *SingleGetQueryParametersParameter) HasShouldcontain() bool`

HasShouldcontain returns a boolean if a field has been set.

### GetShouldnotcontain

`func (o *SingleGetQueryParametersParameter) GetShouldnotcontain() string`

GetShouldnotcontain returns the Shouldnotcontain field if non-nil, zero value otherwise.

### GetShouldnotcontainOk

`func (o *SingleGetQueryParametersParameter) GetShouldnotcontainOk() (*string, bool)`

GetShouldnotcontainOk returns a tuple with the Shouldnotcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldnotcontain

`func (o *SingleGetQueryParametersParameter) SetShouldnotcontain(v string)`

SetShouldnotcontain sets Shouldnotcontain field to given value.

### HasShouldnotcontain

`func (o *SingleGetQueryParametersParameter) HasShouldnotcontain() bool`

HasShouldnotcontain returns a boolean if a field has been set.

### GetPostdata

`func (o *SingleGetQueryParametersParameter) GetPostdata() string`

GetPostdata returns the Postdata field if non-nil, zero value otherwise.

### GetPostdataOk

`func (o *SingleGetQueryParametersParameter) GetPostdataOk() (*string, bool)`

GetPostdataOk returns a tuple with the Postdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostdata

`func (o *SingleGetQueryParametersParameter) SetPostdata(v string)`

SetPostdata sets Postdata field to given value.

### HasPostdata

`func (o *SingleGetQueryParametersParameter) HasPostdata() bool`

HasPostdata returns a boolean if a field has been set.

### GetRequestheaderX

`func (o *SingleGetQueryParametersParameter) GetRequestheaderX() string`

GetRequestheaderX returns the RequestheaderX field if non-nil, zero value otherwise.

### GetRequestheaderXOk

`func (o *SingleGetQueryParametersParameter) GetRequestheaderXOk() (*string, bool)`

GetRequestheaderXOk returns a tuple with the RequestheaderX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestheaderX

`func (o *SingleGetQueryParametersParameter) SetRequestheaderX(v string)`

SetRequestheaderX sets RequestheaderX field to given value.

### HasRequestheaderX

`func (o *SingleGetQueryParametersParameter) HasRequestheaderX() bool`

HasRequestheaderX returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *SingleGetQueryParametersParameter) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *SingleGetQueryParametersParameter) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *SingleGetQueryParametersParameter) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *SingleGetQueryParametersParameter) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *SingleGetQueryParametersParameter) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *SingleGetQueryParametersParameter) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *SingleGetQueryParametersParameter) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *SingleGetQueryParametersParameter) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.

### GetAdditionalurls

`func (o *SingleGetQueryParametersParameter) GetAdditionalurls() string`

GetAdditionalurls returns the Additionalurls field if non-nil, zero value otherwise.

### GetAdditionalurlsOk

`func (o *SingleGetQueryParametersParameter) GetAdditionalurlsOk() (*string, bool)`

GetAdditionalurlsOk returns a tuple with the Additionalurls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalurls

`func (o *SingleGetQueryParametersParameter) SetAdditionalurls(v string)`

SetAdditionalurls sets Additionalurls field to given value.

### HasAdditionalurls

`func (o *SingleGetQueryParametersParameter) HasAdditionalurls() bool`

HasAdditionalurls returns a boolean if a field has been set.

### GetStringtosend

`func (o *SingleGetQueryParametersParameter) GetStringtosend() string`

GetStringtosend returns the Stringtosend field if non-nil, zero value otherwise.

### GetStringtosendOk

`func (o *SingleGetQueryParametersParameter) GetStringtosendOk() (*string, bool)`

GetStringtosendOk returns a tuple with the Stringtosend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtosend

`func (o *SingleGetQueryParametersParameter) SetStringtosend(v string)`

SetStringtosend sets Stringtosend field to given value.

### HasStringtosend

`func (o *SingleGetQueryParametersParameter) HasStringtosend() bool`

HasStringtosend returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *SingleGetQueryParametersParameter) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *SingleGetQueryParametersParameter) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *SingleGetQueryParametersParameter) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *SingleGetQueryParametersParameter) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.

### GetExpectedip

`func (o *SingleGetQueryParametersParameter) GetExpectedip() string`

GetExpectedip returns the Expectedip field if non-nil, zero value otherwise.

### GetExpectedipOk

`func (o *SingleGetQueryParametersParameter) GetExpectedipOk() (*string, bool)`

GetExpectedipOk returns a tuple with the Expectedip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedip

`func (o *SingleGetQueryParametersParameter) SetExpectedip(v string)`

SetExpectedip sets Expectedip field to given value.


### GetNameserver

`func (o *SingleGetQueryParametersParameter) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *SingleGetQueryParametersParameter) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *SingleGetQueryParametersParameter) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


