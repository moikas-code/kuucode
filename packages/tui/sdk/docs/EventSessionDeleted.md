# EventSessionDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Properties** | [**EventSessionUpdatedProperties**](EventSessionUpdatedProperties.md) |  | 

## Methods

### NewEventSessionDeleted

`func NewEventSessionDeleted(type_ string, properties EventSessionUpdatedProperties, ) *EventSessionDeleted`

NewEventSessionDeleted instantiates a new EventSessionDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSessionDeletedWithDefaults

`func NewEventSessionDeletedWithDefaults() *EventSessionDeleted`

NewEventSessionDeletedWithDefaults instantiates a new EventSessionDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventSessionDeleted) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSessionDeleted) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSessionDeleted) SetType(v string)`

SetType sets Type field to given value.


### GetProperties

`func (o *EventSessionDeleted) GetProperties() EventSessionUpdatedProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventSessionDeleted) GetPropertiesOk() (*EventSessionUpdatedProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventSessionDeleted) SetProperties(v EventSessionUpdatedProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


