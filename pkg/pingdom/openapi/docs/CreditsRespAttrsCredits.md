# CreditsRespAttrsCredits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checklimit** | Pointer to **int32** | Total number of check slots on this account | [optional] 
**Availablechecks** | Pointer to **int32** | Free check slots available for new checks | [optional] 
**Useddefault** | Pointer to **int32** | Total number of default check slots | [optional] 
**Usedtransaction** | Pointer to **int32** | Total number of transaction check slots | [optional] 
**Availablesms** | Pointer to **int32** | SMS credits remaining on this account | [optional] 
**Availablesmstests** | Pointer to **int32** | SMS provider tests remaining on this account | [optional] 
**Autofillsms** | Pointer to **bool** | Auto fill SMS | [optional] 
**AutofillsmsAmount** | Pointer to **int32** | The amount of sms to purchase when \&quot;autofillsms_when_left\&quot; is triggered | [optional] 
**AutofillsmsWhenLeft** | Pointer to **int32** | The amount of sms left that is going to trigger purchase of sms | [optional] 
**MaxSmsOverage** | Pointer to **int32** | The amount of overage SMSes that may be used, or null if SMS overage is not enabled. | [optional] 
**Availablerumsites** | Pointer to **int32** | Open RUM site slots available | [optional] 
**Usedrumsites** | Pointer to **int32** | Number of used RUM sites | [optional] 
**Maxrumfilters** | Pointer to **int32** | Number of maximum rum filters | [optional] 
**Maxrumpageviews** | Pointer to **int32** | Number of maximum pageviews per month | [optional] 

## Methods

### NewCreditsRespAttrsCredits

`func NewCreditsRespAttrsCredits() *CreditsRespAttrsCredits`

NewCreditsRespAttrsCredits instantiates a new CreditsRespAttrsCredits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditsRespAttrsCreditsWithDefaults

`func NewCreditsRespAttrsCreditsWithDefaults() *CreditsRespAttrsCredits`

NewCreditsRespAttrsCreditsWithDefaults instantiates a new CreditsRespAttrsCredits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecklimit

`func (o *CreditsRespAttrsCredits) GetChecklimit() int32`

GetChecklimit returns the Checklimit field if non-nil, zero value otherwise.

### GetChecklimitOk

`func (o *CreditsRespAttrsCredits) GetChecklimitOk() (*int32, bool)`

GetChecklimitOk returns a tuple with the Checklimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklimit

`func (o *CreditsRespAttrsCredits) SetChecklimit(v int32)`

SetChecklimit sets Checklimit field to given value.

### HasChecklimit

`func (o *CreditsRespAttrsCredits) HasChecklimit() bool`

HasChecklimit returns a boolean if a field has been set.

### GetAvailablechecks

`func (o *CreditsRespAttrsCredits) GetAvailablechecks() int32`

GetAvailablechecks returns the Availablechecks field if non-nil, zero value otherwise.

### GetAvailablechecksOk

`func (o *CreditsRespAttrsCredits) GetAvailablechecksOk() (*int32, bool)`

GetAvailablechecksOk returns a tuple with the Availablechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablechecks

`func (o *CreditsRespAttrsCredits) SetAvailablechecks(v int32)`

SetAvailablechecks sets Availablechecks field to given value.

### HasAvailablechecks

`func (o *CreditsRespAttrsCredits) HasAvailablechecks() bool`

HasAvailablechecks returns a boolean if a field has been set.

### GetUseddefault

`func (o *CreditsRespAttrsCredits) GetUseddefault() int32`

GetUseddefault returns the Useddefault field if non-nil, zero value otherwise.

### GetUseddefaultOk

`func (o *CreditsRespAttrsCredits) GetUseddefaultOk() (*int32, bool)`

GetUseddefaultOk returns a tuple with the Useddefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseddefault

`func (o *CreditsRespAttrsCredits) SetUseddefault(v int32)`

SetUseddefault sets Useddefault field to given value.

### HasUseddefault

`func (o *CreditsRespAttrsCredits) HasUseddefault() bool`

HasUseddefault returns a boolean if a field has been set.

### GetUsedtransaction

`func (o *CreditsRespAttrsCredits) GetUsedtransaction() int32`

