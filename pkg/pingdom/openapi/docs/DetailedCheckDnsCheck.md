# DetailedCheckDnsCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DetailedDnsAttributesType**](DetailedDnsAttributesType.md) |  | [optional] 
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

### NewDetailedCheckDnsCheck

`func NewDetailedCheckDnsCheck() *DetailedCheckDnsCheck`

NewDetailedCheckDnsCheck instantiates a new DetailedCheckDnsCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedCheckDnsCheckWithDefaults

`func NewDetailedCheckDnsCheckWithDefaults() *DetailedCheckDnsCheck`

NewDetailedCheckDnsCheckWithDefaults instantiates a new DetailedCheckDnsCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DetailedCheckDnsCheck) GetType() DetailedDnsAttributesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetailedCheckDnsCheck) GetTypeOk() (*DetailedDnsAttributesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetailedCheckDnsCheck) SetType(v DetailedDnsAttributesType)`

SetType sets Type field to given value.

### HasType

`func (o *DetailedCheckDnsCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProbeFilters

`func (o *DetailedCheckDnsCheck) GetProbeFilters() []string`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *DetailedCheckDnsCheck) GetProbeFiltersOk() (*[]string, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *DetailedCheckDnsCheck) SetProbeFilters(v []string)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *DetailedCheckDnsCheck) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetSendnotificationwhendown

`func (o *DetailedCheckDnsCheck) GetSendnotificationwhendown() int32`

GetSendnotificationwhendown returns the Sendnotificationwhendown field if non-nil, zero value otherwise.

### GetSendnotificationwhendownOk

`func (o *DetailedCheckDnsCheck) GetSendnotificationwhendownOk() (*int32, bool)`

GetSendnotificationwhendownOk returns a tuple with the Sendnotificationwhendown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendnotificationwhendown

`func (o *DetailedCheckDnsCheck) SetSendnotificationwhendown(v int32)`

SetSendnotificationwhendown sets Sendnotificationwhendown field to given value.

### HasSendnotificationwhendown

`func (o *DetailedCheckDnsCheck) HasSendnotificationwhendown() bool`

HasSendnotificationwhendown returns a boolean if a field has been set.

### GetNotifyagainevery

`func (o *DetailedCheckDnsCheck) GetNotifyagainevery() int32`

GetNotifyagainevery returns the Notifyagainevery field if non-nil, zero value otherwise.

### GetNotifyagaineveryOk

`func (o *DetailedCheckDnsCheck) GetNotifyagaineveryOk() (*int32, bool)`

GetNotifyagaineveryOk returns a tuple with the Notifyagainevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyagainevery

`func (o *DetailedCheckDnsCheck) SetNotifyagainevery(v int32)`

SetNotifyagainevery sets Notifyagainevery field to given value.

### HasNotifyagainevery

`func (o *DetailedCheckDnsCheck) HasNotifyagainevery() bool`

HasNotifyagainevery returns a boolean if a field has been set.

### GetNotifywhenbackup

`func (o *DetailedCheckDnsCheck) GetNotifywhenbackup() bool`

GetNotifywhenbackup returns the Notifywhenbackup field if non-nil, zero value otherwise.

### GetNotifywhenbackupOk

`func (o *DetailedCheckDnsCheck) GetNotifywhenbackupOk() (*bool, bool)`

GetNotifywhenbackupOk returns a tuple with the Notifywhenbackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifywhenbackup

`func (o *DetailedCheckDnsCheck) SetNotifywhenbackup(v bool)`

SetNotifywhenbackup sets Notifywhenbackup field to given value.

### HasNotifywhenbackup

`func (o *DetailedCheckDnsCheck) HasNotifywhenbackup() bool`

HasNotifywhenbackup returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *DetailedCheckDnsCheck) GetResponsetimeThreshold() bool`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *DetailedCheckDnsCheck) GetResponsetimeThresholdOk() (*bool, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *DetailedCheckDnsCheck) SetResponsetimeThreshold(v bool)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *DetailedCheckDnsCheck) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetCustomMessage

`func (o *DetailedCheckDnsCheck) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *DetailedCheckDnsCheck) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *DetailedCheckDnsCheck) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *DetailedCheckDnsCheck) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetIntegrationids

`func (o *DetailedCheckDnsCheck) GetIntegrationids() []int32`

GetIntegrationids returns the Integrationids field if non-nil, zero value otherwise.

### GetIntegrationidsOk

`func (o *DetailedCheckDnsCheck) GetIntegrationidsOk() (*[]int32, bool)`

GetIntegrationidsOk returns a tuple with the Integrationids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationids

`func (o *DetailedCheckDnsCheck) SetIntegrationids(v []int32)`

SetIntegrationids sets Integrationids field to given value.

### HasIntegrationids

`func (o *DetailedCheckDnsCheck) HasIntegrationids() bool`

