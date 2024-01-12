# UdpAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | Target port | 
**Stringtosend** | **string** | String to send | 
**Stringtoexpect** | **string** | String to expect in response | 

## Methods

### NewUdpAttributes

`func NewUdpAttributes(port int32, stringtosend string, stringtoexpect string, ) *UdpAttributes`

NewUdpAttributes instantiates a new UdpAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdpAttributesWithDefaults

`func NewUdpAttributesWithDefaults() *UdpAttributes`

NewUdpAttributesWithDefaults instantiates a new UdpAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *UdpAttributes) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UdpAttributes) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UdpAttributes) SetPort(v int32)`

SetPort sets Port field to given value.


### GetStringtosend

`func (o *UdpAttributes) GetStringtosend() string`

GetStringtosend returns the Stringtosend field if non-nil, zero value otherwise.

### GetStringtosendOk

`func (o *UdpAttributes) GetStringtosendOk() (*string, bool)`

GetStringtosendOk returns a tuple with the Stringtosend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtosend

`func (o *UdpAttributes) SetStringtosend(v string)`

SetStringtosend sets Stringtosend field to given value.


### GetStringtoexpect

`func (o *UdpAttributes) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *UdpAttributes) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *UdpAttributes) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


