# ConfigMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to [**ModeConfig**](ModeConfig.md) |  | [optional] 
**Plan** | Pointer to [**ModeConfig**](ModeConfig.md) |  | [optional] 

## Methods

### NewConfigMode

`func NewConfigMode() *ConfigMode`

NewConfigMode instantiates a new ConfigMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigModeWithDefaults

`func NewConfigModeWithDefaults() *ConfigMode`

NewConfigModeWithDefaults instantiates a new ConfigMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *ConfigMode) GetBuild() ModeConfig`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ConfigMode) GetBuildOk() (*ModeConfig, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ConfigMode) SetBuild(v ModeConfig)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ConfigMode) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetPlan

`func (o *ConfigMode) GetPlan() ModeConfig`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ConfigMode) GetPlanOk() (*ModeConfig, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ConfigMode) SetPlan(v ModeConfig)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ConfigMode) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


