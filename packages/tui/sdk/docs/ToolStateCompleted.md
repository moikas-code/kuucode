# ToolStateCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Input** | **map[string]interface{}** |  | 
**Output** | **string** |  | 
**Title** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Time** | [**ToolStateCompletedTime**](ToolStateCompletedTime.md) |  | 

## Methods

### NewToolStateCompleted

`func NewToolStateCompleted(status string, input map[string]interface{}, output string, title string, metadata map[string]interface{}, time ToolStateCompletedTime, ) *ToolStateCompleted`

NewToolStateCompleted instantiates a new ToolStateCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolStateCompletedWithDefaults

`func NewToolStateCompletedWithDefaults() *ToolStateCompleted`

NewToolStateCompletedWithDefaults instantiates a new ToolStateCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ToolStateCompleted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ToolStateCompleted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ToolStateCompleted) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInput

`func (o *ToolStateCompleted) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ToolStateCompleted) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ToolStateCompleted) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetOutput

`func (o *ToolStateCompleted) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ToolStateCompleted) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ToolStateCompleted) SetOutput(v string)`

SetOutput sets Output field to given value.


### GetTitle

`func (o *ToolStateCompleted) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ToolStateCompleted) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ToolStateCompleted) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMetadata

`func (o *ToolStateCompleted) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ToolStateCompleted) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ToolStateCompleted) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetTime

`func (o *ToolStateCompleted) GetTime() ToolStateCompletedTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ToolStateCompleted) GetTimeOk() (*ToolStateCompletedTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ToolStateCompleted) SetTime(v ToolStateCompletedTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


