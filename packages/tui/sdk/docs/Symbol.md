# Symbol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Kind** | **float32** |  | 
**Location** | [**SymbolLocation**](SymbolLocation.md) |  | 

## Methods

### NewSymbol

`func NewSymbol(name string, kind float32, location SymbolLocation, ) *Symbol`

NewSymbol instantiates a new Symbol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolWithDefaults

`func NewSymbolWithDefaults() *Symbol`

NewSymbolWithDefaults instantiates a new Symbol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Symbol) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Symbol) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Symbol) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *Symbol) GetKind() float32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Symbol) GetKindOk() (*float32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Symbol) SetKind(v float32)`

SetKind sets Kind field to given value.


### GetLocation

`func (o *Symbol) GetLocation() SymbolLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Symbol) GetLocationOk() (*SymbolLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Symbol) SetLocation(v SymbolLocation)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


