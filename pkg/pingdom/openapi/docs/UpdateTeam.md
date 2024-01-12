# UpdateTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Team name | 
**MemberIds** | **[]int64** | IDs of contacts to be the members of this team | 

## Methods

### NewUpdateTeam

`func NewUpdateTeam(name string, memberIds []int64, ) *UpdateTeam`

NewUpdateTeam instantiates a new UpdateTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamWithDefaults

`func NewUpdateTeamWithDefaults() *UpdateTeam`

NewUpdateTeamWithDefaults instantiates a new UpdateTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTeam) SetName(v string)`

SetName sets Name field to given value.


### GetMemberIds

`func (o *UpdateTeam) GetMemberIds() []int64`

GetMemberIds returns the MemberIds field if non-nil, zero value otherwise.

### GetMemberIdsOk

`func (o *UpdateTeam) GetMemberIdsOk() (*[]int64, bool)`

GetMemberIdsOk returns a tuple with the MemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIds

`func (o *UpdateTeam) SetMemberIds(v []int64)`

SetMemberIds sets MemberIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


