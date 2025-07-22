# Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | JSON schema reference for configuration validation | [optional] 
**Theme** | Pointer to **string** | Theme name to use for the interface | [optional] 
**Keybinds** | Pointer to [**KeybindsConfig**](KeybindsConfig.md) |  | [optional] 
**Share** | Pointer to **string** | Control sharing behavior:&#39;manual&#39; allows manual sharing via commands, &#39;auto&#39; enables automatic sharing, &#39;disabled&#39; disables all sharing | [optional] 
**Autoshare** | Pointer to **bool** | @deprecated Use &#39;share&#39; field instead. Share newly created sessions automatically | [optional] 
**Autoupdate** | Pointer to **bool** | Automatically update to the latest version | [optional] 
**DisabledProviders** | Pointer to **[]string** | Disable providers that are loaded automatically | [optional] 
**Model** | Pointer to **string** | Model to use in the format of provider/model, eg anthropic/claude-2 | [optional] 
**SmallModel** | Pointer to **string** | Small model to use for tasks like summarization and title generation in the format of provider/model | [optional] 
**Username** | Pointer to **string** | Custom username to display in conversations instead of system username | [optional] 
**Mode** | Pointer to [**ConfigMode**](ConfigMode.md) |  | [optional] 
**Provider** | Pointer to [**map[string]ConfigProviderValue**](ConfigProviderValue.md) | Custom provider configurations and model overrides | [optional] 
**Mcp** | Pointer to [**map[string]ConfigMcpValue**](ConfigMcpValue.md) | MCP (Model Context Protocol) server configurations | [optional] 
**Instructions** | Pointer to **[]string** | Additional instruction files or patterns to include | [optional] 
**Layout** | Pointer to [**LayoutConfig**](LayoutConfig.md) |  | [optional] 
**Experimental** | Pointer to [**ConfigExperimental**](ConfigExperimental.md) |  | [optional] 

## Methods

### NewConfig

`func NewConfig() *Config`

NewConfig instantiates a new Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigWithDefaults

`func NewConfigWithDefaults() *Config`

NewConfigWithDefaults instantiates a new Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *Config) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Config) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Config) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Config) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTheme

`func (o *Config) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Config) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Config) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Config) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetKeybinds

`func (o *Config) GetKeybinds() KeybindsConfig`

GetKeybinds returns the Keybinds field if non-nil, zero value otherwise.

### GetKeybindsOk

`func (o *Config) GetKeybindsOk() (*KeybindsConfig, bool)`

GetKeybindsOk returns a tuple with the Keybinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeybinds

`func (o *Config) SetKeybinds(v KeybindsConfig)`

SetKeybinds sets Keybinds field to given value.

### HasKeybinds

`func (o *Config) HasKeybinds() bool`

HasKeybinds returns a boolean if a field has been set.

### GetShare

`func (o *Config) GetShare() string`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *Config) GetShareOk() (*string, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *Config) SetShare(v string)`

SetShare sets Share field to given value.

### HasShare

`func (o *Config) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetAutoshare

`func (o *Config) GetAutoshare() bool`

GetAutoshare returns the Autoshare field if non-nil, zero value otherwise.

### GetAutoshareOk

`func (o *Config) GetAutoshareOk() (*bool, bool)`

GetAutoshareOk returns a tuple with the Autoshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoshare

`func (o *Config) SetAutoshare(v bool)`

SetAutoshare sets Autoshare field to given value.

### HasAutoshare

`func (o *Config) HasAutoshare() bool`

HasAutoshare returns a boolean if a field has been set.

### GetAutoupdate

`func (o *Config) GetAutoupdate() bool`

GetAutoupdate returns the Autoupdate field if non-nil, zero value otherwise.

### GetAutoupdateOk

`func (o *Config) GetAutoupdateOk() (*bool, bool)`

GetAutoupdateOk returns a tuple with the Autoupdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoupdate

`func (o *Config) SetAutoupdate(v bool)`

SetAutoupdate sets Autoupdate field to given value.

### HasAutoupdate

`func (o *Config) HasAutoupdate() bool`

HasAutoupdate returns a boolean if a field has been set.

### GetDisabledProviders

`func (o *Config) GetDisabledProviders() []string`

GetDisabledProviders returns the DisabledProviders field if non-nil, zero value otherwise.

### GetDisabledProvidersOk

`func (o *Config) GetDisabledProvidersOk() (*[]string, bool)`

GetDisabledProvidersOk returns a tuple with the DisabledProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledProviders

`func (o *Config) SetDisabledProviders(v []string)`

SetDisabledProviders sets DisabledProviders field to given value.

### HasDisabledProviders

`func (o *Config) HasDisabledProviders() bool`

HasDisabledProviders returns a boolean if a field has been set.

### GetModel

`func (o *Config) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Config) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Config) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Config) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSmallModel

`func (o *Config) GetSmallModel() string`

GetSmallModel returns the SmallModel field if non-nil, zero value otherwise.

### GetSmallModelOk

`func (o *Config) GetSmallModelOk() (*string, bool)`

GetSmallModelOk returns a tuple with the SmallModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallModel

`func (o *Config) SetSmallModel(v string)`

SetSmallModel sets SmallModel field to given value.

### HasSmallModel

`func (o *Config) HasSmallModel() bool`

HasSmallModel returns a boolean if a field has been set.

### GetUsername

`func (o *Config) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Config) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Config) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Config) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetMode

`func (o *Config) GetMode() ConfigMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Config) GetModeOk() (*ConfigMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Config) SetMode(v ConfigMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Config) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProvider

`func (o *Config) GetProvider() map[string]ConfigProviderValue`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Config) GetProviderOk() (*map[string]ConfigProviderValue, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Config) SetProvider(v map[string]ConfigProviderValue)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Config) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetMcp

`func (o *Config) GetMcp() map[string]ConfigMcpValue`

GetMcp returns the Mcp field if non-nil, zero value otherwise.

### GetMcpOk

`func (o *Config) GetMcpOk() (*map[string]ConfigMcpValue, bool)`

GetMcpOk returns a tuple with the Mcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcp

`func (o *Config) SetMcp(v map[string]ConfigMcpValue)`

SetMcp sets Mcp field to given value.

### HasMcp

`func (o *Config) HasMcp() bool`

HasMcp returns a boolean if a field has been set.

### GetInstructions

`func (o *Config) GetInstructions() []string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Config) GetInstructionsOk() (*[]string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Config) SetInstructions(v []string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *Config) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetLayout

`func (o *Config) GetLayout() LayoutConfig`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Config) GetLayoutOk() (*LayoutConfig, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Config) SetLayout(v LayoutConfig)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Config) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetExperimental

`func (o *Config) GetExperimental() ConfigExperimental`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *Config) GetExperimentalOk() (*ConfigExperimental, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *Config) SetExperimental(v ConfigExperimental)`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *Config) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


