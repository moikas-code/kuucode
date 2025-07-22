# TextPartInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Text** | **string** |  | 
**Synthetic** | Pointer to **bool** |  | [optional] 
**Time** | Pointer to [**TextPartTime**](TextPartTime.md) |  | [optional] 

## Methods

### NewTextPartInput

`func NewTextPartInput(type_ string, text string, ) *TextPartInput`

NewTextPartInput instantiates a new TextPartInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextPartInputWithDefaults

`func NewTextPartInputWithDefaults() *TextPartInput`

NewTextPartInputWithDefaults instantiates a new TextPartInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TextPartInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TextPartInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TextPartInput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TextPartInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TextPartInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextPartInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextPartInput) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *TextPartInput) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TextPartInput) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TextPartInput) SetText(v string)`

SetText sets Text field to given value.


### GetSynthetic

`func (o *TextPartInput) GetSynthetic() bool`

GetSynthetic returns the Synthetic field if non-nil, zero value otherwise.

### GetSyntheticOk

`func (o *TextPartInput) GetSyntheticOk() (*bool, bool)`

GetSyntheticOk returns a tuple with the Synthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetic

`func (o *TextPartInput) SetSynthetic(v bool)`

SetSynthetic sets Synthetic field to given value.

### HasSynthetic

`func (o *TextPartInput) HasSynthetic() bool`

HasSynthetic returns a boolean if a field has been set.

### GetTime

`func (o *TextPartInput) GetTime() TextPartTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TextPartInput) GetTimeOk() (*TextPartTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TextPartInput) SetTime(v TextPartTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *TextPartInput) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


