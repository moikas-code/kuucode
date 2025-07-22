# SessionRevert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **string** |  | 
**Part** | **float32** |  | 
**Snapshot** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionRevert

`func NewSessionRevert(messageID string, part float32, ) *SessionRevert`

NewSessionRevert instantiates a new SessionRevert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRevertWithDefaults

`func NewSessionRevertWithDefaults() *SessionRevert`

NewSessionRevertWithDefaults instantiates a new SessionRevert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *SessionRevert) GetMessageID() string`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *SessionRevert) GetMessageIDOk() (*string, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *SessionRevert) SetMessageID(v string)`

SetMessageID sets MessageID field to given value.


### GetPart

`func (o *SessionRevert) GetPart() float32`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *SessionRevert) GetPartOk() (*float32, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *SessionRevert) SetPart(v float32)`

SetPart sets Part field to given value.


### GetSnapshot

`func (o *SessionRevert) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SessionRevert) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SessionRevert) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *SessionRevert) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


