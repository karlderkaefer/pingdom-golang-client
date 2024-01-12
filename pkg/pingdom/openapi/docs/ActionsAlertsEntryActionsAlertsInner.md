# ActionsAlertsEntryActionsAlertsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Name of alerted user | [optional] 
**Userid** | Pointer to **string** | Identifier of alerted user | [optional] 
**Checkid** | Pointer to **string** | Identifier of alerted user | [optional] 
**Time** | Pointer to **string** | Time of alert generation. Format UNIX time | [optional] 
**Via** | Pointer to **string** | Alert medium. apns - iphone, agcm - android | [optional] 
**Status** | Pointer to **string** | Alert status | [optional] 
**Messageshort** | Pointer to **string** | Short description of message | [optional] 
**Messagefull** | Pointer to **string** | Full message body | [optional] 
**Sentto** | Pointer to **string** | Target address, phone number etc | [optional] 
**Charged** | Pointer to **string** | True if your account was charged for this message | [optional] 

## Methods

### NewActionsAlertsEntryActionsAlertsInner

`func NewActionsAlertsEntryActionsAlertsInner() *ActionsAlertsEntryActionsAlertsInner`

NewActionsAlertsEntryActionsAlertsInner instantiates a new ActionsAlertsEntryActionsAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsAlertsEntryActionsAlertsInnerWithDefaults

`func NewActionsAlertsEntryActionsAlertsInnerWithDefaults() *ActionsAlertsEntryActionsAlertsInner`

NewActionsAlertsEntryActionsAlertsInnerWithDefaults instantiates a new ActionsAlertsEntryActionsAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ActionsAlertsEntryActionsAlertsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ActionsAlertsEntryActionsAlertsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ActionsAlertsEntryActionsAlertsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUserid

`func (o *ActionsAlertsEntryActionsAlertsInner) GetUserid() string`

GetUserid returns the Userid field if non-nil, zero value otherwise.

### GetUseridOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetUseridOk() (*string, bool)`

GetUseridOk returns a tuple with the Userid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserid

`func (o *ActionsAlertsEntryActionsAlertsInner) SetUserid(v string)`

SetUserid sets Userid field to given value.

### HasUserid

`func (o *ActionsAlertsEntryActionsAlertsInner) HasUserid() bool`

HasUserid returns a boolean if a field has been set.

### GetCheckid

`func (o *ActionsAlertsEntryActionsAlertsInner) GetCheckid() string`

GetCheckid returns the Checkid field if non-nil, zero value otherwise.

### GetCheckidOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetCheckidOk() (*string, bool)`

GetCheckidOk returns a tuple with the Checkid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckid

`func (o *ActionsAlertsEntryActionsAlertsInner) SetCheckid(v string)`

SetCheckid sets Checkid field to given value.

### HasCheckid

`func (o *ActionsAlertsEntryActionsAlertsInner) HasCheckid() bool`

HasCheckid returns a boolean if a field has been set.

### GetTime

`func (o *ActionsAlertsEntryActionsAlertsInner) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ActionsAlertsEntryActionsAlertsInner) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ActionsAlertsEntryActionsAlertsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetVia

`func (o *ActionsAlertsEntryActionsAlertsInner) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *ActionsAlertsEntryActionsAlertsInner) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *ActionsAlertsEntryActionsAlertsInner) HasVia() bool`

HasVia returns a boolean if a field has been set.

### GetStatus

`func (o *ActionsAlertsEntryActionsAlertsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionsAlertsEntryActionsAlertsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionsAlertsEntryActionsAlertsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessageshort

`func (o *ActionsAlertsEntryActionsAlertsInner) GetMessageshort() string`

GetMessageshort returns the Messageshort field if non-nil, zero value otherwise.

### GetMessageshortOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetMessageshortOk() (*string, bool)`

GetMessageshortOk returns a tuple with the Messageshort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageshort

`func (o *ActionsAlertsEntryActionsAlertsInner) SetMessageshort(v string)`

SetMessageshort sets Messageshort field to given value.

### HasMessageshort

`func (o *ActionsAlertsEntryActionsAlertsInner) HasMessageshort() bool`

HasMessageshort returns a boolean if a field has been set.

### GetMessagefull

`func (o *ActionsAlertsEntryActionsAlertsInner) GetMessagefull() string`

GetMessagefull returns the Messagefull field if non-nil, zero value otherwise.

### GetMessagefullOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetMessagefullOk() (*string, bool)`

GetMessagefullOk returns a tuple with the Messagefull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagefull

`func (o *ActionsAlertsEntryActionsAlertsInner) SetMessagefull(v string)`

SetMessagefull sets Messagefull field to given value.

### HasMessagefull

`func (o *ActionsAlertsEntryActionsAlertsInner) HasMessagefull() bool`

HasMessagefull returns a boolean if a field has been set.

### GetSentto

`func (o *ActionsAlertsEntryActionsAlertsInner) GetSentto() string`

GetSentto returns the Sentto field if non-nil, zero value otherwise.

### GetSenttoOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetSenttoOk() (*string, bool)`

GetSenttoOk returns a tuple with the Sentto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentto

`func (o *ActionsAlertsEntryActionsAlertsInner) SetSentto(v string)`

SetSentto sets Sentto field to given value.

### HasSentto

`func (o *ActionsAlertsEntryActionsAlertsInner) HasSentto() bool`

HasSentto returns a boolean if a field has been set.

### GetCharged

`func (o *ActionsAlertsEntryActionsAlertsInner) GetCharged() string`

GetCharged returns the Charged field if non-nil, zero value otherwise.

### GetChargedOk

`func (o *ActionsAlertsEntryActionsAlertsInner) GetChargedOk() (*string, bool)`

GetChargedOk returns a tuple with the Charged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharged

`func (o *ActionsAlertsEntryActionsAlertsInner) SetCharged(v string)`

SetCharged sets Charged field to given value.

### HasCharged

`func (o *ActionsAlertsEntryActionsAlertsInner) HasCharged() bool`

HasCharged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


