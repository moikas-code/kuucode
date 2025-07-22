# ConfigProviderValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Npm** | Pointer to **string** |  | [optional] 
**Models** | [**map[string]ConfigProviderValueModelsValue**](ConfigProviderValueModelsValue.md) |  | 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConfigProviderValue

`func NewConfigProviderValue(models map[string]ConfigProviderValueModelsValue, ) *ConfigProviderValue`

NewConfigProviderValue instantiates a new ConfigProviderValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigProviderValueWithDefaults

`func NewConfigProviderValueWithDefaults() *ConfigProviderValue`

NewConfigProviderValueWithDefaults instantiates a new ConfigProviderValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *ConfigProviderValue) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ConfigProviderValue) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ConfigProviderValue) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *ConfigProviderValue) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetName

`func (o *ConfigProviderValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigProviderValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigProviderValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigProviderValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnv

`func (o *ConfigProviderValue) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ConfigProviderValue) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ConfigProviderValue) SetEnv(v []string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ConfigProviderValue) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetId

`func (o *ConfigProviderValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigProviderValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigProviderValue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigProviderValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNpm

`func (o *ConfigProviderValue) GetNpm() string`

GetNpm returns the Npm field if non-nil, zero value otherwise.

### GetNpmOk

`func (o *ConfigProviderValue) GetNpmOk() (*string, bool)`

GetNpmOk returns a tuple with the Npm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpm

`func (o *ConfigProviderValue) SetNpm(v string)`

SetNpm sets Npm field to given value.

### HasNpm

`func (o *ConfigProviderValue) HasNpm() bool`

HasNpm returns a boolean if a field has been set.

### GetModels

`func (o *ConfigProviderValue) GetModels() map[string]ConfigProviderValueModelsValue`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ConfigProviderValue) GetModelsOk() (*map[string]ConfigProviderValueModelsValue, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ConfigProviderValue) SetModels(v map[string]ConfigProviderValueModelsValue)`

SetModels sets Models field to given value.


### GetOptions

`func (o *ConfigProviderValue) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigProviderValue) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigProviderValue) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ConfigProviderValue) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


