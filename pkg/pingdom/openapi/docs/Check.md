# Check

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewCheck

`func NewCheck() *Check`

NewCheck instantiates a new Check object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithDefaults

`func NewCheckWithDefaults() *Check`

NewCheckWithDefaults instantiates a new Check object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Check) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Check) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Check) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Check) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Check) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Check) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Check) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Check) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLasterrortime

`func (o *Check) GetLasterrortime() int32`

GetLasterrortime returns the Lasterrortime field if non-nil, zero value otherwise.

### GetLasterrortimeOk

`func (o *Check) GetLasterrortimeOk() (*int32, bool)`

GetLasterrortimeOk returns a tuple with the Lasterrortime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrortime

`func (o *Check) SetLasterrortime(v int32)`

SetLasterrortime sets Lasterrortime field to given value.

### HasLasterrortime

`func (o *Check) HasLasterrortime() bool`

HasLasterrortime returns a boolean if a field has been set.

### GetLasttesttime

`func (o *Check) GetLasttesttime() int32`

GetLasttesttime returns the Lasttesttime field if non-nil, zero value otherwise.

### GetLasttesttimeOk

`func (o *Check) GetLasttesttimeOk() (*int32, bool)`

GetLasttesttimeOk returns a tuple with the Lasttesttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasttesttime

`func (o *Check) SetLasttesttime(v int32)`

SetLasttesttime sets Lasttesttime field to given value.

### HasLasttesttime

`func (o *Check) HasLasttesttime() bool`

HasLasttesttime returns a boolean if a field has been set.

### GetLastresponsetime

`func (o *Check) GetLastresponsetime() int32`

GetLastresponsetime returns the Lastresponsetime field if non-nil, zero value otherwise.

### GetLastresponsetimeOk

`func (o *Check) GetLastresponsetimeOk() (*int32, bool)`

GetLastresponsetimeOk returns a tuple with the Lastresponsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastresponsetime

`func (o *Check) SetLastresponsetime(v int32)`

SetLastresponsetime sets Lastresponsetime field to given value.

### HasLastresponsetime

`func (o *Check) HasLastresponsetime() bool`

HasLastresponsetime returns a boolean if a field has been set.

### GetLastdownstart

`func (o *Check) GetLastdownstart() int32`

GetLastdownstart returns the Lastdownstart field if non-nil, zero value otherwise.

### GetLastdownstartOk

`func (o *Check) GetLastdownstartOk() (*int32, bool)`

GetLastdownstartOk returns a tuple with the Lastdownstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownstart

`func (o *Check) SetLastdownstart(v int32)`

SetLastdownstart sets Lastdownstart field to given value.

### HasLastdownstart

`func (o *Check) HasLastdownstart() bool`

HasLastdownstart returns a boolean if a field has been set.

### GetLastdownend

`func (o *Check) GetLastdownend() int32`

GetLastdownend returns the Lastdownend field if non-nil, zero value otherwise.

### GetLastdownendOk

`func (o *Check) GetLastdownendOk() (*int32, bool)`

GetLastdownendOk returns a tuple with the Lastdownend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownend

`func (o *Check) SetLastdownend(v int32)`

SetLastdownend sets Lastdownend field to given value.

### HasLastdownend

`func (o *Check) HasLastdownend() bool`

HasLastdownend returns a boolean if a field has been set.

### GetStatus

`func (o *Check) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Check) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Check) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Check) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolution

`func (o *Check) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Check) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Check) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *Check) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetHostname

`func (o *Check) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Check) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Check) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Check) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCreated

`func (o *Check) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Check) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Check) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Check) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTags

`func (o *Check) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Check) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Check) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Check) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpv6

`func (o *Check) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Check) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Check) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *Check) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


