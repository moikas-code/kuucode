# ConfigMcpValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of MCP server connection | 
**Command** | **[]string** | Command and arguments to run the MCP server | 
**Environment** | Pointer to **map[string]string** | Environment variables to set when running the MCP server | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the MCP server on startup | [optional] 
**Url** | **string** | URL of the remote MCP server | 
**Headers** | Pointer to **map[string]string** | Headers to send with the request | [optional] 

## Methods

### NewConfigMcpValue

`func NewConfigMcpValue(type_ string, command []string, url string, ) *ConfigMcpValue`

NewConfigMcpValue instantiates a new ConfigMcpValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMcpValueWithDefaults

`func NewConfigMcpValueWithDefaults() *ConfigMcpValue`

NewConfigMcpValueWithDefaults instantiates a new ConfigMcpValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConfigMcpValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigMcpValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigMcpValue) SetType(v string)`

SetType sets Type field to given value.


### GetCommand

`func (o *ConfigMcpValue) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ConfigMcpValue) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ConfigMcpValue) SetCommand(v []string)`

SetCommand sets Command field to given value.


### GetEnvironment

`func (o *ConfigMcpValue) GetEnvironment() map[string]string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConfigMcpValue) GetEnvironmentOk() (*map[string]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConfigMcpValue) SetEnvironment(v map[string]string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConfigMcpValue) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnabled

`func (o *ConfigMcpValue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConfigMcpValue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConfigMcpValue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConfigMcpValue) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUrl

`func (o *ConfigMcpValue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigMcpValue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigMcpValue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *ConfigMcpValue) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConfigMcpValue) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConfigMcpValue) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConfigMcpValue) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


