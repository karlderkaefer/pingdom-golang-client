# DetailedCheckUdpCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DetailedUdpAttributesType**](DetailedUdpAttributesType.md) |  | [optional] 
**ProbeFilters** | Pointer to **[]string** | Filters used for probe selections | [optional] 
**Sendnotificationwhendown** | Pointer to **int32** |  | [optional] 
**Notifyagainevery** | Pointer to **int32** |  | [optional] 
**Notifywhenbackup** | Pointer to **bool** |  | [optional] 
**ResponsetimeThreshold** | Pointer to **bool** |  | [optional] 
**CustomMessage** | Pointer to **string** |  | [optional] 
**Integrationids** | Pointer to **[]int32** |  | [optional] 
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

### NewDetailedCheckUdpCheck

`func NewDetailedCheckUdpCheck() *DetailedCheckUdpCheck`

NewDetailedCheckUdpCheck instantiates a new DetailedCheckUdpCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedCheckUdpCheckWithDefaults

`func NewDetailedCheckUdpCheckWithDefaults() *DetailedCheckUdpCheck`

NewDetailedCheckUdpCheckWithDefaults instantiates a new DetailedCheckUdpCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DetailedCheckUdpCheck) GetType() DetailedUdpAttributesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetailedCheckUdpCheck) GetTypeOk() (*DetailedUdpAttributesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetailedCheckUdpCheck) SetType(v DetailedUdpAttributesType)`

SetType sets Type field to given value.

### HasType

`func (o *DetailedCheckUdpCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProbeFilters

`func (o *DetailedCheckUdpCheck) GetProbeFilters() []string`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *DetailedCheckUdpCheck) GetProbeFiltersOk() (*[]string, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *DetailedCheckUdpCheck) SetProbeFilters(v []string)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *DetailedCheckUdpCheck) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetSendnotificationwhendown

`func (o *DetailedCheckUdpCheck) GetSendnotificationwhendown() int32`

GetSendnotificationwhendown returns the Sendnotificationwhendown field if non-nil, zero value otherwise.

### GetSendnotificationwhendownOk

`func (o *DetailedCheckUdpCheck) GetSendnotificationwhendownOk() (*int32, bool)`

GetSendnotificationwhendownOk returns a tuple with the Sendnotificationwhendown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendnotificationwhendown

`func (o *DetailedCheckUdpCheck) SetSendnotificationwhendown(v int32)`

SetSendnotificationwhendown sets Sendnotificationwhendown field to given value.

### HasSendnotificationwhendown

`func (o *DetailedCheckUdpCheck) HasSendnotificationwhendown() bool`

HasSendnotificationwhendown returns a boolean if a field has been set.

### GetNotifyagainevery

`func (o *DetailedCheckUdpCheck) GetNotifyagainevery() int32`

GetNotifyagainevery returns the Notifyagainevery field if non-nil, zero value otherwise.

### GetNotifyagaineveryOk

`func (o *DetailedCheckUdpCheck) GetNotifyagaineveryOk() (*int32, bool)`

GetNotifyagaineveryOk returns a tuple with the Notifyagainevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyagainevery

`func (o *DetailedCheckUdpCheck) SetNotifyagainevery(v int32)`

SetNotifyagainevery sets Notifyagainevery field to given value.

### HasNotifyagainevery

`func (o *DetailedCheckUdpCheck) HasNotifyagainevery() bool`

HasNotifyagainevery returns a boolean if a field has been set.

### GetNotifywhenbackup

`func (o *DetailedCheckUdpCheck) GetNotifywhenbackup() bool`

GetNotifywhenbackup returns the Notifywhenbackup field if non-nil, zero value otherwise.

### GetNotifywhenbackupOk

`func (o *DetailedCheckUdpCheck) GetNotifywhenbackupOk() (*bool, bool)`

GetNotifywhenbackupOk returns a tuple with the Notifywhenbackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifywhenbackup

`func (o *DetailedCheckUdpCheck) SetNotifywhenbackup(v bool)`

SetNotifywhenbackup sets Notifywhenbackup field to given value.

### HasNotifywhenbackup

`func (o *DetailedCheckUdpCheck) HasNotifywhenbackup() bool`

HasNotifywhenbackup returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *DetailedCheckUdpCheck) GetResponsetimeThreshold() bool`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *DetailedCheckUdpCheck) GetResponsetimeThresholdOk() (*bool, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *DetailedCheckUdpCheck) SetResponsetimeThreshold(v bool)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *DetailedCheckUdpCheck) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetCustomMessage

`func (o *DetailedCheckUdpCheck) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *DetailedCheckUdpCheck) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *DetailedCheckUdpCheck) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *DetailedCheckUdpCheck) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetIntegrationids

`func (o *DetailedCheckUdpCheck) GetIntegrationids() []int32`

GetIntegrationids returns the Integrationids field if non-nil, zero value otherwise.

### GetIntegrationidsOk

`func (o *DetailedCheckUdpCheck) GetIntegrationidsOk() (*[]int32, bool)`

GetIntegrationidsOk returns a tuple with the Integrationids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationids

