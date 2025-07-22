# SymbolSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | [**FilePartSourceText**](FilePartSourceText.md) |  | 
**Type** | **string** |  | 
**Path** | **string** |  | 
**Range** | [**Range**](Range.md) |  | 
**Name** | **string** |  | 
**Kind** | **int32** |  | 

## Methods

### NewSymbolSource

`func NewSymbolSource(text FilePartSourceText, type_ string, path string, range_ Range, name string, kind int32, ) *SymbolSource`

NewSymbolSource instantiates a new SymbolSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolSourceWithDefaults

`func NewSymbolSourceWithDefaults() *SymbolSource`

NewSymbolSourceWithDefaults instantiates a new SymbolSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SymbolSource) GetText() FilePartSourceText`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SymbolSource) GetTextOk() (*FilePartSourceText, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SymbolSource) SetText(v FilePartSourceText)`

SetText sets Text field to given value.


### GetType

`func (o *SymbolSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SymbolSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SymbolSource) SetType(v string)`

SetType sets Type field to given value.


### GetPath

`func (o *SymbolSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SymbolSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SymbolSource) SetPath(v string)`

SetPath sets Path field to given value.


### GetRange

`func (o *SymbolSource) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SymbolSource) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SymbolSource) SetRange(v Range)`

SetRange sets Range field to given value.


### GetName

`func (o *SymbolSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SymbolSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SymbolSource) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *SymbolSource) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SymbolSource) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SymbolSource) SetKind(v int32)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


