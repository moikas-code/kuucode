# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParentID** | Pointer to **string** |  | [optional] 
**Share** | Pointer to [**SessionShare**](SessionShare.md) |  | [optional] 
**Title** | **string** |  | 
**Version** | **string** |  | 
**Time** | [**SessionTime**](SessionTime.md) |  | 
**Revert** | Pointer to [**SessionRevert**](SessionRevert.md) |  | [optional] 

## Methods

### NewSession

`func NewSession(id string, title string, version string, time SessionTime, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.


### GetParentID

`func (o *Session) GetParentID() string`

GetParentID returns the ParentID field if non-nil, zero value otherwise.

### GetParentIDOk

`func (o *Session) GetParentIDOk() (*string, bool)`

GetParentIDOk returns a tuple with the ParentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentID

`func (o *Session) SetParentID(v string)`

SetParentID sets ParentID field to given value.

### HasParentID

`func (o *Session) HasParentID() bool`

HasParentID returns a boolean if a field has been set.

### GetShare

`func (o *Session) GetShare() SessionShare`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *Session) GetShareOk() (*SessionShare, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *Session) SetShare(v SessionShare)`

SetShare sets Share field to given value.

### HasShare

`func (o *Session) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetTitle

`func (o *Session) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Session) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Session) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVersion

`func (o *Session) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Session) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Session) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetTime

`func (o *Session) GetTime() SessionTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Session) GetTimeOk() (*SessionTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Session) SetTime(v SessionTime)`

SetTime sets Time field to given value.


### GetRevert

`func (o *Session) GetRevert() SessionRevert`

GetRevert returns the Revert field if non-nil, zero value otherwise.

### GetRevertOk

`func (o *Session) GetRevertOk() (*SessionRevert, bool)`

GetRevertOk returns a tuple with the Revert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevert

`func (o *Session) SetRevert(v SessionRevert)`

SetRevert sets Revert field to given value.

### HasRevert

`func (o *Session) HasRevert() bool`

HasRevert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


