# PostSessionByIdMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | Pointer to **string** |  | [optional] 
**ProviderID** | **string** |  | 
**ModelID** | **string** |  | 
**Mode** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **map[string]bool** |  | [optional] 
**Parts** | [**[]PostSessionByIdMessageRequestPartsInner**](PostSessionByIdMessageRequestPartsInner.md) |  | 

## Methods

### NewPostSessionByIdMessageRequest

`func NewPostSessionByIdMessageRequest(providerID string, modelID string, parts []PostSessionByIdMessageRequestPartsInner, ) *PostSessionByIdMessageRequest`

NewPostSessionByIdMessageRequest instantiates a new PostSessionByIdMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSessionByIdMessageRequestWithDefaults

`func NewPostSessionByIdMessageRequestWithDefaults() *PostSessionByIdMessageRequest`

NewPostSessionByIdMessageRequestWithDefaults instantiates a new PostSessionByIdMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *PostSessionByIdMessageRequest) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *PostSessionByIdMessageRequest) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *PostSessionByIdMessageRequest) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.

### HasMessageID

`func (o *PostSessionByIdMessageRequest) HasMessageID() bool`

HasMessageID returns a boolean if a field has been set.

### GetProviderID

`func (o *PostSessionByIdMessageRequest) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *PostSessionByIdMessageRequest) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *PostSessionByIdMessageRequest) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.


### GetModelID

`func (o *PostSessionByIdMessageRequest) GetModelID() string`

GetModelID returns the ModelID field if non-nil, zero value otherwise.

### GetModelIDOk

`func (o *PostSessionByIdMessageRequest) GetModelIDOk() (*string, bool)`

GetModelIDOk returns a tuple with the ModelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelID

`func (o *PostSessionByIdMessageRequest) SetModelID(v string)`

SetModelID sets ModelID field to given value.


### GetMode

`func (o *PostSessionByIdMessageRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostSessionByIdMessageRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostSessionByIdMessageRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostSessionByIdMessageRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTools

`func (o *PostSessionByIdMessageRequest) GetTools() map[string]bool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *PostSessionByIdMessageRequest) GetToolsOk() (*map[string]bool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *PostSessionByIdMessageRequest) SetTools(v map[string]bool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *PostSessionByIdMessageRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetParts

`func (o *PostSessionByIdMessageRequest) GetParts() []PostSessionByIdMessageRequestPartsInner`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *PostSessionByIdMessageRequest) GetPartsOk() (*[]PostSessionByIdMessageRequestPartsInner, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *PostSessionByIdMessageRequest) SetParts(v []PostSessionByIdMessageRequestPartsInner)`

SetParts sets Parts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


