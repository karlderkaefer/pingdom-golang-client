# CreateTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Team name | 
**MemberIds** | **[]int32** | Contact ids | 

## Methods

### NewCreateTeam

`func NewCreateTeam(name string, memberIds []int32, ) *CreateTeam`

NewCreateTeam instantiates a new CreateTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeamWithDefaults

`func NewCreateTeamWithDefaults() *CreateTeam`

NewCreateTeamWithDefaults instantiates a new CreateTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTeam) SetName(v string)`

SetName sets Name field to given value.


### GetMemberIds

`func (o *CreateTeam) GetMemberIds() []int32`

GetMemberIds returns the MemberIds field if non-nil, zero value otherwise.

### GetMemberIdsOk

`func (o *CreateTeam) GetMemberIdsOk() (*[]int32, bool)`

GetMemberIdsOk returns a tuple with the MemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIds

`func (o *CreateTeam) SetMemberIds(v []int32)`

SetMemberIds sets MemberIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


