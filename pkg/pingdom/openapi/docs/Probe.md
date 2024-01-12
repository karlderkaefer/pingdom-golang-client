# Probe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** | Is the probe currently active? | [optional] 
**Hostname** | Pointer to **string** | DNS name | [optional] 
**Ip** | Pointer to **string** | IPv4 address | [optional] 
**Ipv6** | Pointer to **string** | IPv6 address (not all probes have this) | [optional] 
**Countryiso** | Pointer to **string** | Country ISO code | [optional] 

## Methods

### NewProbe

`func NewProbe() *Probe`

NewProbe instantiates a new Probe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProbeWithDefaults

`func NewProbeWithDefaults() *Probe`

NewProbeWithDefaults instantiates a new Probe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Probe) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Probe) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Probe) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Probe) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCountry

`func (o *Probe) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Probe) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Probe) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Probe) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *Probe) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Probe) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Probe) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Probe) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetName

`func (o *Probe) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Probe) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Probe) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Probe) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *Probe) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Probe) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Probe) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Probe) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetHostname

`func (o *Probe) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Probe) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Probe) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Probe) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *Probe) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Probe) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Probe) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Probe) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *Probe) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Probe) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Probe) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *Probe) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetCountryiso

`func (o *Probe) GetCountryiso() string`

GetCountryiso returns the Countryiso field if non-nil, zero value otherwise.

### GetCountryisoOk

`func (o *Probe) GetCountryisoOk() (*string, bool)`

GetCountryisoOk returns a tuple with the Countryiso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryiso

`func (o *Probe) SetCountryiso(v string)`

SetCountryiso sets Countryiso field to given value.

### HasCountryiso

`func (o *Probe) HasCountryiso() bool`

HasCountryiso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


