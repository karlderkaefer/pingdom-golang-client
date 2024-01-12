# CreateCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Check name | 
**Host** | **string** | Target host | 
**Type** | **string** | Type of check | 
**Paused** | Pointer to **bool** |  | [optional] [default to false]
**Resolution** | Pointer to **int32** | How often should the check be tested? (minutes) | [optional] [default to 5]
**Userids** | Pointer to **string** | User identifiers. For example userids&#x3D;154325,465231,765871 | [optional] 
**Sendnotificationwhendown** | Pointer to **int32** | Send notification when down X times | [optional] [default to 2]
**Notifyagainevery** | Pointer to **int32** | Notify again every n result. 0 means that no extra notifications will be sent. | [optional] [default to 0]
**Notifywhenbackup** | Pointer to **bool** | Notify when back up again | [optional] [default to true]
**Tags** | Pointer to **[]string** | List of tags for check. The maximum length of a tag is 64 characters. | [optional] 
**ProbeFilters** | Pointer to **[]string** | Filters used for probe selections. Overwrites previous filters for check. To remove all filters from a check, use an empty value. Comma separated key:value pairs. Currently only region is supported. Possible values are &#39;EU&#39;, &#39;NA&#39;, &#39;APAC&#39; and &#39;LATAM&#39;. For example, \&quot;region: NA\&quot;. | [optional] 
**Ipv6** | Pointer to **bool** | Use ipv6 instead of ipv4, if an IP address is provided as host this will be overrided by the IP address version | [optional] [default to false]
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

### NewCreateCheck

`func NewCreateCheck(name string, host string, type_ string, url string, port int32, stringtosend string, stringtoexpect string, nameserver string, expectedip string, ) *CreateCheck`

NewCreateCheck instantiates a new CreateCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckWithDefaults

`func NewCreateCheckWithDefaults() *CreateCheck`

NewCreateCheckWithDefaults instantiates a new CreateCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCheck) SetName(v string)`

SetName sets Name field to given value.


### GetHost

`func (o *CreateCheck) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateCheck) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateCheck) SetHost(v string)`

SetHost sets Host field to given value.


### GetType

`func (o *CreateCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCheck) SetType(v string)`

SetType sets Type field to given value.


### GetPaused

`func (o *CreateCheck) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *CreateCheck) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *CreateCheck) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *CreateCheck) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetResolution

`func (o *CreateCheck) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *CreateCheck) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *CreateCheck) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *CreateCheck) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetUserids

`func (o *CreateCheck) GetUserids() string`

GetUserids returns the Userids field if non-nil, zero value otherwise.

### GetUseridsOk

`func (o *CreateCheck) GetUseridsOk() (*string, bool)`

GetUseridsOk returns a tuple with the Userids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserids

`func (o *CreateCheck) SetUserids(v string)`

SetUserids sets Userids field to given value.

### HasUserids

`func (o *CreateCheck) HasUserids() bool`

HasUserids returns a boolean if a field has been set.

### GetSendnotificationwhendown

`func (o *CreateCheck) GetSendnotificationwhendown() int32`

GetSendnotificationwhendown returns the Sendnotificationwhendown field if non-nil, zero value otherwise.

### GetSendnotificationwhendownOk

`func (o *CreateCheck) GetSendnotificationwhendownOk() (*int32, bool)`

GetSendnotificationwhendownOk returns a tuple with the Sendnotificationwhendown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendnotificationwhendown

`func (o *CreateCheck) SetSendnotificationwhendown(v int32)`

SetSendnotificationwhendown sets Sendnotificationwhendown field to given value.

### HasSendnotificationwhendown

`func (o *CreateCheck) HasSendnotificationwhendown() bool`

HasSendnotificationwhendown returns a boolean if a field has been set.

### GetNotifyagainevery

`func (o *CreateCheck) GetNotifyagainevery() int32`

GetNotifyagainevery returns the Notifyagainevery field if non-nil, zero value otherwise.

### GetNotifyagaineveryOk

`func (o *CreateCheck) GetNotifyagaineveryOk() (*int32, bool)`

GetNotifyagaineveryOk returns a tuple with the Notifyagainevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyagainevery

`func (o *CreateCheck) SetNotifyagainevery(v int32)`

SetNotifyagainevery sets Notifyagainevery field to given value.

### HasNotifyagainevery

`func (o *CreateCheck) HasNotifyagainevery() bool`

HasNotifyagainevery returns a boolean if a field has been set.

### GetNotifywhenbackup

`func (o *CreateCheck) GetNotifywhenbackup() bool`

GetNotifywhenbackup returns the Notifywhenbackup field if non-nil, zero value otherwise.

### GetNotifywhenbackupOk

`func (o *CreateCheck) GetNotifywhenbackupOk() (*bool, bool)`

GetNotifywhenbackupOk returns a tuple with the Notifywhenbackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifywhenbackup

`func (o *CreateCheck) SetNotifywhenbackup(v bool)`

SetNotifywhenbackup sets Notifywhenbackup field to given value.

### HasNotifywhenbackup

`func (o *CreateCheck) HasNotifywhenbackup() bool`

HasNotifywhenbackup returns a boolean if a field has been set.

### GetTags

`func (o *CreateCheck) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCheck) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCheck) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProbeFilters

