# HttpAuthenticationsCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Basic Authentication password | [optional] 
**UserName** | Pointer to **string** | Basic Authentication Username | [optional] 

## Methods

### NewHttpAuthenticationsCredentials

`func NewHttpAuthenticationsCredentials() *HttpAuthenticationsCredentials`

NewHttpAuthenticationsCredentials instantiates a new HttpAuthenticationsCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpAuthenticationsCredentialsWithDefaults

`func NewHttpAuthenticationsCredentialsWithDefaults() *HttpAuthenticationsCredentials`

NewHttpAuthenticationsCredentialsWithDefaults instantiates a new HttpAuthenticationsCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *HttpAuthenticationsCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HttpAuthenticationsCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HttpAuthenticationsCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HttpAuthenticationsCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserName

`func (o *HttpAuthenticationsCredentials) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HttpAuthenticationsCredentials) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HttpAuthenticationsCredentials) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *HttpAuthenticationsCredentials) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


