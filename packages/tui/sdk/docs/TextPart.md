# TextPart

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

## Methods

### NewTextPart

`func NewTextPart(id string, sessionID string, messageID string, type_ string, text string, ) *TextPart`

NewTextPart instantiates a new TextPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextPartWithDefaults

`func NewTextPartWithDefaults() *TextPart`

NewTextPartWithDefaults instantiates a new TextPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TextPart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextPart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextPart) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *TextPart) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *TextPart) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *TextPart) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetMessageID

`func (o *TextPart) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *TextPart) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *TextPart) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.


### GetType

`func (o *TextPart) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextPart) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextPart) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *TextPart) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextPart) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextPart) SetText(v string)`

SetText sets Text field to given value.


### GetSynthetic

`func (o *TextPart) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *TextPart) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *TextPart) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *TextPart) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTime

`func (o *TextPart) GetTime() TextPartTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TextPart) GetTimeOk() (*TextPartTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TextPart) SetTime(v TextPartTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *TextPart) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


