# DetailedCheckTcpCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DetailedTcpAttributesType**](DetailedTcpAttributesType.md) |  | [optional] 
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

### NewDetailedCheckTcpCheck

`func NewDetailedCheckTcpCheck() *DetailedCheckTcpCheck`

NewDetailedCheckTcpCheck instantiates a new DetailedCheckTcpCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedCheckTcpCheckWithDefaults

`func NewDetailedCheckTcpCheckWithDefaults() *DetailedCheckTcpCheck`

NewDetailedCheckTcpCheckWithDefaults instantiates a new DetailedCheckTcpCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DetailedCheckTcpCheck) GetType() DetailedTcpAttributesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetailedCheckTcpCheck) GetTypeOk() (*DetailedTcpAttributesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetailedCheckTcpCheck) SetType(v DetailedTcpAttributesType)`

SetType sets Type field to given value.

### HasType

`func (o *DetailedCheckTcpCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProbeFilters

`func (o *DetailedCheckTcpCheck) GetProbeFilters() []string`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *DetailedCheckTcpCheck) GetProbeFiltersOk() (*[]string, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *DetailedCheckTcpCheck) SetProbeFilters(v []string)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *DetailedCheckTcpCheck) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetSendnotificationwhendown

`func (o *DetailedCheckTcpCheck) GetSendnotificationwhendown() int32`

GetSendnotificationwhendown returns the Sendnotificationwhendown field if non-nil, zero value otherwise.

### GetSendnotificationwhendownOk

`func (o *DetailedCheckTcpCheck) GetSendnotificationwhendownOk() (*int32, bool)`

GetSendnotificationwhendownOk returns a tuple with the Sendnotificationwhendown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendnotificationwhendown

`func (o *DetailedCheckTcpCheck) SetSendnotificationwhendown(v int32)`

SetSendnotificationwhendown sets Sendnotificationwhendown field to given value.

### HasSendnotificationwhendown

`func (o *DetailedCheckTcpCheck) HasSendnotificationwhendown() bool`

HasSendnotificationwhendown returns a boolean if a field has been set.

### GetNotifyagainevery

`func (o *DetailedCheckTcpCheck) GetNotifyagainevery() int32`

GetNotifyagainevery returns the Notifyagainevery field if non-nil, zero value otherwise.

### GetNotifyagaineveryOk

`func (o *DetailedCheckTcpCheck) GetNotifyagaineveryOk() (*int32, bool)`

GetNotifyagaineveryOk returns a tuple with the Notifyagainevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyagainevery

`func (o *DetailedCheckTcpCheck) SetNotifyagainevery(v int32)`

SetNotifyagainevery sets Notifyagainevery field to given value.

### HasNotifyagainevery

`func (o *DetailedCheckTcpCheck) HasNotifyagainevery() bool`

HasNotifyagainevery returns a boolean if a field has been set.

### GetNotifywhenbackup

`func (o *DetailedCheckTcpCheck) GetNotifywhenbackup() bool`

GetNotifywhenbackup returns the Notifywhenbackup field if non-nil, zero value otherwise.

### GetNotifywhenbackupOk

`func (o *DetailedCheckTcpCheck) GetNotifywhenbackupOk() (*bool, bool)`

GetNotifywhenbackupOk returns a tuple with the Notifywhenbackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifywhenbackup

`func (o *DetailedCheckTcpCheck) SetNotifywhenbackup(v bool)`

SetNotifywhenbackup sets Notifywhenbackup field to given value.

### HasNotifywhenbackup

`func (o *DetailedCheckTcpCheck) HasNotifywhenbackup() bool`

HasNotifywhenbackup returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *DetailedCheckTcpCheck) GetResponsetimeThreshold() bool`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *DetailedCheckTcpCheck) GetResponsetimeThresholdOk() (*bool, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *DetailedCheckTcpCheck) SetResponsetimeThreshold(v bool)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *DetailedCheckTcpCheck) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetCustomMessage

`func (o *DetailedCheckTcpCheck) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *DetailedCheckTcpCheck) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *DetailedCheckTcpCheck) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *DetailedCheckTcpCheck) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetIntegrationids

`func (o *DetailedCheckTcpCheck) GetIntegrationids() []int32`

GetIntegrationids returns the Integrationids field if non-nil, zero value otherwise.

### GetIntegrationidsOk

`func (o *DetailedCheckTcpCheck) GetIntegrationidsOk() (*[]int32, bool)`

GetIntegrationidsOk returns a tuple with the Integrationids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationids

`func (o *DetailedCheckTcpCheck) SetIntegrationids(v []int32)`

SetIntegrationids sets Integrationids field to given value.

### HasIntegrationids

`func (o *DetailedCheckTcpCheck) HasIntegrationids() bool`

HasIntegrationids returns a boolean if a field has been set.

