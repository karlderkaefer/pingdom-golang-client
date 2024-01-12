# MetadataGET

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | Pointer to **int32** | Width of the browser window | [optional] 
**Height** | Pointer to **int32** | Height of the browser window | [optional] 
**DisableWebSecurity** | Pointer to **bool** | Setting this field to false will disable the same-origin policy during check execution | [optional] 
**Authentications** | Pointer to [**MetadataGETAuthentications**](MetadataGETAuthentications.md) |  | [optional] 

## Methods

### NewMetadataGET

`func NewMetadataGET() *MetadataGET`

NewMetadataGET instantiates a new MetadataGET object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataGETWithDefaults

`func NewMetadataGETWithDefaults() *MetadataGET`

NewMetadataGETWithDefaults instantiates a new MetadataGET object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *MetadataGET) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MetadataGET) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MetadataGET) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MetadataGET) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *MetadataGET) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MetadataGET) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MetadataGET) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *MetadataGET) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDisableWebSecurity

`func (o *MetadataGET) GetDisableWebSecurity() bool`

GetDisableWebSecurity returns the DisableWebSecurity field if non-nil, zero value otherwise.

### GetDisableWebSecurityOk

`func (o *MetadataGET) GetDisableWebSecurityOk() (*bool, bool)`

GetDisableWebSecurityOk returns a tuple with the DisableWebSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableWebSecurity

`func (o *MetadataGET) SetDisableWebSecurity(v bool)`

SetDisableWebSecurity sets DisableWebSecurity field to given value.

### HasDisableWebSecurity

`func (o *MetadataGET) HasDisableWebSecurity() bool`

HasDisableWebSecurity returns a boolean if a field has been set.

### GetAuthentications

`func (o *MetadataGET) GetAuthentications() MetadataGETAuthentications`

GetAuthentications returns the Authentications field if non-nil, zero value otherwise.

### GetAuthenticationsOk

`func (o *MetadataGET) GetAuthenticationsOk() (*MetadataGETAuthentications, bool)`

GetAuthenticationsOk returns a tuple with the Authentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentications

`func (o *MetadataGET) SetAuthentications(v MetadataGETAuthentications)`

SetAuthentications sets Authentications field to given value.

### HasAuthentications

`func (o *MetadataGET) HasAuthentications() bool`

HasAuthentications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


