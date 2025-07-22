# Mode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Model** | Pointer to [**ModeModel**](ModeModel.md) |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Tools** | **map[string]bool** |  | 

## Methods

### NewMode

`func NewMode(name string, tools map[string]bool, ) *Mode`

NewMode instantiates a new Mode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModeWithDefaults

`func NewModeWithDefaults() *Mode`

NewModeWithDefaults instantiates a new Mode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Mode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mode) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *Mode) GetModel() ModeModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Mode) GetModelOk() (*ModeModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Mode) SetModel(v ModeModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *Mode) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrompt

`func (o *Mode) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *Mode) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *Mode) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *Mode) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTools

`func (o *Mode) GetTools() map[string]bool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Mode) GetToolsOk() (*map[string]bool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Mode) SetTools(v map[string]bool)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


