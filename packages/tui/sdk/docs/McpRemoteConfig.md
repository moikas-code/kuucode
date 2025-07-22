# McpRemoteConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of MCP server connection | 
**Url** | **string** | URL of the remote MCP server | 
**Enabled** | Pointer to **bool** | Enable or disable the MCP server on startup | [optional] 
**Headers** | Pointer to **map[string]string** | Headers to send with the request | [optional] 

## Methods

### NewMcpRemoteConfig

`func NewMcpRemoteConfig(type_ string, url string, ) *McpRemoteConfig`

NewMcpRemoteConfig instantiates a new McpRemoteConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpRemoteConfigWithDefaults

`func NewMcpRemoteConfigWithDefaults() *McpRemoteConfig`

NewMcpRemoteConfigWithDefaults instantiates a new McpRemoteConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *McpRemoteConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *McpRemoteConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *McpRemoteConfig) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *McpRemoteConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpRemoteConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpRemoteConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *McpRemoteConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *McpRemoteConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *McpRemoteConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *McpRemoteConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHeaders

`func (o *McpRemoteConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *McpRemoteConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *McpRemoteConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *McpRemoteConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


