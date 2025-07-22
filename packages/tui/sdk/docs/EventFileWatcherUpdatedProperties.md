# EventFileWatcherUpdatedProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | **string** |  | 
**Event** | [**EventFileWatcherUpdatedPropertiesEvent**](EventFileWatcherUpdatedPropertiesEvent.md) |  | 

## Methods

### NewEventFileWatcherUpdatedProperties

`func NewEventFileWatcherUpdatedProperties(file string, event EventFileWatcherUpdatedPropertiesEvent, ) *EventFileWatcherUpdatedProperties`

NewEventFileWatcherUpdatedProperties instantiates a new EventFileWatcherUpdatedProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFileWatcherUpdatedPropertiesWithDefaults

`func NewEventFileWatcherUpdatedPropertiesWithDefaults() *EventFileWatcherUpdatedProperties`

NewEventFileWatcherUpdatedPropertiesWithDefaults instantiates a new EventFileWatcherUpdatedProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *EventFileWatcherUpdatedProperties) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *EventFileWatcherUpdatedProperties) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *EventFileWatcherUpdatedProperties) SetFile(v string)`

SetFile sets File field to given value.


### GetEvent

`func (o *EventFileWatcherUpdatedProperties) GetEvent() EventFileWatcherUpdatedPropertiesEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventFileWatcherUpdatedProperties) GetEventOk() (*EventFileWatcherUpdatedPropertiesEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventFileWatcherUpdatedProperties) SetEvent(v EventFileWatcherUpdatedPropertiesEvent)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


