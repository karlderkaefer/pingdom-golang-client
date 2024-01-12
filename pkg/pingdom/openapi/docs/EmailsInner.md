# EmailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | Pointer to **string** | Contact target&#39;s severity level | [optional] 
**Address** | Pointer to **string** | Email address | [optional] 

## Methods

### NewEmailsInner

`func NewEmailsInner() *EmailsInner`

NewEmailsInner instantiates a new EmailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailsInnerWithDefaults

`func NewEmailsInnerWithDefaults() *EmailsInner`

NewEmailsInnerWithDefaults instantiates a new EmailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *EmailsInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EmailsInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EmailsInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EmailsInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetAddress

`func (o *EmailsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmailsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmailsInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EmailsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


