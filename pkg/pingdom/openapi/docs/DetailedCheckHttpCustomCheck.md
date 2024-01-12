# DetailedCheckHttpCustomCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DetailedHttpCustomAttributesType**](DetailedHttpCustomAttributesType.md) |  | [optional] 
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
**VerifyCertificate** | Pointer to **bool** | Treat target site as down if an invalid/unverifiable certificate is found. | [optional] [default to true]
**SslDownDaysBefore** | Pointer to **int32** | Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if &#x60;verify_certificate&#x60; is set to &#x60;false&#x60;. | [optional] [default to 0]

## Methods

### NewDetailedCheckHttpCustomCheck

`func NewDetailedCheckHttpCustomCheck() *DetailedCheckHttpCustomCheck`

NewDetailedCheckHttpCustomCheck instantiates a new DetailedCheckHttpCustomCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedCheckHttpCustomCheckWithDefaults

`func NewDetailedCheckHttpCustomCheckWithDefaults() *DetailedCheckHttpCustomCheck`

NewDetailedCheckHttpCustomCheckWithDefaults instantiates a new DetailedCheckHttpCustomCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DetailedCheckHttpCustomCheck) GetType() DetailedHttpCustomAttributesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetailedCheckHttpCustomCheck) GetTypeOk() (*DetailedHttpCustomAttributesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetailedCheckHttpCustomCheck) SetType(v DetailedHttpCustomAttributesType)`

SetType sets Type field to given value.

### HasType

`func (o *DetailedCheckHttpCustomCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProbeFilters

`func (o *DetailedCheckHttpCustomCheck) GetProbeFilters() []string`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *DetailedCheckHttpCustomCheck) GetProbeFiltersOk() (*[]string, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *DetailedCheckHttpCustomCheck) SetProbeFilters(v []string)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *DetailedCheckHttpCustomCheck) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetSendnotificationwhendown

`func (o *DetailedCheckHttpCustomCheck) GetSendnotificationwhendown() int32`

GetSendnotificationwhendown returns the Sendnotificationwhendown field if non-nil, zero value otherwise.

### GetSendnotificationwhendownOk

`func (o *DetailedCheckHttpCustomCheck) GetSendnotificationwhendownOk() (*int32, bool)`

GetSendnotificationwhendownOk returns a tuple with the Sendnotificationwhendown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendnotificationwhendown

`func (o *DetailedCheckHttpCustomCheck) SetSendnotificationwhendown(v int32)`

SetSendnotificationwhendown sets Sendnotificationwhendown field to given value.

### HasSendnotificationwhendown

`func (o *DetailedCheckHttpCustomCheck) HasSendnotificationwhendown() bool`

HasSendnotificationwhendown returns a boolean if a field has been set.

### GetNotifyagainevery

`func (o *DetailedCheckHttpCustomCheck) GetNotifyagainevery() int32`

GetNotifyagainevery returns the Notifyagainevery field if non-nil, zero value otherwise.

### GetNotifyagaineveryOk

`func (o *DetailedCheckHttpCustomCheck) GetNotifyagaineveryOk() (*int32, bool)`

GetNotifyagaineveryOk returns a tuple with the Notifyagainevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyagainevery

`func (o *DetailedCheckHttpCustomCheck) SetNotifyagainevery(v int32)`

SetNotifyagainevery sets Notifyagainevery field to given value.

### HasNotifyagainevery

`func (o *DetailedCheckHttpCustomCheck) HasNotifyagainevery() bool`

HasNotifyagainevery returns a boolean if a field has been set.

### GetNotifywhenbackup

`func (o *DetailedCheckHttpCustomCheck) GetNotifywhenbackup() bool`

GetNotifywhenbackup returns the Notifywhenbackup field if non-nil, zero value otherwise.

### GetNotifywhenbackupOk

`func (o *DetailedCheckHttpCustomCheck) GetNotifywhenbackupOk() (*bool, bool)`

GetNotifywhenbackupOk returns a tuple with the Notifywhenbackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifywhenbackup

`func (o *DetailedCheckHttpCustomCheck) SetNotifywhenbackup(v bool)`

SetNotifywhenbackup sets Notifywhenbackup field to given value.

### HasNotifywhenbackup

`func (o *DetailedCheckHttpCustomCheck) HasNotifywhenbackup() bool`

HasNotifywhenbackup returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *DetailedCheckHttpCustomCheck) GetResponsetimeThreshold() bool`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *DetailedCheckHttpCustomCheck) GetResponsetimeThresholdOk() (*bool, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *DetailedCheckHttpCustomCheck) SetResponsetimeThreshold(v bool)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *DetailedCheckHttpCustomCheck) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetCustomMessage

`func (o *DetailedCheckHttpCustomCheck) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *DetailedCheckHttpCustomCheck) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *DetailedCheckHttpCustomCheck) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *DetailedCheckHttpCustomCheck) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetIntegrationids

`func (o *DetailedCheckHttpCustomCheck) GetIntegrationids() []int32`

GetIntegrationids returns the Integrationids field if non-nil, zero value otherwise.

### GetIntegrationidsOk

`func (o *DetailedCheckHttpCustomCheck) GetIntegrationidsOk() (*[]int32, bool)`

GetIntegrationidsOk returns a tuple with the Integrationids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationids

`func (o *DetailedCheckHttpCustomCheck) SetIntegrationids(v []int32)`

SetIntegrationids sets Integrationids field to given value.

### HasIntegrationids

`func (o *DetailedCheckHttpCustomCheck) HasIntegrationids() bool`

HasIntegrationids returns a boolean if a field has been set.

