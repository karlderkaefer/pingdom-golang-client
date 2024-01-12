# ResultsRespAttrs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ResultsRespAttrsResultsInner**](ResultsRespAttrsResultsInner.md) |  | [optional] 
**Activeprobes** | Pointer to **[]float32** | For your convenience, a list of used probes that produced the showed results | [optional] 

## Methods

### NewResultsRespAttrs

`func NewResultsRespAttrs() *ResultsRespAttrs`

NewResultsRespAttrs instantiates a new ResultsRespAttrs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsRespAttrsWithDefaults

`func NewResultsRespAttrsWithDefaults() *ResultsRespAttrs`

NewResultsRespAttrsWithDefaults instantiates a new ResultsRespAttrs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ResultsRespAttrs) GetResults() []ResultsRespAttrsResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResultsRespAttrs) GetResultsOk() (*[]ResultsRespAttrsResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResultsRespAttrs) SetResults(v []ResultsRespAttrsResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *ResultsRespAttrs) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetActiveprobes

`func (o *ResultsRespAttrs) GetActiveprobes() []float32`

GetActiveprobes returns the Activeprobes field if non-nil, zero value otherwise.

### GetActiveprobesOk

`func (o *ResultsRespAttrs) GetActiveprobesOk() (*[]float32, bool)`

GetActiveprobesOk returns a tuple with the Activeprobes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveprobes

`func (o *ResultsRespAttrs) SetActiveprobes(v []float32)`

SetActiveprobes sets Activeprobes field to given value.

### HasActiveprobes

`func (o *ResultsRespAttrs) HasActiveprobes() bool`

HasActiveprobes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