GetUsedtransaction returns the Usedtransaction field if non-nil, zero value otherwise.

### GetUsedtransactionOk

`func (o *CreditsRespAttrsCredits) GetUsedtransactionOk() (*int32, bool)`

GetUsedtransactionOk returns a tuple with the Usedtransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedtransaction

`func (o *CreditsRespAttrsCredits) SetUsedtransaction(v int32)`

SetUsedtransaction sets Usedtransaction field to given value.

### HasUsedtransaction

`func (o *CreditsRespAttrsCredits) HasUsedtransaction() bool`

HasUsedtransaction returns a boolean if a field has been set.

### GetAvailablesms

`func (o *CreditsRespAttrsCredits) GetAvailablesms() int32`

GetAvailablesms returns the Availablesms field if non-nil, zero value otherwise.

### GetAvailablesmsOk

`func (o *CreditsRespAttrsCredits) GetAvailablesmsOk() (*int32, bool)`

GetAvailablesmsOk returns a tuple with the Availablesms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablesms

`func (o *CreditsRespAttrsCredits) SetAvailablesms(v int32)`

SetAvailablesms sets Availablesms field to given value.

### HasAvailablesms

`func (o *CreditsRespAttrsCredits) HasAvailablesms() bool`

HasAvailablesms returns a boolean if a field has been set.

### GetAvailablesmstests

`func (o *CreditsRespAttrsCredits) GetAvailablesmstests() int32`

GetAvailablesmstests returns the Availablesmstests field if non-nil, zero value otherwise.

### GetAvailablesmstestsOk

`func (o *CreditsRespAttrsCredits) GetAvailablesmstestsOk() (*int32, bool)`

GetAvailablesmstestsOk returns a tuple with the Availablesmstests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablesmstests

`func (o *CreditsRespAttrsCredits) SetAvailablesmstests(v int32)`

SetAvailablesmstests sets Availablesmstests field to given value.

### HasAvailablesmstests

`func (o *CreditsRespAttrsCredits) HasAvailablesmstests() bool`

HasAvailablesmstests returns a boolean if a field has been set.

### GetAutofillsms

`func (o *CreditsRespAttrsCredits) GetAutofillsms() bool`

GetAutofillsms returns the Autofillsms field if non-nil, zero value otherwise.

### GetAutofillsmsOk

`func (o *CreditsRespAttrsCredits) GetAutofillsmsOk() (*bool, bool)`

GetAutofillsmsOk returns a tuple with the Autofillsms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillsms

`func (o *CreditsRespAttrsCredits) SetAutofillsms(v bool)`

SetAutofillsms sets Autofillsms field to given value.

### HasAutofillsms

`func (o *CreditsRespAttrsCredits) HasAutofillsms() bool`

HasAutofillsms returns a boolean if a field has been set.

### GetAutofillsmsAmount

`func (o *CreditsRespAttrsCredits) GetAutofillsmsAmount() int32`

GetAutofillsmsAmount returns the AutofillsmsAmount field if non-nil, zero value otherwise.

### GetAutofillsmsAmountOk

`func (o *CreditsRespAttrsCredits) GetAutofillsmsAmountOk() (*int32, bool)`

GetAutofillsmsAmountOk returns a tuple with the AutofillsmsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillsmsAmount

`func (o *CreditsRespAttrsCredits) SetAutofillsmsAmount(v int32)`

SetAutofillsmsAmount sets AutofillsmsAmount field to given value.

### HasAutofillsmsAmount

`func (o *CreditsRespAttrsCredits) HasAutofillsmsAmount() bool`

HasAutofillsmsAmount returns a boolean if a field has been set.

### GetAutofillsmsWhenLeft

`func (o *CreditsRespAttrsCredits) GetAutofillsmsWhenLeft() int32`

GetAutofillsmsWhenLeft returns the AutofillsmsWhenLeft field if non-nil, zero value otherwise.

### GetAutofillsmsWhenLeftOk

`func (o *CreditsRespAttrsCredits) GetAutofillsmsWhenLeftOk() (*int32, bool)`

GetAutofillsmsWhenLeftOk returns a tuple with the AutofillsmsWhenLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillsmsWhenLeft

`func (o *CreditsRespAttrsCredits) SetAutofillsmsWhenLeft(v int32)`

