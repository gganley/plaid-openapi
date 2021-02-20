# InstitutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemLogins** | [**ProductStatus**](ProductStatus.md) |  | 
**TransactionsUpdates** | [**ProductStatus**](ProductStatus.md) |  | 
**Auth** | [**ProductStatus**](ProductStatus.md) |  | 
**Balance** | [**ProductStatus**](ProductStatus.md) |  | 
**Identity** | [**ProductStatus**](ProductStatus.md) |  | 
**InvestmentsUpdates** | [**ProductStatus**](ProductStatus.md) |  | 
**HealthIncidents** | Pointer to [**[]HealthIncident**](HealthIncident.md) | Details of recent health incidents associated with the institution. | [optional] 

## Methods

### NewInstitutionStatus

`func NewInstitutionStatus(itemLogins ProductStatus, transactionsUpdates ProductStatus, auth ProductStatus, balance ProductStatus, identity ProductStatus, investmentsUpdates ProductStatus, ) *InstitutionStatus`

NewInstitutionStatus instantiates a new InstitutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionStatusWithDefaults

`func NewInstitutionStatusWithDefaults() *InstitutionStatus`

NewInstitutionStatusWithDefaults instantiates a new InstitutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemLogins

`func (o *InstitutionStatus) GetItemLogins() ProductStatus`

GetItemLogins returns the ItemLogins field if non-nil, zero value otherwise.

### GetItemLoginsOk

`func (o *InstitutionStatus) GetItemLoginsOk() (*ProductStatus, bool)`

GetItemLoginsOk returns a tuple with the ItemLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLogins

`func (o *InstitutionStatus) SetItemLogins(v ProductStatus)`

SetItemLogins sets ItemLogins field to given value.


### GetTransactionsUpdates

`func (o *InstitutionStatus) GetTransactionsUpdates() ProductStatus`

GetTransactionsUpdates returns the TransactionsUpdates field if non-nil, zero value otherwise.

### GetTransactionsUpdatesOk

`func (o *InstitutionStatus) GetTransactionsUpdatesOk() (*ProductStatus, bool)`

GetTransactionsUpdatesOk returns a tuple with the TransactionsUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsUpdates

`func (o *InstitutionStatus) SetTransactionsUpdates(v ProductStatus)`

SetTransactionsUpdates sets TransactionsUpdates field to given value.


### GetAuth

`func (o *InstitutionStatus) GetAuth() ProductStatus`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *InstitutionStatus) GetAuthOk() (*ProductStatus, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *InstitutionStatus) SetAuth(v ProductStatus)`

SetAuth sets Auth field to given value.


### GetBalance

`func (o *InstitutionStatus) GetBalance() ProductStatus`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *InstitutionStatus) GetBalanceOk() (*ProductStatus, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *InstitutionStatus) SetBalance(v ProductStatus)`

SetBalance sets Balance field to given value.


### GetIdentity

`func (o *InstitutionStatus) GetIdentity() ProductStatus`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *InstitutionStatus) GetIdentityOk() (*ProductStatus, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *InstitutionStatus) SetIdentity(v ProductStatus)`

SetIdentity sets Identity field to given value.


### GetInvestmentsUpdates

`func (o *InstitutionStatus) GetInvestmentsUpdates() ProductStatus`

GetInvestmentsUpdates returns the InvestmentsUpdates field if non-nil, zero value otherwise.

### GetInvestmentsUpdatesOk

`func (o *InstitutionStatus) GetInvestmentsUpdatesOk() (*ProductStatus, bool)`

GetInvestmentsUpdatesOk returns a tuple with the InvestmentsUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentsUpdates

`func (o *InstitutionStatus) SetInvestmentsUpdates(v ProductStatus)`

SetInvestmentsUpdates sets InvestmentsUpdates field to given value.


### GetHealthIncidents

`func (o *InstitutionStatus) GetHealthIncidents() []HealthIncident`

GetHealthIncidents returns the HealthIncidents field if non-nil, zero value otherwise.

### GetHealthIncidentsOk

`func (o *InstitutionStatus) GetHealthIncidentsOk() (*[]HealthIncident, bool)`

GetHealthIncidentsOk returns a tuple with the HealthIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthIncidents

`func (o *InstitutionStatus) SetHealthIncidents(v []HealthIncident)`

SetHealthIncidents sets HealthIncidents field to given value.

### HasHealthIncidents

`func (o *InstitutionStatus) HasHealthIncidents() bool`

HasHealthIncidents returns a boolean if a field has been set.

### SetHealthIncidentsNil

`func (o *InstitutionStatus) SetHealthIncidentsNil(b bool)`

 SetHealthIncidentsNil sets the value for HealthIncidents to be an explicit nil

### UnsetHealthIncidents
`func (o *InstitutionStatus) UnsetHealthIncidents()`

UnsetHealthIncidents ensures that no value is present for HealthIncidents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


