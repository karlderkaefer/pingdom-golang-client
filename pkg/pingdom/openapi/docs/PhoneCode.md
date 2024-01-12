# PhoneCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countryid** | Pointer to **int32** | Country id (Can be mapped against countries.id) | [optional] 
**Name** | Pointer to **string** | Country name | [optional] 
**Phonecode** | Pointer to **string** | Area phone code | [optional] 

## Methods

### NewPhoneCode

`func NewPhoneCode() *PhoneCode`

NewPhoneCode instantiates a new PhoneCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneCodeWithDefaults

`func NewPhoneCodeWithDefaults() *PhoneCode`

NewPhoneCodeWithDefaults instantiates a new PhoneCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryid

`func (o *PhoneCode) GetCountryid() int32`

GetCountryid returns the Countryid field if non-nil, zero value otherwise.

### GetCountryidOk

`func (o *PhoneCode) GetCountryidOk() (*int32, bool)`

GetCountryidOk returns a tuple with the Countryid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryid

`func (o *PhoneCode) SetCountryid(v int32)`

SetCountryid sets Countryid field to given value.

### HasCountryid

`func (o *PhoneCode) HasCountryid() bool`

HasCountryid returns a boolean if a field has been set.

### GetName

`func (o *PhoneCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhoneCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhoneCode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhoneCode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhonecode

`func (o *PhoneCode) GetPhonecode() string`

GetPhonecode returns the Phonecode field if non-nil, zero value otherwise.

### GetPhonecodeOk

`func (o *PhoneCode) GetPhonecodeOk() (*string, bool)`

GetPhonecodeOk returns a tuple with the Phonecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonecode

`func (o *PhoneCode) SetPhonecode(v string)`

SetPhonecode sets Phonecode field to given value.

### HasPhonecode

`func (o *PhoneCode) HasPhonecode() bool`

HasPhonecode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


