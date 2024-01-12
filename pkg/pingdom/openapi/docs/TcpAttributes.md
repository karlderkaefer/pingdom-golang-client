# TcpAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | Target port | 
**Stringtosend** | Pointer to **string** | String to send | [optional] 
**Stringtoexpect** | Pointer to **string** | String to expect in response | [optional] 

## Methods

### NewTcpAttributes

`func NewTcpAttributes(port int32, ) *TcpAttributes`

NewTcpAttributes instantiates a new TcpAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcpAttributesWithDefaults

`func NewTcpAttributesWithDefaults() *TcpAttributes`

NewTcpAttributesWithDefaults instantiates a new TcpAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *TcpAttributes) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TcpAttributes) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TcpAttributes) SetPort(v int32)`

SetPort sets Port field to given value.


### GetStringtosend

`func (o *TcpAttributes) GetStringtosend() string`

GetStringtosend returns the Stringtosend field if non-nil, zero value otherwise.

### GetStringtosendOk

`func (o *TcpAttributes) GetStringtosendOk() (*string, bool)`

GetStringtosendOk returns a tuple with the Stringtosend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtosend

`func (o *TcpAttributes) SetStringtosend(v string)`

SetStringtosend sets Stringtosend field to given value.

### HasStringtosend

`func (o *TcpAttributes) HasStringtosend() bool`

HasStringtosend returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *TcpAttributes) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *TcpAttributes) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *TcpAttributes) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *TcpAttributes) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


