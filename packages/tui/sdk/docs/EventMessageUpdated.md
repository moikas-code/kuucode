# EventMessageUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Properties** | [**EventMessageUpdatedProperties**](EventMessageUpdatedProperties.md) |  | 

## Methods

### NewEventMessageUpdated

`func NewEventMessageUpdated(type_ string, properties EventMessageUpdatedProperties, ) *EventMessageUpdated`

NewEventMessageUpdated instantiates a new EventMessageUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMessageUpdatedWithDefaults

`func NewEventMessageUpdatedWithDefaults() *EventMessageUpdated`

NewEventMessageUpdatedWithDefaults instantiates a new EventMessageUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventMessageUpdated) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventMessageUpdated) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventMessageUpdated) SetType(v string)`

SetType sets Type field to given value.


### GetProperties

`func (o *EventMessageUpdated) GetProperties() EventMessageUpdatedProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventMessageUpdated) GetPropertiesOk() (*EventMessageUpdatedProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventMessageUpdated) SetProperties(v EventMessageUpdatedProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


