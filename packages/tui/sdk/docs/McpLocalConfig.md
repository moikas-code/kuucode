# McpLocalConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of MCP server connection | 
**Command** | **[]string** | Command and arguments to run the MCP server | 
**Environment** | Pointer to **map[string]string** | Environment variables to set when running the MCP server | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the MCP server on startup | [optional] 

## Methods

### NewMcpLocalConfig

`func NewMcpLocalConfig(type_ string, command []string, ) *McpLocalConfig`

NewMcpLocalConfig instantiates a new McpLocalConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpLocalConfigWithDefaults

`func NewMcpLocalConfigWithDefaults() *McpLocalConfig`

NewMcpLocalConfigWithDefaults instantiates a new McpLocalConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *McpLocalConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *McpLocalConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *McpLocalConfig) SetType(v string)`

SetType sets Type field to given value.


### GetCommand

`func (o *McpLocalConfig) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *McpLocalConfig) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *McpLocalConfig) SetCommand(v []string)`

SetCommand sets Command field to given value.


### GetEnvironment

`func (o *McpLocalConfig) GetEnvironment() map[string]string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *McpLocalConfig) GetEnvironmentOk() (*map[string]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *McpLocalConfig) SetEnvironment(v map[string]string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *McpLocalConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnabled

`func (o *McpLocalConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *McpLocalConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *McpLocalConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *McpLocalConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


