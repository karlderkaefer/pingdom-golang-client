# HTTPCustom

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
**Encryption** | Pointer to **bool** | (httpcustom specific) Connection encryption | [optional] [default to false]
**Port** | Pointer to **int32** | (httpcustom specific) Target port | [optional] [default to 80]
**Auth** | Pointer to **string** | (httpcustom specific) Username and password for target HTTP authentication. | [optional] 
**Additionalurls** | Pointer to **string** | (httpcustom specific) ;-separated list of addidional URLs with hostname included. | [optional] 

## Methods

### NewHTTPCustom

`func NewHTTPCustom(host string, type_ string, url string, ) *HTTPCustom`

NewHTTPCustom instantiates a new HTTPCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPCustomWithDefaults

`func NewHTTPCustomWithDefaults() *HTTPCustom`

NewHTTPCustomWithDefaults instantiates a new HTTPCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HTTPCustom) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HTTPCustom) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HTTPCustom) SetHost(v string)`

SetHost sets Host field to given value.


### GetType

`func (o *HTTPCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HTTPCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HTTPCustom) SetType(v string)`

SetType sets Type field to given value.


### GetProbeid

`func (o *HTTPCustom) GetProbeid() int32`

GetProbeid returns the Probeid field if non-nil, zero value otherwise.

### GetProbeidOk

`func (o *HTTPCustom) GetProbeidOk() (*int32, bool)`

GetProbeidOk returns a tuple with the Probeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeid

`func (o *HTTPCustom) SetProbeid(v int32)`

SetProbeid sets Probeid field to given value.

### HasProbeid

`func (o *HTTPCustom) HasProbeid() bool`

HasProbeid returns a boolean if a field has been set.

### GetProbeFilters

`func (o *HTTPCustom) GetProbeFilters() int32`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *HTTPCustom) GetProbeFiltersOk() (*int32, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *HTTPCustom) SetProbeFilters(v int32)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *HTTPCustom) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetIpv6

`func (o *HTTPCustom) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *HTTPCustom) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *HTTPCustom) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *HTTPCustom) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *HTTPCustom) GetResponsetimeThreshold() int32`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *HTTPCustom) GetResponsetimeThresholdOk() (*int32, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *HTTPCustom) SetResponsetimeThreshold(v int32)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *HTTPCustom) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetUrl

`func (o *HTTPCustom) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HTTPCustom) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HTTPCustom) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEncryption

`func (o *HTTPCustom) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *HTTPCustom) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *HTTPCustom) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *HTTPCustom) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *HTTPCustom) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HTTPCustom) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HTTPCustom) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HTTPCustom) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuth

`func (o *HTTPCustom) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *HTTPCustom) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *HTTPCustom) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *HTTPCustom) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAdditionalurls

`func (o *HTTPCustom) GetAdditionalurls() string`

GetAdditionalurls returns the Additionalurls field if non-nil, zero value otherwise.

### GetAdditionalurlsOk

`func (o *HTTPCustom) GetAdditionalurlsOk() (*string, bool)`

GetAdditionalurlsOk returns a tuple with the Additionalurls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalurls

`func (o *HTTPCustom) SetAdditionalurls(v string)`

SetAdditionalurls sets Additionalurls field to given value.

### HasAdditionalurls

`func (o *HTTPCustom) HasAdditionalurls() bool`

HasAdditionalurls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


