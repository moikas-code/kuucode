# Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ReleaseDate** | **string** |  | 
**Attachment** | **bool** |  | 
**Reasoning** | **bool** |  | 
**Temperature** | **bool** |  | 
**ToolCall** | **bool** |  | 
**Cost** | [**ConfigProviderValueModelsValueCost**](ConfigProviderValueModelsValueCost.md) |  | 
**Limit** | [**ConfigProviderValueModelsValueLimit**](ConfigProviderValueModelsValueLimit.md) |  | 
**Options** | **map[string]interface{}** |  | 

## Methods

### NewModel

`func NewModel(id string, name string, releaseDate string, attachment bool, reasoning bool, temperature bool, toolCall bool, cost ConfigProviderValueModelsValueCost, limit ConfigProviderValueModelsValueLimit, options map[string]interface{}, ) *Model`

NewModel instantiates a new Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelWithDefaults

`func NewModelWithDefaults() *Model`

NewModelWithDefaults instantiates a new Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Model) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Model) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Model) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Model) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Model) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Model) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *Model) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *Model) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *Model) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetAttachment

`func (o *Model) GetAttachment() bool`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *Model) GetAttachmentOk() (*bool, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *Model) SetAttachment(v bool)`

SetAttachment sets Attachment field to given value.


### GetReasoning

`func (o *Model) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *Model) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *Model) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.


### GetTemperature

`func (o *Model) GetTemperature() bool`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *Model) GetTemperatureOk() (*bool, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *Model) SetTemperature(v bool)`

SetTemperature sets Temperature field to given value.


### GetToolCall

`func (o *Model) GetToolCall() bool`

GetToolCall returns the ToolCall field if non-nil, zero value otherwise.

### GetToolCallOk

`func (o *Model) GetToolCallOk() (*bool, bool)`

GetToolCallOk returns a tuple with the ToolCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCall

`func (o *Model) SetToolCall(v bool)`

SetToolCall sets ToolCall field to given value.


### GetCost

`func (o *Model) GetCost() ConfigProviderValueModelsValueCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Model) GetCostOk() (*ConfigProviderValueModelsValueCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Model) SetCost(v ConfigProviderValueModelsValueCost)`

SetCost sets Cost field to given value.


### GetLimit

`func (o *Model) GetLimit() ConfigProviderValueModelsValueLimit`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Model) GetLimitOk() (*ConfigProviderValueModelsValueLimit, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Model) SetLimit(v ConfigProviderValueModelsValueLimit)`

SetLimit sets Limit field to given value.


### GetOptions

`func (o *Model) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Model) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Model) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


