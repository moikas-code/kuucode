# ToolState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Input** | **map[string]map[string]interface{}** |  | 
**Title** | **string** |  | 
**Metadata** | **map[string]map[string]interface{}** |  | 
**Time** | [**ToolStateCompletedTime**](ToolStateCompletedTime.md) |  | 
**Output** | **string** |  | 
**Error** | **string** |  | 

## Methods

### NewToolState

`func NewToolState(status string, input map[string]map[string]interface{}, title string, metadata map[string]map[string]interface{}, time ToolStateCompletedTime, output string, error_ string, ) *ToolState`

NewToolState instantiates a new ToolState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolStateWithDefaults

`func NewToolStateWithDefaults() *ToolState`

NewToolStateWithDefaults instantiates a new ToolState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ToolState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ToolState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ToolState) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInput

`func (o *ToolState) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ToolState) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ToolState) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.


### GetTitle

`func (o *ToolState) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ToolState) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ToolState) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMetadata

`func (o *ToolState) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ToolState) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ToolState) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetTime

`func (o *ToolState) GetTime() ToolStateCompletedTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ToolState) GetTimeOk() (*ToolStateCompletedTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ToolState) SetTime(v ToolStateCompletedTime)`

SetTime sets Time field to given value.


### GetOutput

`func (o *ToolState) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ToolState) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ToolState) SetOutput(v string)`

SetOutput sets Output field to given value.


### GetError

`func (o *ToolState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ToolState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ToolState) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


