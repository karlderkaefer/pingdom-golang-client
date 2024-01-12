# ModifyCheckSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Check name | [optional] 
**Host** | Pointer to **string** | Target host | [optional] 
**Paused** | Pointer to **bool** |  | [optional] [default to false]
**Resolution** | Pointer to **int32** | How often should the check be tested? (minutes) | [optional] [default to 5]
**Userids** | Pointer to **string** | User identifiers. For example userids&#x3D;154325,465231,765871 | [optional] 
**Sendnotificationwhendown** | Pointer to **int32** | Send notification when down X times | [optional] [default to 2]
**Notifyagainevery** | Pointer to **int32** | Notify again every n result. 0 means that no extra notifications will be sent. | [optional] [default to 0]
**Notifywhenbackup** | Pointer to **bool** | Notify when back up again | [optional] [default to true]
**Checkids** | Pointer to **string** | Identifiers of checks to modify in bulk. For example checkids&#x3D;1234,5678 | [optional] 
**Tags** | Pointer to **[]string** | List of tags for check. The maximum length of a tag is 64 characters. | [optional] 
**Addtags** | Pointer to **[]string** | Check tags to add in addition to current check tags | [optional] 
**ProbeFilters** | Pointer to **[]string** | Filters used for probe selections. Overwrites previous filters for check. To remove all filters from a check, use an empty value. Comma separated key:value pairs. Currently only region is supported. Possible values are &#39;EU&#39;, &#39;NA&#39;, &#39;APAC&#39; and &#39;LATAM&#39;. For example, \&quot;region: NA\&quot;. | [optional] 
**Ipv6** | Pointer to **bool** | Use ipv6 instead of ipv4, if an IP address is provided as host this will be overrided by the IP address version | [optional] 
**ResponsetimeThreshold** | Pointer to **int32** | Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.) | [optional] [default to 30000]
**Integrationids** | Pointer to **[]int32** | Integration identifiers. For example integrationids:[1,2,3]. | [optional] 
**Teamids** | Pointer to **string** | Teams to alert. Comma separated Integers. | [optional] 
**CustomMessage** | Pointer to **string** | Custom message that will be added to email and webhook alerts. | [optional] 
**Auth** | Pointer to **string** | Username and password colon separated for target SMTP authentication | [optional] 
**Url** | **string** | Path to target on server | 
**Encryption** | Pointer to **bool** | Connection encryption | [optional] 
**Port** | **int32** | Target port | 
**Shouldcontain** | Pointer to **string** | Target site should contain this string. Note! This parameter cannot be used together with the parameter “shouldnotcontain”, use only one of them in your request. | [optional] 
**Shouldnotcontain** | Pointer to **string** | Target site should NOT contain this string. Note! This parameter cannot be used together with the parameter “shouldcontain”, use only one of them in your request. | [optional] 
**Postdata** | Pointer to **string** | Data that should be posted to the web page, for example submission data for a sign-up or login form. The data needs to be formatted in the same way as a web browser would send it to the web server | [optional] 
**Requestheaders** | Pointer to **[]string** | Custom HTTP header. The entry value should contain a one-element string array. The element should contain &#x60;headerName&#x60; and &#x60;headerValue&#x60; colon-separated. To add more than one header send other parameters named &#x60;requestheaders{number}&#x60;. | [optional] 
**VerifyCertificate** | Pointer to **bool** | Treat target site as down if an invalid/unverifiable certificate is found. | [optional] [default to true]
**SslDownDaysBefore** | Pointer to **int32** | Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if &#x60;verify_certificate&#x60; is set to &#x60;false&#x60;. | [optional] [default to 0]
**Additionalurls** | Pointer to **string** | Full URL (including hostname) to target additional XML file | [optional] 
**Stringtosend** | **string** | String to send | 
**Stringtoexpect** | **string** | String to expect in response | 
**Nameserver** | **string** | DNS server to use | 
**Expectedip** | **string** | Expected IP | 

## Methods

### NewModifyCheckSettings

`func NewModifyCheckSettings(url string, port int32, stringtosend string, stringtoexpect string, nameserver string, expectedip string, ) *ModifyCheckSettings`

