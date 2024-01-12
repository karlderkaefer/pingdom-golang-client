# MetadataGETAuthentications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpAuthentications** | Pointer to [**[]HttpAuthentications**](HttpAuthentications.md) | HTTP (browser-level) authentications. Currently, only Basic Authentication is supported | [optional] 

## Methods

### NewMetadataGETAuthentications

`func NewMetadataGETAuthentications() *MetadataGETAuthentications`

NewMetadataGETAuthentications instantiates a new MetadataGETAuthentications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataGETAuthenticationsWithDefaults

`func NewMetadataGETAuthenticationsWithDefaults() *MetadataGETAuthentications`

NewMetadataGETAuthenticationsWithDefaults instantiates a new MetadataGETAuthentications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpAuthentications

`func (o *MetadataGETAuthentications) GetHttpAuthentications() []HttpAuthentications`

GetHttpAuthentications returns the HttpAuthentications field if non-nil, zero value otherwise.

### GetHttpAuthenticationsOk

`func (o *MetadataGETAuthentications) GetHttpAuthenticationsOk() (*[]HttpAuthentications, bool)`

GetHttpAuthenticationsOk returns a tuple with the HttpAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuthentications

`func (o *MetadataGETAuthentications) SetHttpAuthentications(v []HttpAuthentications)`

SetHttpAuthentications sets HttpAuthentications field to given value.

### HasHttpAuthentications

`func (o *MetadataGETAuthentications) HasHttpAuthentications() bool`

HasHttpAuthentications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


