# StepFinishPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionID** | **string** |  | 
**MessageID** | **string** |  | 
**Type** | **string** |  | 
**Cost** | **float32** |  | 
**Tokens** | [**AssistantMessageTokens**](AssistantMessageTokens.md) |  | 

## Methods

### NewStepFinishPart

`func NewStepFinishPart(id string, sessionID string, messageID string, type_ string, cost float32, tokens AssistantMessageTokens, ) *StepFinishPart`

NewStepFinishPart instantiates a new StepFinishPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepFinishPartWithDefaults

`func NewStepFinishPartWithDefaults() *StepFinishPart`

NewStepFinishPartWithDefaults instantiates a new StepFinishPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StepFinishPart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StepFinishPart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StepFinishPart) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *StepFinishPart) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *StepFinishPart) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *StepFinishPart) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetMessageID

`func (o *StepFinishPart) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *StepFinishPart) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *StepFinishPart) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.


### GetType

`func (o *StepFinishPart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StepFinishPart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StepFinishPart) SetType(v string)`

SetType sets Type field to given value.


### GetCost

`func (o *StepFinishPart) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *StepFinishPart) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *StepFinishPart) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetTokens

`func (o *StepFinishPart) GetTokens() AssistantMessageTokens`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *StepFinishPart) GetTokensOk() (*AssistantMessageTokens, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *StepFinishPart) SetTokens(v AssistantMessageTokens)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


