# ToolPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionID** | **string** |  | 
**MessageID** | **string** |  | 
**Type** | **string** |  | 
**CallID** | **string** |  | 
**Tool** | **string** |  | 
**State** | [**ToolState**](ToolState.md) |  | 

## Methods

### NewToolPart

`func NewToolPart(id string, sessionID string, messageID string, type_ string, callID string, tool string, state ToolState, ) *ToolPart`

NewToolPart instantiates a new ToolPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolPartWithDefaults

`func NewToolPartWithDefaults() *ToolPart`

NewToolPartWithDefaults instantiates a new ToolPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolPart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolPart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolPart) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *ToolPart) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *ToolPart) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *ToolPart) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetMessageID

`func (o *ToolPart) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *ToolPart) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *ToolPart) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.


### GetType

`func (o *ToolPart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolPart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolPart) SetType(v string)`

SetType sets Type field to given value.


### GetCallID

`func (o *ToolPart) GetCallID() string`

GetCallID returns the CallID field if non-nil, zero value otherwise.

### GetCallIDOk

`func (o *ToolPart) GetCallIDOk() (*string, bool)`

GetCallIDOk returns a tuple with the CallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallID

`func (o *ToolPart) SetCallID(v string)`

SetCallID sets CallID field to given value.


### GetTool

`func (o *ToolPart) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *ToolPart) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *ToolPart) SetTool(v string)`

SetTool sets Tool field to given value.


### GetState

`func (o *ToolPart) GetState() ToolState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ToolPart) GetStateOk() (*ToolState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ToolPart) SetState(v ToolState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


