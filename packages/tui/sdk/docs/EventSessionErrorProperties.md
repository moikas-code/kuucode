# EventSessionErrorProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionID** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**AssistantMessageError**](AssistantMessageError.md) |  | [optional] 

## Methods

### NewEventSessionErrorProperties

`func NewEventSessionErrorProperties() *EventSessionErrorProperties`

NewEventSessionErrorProperties instantiates a new EventSessionErrorProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSessionErrorPropertiesWithDefaults

`func NewEventSessionErrorPropertiesWithDefaults() *EventSessionErrorProperties`

NewEventSessionErrorPropertiesWithDefaults instantiates a new EventSessionErrorProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionID

`func (o *EventSessionErrorProperties) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *EventSessionErrorProperties) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *EventSessionErrorProperties) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.

### HasSessionID

`func (o *EventSessionErrorProperties) HasSessionID() bool`

HasSessionID returns a boolean if a field has been set.

### GetError

`func (o *EventSessionErrorProperties) GetError() AssistantMessageError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventSessionErrorProperties) GetErrorOk() (*AssistantMessageError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventSessionErrorProperties) SetError(v AssistantMessageError)`

SetError sets Error field to given value.

### HasError

`func (o *EventSessionErrorProperties) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


