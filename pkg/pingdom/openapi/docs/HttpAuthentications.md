# HttpAuthentications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**HttpAuthenticationsCredentials**](HttpAuthenticationsCredentials.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpAuthentications

`func NewHttpAuthentications() *HttpAuthentications`

NewHttpAuthentications instantiates a new HttpAuthentications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpAuthenticationsWithDefaults

`func NewHttpAuthenticationsWithDefaults() *HttpAuthentications`

NewHttpAuthenticationsWithDefaults instantiates a new HttpAuthentications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *HttpAuthentications) GetCredentials() HttpAuthenticationsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HttpAuthentications) GetCredentialsOk() (*HttpAuthenticationsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HttpAuthentications) SetCredentials(v HttpAuthenticationsCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *HttpAuthentications) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHost

`func (o *HttpAuthentications) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HttpAuthentications) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HttpAuthentications) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HttpAuthentications) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