SetAutofillsmsWhenLeft sets AutofillsmsWhenLeft field to given value.

### HasAutofillsmsWhenLeft

`func (o *CreditsRespAttrsCredits) HasAutofillsmsWhenLeft() bool`

HasAutofillsmsWhenLeft returns a boolean if a field has been set.

### GetMaxSmsOverage

`func (o *CreditsRespAttrsCredits) GetMaxSmsOverage() int32`

GetMaxSmsOverage returns the MaxSmsOverage field if non-nil, zero value otherwise.

### GetMaxSmsOverageOk

`func (o *CreditsRespAttrsCredits) GetMaxSmsOverageOk() (*int32, bool)`

GetMaxSmsOverageOk returns a tuple with the MaxSmsOverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSmsOverage

`func (o *CreditsRespAttrsCredits) SetMaxSmsOverage(v int32)`

SetMaxSmsOverage sets MaxSmsOverage field to given value.

### HasMaxSmsOverage

`func (o *CreditsRespAttrsCredits) HasMaxSmsOverage() bool`

HasMaxSmsOverage returns a boolean if a field has been set.

### GetAvailablerumsites

`func (o *CreditsRespAttrsCredits) GetAvailablerumsites() int32`

GetAvailablerumsites returns the Availablerumsites field if non-nil, zero value otherwise.

### GetAvailablerumsitesOk

`func (o *CreditsRespAttrsCredits) GetAvailablerumsitesOk() (*int32, bool)`

GetAvailablerumsitesOk returns a tuple with the Availablerumsites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablerumsites

`func (o *CreditsRespAttrsCredits) SetAvailablerumsites(v int32)`

SetAvailablerumsites sets Availablerumsites field to given value.

### HasAvailablerumsites

`func (o *CreditsRespAttrsCredits) HasAvailablerumsites() bool`

HasAvailablerumsites returns a boolean if a field has been set.

### GetUsedrumsites

`func (o *CreditsRespAttrsCredits) GetUsedrumsites() int32`

GetUsedrumsites returns the Usedrumsites field if non-nil, zero value otherwise.

### GetUsedrumsitesOk

`func (o *CreditsRespAttrsCredits) GetUsedrumsitesOk() (*int32, bool)`

GetUsedrumsitesOk returns a tuple with the Usedrumsites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedrumsites

`func (o *CreditsRespAttrsCredits) SetUsedrumsites(v int32)`

SetUsedrumsites sets Usedrumsites field to given value.

### HasUsedrumsites

`func (o *CreditsRespAttrsCredits) HasUsedrumsites() bool`

HasUsedrumsites returns a boolean if a field has been set.

### GetMaxrumfilters

`func (o *CreditsRespAttrsCredits) GetMaxrumfilters() int32`

GetMaxrumfilters returns the Maxrumfilters field if non-nil, zero value otherwise.

### GetMaxrumfiltersOk

`func (o *CreditsRespAttrsCredits) GetMaxrumfiltersOk() (*int32, bool)`

GetMaxrumfiltersOk returns a tuple with the Maxrumfilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxrumfilters

`func (o *CreditsRespAttrsCredits) SetMaxrumfilters(v int32)`

SetMaxrumfilters sets Maxrumfilters field to given value.

### HasMaxrumfilters

`func (o *CreditsRespAttrsCredits) HasMaxrumfilters() bool`

HasMaxrumfilters returns a boolean if a field has been set.

### GetMaxrumpageviews

`func (o *CreditsRespAttrsCredits) GetMaxrumpageviews() int32`

GetMaxrumpageviews returns the Maxrumpageviews field if non-nil, zero value otherwise.

### GetMaxrumpageviewsOk

`func (o *CreditsRespAttrsCredits) GetMaxrumpageviewsOk() (*int32, bool)`

GetMaxrumpageviewsOk returns a tuple with the Maxrumpageviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxrumpageviews

`func (o *CreditsRespAttrsCredits) SetMaxrumpageviews(v int32)`

SetMaxrumpageviews sets Maxrumpageviews field to given value.

### HasMaxrumpageviews

`func (o *CreditsRespAttrsCredits) HasMaxrumpageviews() bool`

HasMaxrumpageviews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


