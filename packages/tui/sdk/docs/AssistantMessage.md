# AssistantMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionID** | **string** |  | 
**Role** | **string** |  | 
**Time** | [**AssistantMessageTime**](AssistantMessageTime.md) |  | 
**Error** | Pointer to [**AssistantMessageError**](AssistantMessageError.md) |  | [optional] 
**System** | **[]string** |  | 
**ModelID** | **string** |  | 
**ProviderID** | **string** |  | 
**Path** | [**AssistantMessagePath**](AssistantMessagePath.md) |  | 
**Summary** | Pointer to **bool** |  | [optional] 
**Cost** | **float32** |  | 
**Tokens** | [**AssistantMessageTokens**](AssistantMessageTokens.md) |  | 

## Methods

### NewAssistantMessage

`func NewAssistantMessage(id string, sessionID string, role string, time AssistantMessageTime, system []string, modelID string, providerID string, path AssistantMessagePath, cost float32, tokens AssistantMessageTokens, ) *AssistantMessage`

NewAssistantMessage instantiates a new AssistantMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantMessageWithDefaults

`func NewAssistantMessageWithDefaults() *AssistantMessage`

NewAssistantMessageWithDefaults instantiates a new AssistantMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssistantMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssistantMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssistantMessage) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *AssistantMessage) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *AssistantMessage) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *AssistantMessage) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetRole

`func (o *AssistantMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssistantMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssistantMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetTime

`func (o *AssistantMessage) GetTime() AssistantMessageTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AssistantMessage) GetTimeOk() (*AssistantMessageTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AssistantMessage) SetTime(v AssistantMessageTime)`

SetTime sets Time field to given value.


### GetError

`func (o *AssistantMessage) GetError() AssistantMessageError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AssistantMessage) GetErrorOk() (*AssistantMessageError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AssistantMessage) SetError(v AssistantMessageError)`

SetError sets Error field to given value.

### HasError

`func (o *AssistantMessage) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSystem

`func (o *AssistantMessage) GetSystem() []string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AssistantMessage) GetSystemOk() (*[]string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AssistantMessage) SetSystem(v []string)`

SetSystem sets System field to given value.


### GetModelID

`func (o *AssistantMessage) GetModelID() string`

GetModelID returns the ModelID field if non-nil, zero value otherwise.

### GetModelIDOk

`func (o *AssistantMessage) GetModelIDOk() (*string, bool)`

GetModelIDOk returns a tuple with the ModelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelID

`func (o *AssistantMessage) SetModelID(v string)`

SetModelID sets ModelID field to given value.


### GetProviderID

`func (o *AssistantMessage) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *AssistantMessage) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *AssistantMessage) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.


### GetPath

`func (o *AssistantMessage) GetPath() AssistantMessagePath`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AssistantMessage) GetPathOk() (*AssistantMessagePath, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AssistantMessage) SetPath(v AssistantMessagePath)`

SetPath sets Path field to given value.


### GetSummary

`func (o *AssistantMessage) GetSummary() bool`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AssistantMessage) GetSummaryOk() (*bool, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AssistantMessage) SetSummary(v bool)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AssistantMessage) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCost

`func (o *AssistantMessage) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AssistantMessage) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AssistantMessage) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetTokens

`func (o *AssistantMessage) GetTokens() AssistantMessageTokens`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AssistantMessage) GetTokensOk() (*AssistantMessageTokens, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AssistantMessage) SetTokens(v AssistantMessageTokens)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