`func (o *CreateCheck) GetProbeFilters() []string`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *CreateCheck) GetProbeFiltersOk() (*[]string, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *CreateCheck) SetProbeFilters(v []string)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *CreateCheck) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetIpv6

`func (o *CreateCheck) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateCheck) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateCheck) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateCheck) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *CreateCheck) GetResponsetimeThreshold() int32`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *CreateCheck) GetResponsetimeThresholdOk() (*int32, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *CreateCheck) SetResponsetimeThreshold(v int32)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *CreateCheck) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetIntegrationids

`func (o *CreateCheck) GetIntegrationids() []int32`

GetIntegrationids returns the Integrationids field if non-nil, zero value otherwise.

### GetIntegrationidsOk

`func (o *CreateCheck) GetIntegrationidsOk() (*[]int32, bool)`

GetIntegrationidsOk returns a tuple with the Integrationids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationids

`func (o *CreateCheck) SetIntegrationids(v []int32)`

SetIntegrationids sets Integrationids field to given value.

### HasIntegrationids

`func (o *CreateCheck) HasIntegrationids() bool`

HasIntegrationids returns a boolean if a field has been set.

### GetTeamids

`func (o *CreateCheck) GetTeamids() string`

GetTeamids returns the Teamids field if non-nil, zero value otherwise.

### GetTeamidsOk

`func (o *CreateCheck) GetTeamidsOk() (*string, bool)`

GetTeamidsOk returns a tuple with the Teamids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamids

`func (o *CreateCheck) SetTeamids(v string)`

SetTeamids sets Teamids field to given value.

### HasTeamids

`func (o *CreateCheck) HasTeamids() bool`

HasTeamids returns a boolean if a field has been set.

### GetCustomMessage

`func (o *CreateCheck) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *CreateCheck) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *CreateCheck) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *CreateCheck) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetAuth

`func (o *CreateCheck) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CreateCheck) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CreateCheck) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *CreateCheck) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetUrl

