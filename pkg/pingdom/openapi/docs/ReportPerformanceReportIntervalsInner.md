# ReportPerformanceReportIntervalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageResponse** | Pointer to **int64** | Average response times in milliseconds | [optional] 
**From** | Pointer to **time.Time** | Interval start. The format is &#x60;RFC 3339&#x60; | [optional] 
**Downtime** | Pointer to **int64** | Amount of time when the check was down within given interval (only with the &#x60;include_uptime&#x60; flag) | [optional] 
**Uptime** | Pointer to **int64** | Amount of time when the check was up within given interval (only with the &#x60;include_uptime&#x60; flag) | [optional] 
**Unmonitored** | Pointer to **int64** | Amount of time when there is no specific data about check status (up/down) within given interval (only with the &#x60;include_uptime&#x60; flag) | [optional] 
**Steps** | Pointer to [**[]ReportPerformanceReportIntervalsInnerStepsInner**](ReportPerformanceReportIntervalsInnerStepsInner.md) |  | [optional] 

## Methods

### NewReportPerformanceReportIntervalsInner

`func NewReportPerformanceReportIntervalsInner() *ReportPerformanceReportIntervalsInner`

NewReportPerformanceReportIntervalsInner instantiates a new ReportPerformanceReportIntervalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportPerformanceReportIntervalsInnerWithDefaults

`func NewReportPerformanceReportIntervalsInnerWithDefaults() *ReportPerformanceReportIntervalsInner`

NewReportPerformanceReportIntervalsInnerWithDefaults instantiates a new ReportPerformanceReportIntervalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageResponse

`func (o *ReportPerformanceReportIntervalsInner) GetAverageResponse() int64`

GetAverageResponse returns the AverageResponse field if non-nil, zero value otherwise.

### GetAverageResponseOk

`func (o *ReportPerformanceReportIntervalsInner) GetAverageResponseOk() (*int64, bool)`

GetAverageResponseOk returns a tuple with the AverageResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageResponse

`func (o *ReportPerformanceReportIntervalsInner) SetAverageResponse(v int64)`

SetAverageResponse sets AverageResponse field to given value.

### HasAverageResponse

`func (o *ReportPerformanceReportIntervalsInner) HasAverageResponse() bool`

HasAverageResponse returns a boolean if a field has been set.

### GetFrom

`func (o *ReportPerformanceReportIntervalsInner) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ReportPerformanceReportIntervalsInner) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ReportPerformanceReportIntervalsInner) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ReportPerformanceReportIntervalsInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetDowntime

`func (o *ReportPerformanceReportIntervalsInner) GetDowntime() int64`

GetDowntime returns the Downtime field if non-nil, zero value otherwise.

### GetDowntimeOk

`func (o *ReportPerformanceReportIntervalsInner) GetDowntimeOk() (*int64, bool)`

GetDowntimeOk returns a tuple with the Downtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntime

`func (o *ReportPerformanceReportIntervalsInner) SetDowntime(v int64)`

SetDowntime sets Downtime field to given value.

### HasDowntime

`func (o *ReportPerformanceReportIntervalsInner) HasDowntime() bool`

HasDowntime returns a boolean if a field has been set.

### GetUptime

`func (o *ReportPerformanceReportIntervalsInner) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ReportPerformanceReportIntervalsInner) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ReportPerformanceReportIntervalsInner) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ReportPerformanceReportIntervalsInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUnmonitored

`func (o *ReportPerformanceReportIntervalsInner) GetUnmonitored() int64`

GetUnmonitored returns the Unmonitored field if non-nil, zero value otherwise.

### GetUnmonitoredOk

`func (o *ReportPerformanceReportIntervalsInner) GetUnmonitoredOk() (*int64, bool)`

GetUnmonitoredOk returns a tuple with the Unmonitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmonitored

`func (o *ReportPerformanceReportIntervalsInner) SetUnmonitored(v int64)`

SetUnmonitored sets Unmonitored field to given value.

### HasUnmonitored

`func (o *ReportPerformanceReportIntervalsInner) HasUnmonitored() bool`

HasUnmonitored returns a boolean if a field has been set.

### GetSteps

`func (o *ReportPerformanceReportIntervalsInner) GetSteps() []ReportPerformanceReportIntervalsInnerStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ReportPerformanceReportIntervalsInner) GetStepsOk() (*[]ReportPerformanceReportIntervalsInnerStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ReportPerformanceReportIntervalsInner) SetSteps(v []ReportPerformanceReportIntervalsInnerStepsInner)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ReportPerformanceReportIntervalsInner) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