NewModifyCheckSettings instantiates a new ModifyCheckSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCheckSettingsWithDefaults

`func NewModifyCheckSettingsWithDefaults() *ModifyCheckSettings`

NewModifyCheckSettingsWithDefaults instantiates a new ModifyCheckSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModifyCheckSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyCheckSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyCheckSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyCheckSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *ModifyCheckSettings) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ModifyCheckSettings) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ModifyCheckSettings) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ModifyCheckSettings) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPaused

`func (o *ModifyCheckSettings) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ModifyCheckSettings) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ModifyCheckSettings) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ModifyCheckSettings) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetResolution

`func (o *ModifyCheckSettings) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ModifyCheckSettings) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ModifyCheckSettings) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ModifyCheckSettings) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetUserids

`func (o *ModifyCheckSettings) GetUserids() string`

GetUserids returns the Userids field if non-nil, zero value otherwise.

### GetUseridsOk

`func (o *ModifyCheckSettings) GetUseridsOk() (*string, bool)`

GetUseridsOk returns a tuple with the Userids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserids

`func (o *ModifyCheckSettings) SetUserids(v string)`

SetUserids sets Userids field to given value.

### HasUserids

`func (o *ModifyCheckSettings) HasUserids() bool`

HasUserids returns a boolean if a field has been set.

### GetSendnotificationwhendown

`func (o *ModifyCheckSettings) GetSendnotificationwhendown() int32`

GetSendnotificationwhendown returns the Sendnotificationwhendown field if non-nil, zero value otherwise.

### GetSendnotificationwhendownOk

`func (o *ModifyCheckSettings) GetSendnotificationwhendownOk() (*int32, bool)`

GetSendnotificationwhendownOk returns a tuple with the Sendnotificationwhendown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendnotificationwhendown

`func (o *ModifyCheckSettings) SetSendnotificationwhendown(v int32)`

SetSendnotificationwhendown sets Sendnotificationwhendown field to given value.

### HasSendnotificationwhendown

`func (o *ModifyCheckSettings) HasSendnotificationwhendown() bool`

HasSendnotificationwhendown returns a boolean if a field has been set.

### GetNotifyagainevery

`func (o *ModifyCheckSettings) GetNotifyagainevery() int32`

GetNotifyagainevery returns the Notifyagainevery field if non-nil, zero value otherwise.

### GetNotifyagaineveryOk

`func (o *ModifyCheckSettings) GetNotifyagaineveryOk() (*int32, bool)`

GetNotifyagaineveryOk returns a tuple with the Notifyagainevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyagainevery

`func (o *ModifyCheckSettings) SetNotifyagainevery(v int32)`

SetNotifyagainevery sets Notifyagainevery field to given value.

### HasNotifyagainevery

`func (o *ModifyCheckSettings) HasNotifyagainevery() bool`

HasNotifyagainevery returns a boolean if a field has been set.

### GetNotifywhenbackup

`func (o *ModifyCheckSettings) GetNotifywhenbackup() bool`

GetNotifywhenbackup returns the Notifywhenbackup field if non-nil, zero value otherwise.

### GetNotifywhenbackupOk

`func (o *ModifyCheckSettings) GetNotifywhenbackupOk() (*bool, bool)`

GetNotifywhenbackupOk returns a tuple with the Notifywhenbackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifywhenbackup

`func (o *ModifyCheckSettings) SetNotifywhenbackup(v bool)`

SetNotifywhenbackup sets Notifywhenbackup field to given value.

### HasNotifywhenbackup

`func (o *ModifyCheckSettings) HasNotifywhenbackup() bool`

HasNotifywhenbackup returns a boolean if a field has been set.

### GetCheckids

`func (o *ModifyCheckSettings) GetCheckids() string`

GetCheckids returns the Checkids field if non-nil, zero value otherwise.

### GetCheckidsOk

`func (o *ModifyCheckSettings) GetCheckidsOk() (*string, bool)`

GetCheckidsOk returns a tuple with the Checkids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckids

`func (o *ModifyCheckSettings) SetCheckids(v string)`

SetCheckids sets Checkids field to given value.

### HasCheckids

