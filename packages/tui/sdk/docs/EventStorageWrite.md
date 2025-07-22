# EventStorageWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Properties** | [**EventStorageWriteProperties**](EventStorageWriteProperties.md) |  | 

## Methods

### NewEventStorageWrite

`func NewEventStorageWrite(type_ string, properties EventStorageWriteProperties, ) *EventStorageWrite`

NewEventStorageWrite instantiates a new EventStorageWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStorageWriteWithDefaults

`func NewEventStorageWriteWithDefaults() *EventStorageWrite`

NewEventStorageWriteWithDefaults instantiates a new EventStorageWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventStorageWrite) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventStorageWrite) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventStorageWrite) SetType(v string)`

SetType sets Type field to given value.


### GetProperties

`func (o *EventStorageWrite) GetProperties() EventStorageWriteProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventStorageWrite) GetPropertiesOk() (*EventStorageWriteProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventStorageWrite) SetProperties(v EventStorageWriteProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


