# PostLogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | Service name for the log entry | 
**Level** | **string** | Log level | 
**Message** | **string** | Log message | 
**Extra** | Pointer to **map[string]interface{}** | Additional metadata for the log entry | [optional] 

## Methods

### NewPostLogRequest

`func NewPostLogRequest(service string, level string, message string, ) *PostLogRequest`

NewPostLogRequest instantiates a new PostLogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLogRequestWithDefaults

`func NewPostLogRequestWithDefaults() *PostLogRequest`

NewPostLogRequestWithDefaults instantiates a new PostLogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *PostLogRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *PostLogRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *PostLogRequest) SetService(v string)`

SetService sets Service field to given value.


### GetLevel

`func (o *PostLogRequest) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PostLogRequest) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PostLogRequest) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *PostLogRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostLogRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostLogRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExtra

`func (o *PostLogRequest) GetExtra() map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *PostLogRequest) GetExtraOk() (*map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *PostLogRequest) SetExtra(v map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *PostLogRequest) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


