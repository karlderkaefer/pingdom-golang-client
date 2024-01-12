# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | **int32** | Width of the browser window | 
**Height** | **int32** | Height of the browser window | 
**DisableWebSecurity** | **bool** | Setting this field to false will disable the same-origin policy during check execution | 
**Authentications** | [**MetadataGETAuthentications**](MetadataGETAuthentications.md) |  | 

## Methods

### NewMetadata

`func NewMetadata(width int32, height int32, disableWebSecurity bool, authentications MetadataGETAuthentications, ) *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *Metadata) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Metadata) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Metadata) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *Metadata) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Metadata) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Metadata) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDisableWebSecurity

`func (o *Metadata) GetDisableWebSecurity() bool`

GetDisableWebSecurity returns the DisableWebSecurity field if non-nil, zero value otherwise.

### GetDisableWebSecurityOk

`func (o *Metadata) GetDisableWebSecurityOk() (*bool, bool)`

GetDisableWebSecurityOk returns a tuple with the DisableWebSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableWebSecurity

`func (o *Metadata) SetDisableWebSecurity(v bool)`

SetDisableWebSecurity sets DisableWebSecurity field to given value.


### GetAuthentications

`func (o *Metadata) GetAuthentications() MetadataGETAuthentications`

GetAuthentications returns the Authentications field if non-nil, zero value otherwise.

### GetAuthenticationsOk

`func (o *Metadata) GetAuthenticationsOk() (*MetadataGETAuthentications, bool)`

GetAuthenticationsOk returns a tuple with the Authentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentications

`func (o *Metadata) SetAuthentications(v MetadataGETAuthentications)`

SetAuthentications sets Authentications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


