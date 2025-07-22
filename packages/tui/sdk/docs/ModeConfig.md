# ModeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewModeConfig

`func NewModeConfig() *ModeConfig`

NewModeConfig instantiates a new ModeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModeConfigWithDefaults

`func NewModeConfigWithDefaults() *ModeConfig`

NewModeConfigWithDefaults instantiates a new ModeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ModeConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModeConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModeConfig) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModeConfig) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrompt

`func (o *ModeConfig) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ModeConfig) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ModeConfig) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *ModeConfig) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTools

`func (o *ModeConfig) GetTools() map[string]bool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ModeConfig) GetToolsOk() (*map[string]bool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ModeConfig) SetTools(v map[string]bool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ModeConfig) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


