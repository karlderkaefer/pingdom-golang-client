# CheckWithStringType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Lasterrortime** | Pointer to **int32** | Timestamp of last error (if any). Format is UNIX timestamp | [optional] 
**Lasttesttime** | Pointer to **int32** | Timestamp of last test (if any). Format is UNIX timestamp | [optional] 
**Lastresponsetime** | Pointer to **int32** | Response time (in milliseconds) of last test. | [optional] 
**Lastdownstart** | Pointer to **int32** | Timestamp of start of last check down (if any). Format is UNIX timestamp. | [optional] 
**Lastdownend** | Pointer to **int32** | Timestamp of end of last check down (if any). Format is UNIX timestamp. During a downtime it will be lasttesttime. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to **int32** | How often should the check be tested? (minutes) | [optional] 
**Hostname** | Pointer to **string** | Target host | [optional] 
**Created** | Pointer to **int32** | Creating time. Format is UNIX timestamp | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | List of tags for check | [optional] 
**Ipv6** | Pointer to **bool** | Use ipv6 instead of ipv4 | [optional] 

## Methods

### NewCheckWithStringType

`func NewCheckWithStringType() *CheckWithStringType`

NewCheckWithStringType instantiates a new CheckWithStringType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithStringTypeWithDefaults

`func NewCheckWithStringTypeWithDefaults() *CheckWithStringType`

NewCheckWithStringTypeWithDefaults instantiates a new CheckWithStringType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckWithStringType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckWithStringType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckWithStringType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckWithStringType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CheckWithStringType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckWithStringType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckWithStringType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CheckWithStringType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CheckWithStringType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckWithStringType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckWithStringType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckWithStringType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLasterrortime

`func (o *CheckWithStringType) GetLasterrortime() int32`

GetLasterrortime returns the Lasterrortime field if non-nil, zero value otherwise.

### GetLasterrortimeOk

`func (o *CheckWithStringType) GetLasterrortimeOk() (*int32, bool)`

GetLasterrortimeOk returns a tuple with the Lasterrortime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrortime

`func (o *CheckWithStringType) SetLasterrortime(v int32)`

SetLasterrortime sets Lasterrortime field to given value.

### HasLasterrortime

`func (o *CheckWithStringType) HasLasterrortime() bool`

HasLasterrortime returns a boolean if a field has been set.

### GetLasttesttime

`func (o *CheckWithStringType) GetLasttesttime() int32`

GetLasttesttime returns the Lasttesttime field if non-nil, zero value otherwise.

### GetLasttesttimeOk

`func (o *CheckWithStringType) GetLasttesttimeOk() (*int32, bool)`

GetLasttesttimeOk returns a tuple with the Lasttesttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasttesttime

`func (o *CheckWithStringType) SetLasttesttime(v int32)`

SetLasttesttime sets Lasttesttime field to given value.

### HasLasttesttime

`func (o *CheckWithStringType) HasLasttesttime() bool`

HasLasttesttime returns a boolean if a field has been set.

### GetLastresponsetime

`func (o *CheckWithStringType) GetLastresponsetime() int32`

GetLastresponsetime returns the Lastresponsetime field if non-nil, zero value otherwise.

### GetLastresponsetimeOk

`func (o *CheckWithStringType) GetLastresponsetimeOk() (*int32, bool)`

GetLastresponsetimeOk returns a tuple with the Lastresponsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastresponsetime

`func (o *CheckWithStringType) SetLastresponsetime(v int32)`

SetLastresponsetime sets Lastresponsetime field to given value.

### HasLastresponsetime

`func (o *CheckWithStringType) HasLastresponsetime() bool`

HasLastresponsetime returns a boolean if a field has been set.

### GetLastdownstart

`func (o *CheckWithStringType) GetLastdownstart() int32`

GetLastdownstart returns the Lastdownstart field if non-nil, zero value otherwise.

### GetLastdownstartOk

`func (o *CheckWithStringType) GetLastdownstartOk() (*int32, bool)`

GetLastdownstartOk returns a tuple with the Lastdownstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownstart

`func (o *CheckWithStringType) SetLastdownstart(v int32)`

SetLastdownstart sets Lastdownstart field to given value.

### HasLastdownstart

`func (o *CheckWithStringType) HasLastdownstart() bool`

HasLastdownstart returns a boolean if a field has been set.

### GetLastdownend

`func (o *CheckWithStringType) GetLastdownend() int32`

GetLastdownend returns the Lastdownend field if non-nil, zero value otherwise.

### GetLastdownendOk

`func (o *CheckWithStringType) GetLastdownendOk() (*int32, bool)`

GetLastdownendOk returns a tuple with the Lastdownend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownend

`func (o *CheckWithStringType) SetLastdownend(v int32)`

SetLastdownend sets Lastdownend field to given value.

### HasLastdownend

`func (o *CheckWithStringType) HasLastdownend() bool`

HasLastdownend returns a boolean if a field has been set.

### GetStatus

`func (o *CheckWithStringType) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckWithStringType) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckWithStringType) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckWithStringType) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolution

`func (o *CheckWithStringType) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *CheckWithStringType) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *CheckWithStringType) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *CheckWithStringType) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetHostname

`func (o *CheckWithStringType) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CheckWithStringType) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CheckWithStringType) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CheckWithStringType) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCreated

`func (o *CheckWithStringType) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CheckWithStringType) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CheckWithStringType) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CheckWithStringType) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTags

`func (o *CheckWithStringType) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckWithStringType) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckWithStringType) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckWithStringType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpv6

`func (o *CheckWithStringType) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CheckWithStringType) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CheckWithStringType) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CheckWithStringType) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


