# EventMessagePartUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Properties** | [**EventMessagePartUpdatedProperties**](EventMessagePartUpdatedProperties.md) |  | 

## Methods

### NewEventMessagePartUpdated

`func NewEventMessagePartUpdated(type_ string, properties EventMessagePartUpdatedProperties, ) *EventMessagePartUpdated`

NewEventMessagePartUpdated instantiates a new EventMessagePartUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMessagePartUpdatedWithDefaults

`func NewEventMessagePartUpdatedWithDefaults() *EventMessagePartUpdated`

NewEventMessagePartUpdatedWithDefaults instantiates a new EventMessagePartUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventMessagePartUpdated) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventMessagePartUpdated) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventMessagePartUpdated) SetType(v string)`

SetType sets Type field to given value.


### GetProperties

`func (o *EventMessagePartUpdated) GetProperties() EventMessagePartUpdatedProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventMessagePartUpdated) GetPropertiesOk() (*EventMessagePartUpdatedProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventMessagePartUpdated) SetProperties(v EventMessagePartUpdatedProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


