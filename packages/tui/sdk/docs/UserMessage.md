# UserMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionID** | **string** |  | 
**Role** | **string** |  | 
**Time** | [**PermissionInfoTime**](PermissionInfoTime.md) |  | 

## Methods

### NewUserMessage

`func NewUserMessage(id string, sessionID string, role string, time PermissionInfoTime, ) *UserMessage`

NewUserMessage instantiates a new UserMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMessageWithDefaults

`func NewUserMessageWithDefaults() *UserMessage`

NewUserMessageWithDefaults instantiates a new UserMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserMessage) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *UserMessage) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *UserMessage) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *UserMessage) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetRole

`func (o *UserMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetTime

`func (o *UserMessage) GetTime() PermissionInfoTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UserMessage) GetTimeOk() (*PermissionInfoTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UserMessage) SetTime(v PermissionInfoTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


