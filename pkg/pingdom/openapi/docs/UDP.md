# UDP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Target host | 
**Type** | **string** |  | 
**Probeid** | Pointer to **int32** | Probe identifier | [optional] 
**ProbeFilters** | Pointer to **int32** | Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are &#39;EU&#39;, &#39;NA&#39;, &#39;APAC&#39; and &#39;LATAM&#39;. | [optional] 
**Ipv6** | Pointer to **bool** | Use ipv6 instead of ipv4 | [optional] 
**ResponsetimeThreshold** | Pointer to **int32** | Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.) | [optional] [default to 30000]
**Port** | **int32** | (udp specific) Target port | 
**Stringtosend** | Pointer to **string** | (udp specific) String to send | [optional] 
**Stringtoexpect** | Pointer to **string** | (udp specific) String to expect in response | [optional] 

## Methods

### NewUDP

`func NewUDP(host string, type_ string, port int32, ) *UDP`

NewUDP instantiates a new UDP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUDPWithDefaults

`func NewUDPWithDefaults() *UDP`

NewUDPWithDefaults instantiates a new UDP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *UDP) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UDP) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UDP) SetHost(v string)`

SetHost sets Host field to given value.


### GetType

`func (o *UDP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UDP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UDP) SetType(v string)`

SetType sets Type field to given value.


### GetProbeid

`func (o *UDP) GetProbeid() int32`

GetProbeid returns the Probeid field if non-nil, zero value otherwise.

### GetProbeidOk

`func (o *UDP) GetProbeidOk() (*int32, bool)`

GetProbeidOk returns a tuple with the Probeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeid

`func (o *UDP) SetProbeid(v int32)`

SetProbeid sets Probeid field to given value.

### HasProbeid

`func (o *UDP) HasProbeid() bool`

HasProbeid returns a boolean if a field has been set.

### GetProbeFilters

`func (o *UDP) GetProbeFilters() int32`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *UDP) GetProbeFiltersOk() (*int32, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *UDP) SetProbeFilters(v int32)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *UDP) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetIpv6

`func (o *UDP) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *UDP) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *UDP) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *UDP) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *UDP) GetResponsetimeThreshold() int32`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *UDP) GetResponsetimeThresholdOk() (*int32, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *UDP) SetResponsetimeThreshold(v int32)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *UDP) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetPort

`func (o *UDP) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UDP) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UDP) SetPort(v int32)`

SetPort sets Port field to given value.


### GetStringtosend

`func (o *UDP) GetStringtosend() string`

GetStringtosend returns the Stringtosend field if non-nil, zero value otherwise.

### GetStringtosendOk

`func (o *UDP) GetStringtosendOk() (*string, bool)`

GetStringtosendOk returns a tuple with the Stringtosend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtosend

`func (o *UDP) SetStringtosend(v string)`

SetStringtosend sets Stringtosend field to given value.

### HasStringtosend

`func (o *UDP) HasStringtosend() bool`

HasStringtosend returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *UDP) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *UDP) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *UDP) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *UDP) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


