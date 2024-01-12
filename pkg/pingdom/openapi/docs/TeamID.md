# TeamID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | Pointer to [**AlertingTeamID**](AlertingTeamID.md) |  | [optional] 

## Methods

### NewTeamID

`func NewTeamID() *TeamID`

NewTeamID instantiates a new TeamID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamIDWithDefaults

`func NewTeamIDWithDefaults() *TeamID`

NewTeamIDWithDefaults instantiates a new TeamID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *TeamID) GetTeam() AlertingTeamID`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamID) GetTeamOk() (*AlertingTeamID, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamID) SetTeam(v AlertingTeamID)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *TeamID) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


