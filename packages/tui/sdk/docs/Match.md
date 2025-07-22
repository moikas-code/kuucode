# Match

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | [**MatchPath**](MatchPath.md) |  | 
**Lines** | [**MatchPath**](MatchPath.md) |  | 
**LineNumber** | **float32** |  | 
**AbsoluteOffset** | **float32** |  | 
**Submatches** | [**[]MatchSubmatchesInner**](MatchSubmatchesInner.md) |  | 

## Methods

### NewMatch

`func NewMatch(path MatchPath, lines MatchPath, lineNumber float32, absoluteOffset float32, submatches []MatchSubmatchesInner, ) *Match`

NewMatch instantiates a new Match object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchWithDefaults

`func NewMatchWithDefaults() *Match`

NewMatchWithDefaults instantiates a new Match object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *Match) GetPath() MatchPath`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Match) GetPathOk() (*MatchPath, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Match) SetPath(v MatchPath)`

SetPath sets Path field to given value.


### GetLines

`func (o *Match) GetLines() MatchPath`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Match) GetLinesOk() (*MatchPath, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Match) SetLines(v MatchPath)`

SetLines sets Lines field to given value.


### GetLineNumber

`func (o *Match) GetLineNumber() float32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *Match) GetLineNumberOk() (*float32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *Match) SetLineNumber(v float32)`

SetLineNumber sets LineNumber field to given value.


### GetAbsoluteOffset

`func (o *Match) GetAbsoluteOffset() float32`

GetAbsoluteOffset returns the AbsoluteOffset field if non-nil, zero value otherwise.

### GetAbsoluteOffsetOk

`func (o *Match) GetAbsoluteOffsetOk() (*float32, bool)`

GetAbsoluteOffsetOk returns a tuple with the AbsoluteOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteOffset

`func (o *Match) SetAbsoluteOffset(v float32)`

SetAbsoluteOffset sets AbsoluteOffset field to given value.


### GetSubmatches

`func (o *Match) GetSubmatches() []MatchSubmatchesInner`

GetSubmatches returns the Submatches field if non-nil, zero value otherwise.

### GetSubmatchesOk

`func (o *Match) GetSubmatchesOk() (*[]MatchSubmatchesInner, bool)`

GetSubmatchesOk returns a tuple with the Submatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmatches

`func (o *Match) SetSubmatches(v []MatchSubmatchesInner)`

SetSubmatches sets Submatches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