`func (o *ModifyCheckSettings) HasCheckids() bool`

HasCheckids returns a boolean if a field has been set.

### GetTags

`func (o *ModifyCheckSettings) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModifyCheckSettings) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModifyCheckSettings) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModifyCheckSettings) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAddtags

`func (o *ModifyCheckSettings) GetAddtags() []string`

GetAddtags returns the Addtags field if non-nil, zero value otherwise.

### GetAddtagsOk

`func (o *ModifyCheckSettings) GetAddtagsOk() (*[]string, bool)`

GetAddtagsOk returns a tuple with the Addtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddtags

`func (o *ModifyCheckSettings) SetAddtags(v []string)`

SetAddtags sets Addtags field to given value.

### HasAddtags

`func (o *ModifyCheckSettings) HasAddtags() bool`

HasAddtags returns a boolean if a field has been set.

### GetProbeFilters

`func (o *ModifyCheckSettings) GetProbeFilters() []string`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *ModifyCheckSettings) GetProbeFiltersOk() (*[]string, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *ModifyCheckSettings) SetProbeFilters(v []string)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *ModifyCheckSettings) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetIpv6

`func (o *ModifyCheckSettings) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ModifyCheckSettings) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ModifyCheckSettings) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ModifyCheckSettings) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *ModifyCheckSettings) GetResponsetimeThreshold() int32`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *ModifyCheckSettings) GetResponsetimeThresholdOk() (*int32, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *ModifyCheckSettings) SetResponsetimeThreshold(v int32)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *ModifyCheckSettings) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetIntegrationids

`func (o *ModifyCheckSettings) GetIntegrationids() []int32`

GetIntegrationids returns the Integrationids field if non-nil, zero value otherwise.

### GetIntegrationidsOk

`func (o *ModifyCheckSettings) GetIntegrationidsOk() (*[]int32, bool)`

GetIntegrationidsOk returns a tuple with the Integrationids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationids

`func (o *ModifyCheckSettings) SetIntegrationids(v []int32)`

SetIntegrationids sets Integrationids field to given value.

### HasIntegrationids

`func (o *ModifyCheckSettings) HasIntegrationids() bool`

HasIntegrationids returns a boolean if a field has been set.

### GetTeamids

`func (o *ModifyCheckSettings) GetTeamids() string`

GetTeamids returns the Teamids field if non-nil, zero value otherwise.

### GetTeamidsOk

`func (o *ModifyCheckSettings) GetTeamidsOk() (*string, bool)`

GetTeamidsOk returns a tuple with the Teamids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamids

`func (o *ModifyCheckSettings) SetTeamids(v string)`

SetTeamids sets Teamids field to given value.

### HasTeamids

`func (o *ModifyCheckSettings) HasTeamids() bool`

HasTeamids returns a boolean if a field has been set.

### GetCustomMessage

`func (o *ModifyCheckSettings) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *ModifyCheckSettings) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *ModifyCheckSettings) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *ModifyCheckSettings) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetAuth

`func (o *ModifyCheckSettings) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ModifyCheckSettings) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ModifyCheckSettings) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ModifyCheckSettings) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetUrl

`func (o *ModifyCheckSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModifyCheckSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModifyCheckSettings) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEncryption

`func (o *ModifyCheckSettings) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ModifyCheckSettings) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ModifyCheckSettings) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *ModifyCheckSettings) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *ModifyCheckSettings) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModifyCheckSettings) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModifyCheckSettings) SetPort(v int32)`

SetPort sets Port field to given value.


### GetShouldcontain

`func (o *ModifyCheckSettings) GetShouldcontain() string`

GetShouldcontain returns the Shouldcontain field if non-nil, zero value otherwise.

### GetShouldcontainOk

`func (o *ModifyCheckSettings) GetShouldcontainOk() (*string, bool)`

GetShouldcontainOk returns a tuple with the Shouldcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldcontain

`func (o *ModifyCheckSettings) SetShouldcontain(v string)`

SetShouldcontain sets Shouldcontain field to given value.

### HasShouldcontain

`func (o *ModifyCheckSettings) HasShouldcontain() bool`

HasShouldcontain returns a boolean if a field has been set.