### GetId

`func (o *DetailedCheckTcpCheck) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedCheckTcpCheck) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedCheckTcpCheck) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedCheckTcpCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DetailedCheckTcpCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedCheckTcpCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedCheckTcpCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedCheckTcpCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLasterrortime

`func (o *DetailedCheckTcpCheck) GetLasterrortime() int32`

GetLasterrortime returns the Lasterrortime field if non-nil, zero value otherwise.

### GetLasterrortimeOk

`func (o *DetailedCheckTcpCheck) GetLasterrortimeOk() (*int32, bool)`

GetLasterrortimeOk returns a tuple with the Lasterrortime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrortime

`func (o *DetailedCheckTcpCheck) SetLasterrortime(v int32)`

SetLasterrortime sets Lasterrortime field to given value.

### HasLasterrortime

`func (o *DetailedCheckTcpCheck) HasLasterrortime() bool`

HasLasterrortime returns a boolean if a field has been set.

### GetLasttesttime

`func (o *DetailedCheckTcpCheck) GetLasttesttime() int32`

GetLasttesttime returns the Lasttesttime field if non-nil, zero value otherwise.

### GetLasttesttimeOk

`func (o *DetailedCheckTcpCheck) GetLasttesttimeOk() (*int32, bool)`

GetLasttesttimeOk returns a tuple with the Lasttesttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasttesttime

`func (o *DetailedCheckTcpCheck) SetLasttesttime(v int32)`

SetLasttesttime sets Lasttesttime field to given value.

### HasLasttesttime

`func (o *DetailedCheckTcpCheck) HasLasttesttime() bool`

HasLasttesttime returns a boolean if a field has been set.

### GetLastresponsetime

`func (o *DetailedCheckTcpCheck) GetLastresponsetime() int32`

GetLastresponsetime returns the Lastresponsetime field if non-nil, zero value otherwise.

### GetLastresponsetimeOk

`func (o *DetailedCheckTcpCheck) GetLastresponsetimeOk() (*int32, bool)`

GetLastresponsetimeOk returns a tuple with the Lastresponsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastresponsetime

`func (o *DetailedCheckTcpCheck) SetLastresponsetime(v int32)`

SetLastresponsetime sets Lastresponsetime field to given value.

### HasLastresponsetime

`func (o *DetailedCheckTcpCheck) HasLastresponsetime() bool`

HasLastresponsetime returns a boolean if a field has been set.

### GetLastdownstart

`func (o *DetailedCheckTcpCheck) GetLastdownstart() int32`

GetLastdownstart returns the Lastdownstart field if non-nil, zero value otherwise.

### GetLastdownstartOk

`func (o *DetailedCheckTcpCheck) GetLastdownstartOk() (*int32, bool)`

GetLastdownstartOk returns a tuple with the Lastdownstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownstart

`func (o *DetailedCheckTcpCheck) SetLastdownstart(v int32)`

SetLastdownstart sets Lastdownstart field to given value.

### HasLastdownstart

`func (o *DetailedCheckTcpCheck) HasLastdownstart() bool`

HasLastdownstart returns a boolean if a field has been set.

### GetLastdownend

`func (o *DetailedCheckTcpCheck) GetLastdownend() int32`

GetLastdownend returns the Lastdownend field if non-nil, zero value otherwise.

### GetLastdownendOk

`func (o *DetailedCheckTcpCheck) GetLastdownendOk() (*int32, bool)`

GetLastdownendOk returns a tuple with the Lastdownend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownend

`func (o *DetailedCheckTcpCheck) SetLastdownend(v int32)`

SetLastdownend sets Lastdownend field to given value.

### HasLastdownend

`func (o *DetailedCheckTcpCheck) HasLastdownend() bool`

HasLastdownend returns a boolean if a field has been set.

### GetStatus

`func (o *DetailedCheckTcpCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailedCheckTcpCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailedCheckTcpCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetailedCheckTcpCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolution

`func (o *DetailedCheckTcpCheck) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *DetailedCheckTcpCheck) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *DetailedCheckTcpCheck) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *DetailedCheckTcpCheck) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetHostname

`func (o *DetailedCheckTcpCheck) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DetailedCheckTcpCheck) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DetailedCheckTcpCheck) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DetailedCheckTcpCheck) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCreated

`func (o *DetailedCheckTcpCheck) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DetailedCheckTcpCheck) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DetailedCheckTcpCheck) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DetailedCheckTcpCheck) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTags

`func (o *DetailedCheckTcpCheck) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetailedCheckTcpCheck) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetailedCheckTcpCheck) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DetailedCheckTcpCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpv6

`func (o *DetailedCheckTcpCheck) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DetailedCheckTcpCheck) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DetailedCheckTcpCheck) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *DetailedCheckTcpCheck) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


