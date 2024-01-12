# DnsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nameserver** | **string** | DNS server to use | 
**Expectedip** | **string** | Expected IP | 

## Methods

### NewDnsAttributes

`func NewDnsAttributes(nameserver string, expectedip string, ) *DnsAttributes`

NewDnsAttributes instantiates a new DnsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsAttributesWithDefaults

`func NewDnsAttributesWithDefaults() *DnsAttributes`

NewDnsAttributesWithDefaults instantiates a new DnsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameserver

`func (o *DnsAttributes) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *DnsAttributes) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *DnsAttributes) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.


### GetExpectedip

`func (o *DnsAttributes) GetExpectedip() string`

GetExpectedip returns the Expectedip field if non-nil, zero value otherwise.

### GetExpectedipOk

`func (o *DnsAttributes) GetExpectedipOk() (*string, bool)`

GetExpectedipOk returns a tuple with the Expectedip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedip

`func (o *DnsAttributes) SetExpectedip(v string)`

SetExpectedip sets Expectedip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


