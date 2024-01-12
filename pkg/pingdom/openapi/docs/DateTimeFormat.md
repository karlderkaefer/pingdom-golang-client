# DateTimeFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Date/time identifier | [optional] 
**Description** | Pointer to **string** | Date/time description | [optional] 

## Methods

### NewDateTimeFormat

`func NewDateTimeFormat() *DateTimeFormat`

NewDateTimeFormat instantiates a new DateTimeFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeFormatWithDefaults

`func NewDateTimeFormatWithDefaults() *DateTimeFormat`

NewDateTimeFormatWithDefaults instantiates a new DateTimeFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DateTimeFormat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DateTimeFormat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DateTimeFormat) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DateTimeFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *DateTimeFormat) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DateTimeFormat) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DateTimeFormat) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DateTimeFormat) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