`func (o *CreateCheck) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateCheck) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateCheck) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEncryption

`func (o *CreateCheck) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CreateCheck) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CreateCheck) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *CreateCheck) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *CreateCheck) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateCheck) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateCheck) SetPort(v int32)`

SetPort sets Port field to given value.


### GetShouldcontain

`func (o *CreateCheck) GetShouldcontain() string`

GetShouldcontain returns the Shouldcontain field if non-nil, zero value otherwise.

### GetShouldcontainOk

`func (o *CreateCheck) GetShouldcontainOk() (*string, bool)`

GetShouldcontainOk returns a tuple with the Shouldcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldcontain

`func (o *CreateCheck) SetShouldcontain(v string)`

SetShouldcontain sets Shouldcontain field to given value.

### HasShouldcontain

`func (o *CreateCheck) HasShouldcontain() bool`

HasShouldcontain returns a boolean if a field has been set.

### GetShouldnotcontain

`func (o *CreateCheck) GetShouldnotcontain() string`

GetShouldnotcontain returns the Shouldnotcontain field if non-nil, zero value otherwise.

### GetShouldnotcontainOk

`func (o *CreateCheck) GetShouldnotcontainOk() (*string, bool)`

GetShouldnotcontainOk returns a tuple with the Shouldnotcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldnotcontain

`func (o *CreateCheck) SetShouldnotcontain(v string)`

SetShouldnotcontain sets Shouldnotcontain field to given value.

### HasShouldnotcontain

`func (o *CreateCheck) HasShouldnotcontain() bool`

HasShouldnotcontain returns a boolean if a field has been set.

### GetPostdata

`func (o *CreateCheck) GetPostdata() string`

GetPostdata returns the Postdata field if non-nil, zero value otherwise.

### GetPostdataOk

`func (o *CreateCheck) GetPostdataOk() (*string, bool)`

GetPostdataOk returns a tuple with the Postdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostdata

`func (o *CreateCheck) SetPostdata(v string)`

SetPostdata sets Postdata field to given value.

### HasPostdata

`func (o *CreateCheck) HasPostdata() bool`

HasPostdata returns a boolean if a field has been set.

### GetRequestheaders

`func (o *CreateCheck) GetRequestheaders() []string`

GetRequestheaders returns the Requestheaders field if non-nil, zero value otherwise.

### GetRequestheadersOk

`func (o *CreateCheck) GetRequestheadersOk() (*[]string, bool)`

GetRequestheadersOk returns a tuple with the Requestheaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestheaders

`func (o *CreateCheck) SetRequestheaders(v []string)`

SetRequestheaders sets Requestheaders field to given value.

### HasRequestheaders

`func (o *CreateCheck) HasRequestheaders() bool`

HasRequestheaders returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *CreateCheck) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *CreateCheck) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *CreateCheck) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *CreateCheck) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *CreateCheck) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *CreateCheck) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *CreateCheck) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *CreateCheck) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.

### GetAdditionalurls

`func (o *CreateCheck) GetAdditionalurls() string`

GetAdditionalurls returns the Additionalurls field if non-nil, zero value otherwise.

### GetAdditionalurlsOk

`func (o *CreateCheck) GetAdditionalurlsOk() (*string, bool)`

GetAdditionalurlsOk returns a tuple with the Additionalurls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalurls

`func (o *CreateCheck) SetAdditionalurls(v string)`

SetAdditionalurls sets Additionalurls field to given value.

### HasAdditionalurls

`func (o *CreateCheck) HasAdditionalurls() bool`

HasAdditionalurls returns a boolean if a field has been set.

### GetStringtosend

`func (o *CreateCheck) GetStringtosend() string`

GetStringtosend returns the Stringtosend field if non-nil, zero value otherwise.

### GetStringtosendOk

`func (o *CreateCheck) GetStringtosendOk() (*string, bool)`

GetStringtosendOk returns a tuple with the Stringtosend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtosend

`func (o *CreateCheck) SetStringtosend(v string)`

SetStringtosend sets Stringtosend field to given value.


### GetStringtoexpect

`func (o *CreateCheck) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *CreateCheck) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *CreateCheck) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.


### GetNameserver

`func (o *CreateCheck) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *CreateCheck) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *CreateCheck) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.


### GetExpectedip

`func (o *CreateCheck) GetExpectedip() string`

GetExpectedip returns the Expectedip field if non-nil, zero value otherwise.

### GetExpectedipOk

`func (o *CreateCheck) GetExpectedipOk() (*string, bool)`

GetExpectedipOk returns a tuple with the Expectedip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedip

`func (o *CreateCheck) SetExpectedip(v string)`

SetExpectedip sets Expectedip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


