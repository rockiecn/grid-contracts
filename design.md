# Introduction

Design doc for GRID contracts.

## Registry

Registry contract is used to manage all providers' details, including host information, resource number and prices.

### Data structures

struct resources:

Records information about all resources, including cpu, memory, storage, gpu. Two instances in all, one is for total resources, the other is for available resources.

struct info:

Records nodes' detailed information, local information and resource information.

### Methods

Get:

Used to acquire a node's information. A node's address is needed when called.

Set:

Used to set a node's information. The caller must be the node's address and information data must be passed when called.

Update:

Used to update a node's available resource information. The caller must be market contract. Each time an order is created, the related resources must be deduct from the available resource in the node information. When an order is completed or cancelled, the resources in this order must be given back to the available resources.

## Market

The market contract is used to record the leasing relationship between users and providers. It implemented the tasks of order creating and profit assignment by running the contract methods.

### Data structures

struct order:

Order infomation, including pricePerHour and resources struct recording the price and number of each resource.

deposit records the number of token a user payed for this order.

remuneration records the number of token should be paid to the provider.

userCheck and providerCheck are used to represent if this order is confirmed by the user or provider.

activateHeight is used to record the height of the blockchain when this order is activated.

settleHeight is used to record the height of the blockchain when this order is settled.

probation is used to record the duration of the probationary period.

duration is used to record the service duration of this order.

isRunning is used to record this running status of this order. It could be unactivated, running, cancelled or completed.

mapping(address => mapping(address => order)) orders:

This data structure is used to record the order information between users and providers.

key1 is the user address, key2 is the provider address.

### Methods

createOrder:

The order creation method for users. Enough balance must be in the user's token balance before calling this method. Only one order can exist for a user-provider pair at a same time. If a user wants to create a order, he must cancel the current order or wait for it to be completed.
The parameters for this method are order details and the provider address.

cancelOrder:

This method is used to cancel the existing order for a user-provider pair.

The cancelOrder method will call the settle method to settle the service of this order, and set the order's isRunning status to cancelled. At the end of this method, the erc20's transfer method will be called to fetch the remaining token in market contract for both the user and the provider.

The user address or the provider address is passed along with the msg.sender in transaction to locate the order to be cancelled.

settle:

This method used to settle the toekn for this order. It calculates the number of token should be paid to the provider, and the remaining token of the order. The calculation will consider the probation, the current block height and the settlement height. If the service length is over, the status of this order should be set to completed.

The user address and provider address should be passed when call this method to locate the order.

activateOrder:

This method is called by providers, and it is used to activate the order status to start the service.

The user address should be passed to this method.

withdraw:

This method is used to withdraw the token in this order for a user or a provider. The status of this order must be completed or cancelled when call this method. It will call the settle method to settle the service for this order in advance, and transfer the token for the calling user or provider.

The amount of token to withdraw, caller address should be passed when call this method.

checkPermission:

This method is used to check if the user has the permission to create a order, in order to avoid the Sybil Attack.

## Pledge

Pledge logic design:

Pledge token, used to record the total pledge amount and the specific pledge amount for each node.

Pay token, used to pay for each order, and calculate the profit according to the total unit profit, node unit profit, total pledge amount in pledge token and the node's pledge amount in pledge token.

Data:

Total unit token profit for the pay token:

To record the profit for a single token in pledge token

Node unit token profit for a node:

To record the unit profit that has already been checked out.

About mint:

The pay token(erc20 token) will mint some profit tokens when an order is created, and these profit tokens will be assigned to each node who has pledged in the pledge token.

## ERC20

