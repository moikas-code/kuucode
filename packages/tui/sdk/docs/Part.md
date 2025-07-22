# Part

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionID** | **string** |  | 
**MessageID** | **string** |  | 
**Type** | **string** |  | 
**Text** | **string** |  | 
**Synthetic** | Pointer to **bool** |  | [optional] 
**Time** | Pointer to [**TextPartTime**](TextPartTime.md) |  | [optional] 
**Mime** | **string** |  | 
**Filename** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**Source** | Pointer to [**FilePartSource**](FilePartSource.md) |  | [optional] 
**CallID** | **string** |  | 
**Tool** | **string** |  | 
**State** | [**ToolState**](ToolState.md) |  | 
**Cost** | **float32** |  | 
**Tokens** | [**AssistantMessageTokens**](AssistantMessageTokens.md) |  | 
**Snapshot** | **string** |  | 

## Methods

### NewPart

`func NewPart(id string, sessionID string, messageID string, type_ string, text string, mime string, url string, callID string, tool string, state ToolState, cost float32, tokens AssistantMessageTokens, snapshot string, ) *Part`

NewPart instantiates a new Part object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartWithDefaults

`func NewPartWithDefaults() *Part`

NewPartWithDefaults instantiates a new Part object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Part) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Part) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Part) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *Part) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *Part) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *Part) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetMessageID

`func (o *Part) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *Part) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *Part) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.


### GetType

`func (o *Part) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Part) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Part) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *Part) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Part) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Part) SetText(v string)`

SetText sets Text field to given value.


### GetSynthetic

`func (o *Part) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *Part) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *Part) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *Part) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTime

`func (o *Part) GetTime() TextPartTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Part) GetTimeOk() (*TextPartTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Part) SetTime(v TextPartTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *Part) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetMime

`func (o *Part) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *Part) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *Part) SetMime(v string)`

SetMime sets Mime field to given value.


### GetFilename

`func (o *Part) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Part) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Part) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Part) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetUrl

`func (o *Part) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Part) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Part) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *Part) GetSource() FilePartSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Part) GetSourceOk() (*FilePartSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Part) SetSource(v FilePartSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Part) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCallID

`func (o *Part) GetCallID() string`

GetCallID returns the CallID field if non-nil, zero value otherwise.

### GetCallIDOk

`func (o *Part) GetCallIDOk() (*string, bool)`

GetCallIDOk returns a tuple with the CallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallID

`func (o *Part) SetCallID(v string)`

SetCallID sets CallID field to given value.


### GetTool

`func (o *Part) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *Part) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *Part) SetTool(v string)`

SetTool sets Tool field to given value.


### GetState

`func (o *Part) GetState() ToolState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Part) GetStateOk() (*ToolState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Part) SetState(v ToolState)`

SetState sets State field to given value.


### GetCost

`func (o *Part) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Part) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Part) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetTokens

`func (o *Part) GetTokens() AssistantMessageTokens`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *Part) GetTokensOk() (*AssistantMessageTokens, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *Part) SetTokens(v AssistantMessageTokens)`

SetTokens sets Tokens field to given value.


### GetSnapshot

`func (o *Part) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *Part) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *Part) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