`func (o *DetailedCheckUdpCheck) SetIntegrationids(v []int32)`

SetIntegrationids sets Integrationids field to given value.

### HasIntegrationids

`func (o *DetailedCheckUdpCheck) HasIntegrationids() bool`

HasIntegrationids returns a boolean if a field has been set.

### GetId

`func (o *DetailedCheckUdpCheck) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedCheckUdpCheck) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedCheckUdpCheck) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedCheckUdpCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DetailedCheckUdpCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedCheckUdpCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedCheckUdpCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedCheckUdpCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLasterrortime

`func (o *DetailedCheckUdpCheck) GetLasterrortime() int32`

GetLasterrortime returns the Lasterrortime field if non-nil, zero value otherwise.

### GetLasterrortimeOk

`func (o *DetailedCheckUdpCheck) GetLasterrortimeOk() (*int32, bool)`

GetLasterrortimeOk returns a tuple with the Lasterrortime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrortime

`func (o *DetailedCheckUdpCheck) SetLasterrortime(v int32)`

SetLasterrortime sets Lasterrortime field to given value.

### HasLasterrortime

`func (o *DetailedCheckUdpCheck) HasLasterrortime() bool`

HasLasterrortime returns a boolean if a field has been set.

### GetLasttesttime

`func (o *DetailedCheckUdpCheck) GetLasttesttime() int32`

GetLasttesttime returns the Lasttesttime field if non-nil, zero value otherwise.

### GetLasttesttimeOk

`func (o *DetailedCheckUdpCheck) GetLasttesttimeOk() (*int32, bool)`

GetLasttesttimeOk returns a tuple with the Lasttesttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasttesttime

`func (o *DetailedCheckUdpCheck) SetLasttesttime(v int32)`

SetLasttesttime sets Lasttesttime field to given value.

### HasLasttesttime

`func (o *DetailedCheckUdpCheck) HasLasttesttime() bool`

HasLasttesttime returns a boolean if a field has been set.

### GetLastresponsetime

`func (o *DetailedCheckUdpCheck) GetLastresponsetime() int32`

GetLastresponsetime returns the Lastresponsetime field if non-nil, zero value otherwise.

### GetLastresponsetimeOk

`func (o *DetailedCheckUdpCheck) GetLastresponsetimeOk() (*int32, bool)`

GetLastresponsetimeOk returns a tuple with the Lastresponsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastresponsetime

`func (o *DetailedCheckUdpCheck) SetLastresponsetime(v int32)`

SetLastresponsetime sets Lastresponsetime field to given value.

### HasLastresponsetime

`func (o *DetailedCheckUdpCheck) HasLastresponsetime() bool`

HasLastresponsetime returns a boolean if a field has been set.

### GetLastdownstart

`func (o *DetailedCheckUdpCheck) GetLastdownstart() int32`

GetLastdownstart returns the Lastdownstart field if non-nil, zero value otherwise.

### GetLastdownstartOk

`func (o *DetailedCheckUdpCheck) GetLastdownstartOk() (*int32, bool)`

GetLastdownstartOk returns a tuple with the Lastdownstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownstart

`func (o *DetailedCheckUdpCheck) SetLastdownstart(v int32)`

SetLastdownstart sets Lastdownstart field to given value.

### HasLastdownstart

`func (o *DetailedCheckUdpCheck) HasLastdownstart() bool`

HasLastdownstart returns a boolean if a field has been set.

### GetLastdownend

`func (o *DetailedCheckUdpCheck) GetLastdownend() int32`

GetLastdownend returns the Lastdownend field if non-nil, zero value otherwise.

### GetLastdownendOk

`func (o *DetailedCheckUdpCheck) GetLastdownendOk() (*int32, bool)`

GetLastdownendOk returns a tuple with the Lastdownend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownend

`func (o *DetailedCheckUdpCheck) SetLastdownend(v int32)`

SetLastdownend sets Lastdownend field to given value.

### HasLastdownend

`func (o *DetailedCheckUdpCheck) HasLastdownend() bool`

HasLastdownend returns a boolean if a field has been set.

### GetStatus

`func (o *DetailedCheckUdpCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailedCheckUdpCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailedCheckUdpCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetailedCheckUdpCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolution

`func (o *DetailedCheckUdpCheck) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *DetailedCheckUdpCheck) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *DetailedCheckUdpCheck) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *DetailedCheckUdpCheck) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetHostname

`func (o *DetailedCheckUdpCheck) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DetailedCheckUdpCheck) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DetailedCheckUdpCheck) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DetailedCheckUdpCheck) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCreated

`func (o *DetailedCheckUdpCheck) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DetailedCheckUdpCheck) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DetailedCheckUdpCheck) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DetailedCheckUdpCheck) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTags

`func (o *DetailedCheckUdpCheck) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetailedCheckUdpCheck) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetailedCheckUdpCheck) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DetailedCheckUdpCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpv6

`func (o *DetailedCheckUdpCheck) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DetailedCheckUdpCheck) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DetailedCheckUdpCheck) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *DetailedCheckUdpCheck) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


