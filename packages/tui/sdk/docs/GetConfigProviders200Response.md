# GetConfigProviders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | [**[]Provider**](Provider.md) |  | 
**Default** | **map[string]string** |  | 

## Methods

### NewGetConfigProviders200Response

`func NewGetConfigProviders200Response(providers []Provider, default_ map[string]string, ) *GetConfigProviders200Response`

NewGetConfigProviders200Response instantiates a new GetConfigProviders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfigProviders200ResponseWithDefaults

`func NewGetConfigProviders200ResponseWithDefaults() *GetConfigProviders200Response`

NewGetConfigProviders200ResponseWithDefaults instantiates a new GetConfigProviders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *GetConfigProviders200Response) GetProviders() []Provider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *GetConfigProviders200Response) GetProvidersOk() (*[]Provider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *GetConfigProviders200Response) SetProviders(v []Provider)`

SetProviders sets Providers field to given value.


### GetDefault

`func (o *GetConfigProviders200Response) GetDefault() map[string]string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *GetConfigProviders200Response) GetDefaultOk() (*map[string]string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *GetConfigProviders200Response) SetDefault(v map[string]string)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


