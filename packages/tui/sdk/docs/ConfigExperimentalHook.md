# ConfigExperimentalHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileEdited** | Pointer to [**map[string][]ConfigExperimentalHookFileEditedValueInner**](array.md) |  | [optional] 
**SessionCompleted** | Pointer to [**[]ConfigExperimentalHookFileEditedValueInner**](ConfigExperimentalHookFileEditedValueInner.md) |  | [optional] 

## Methods

### NewConfigExperimentalHook

`func NewConfigExperimentalHook() *ConfigExperimentalHook`

NewConfigExperimentalHook instantiates a new ConfigExperimentalHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigExperimentalHookWithDefaults

`func NewConfigExperimentalHookWithDefaults() *ConfigExperimentalHook`

NewConfigExperimentalHookWithDefaults instantiates a new ConfigExperimentalHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileEdited

`func (o *ConfigExperimentalHook) GetFileEdited() map[string][]ConfigExperimentalHookFileEditedValueInner`

GetFileEdited returns the FileEdited field if non-nil, zero value otherwise.

### GetFileEditedOk

`func (o *ConfigExperimentalHook) GetFileEditedOk() (*map[string][]ConfigExperimentalHookFileEditedValueInner, bool)`

GetFileEditedOk returns a tuple with the FileEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileEdited

`func (o *ConfigExperimentalHook) SetFileEdited(v map[string][]ConfigExperimentalHookFileEditedValueInner)`

SetFileEdited sets FileEdited field to given value.

### HasFileEdited

`func (o *ConfigExperimentalHook) HasFileEdited() bool`

HasFileEdited returns a boolean if a field has been set.

### GetSessionCompleted

`func (o *ConfigExperimentalHook) GetSessionCompleted() []ConfigExperimentalHookFileEditedValueInner`

GetSessionCompleted returns the SessionCompleted field if non-nil, zero value otherwise.

### GetSessionCompletedOk

`func (o *ConfigExperimentalHook) GetSessionCompletedOk() (*[]ConfigExperimentalHookFileEditedValueInner, bool)`

GetSessionCompletedOk returns a tuple with the SessionCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCompleted

`func (o *ConfigExperimentalHook) SetSessionCompleted(v []ConfigExperimentalHookFileEditedValueInner)`

SetSessionCompleted sets SessionCompleted field to given value.

### HasSessionCompleted

`func (o *ConfigExperimentalHook) HasSessionCompleted() bool`

HasSessionCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


