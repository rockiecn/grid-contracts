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

 


## ERC20



## Pledge


