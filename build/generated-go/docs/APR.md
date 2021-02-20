# APR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AprPercentage** | **float32** | Annual Percentage Rate applied.  | 
**AprType** | **string** | The type of balance to which the APR applies. | 
**BalanceSubjectToApr** | Pointer to **NullableFloat32** | Amount of money that is subjected to the APR if a balance was carried beyond payment due date. How it is calculated can vary by card issuer. It is often calculated as an average daily balance. | [optional] 
**InterestChargeAmount** | Pointer to **NullableFloat32** | Amount of money charged due to interest from last statement. | [optional] 

## Methods

### NewAPR

`func NewAPR(aprPercentage float32, aprType string, ) *APR`

NewAPR instantiates a new APR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPRWithDefaults

`func NewAPRWithDefaults() *APR`

NewAPRWithDefaults instantiates a new APR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAprPercentage

`func (o *APR) GetAprPercentage() float32`

GetAprPercentage returns the AprPercentage field if non-nil, zero value otherwise.

### GetAprPercentageOk

`func (o *APR) GetAprPercentageOk() (*float32, bool)`

GetAprPercentageOk returns a tuple with the AprPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprPercentage

`func (o *APR) SetAprPercentage(v float32)`

SetAprPercentage sets AprPercentage field to given value.


### GetAprType

`func (o *APR) GetAprType() string`

GetAprType returns the AprType field if non-nil, zero value otherwise.

### GetAprTypeOk

`func (o *APR) GetAprTypeOk() (*string, bool)`

GetAprTypeOk returns a tuple with the AprType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprType

`func (o *APR) SetAprType(v string)`

SetAprType sets AprType field to given value.


### GetBalanceSubjectToApr

`func (o *APR) GetBalanceSubjectToApr() float32`

GetBalanceSubjectToApr returns the BalanceSubjectToApr field if non-nil, zero value otherwise.

### GetBalanceSubjectToAprOk

`func (o *APR) GetBalanceSubjectToAprOk() (*float32, bool)`

GetBalanceSubjectToAprOk returns a tuple with the BalanceSubjectToApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSubjectToApr

`func (o *APR) SetBalanceSubjectToApr(v float32)`

SetBalanceSubjectToApr sets BalanceSubjectToApr field to given value.

### HasBalanceSubjectToApr

`func (o *APR) HasBalanceSubjectToApr() bool`

HasBalanceSubjectToApr returns a boolean if a field has been set.

### SetBalanceSubjectToAprNil

`func (o *APR) SetBalanceSubjectToAprNil(b bool)`

 SetBalanceSubjectToAprNil sets the value for BalanceSubjectToApr to be an explicit nil

### UnsetBalanceSubjectToApr
`func (o *APR) UnsetBalanceSubjectToApr()`

UnsetBalanceSubjectToApr ensures that no value is present for BalanceSubjectToApr, not even an explicit nil
### GetInterestChargeAmount

`func (o *APR) GetInterestChargeAmount() float32`

GetInterestChargeAmount returns the InterestChargeAmount field if non-nil, zero value otherwise.

### GetInterestChargeAmountOk

`func (o *APR) GetInterestChargeAmountOk() (*float32, bool)`

GetInterestChargeAmountOk returns a tuple with the InterestChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestChargeAmount

`func (o *APR) SetInterestChargeAmount(v float32)`

SetInterestChargeAmount sets InterestChargeAmount field to given value.

### HasInterestChargeAmount

`func (o *APR) HasInterestChargeAmount() bool`

HasInterestChargeAmount returns a boolean if a field has been set.

### SetInterestChargeAmountNil

`func (o *APR) SetInterestChargeAmountNil(b bool)`

 SetInterestChargeAmountNil sets the value for InterestChargeAmount to be an explicit nil

### UnsetInterestChargeAmount
`func (o *APR) UnsetInterestChargeAmount()`

UnsetInterestChargeAmount ensures that no value is present for InterestChargeAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


