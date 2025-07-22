# PermissionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionID** | **string** |  | 
**Title** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Time** | [**PermissionInfoTime**](PermissionInfoTime.md) |  | 

## Methods

### NewPermissionInfo

`func NewPermissionInfo(id string, sessionID string, title string, metadata map[string]interface{}, time PermissionInfoTime, ) *PermissionInfo`

NewPermissionInfo instantiates a new PermissionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionInfoWithDefaults

`func NewPermissionInfoWithDefaults() *PermissionInfo`

NewPermissionInfoWithDefaults instantiates a new PermissionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionInfo) SetId(v string)`

SetId sets Id field to given value.


### GetSessionID

`func (o *PermissionInfo) GetSessionID() string`

GetSessionID returns the SessionID field if non-nil, zero value otherwise.

### GetSessionIDOk

`func (o *PermissionInfo) GetSessionIDOk() (*string, bool)`

GetSessionIDOk returns a tuple with the SessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionID

`func (o *PermissionInfo) SetSessionID(v string)`

SetSessionID sets SessionID field to given value.


### GetTitle

`func (o *PermissionInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PermissionInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PermissionInfo) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMetadata

`func (o *PermissionInfo) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionInfo) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionInfo) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetTime

`func (o *PermissionInfo) GetTime() PermissionInfoTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PermissionInfo) GetTimeOk() (*PermissionInfoTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PermissionInfo) SetTime(v PermissionInfoTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


