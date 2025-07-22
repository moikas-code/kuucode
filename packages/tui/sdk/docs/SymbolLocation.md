# SymbolLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** |  | 
**Range** | [**Range**](Range.md) |  | 

## Methods

### NewSymbolLocation

`func NewSymbolLocation(uri string, range_ Range, ) *SymbolLocation`

NewSymbolLocation instantiates a new SymbolLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolLocationWithDefaults

`func NewSymbolLocationWithDefaults() *SymbolLocation`

NewSymbolLocationWithDefaults instantiates a new SymbolLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *SymbolLocation) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SymbolLocation) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SymbolLocation) SetUri(v string)`

SetUri sets Uri field to given value.


### GetRange

`func (o *SymbolLocation) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SymbolLocation) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SymbolLocation) SetRange(v Range)`

SetRange sets Range field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


