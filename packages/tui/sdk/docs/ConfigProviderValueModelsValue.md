# ConfigProviderValueModelsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Attachment** | Pointer to **bool** |  | [optional] 
**Reasoning** | Pointer to **bool** |  | [optional] 
**Temperature** | Pointer to **bool** |  | [optional] 
**ToolCall** | Pointer to **bool** |  | [optional] 
**Cost** | Pointer to [**ConfigProviderValueModelsValueCost**](ConfigProviderValueModelsValueCost.md) |  | [optional] 
**Limit** | Pointer to [**ConfigProviderValueModelsValueLimit**](ConfigProviderValueModelsValueLimit.md) |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConfigProviderValueModelsValue

`func NewConfigProviderValueModelsValue() *ConfigProviderValueModelsValue`

NewConfigProviderValueModelsValue instantiates a new ConfigProviderValueModelsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigProviderValueModelsValueWithDefaults

`func NewConfigProviderValueModelsValueWithDefaults() *ConfigProviderValueModelsValue`

NewConfigProviderValueModelsValueWithDefaults instantiates a new ConfigProviderValueModelsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigProviderValueModelsValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigProviderValueModelsValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigProviderValueModelsValue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigProviderValueModelsValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConfigProviderValueModelsValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigProviderValueModelsValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigProviderValueModelsValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigProviderValueModelsValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ConfigProviderValueModelsValue) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ConfigProviderValueModelsValue) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ConfigProviderValueModelsValue) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ConfigProviderValueModelsValue) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetAttachment

`func (o *ConfigProviderValueModelsValue) GetAttachment() bool`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *ConfigProviderValueModelsValue) GetAttachmentOk() (*bool, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *ConfigProviderValueModelsValue) SetAttachment(v bool)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *ConfigProviderValueModelsValue) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetReasoning

`func (o *ConfigProviderValueModelsValue) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ConfigProviderValueModelsValue) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ConfigProviderValueModelsValue) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *ConfigProviderValueModelsValue) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetTemperature

`func (o *ConfigProviderValueModelsValue) GetTemperature() bool`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ConfigProviderValueModelsValue) GetTemperatureOk() (*bool, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ConfigProviderValueModelsValue) SetTemperature(v bool)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ConfigProviderValueModelsValue) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetToolCall

`func (o *ConfigProviderValueModelsValue) GetToolCall() bool`

GetToolCall returns the ToolCall field if non-nil, zero value otherwise.

### GetToolCallOk

`func (o *ConfigProviderValueModelsValue) GetToolCallOk() (*bool, bool)`

GetToolCallOk returns a tuple with the ToolCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCall

`func (o *ConfigProviderValueModelsValue) SetToolCall(v bool)`

SetToolCall sets ToolCall field to given value.

### HasToolCall

`func (o *ConfigProviderValueModelsValue) HasToolCall() bool`

HasToolCall returns a boolean if a field has been set.

### GetCost

`func (o *ConfigProviderValueModelsValue) GetCost() ConfigProviderValueModelsValueCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ConfigProviderValueModelsValue) GetCostOk() (*ConfigProviderValueModelsValueCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ConfigProviderValueModelsValue) SetCost(v ConfigProviderValueModelsValueCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ConfigProviderValueModelsValue) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetLimit

`func (o *ConfigProviderValueModelsValue) GetLimit() ConfigProviderValueModelsValueLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ConfigProviderValueModelsValue) GetLimitOk() (*ConfigProviderValueModelsValueLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ConfigProviderValueModelsValue) SetLimit(v ConfigProviderValueModelsValueLimit)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ConfigProviderValueModelsValue) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOptions

`func (o *ConfigProviderValueModelsValue) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigProviderValueModelsValue) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigProviderValueModelsValue) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ConfigProviderValueModelsValue) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


