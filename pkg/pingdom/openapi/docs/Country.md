# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Country id | [optional] 
**Iso** | Pointer to **string** | Country ISO code | [optional] 
**Name** | Pointer to **string** | Country name | [optional] 

## Methods

### NewCountry

`func NewCountry() *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Country) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Country) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Country) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Country) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIso

`func (o *Country) GetIso() string`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *Country) GetIsoOk() (*string, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *Country) SetIso(v string)`

SetIso sets Iso field to given value.

### HasIso

`func (o *Country) HasIso() bool`

HasIso returns a boolean if a field has been set.

### GetName

`func (o *Country) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Country) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Country) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Country) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