### GetId

`func (o *DetailedCheckHttpCustomCheck) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedCheckHttpCustomCheck) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedCheckHttpCustomCheck) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedCheckHttpCustomCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DetailedCheckHttpCustomCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedCheckHttpCustomCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedCheckHttpCustomCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedCheckHttpCustomCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLasterrortime

`func (o *DetailedCheckHttpCustomCheck) GetLasterrortime() int32`

GetLasterrortime returns the Lasterrortime field if non-nil, zero value otherwise.

### GetLasterrortimeOk

`func (o *DetailedCheckHttpCustomCheck) GetLasterrortimeOk() (*int32, bool)`

GetLasterrortimeOk returns a tuple with the Lasterrortime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrortime

`func (o *DetailedCheckHttpCustomCheck) SetLasterrortime(v int32)`

SetLasterrortime sets Lasterrortime field to given value.

### HasLasterrortime

`func (o *DetailedCheckHttpCustomCheck) HasLasterrortime() bool`

HasLasterrortime returns a boolean if a field has been set.

### GetLasttesttime

`func (o *DetailedCheckHttpCustomCheck) GetLasttesttime() int32`

GetLasttesttime returns the Lasttesttime field if non-nil, zero value otherwise.

### GetLasttesttimeOk

`func (o *DetailedCheckHttpCustomCheck) GetLasttesttimeOk() (*int32, bool)`

GetLasttesttimeOk returns a tuple with the Lasttesttime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasttesttime

`func (o *DetailedCheckHttpCustomCheck) SetLasttesttime(v int32)`

SetLasttesttime sets Lasttesttime field to given value.

### HasLasttesttime

`func (o *DetailedCheckHttpCustomCheck) HasLasttesttime() bool`

HasLasttesttime returns a boolean if a field has been set.

### GetLastresponsetime

`func (o *DetailedCheckHttpCustomCheck) GetLastresponsetime() int32`

GetLastresponsetime returns the Lastresponsetime field if non-nil, zero value otherwise.

### GetLastresponsetimeOk

`func (o *DetailedCheckHttpCustomCheck) GetLastresponsetimeOk() (*int32, bool)`

GetLastresponsetimeOk returns a tuple with the Lastresponsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastresponsetime

`func (o *DetailedCheckHttpCustomCheck) SetLastresponsetime(v int32)`

SetLastresponsetime sets Lastresponsetime field to given value.

### HasLastresponsetime

`func (o *DetailedCheckHttpCustomCheck) HasLastresponsetime() bool`

HasLastresponsetime returns a boolean if a field has been set.

### GetLastdownstart

`func (o *DetailedCheckHttpCustomCheck) GetLastdownstart() int32`

GetLastdownstart returns the Lastdownstart field if non-nil, zero value otherwise.

### GetLastdownstartOk

`func (o *DetailedCheckHttpCustomCheck) GetLastdownstartOk() (*int32, bool)`

GetLastdownstartOk returns a tuple with the Lastdownstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownstart

`func (o *DetailedCheckHttpCustomCheck) SetLastdownstart(v int32)`

SetLastdownstart sets Lastdownstart field to given value.

### HasLastdownstart

`func (o *DetailedCheckHttpCustomCheck) HasLastdownstart() bool`

HasLastdownstart returns a boolean if a field has been set.

### GetLastdownend

`func (o *DetailedCheckHttpCustomCheck) GetLastdownend() int32`

GetLastdownend returns the Lastdownend field if non-nil, zero value otherwise.

### GetLastdownendOk

`func (o *DetailedCheckHttpCustomCheck) GetLastdownendOk() (*int32, bool)`

GetLastdownendOk returns a tuple with the Lastdownend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastdownend

`func (o *DetailedCheckHttpCustomCheck) SetLastdownend(v int32)`

SetLastdownend sets Lastdownend field to given value.

### HasLastdownend

`func (o *DetailedCheckHttpCustomCheck) HasLastdownend() bool`

HasLastdownend returns a boolean if a field has been set.

### GetStatus

`func (o *DetailedCheckHttpCustomCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailedCheckHttpCustomCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailedCheckHttpCustomCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetailedCheckHttpCustomCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResolution

`func (o *DetailedCheckHttpCustomCheck) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *DetailedCheckHttpCustomCheck) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *DetailedCheckHttpCustomCheck) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *DetailedCheckHttpCustomCheck) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetHostname

`func (o *DetailedCheckHttpCustomCheck) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DetailedCheckHttpCustomCheck) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DetailedCheckHttpCustomCheck) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DetailedCheckHttpCustomCheck) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetCreated

`func (o *DetailedCheckHttpCustomCheck) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DetailedCheckHttpCustomCheck) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DetailedCheckHttpCustomCheck) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DetailedCheckHttpCustomCheck) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTags

`func (o *DetailedCheckHttpCustomCheck) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetailedCheckHttpCustomCheck) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetailedCheckHttpCustomCheck) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DetailedCheckHttpCustomCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIpv6

`func (o *DetailedCheckHttpCustomCheck) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DetailedCheckHttpCustomCheck) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DetailedCheckHttpCustomCheck) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *DetailedCheckHttpCustomCheck) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *DetailedCheckHttpCustomCheck) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *DetailedCheckHttpCustomCheck) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *DetailedCheckHttpCustomCheck) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *DetailedCheckHttpCustomCheck) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *DetailedCheckHttpCustomCheck) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *DetailedCheckHttpCustomCheck) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *DetailedCheckHttpCustomCheck) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *DetailedCheckHttpCustomCheck) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


