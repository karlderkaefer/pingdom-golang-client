# ReportPerformanceReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckId** | Pointer to **int64** | ID of the check | [optional] 
**Name** | Pointer to **string** | Name of the check | [optional] 
**Resolution** | Pointer to **string** | Size of a time bucket for which the results are calculated | [optional] 
**Intervals** | Pointer to [**[]ReportPerformanceReportIntervalsInner**](ReportPerformanceReportIntervalsInner.md) | Intervals for which the measurements were performed. | [optional] 

## Methods

### NewReportPerformanceReport

`func NewReportPerformanceReport() *ReportPerformanceReport`

NewReportPerformanceReport instantiates a new ReportPerformanceReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportPerformanceReportWithDefaults

`func NewReportPerformanceReportWithDefaults() *ReportPerformanceReport`

NewReportPerformanceReportWithDefaults instantiates a new ReportPerformanceReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckId

`func (o *ReportPerformanceReport) GetCheckId() int64`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *ReportPerformanceReport) GetCheckIdOk() (*int64, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *ReportPerformanceReport) SetCheckId(v int64)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *ReportPerformanceReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetName

`func (o *ReportPerformanceReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportPerformanceReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportPerformanceReport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportPerformanceReport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolution

`func (o *ReportPerformanceReport) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ReportPerformanceReport) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ReportPerformanceReport) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ReportPerformanceReport) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetIntervals

`func (o *ReportPerformanceReport) GetIntervals() []ReportPerformanceReportIntervalsInner`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *ReportPerformanceReport) GetIntervalsOk() (*[]ReportPerformanceReportIntervalsInner, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *ReportPerformanceReport) SetIntervals(v []ReportPerformanceReportIntervalsInner)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *ReportPerformanceReport) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


