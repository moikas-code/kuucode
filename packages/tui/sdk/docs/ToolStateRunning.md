# ToolStateRunning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Input** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Time** | [**ToolStateRunningTime**](ToolStateRunningTime.md) |  | 

## Methods

### NewToolStateRunning

`func NewToolStateRunning(status string, time ToolStateRunningTime, ) *ToolStateRunning`

NewToolStateRunning instantiates a new ToolStateRunning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolStateRunningWithDefaults

`func NewToolStateRunningWithDefaults() *ToolStateRunning`

NewToolStateRunningWithDefaults instantiates a new ToolStateRunning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ToolStateRunning) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ToolStateRunning) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ToolStateRunning) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInput

`func (o *ToolStateRunning) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ToolStateRunning) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ToolStateRunning) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ToolStateRunning) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *ToolStateRunning) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ToolStateRunning) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetTitle

`func (o *ToolStateRunning) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ToolStateRunning) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ToolStateRunning) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ToolStateRunning) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMetadata

`func (o *ToolStateRunning) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ToolStateRunning) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ToolStateRunning) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ToolStateRunning) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTime

`func (o *ToolStateRunning) GetTime() ToolStateRunningTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ToolStateRunning) GetTimeOk() (*ToolStateRunningTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ToolStateRunning) SetTime(v ToolStateRunningTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


