# ProviderAuthError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Data** | [**ProviderAuthErrorData**](ProviderAuthErrorData.md) |  | 

## Methods

### NewProviderAuthError

`func NewProviderAuthError(name string, data ProviderAuthErrorData, ) *ProviderAuthError`

NewProviderAuthError instantiates a new ProviderAuthError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderAuthErrorWithDefaults

`func NewProviderAuthErrorWithDefaults() *ProviderAuthError`

NewProviderAuthErrorWithDefaults instantiates a new ProviderAuthError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProviderAuthError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderAuthError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderAuthError) SetName(v string)`

SetName sets Name field to given value.


### GetData

`func (o *ProviderAuthError) GetData() ProviderAuthErrorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProviderAuthError) GetDataOk() (*ProviderAuthErrorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProviderAuthError) SetData(v ProviderAuthErrorData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


