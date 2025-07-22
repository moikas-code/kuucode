# ToolStateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Input** | **map[string]interface{}** |  | 
**Error** | **string** |  | 
**Time** | [**ToolStateCompletedTime**](ToolStateCompletedTime.md) |  | 

## Methods

### NewToolStateError

`func NewToolStateError(status string, input map[string]interface{}, error_ string, time ToolStateCompletedTime, ) *ToolStateError`

NewToolStateError instantiates a new ToolStateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolStateErrorWithDefaults

`func NewToolStateErrorWithDefaults() *ToolStateError`

NewToolStateErrorWithDefaults instantiates a new ToolStateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ToolStateError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ToolStateError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ToolStateError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInput

`func (o *ToolStateError) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ToolStateError) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ToolStateError) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetError

`func (o *ToolStateError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ToolStateError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ToolStateError) SetError(v string)`

SetError sets Error field to given value.


### GetTime

`func (o *ToolStateError) GetTime() ToolStateCompletedTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ToolStateError) GetTimeOk() (*ToolStateCompletedTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ToolStateError) SetTime(v ToolStateCompletedTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