### GetShouldnotcontain

`func (o *ModifyCheckSettings) GetShouldnotcontain() string`

GetShouldnotcontain returns the Shouldnotcontain field if non-nil, zero value otherwise.

### GetShouldnotcontainOk

`func (o *ModifyCheckSettings) GetShouldnotcontainOk() (*string, bool)`

GetShouldnotcontainOk returns a tuple with the Shouldnotcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldnotcontain

`func (o *ModifyCheckSettings) SetShouldnotcontain(v string)`

SetShouldnotcontain sets Shouldnotcontain field to given value.

### HasShouldnotcontain

`func (o *ModifyCheckSettings) HasShouldnotcontain() bool`

HasShouldnotcontain returns a boolean if a field has been set.

### GetPostdata

`func (o *ModifyCheckSettings) GetPostdata() string`

GetPostdata returns the Postdata field if non-nil, zero value otherwise.

### GetPostdataOk

`func (o *ModifyCheckSettings) GetPostdataOk() (*string, bool)`

GetPostdataOk returns a tuple with the Postdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostdata

`func (o *ModifyCheckSettings) SetPostdata(v string)`

SetPostdata sets Postdata field to given value.

### HasPostdata

`func (o *ModifyCheckSettings) HasPostdata() bool`

HasPostdata returns a boolean if a field has been set.

### GetRequestheaders

`func (o *ModifyCheckSettings) GetRequestheaders() []string`

GetRequestheaders returns the Requestheaders field if non-nil, zero value otherwise.

### GetRequestheadersOk

`func (o *ModifyCheckSettings) GetRequestheadersOk() (*[]string, bool)`

GetRequestheadersOk returns a tuple with the Requestheaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestheaders

`func (o *ModifyCheckSettings) SetRequestheaders(v []string)`

SetRequestheaders sets Requestheaders field to given value.

### HasRequestheaders

`func (o *ModifyCheckSettings) HasRequestheaders() bool`

HasRequestheaders returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *ModifyCheckSettings) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *ModifyCheckSettings) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *ModifyCheckSettings) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *ModifyCheckSettings) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *ModifyCheckSettings) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *ModifyCheckSettings) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *ModifyCheckSettings) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *ModifyCheckSettings) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.

### GetAdditionalurls

`func (o *ModifyCheckSettings) GetAdditionalurls() string`

GetAdditionalurls returns the Additionalurls field if non-nil, zero value otherwise.

### GetAdditionalurlsOk

`func (o *ModifyCheckSettings) GetAdditionalurlsOk() (*string, bool)`

GetAdditionalurlsOk returns a tuple with the Additionalurls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalurls

`func (o *ModifyCheckSettings) SetAdditionalurls(v string)`

SetAdditionalurls sets Additionalurls field to given value.

### HasAdditionalurls

`func (o *ModifyCheckSettings) HasAdditionalurls() bool`

HasAdditionalurls returns a boolean if a field has been set.

### GetStringtosend

`func (o *ModifyCheckSettings) GetStringtosend() string`

GetStringtosend returns the Stringtosend field if non-nil, zero value otherwise.

### GetStringtosendOk

`func (o *ModifyCheckSettings) GetStringtosendOk() (*string, bool)`

GetStringtosendOk returns a tuple with the Stringtosend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtosend

`func (o *ModifyCheckSettings) SetStringtosend(v string)`

SetStringtosend sets Stringtosend field to given value.


### GetStringtoexpect

`func (o *ModifyCheckSettings) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *ModifyCheckSettings) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *ModifyCheckSettings) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.


### GetNameserver

`func (o *ModifyCheckSettings) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *ModifyCheckSettings) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *ModifyCheckSettings) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.


### GetExpectedip

`func (o *ModifyCheckSettings) GetExpectedip() string`

GetExpectedip returns the Expectedip field if non-nil, zero value otherwise.

### GetExpectedipOk

`func (o *ModifyCheckSettings) GetExpectedipOk() (*string, bool)`

GetExpectedipOk returns a tuple with the Expectedip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedip

`func (o *ModifyCheckSettings) SetExpectedip(v string)`

SetExpectedip sets Expectedip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