HasIntegrationids returns a boolean if a field has been set.

### GetId

`func (o *DetailedCheckDnsCheck) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedCheckDnsCheck) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedCheckDnsCheck) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedCheckDnsCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DetailedCheckDnsCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedCheckDnsCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedCheckDnsCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedCheckDnsCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLasterrortime

`func (o *DetailedCheckDnsCheck) GetLasterrortime() int32`

GetLasterrortime returns the Lasterrortime field if non-nil, zero value otherwise.

### GetLasterrortimeOk

`func (o *DetailedCheckDnsCheck) GetLasterrortimeOk() (*int32, bool)`

GetLasterrortimeOk returns a tuple with the Lasterrortime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrortime

`func (o *DetailedCheckDnsCheck) SetLasterrortime(v int32)`

SetLasterrortime sets Lasterrortime field to given value.

### HasLasterrortime

`func (o *DetailedCheckDnsCheck) HasLasterrortime() bool`

HasLasterrortime returns a boolean if a field has been set.

### GetLasttesttime

`func (o *DetailedCheckDnsCheck) GetLasttesttime() int32`

GetLasttesttime returns the Lasttesttime field if non-nil, zero value otherwise.

### GetLasttesttimeOk

`func (o *DetailedCheckDnsCheck) GetLasttesttimeOk() (*int32, bool)`

GetLasttesttimeOk returns a tuple with the Lasttesttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasttesttime

`func (o *DetailedCheckDnsCheck) SetLasttesttime(v int32)`

SetLasttesttime sets Lasttesttime field to given value.

### HasLasttesttime

`func (o *DetailedCheckDnsCheck) HasLasttesttime() bool`

HasLasttesttime returns a boolean if a field has been set.

### GetLastresponsetime

`func (o *DetailedCheckDnsCheck) GetLastresponsetime() int32`

GetLastresponsetime returns the Lastresponsetime field if non-nil, zero value otherwise.

### GetLastresponsetimeOk

`func (o *DetailedCheckDnsCheck) GetLastresponsetimeOk() (*int32, bool)`

GetLastresponsetimeOk returns a tuple with the Lastresponsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastresponsetime

`func (o *DetailedCheckDnsCheck) SetLastresponsetime(v int32)`

SetLastresponsetime sets Lastresponsetime field to given value.

### HasLastresponsetime

`func (o *DetailedCheckDnsCheck) HasLastresponsetime() bool`

HasLastresponsetime returns a boolean if a field has been set.

### GetLastdownstart

`func (o *DetailedCheckDnsCheck) GetLastdownstart() int32`

GetLastdownstart returns the Lastdownstart field if non-nil, zero value otherwise.

### GetLastdownstartOk

`func (o *DetailedCheckDnsCheck) GetLastdownstartOk() (*int32, bool)`

GetLastdownstartOk returns a tuple with the Lastdownstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownstart

`func (o *DetailedCheckDnsCheck) SetLastdownstart(v int32)`

SetLastdownstart sets Lastdownstart field to given value.

### HasLastdownstart

`func (o *DetailedCheckDnsCheck) HasLastdownstart() bool`

HasLastdownstart returns a boolean if a field has been set.

### GetLastdownend

`func (o *DetailedCheckDnsCheck) GetLastdownend() int32`

GetLastdownend returns the Lastdownend field if non-nil, zero value otherwise.

### GetLastdownendOk

`func (o *DetailedCheckDnsCheck) GetLastdownendOk() (*int32, bool)`

GetLastdownendOk returns a tuple with the Lastdownend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownend

`func (o *DetailedCheckDnsCheck) SetLastdownend(v int32)`

SetLastdownend sets Lastdownend field to given value.

### HasLastdownend

`func (o *DetailedCheckDnsCheck) HasLastdownend() bool`

HasLastdownend returns a boolean if a field has been set.

### GetStatus

`func (o *DetailedCheckDnsCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailedCheckDnsCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailedCheckDnsCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetailedCheckDnsCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolution

`func (o *DetailedCheckDnsCheck) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *DetailedCheckDnsCheck) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *DetailedCheckDnsCheck) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *DetailedCheckDnsCheck) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetHostname

`func (o *DetailedCheckDnsCheck) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DetailedCheckDnsCheck) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DetailedCheckDnsCheck) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DetailedCheckDnsCheck) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCreated

`func (o *DetailedCheckDnsCheck) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DetailedCheckDnsCheck) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DetailedCheckDnsCheck) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DetailedCheckDnsCheck) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTags

`func (o *DetailedCheckDnsCheck) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetailedCheckDnsCheck) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetailedCheckDnsCheck) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DetailedCheckDnsCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpv6

`func (o *DetailedCheckDnsCheck) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DetailedCheckDnsCheck) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DetailedCheckDnsCheck) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *DetailedCheckDnsCheck) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


