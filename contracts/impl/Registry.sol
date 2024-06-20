// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Registry
 * @dev For cp registry
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
*/
contract Registry {
    // record the number of each resource type
    struct Resources{
        uint64 nCPU;
        uint64 nGPU;
        uint64 nMEM;
        uint64 nDISK;
    }

    // cp info
    struct Info{
        address addr;
        string ip;
        string domain;
        string port;

        Resources total; // total resources
        Resources avail; // available resources
    }

    // store all cp info
    mapping(address => Info) registry;
    // store all keys of cp(address)
    address[] private keys;

    /**
     * @dev register cp info
     * @param cpu value of cpu
    */
    function register(string memory ip, string memory domain, string memory port, uint64 cpu,  uint64 gpu,  uint64 mem,  uint64 disk) public {
        registry[msg.sender].addr=msg.sender;
        registry[msg.sender].ip=ip;
        registry[msg.sender].domain=domain;
        registry[msg.sender].port=port;

        // total resource
        registry[msg.sender].total.nCPU=cpu;
        registry[msg.sender].total.nGPU=gpu;
        registry[msg.sender].total.nMEM=mem;
        registry[msg.sender].total.nDISK=disk;

        // available resource
        registry[msg.sender].avail.nCPU=cpu;
        registry[msg.sender].avail.nGPU=gpu;
        registry[msg.sender].avail.nMEM=mem;
        registry[msg.sender].avail.nDISK=disk;

        // store the provider key
        keys.push(msg.sender);
    }

    /**
     * @dev get cp info
     * @return all info of a cp
    */
    function get(address a) public view returns (Info memory){
        return registry[a];
    }

    /**
     * @dev get the cp keys
    */
    function getKeys() external view returns(address[] memory) {
        return keys;
    }

    /**
     * @dev update avail with order info, only called by market contract
    */
    //function update(uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) public {
    function update(address provider, uint64 nCPU, uint64 nGPU, uint64 nMEM, uint64 nDISK) external {
        require(nCPU <= registry[provider].avail.nCPU, "not enough cpu");
        require(nGPU <= registry[provider].avail.nGPU, "not enough gpu");
        require(nMEM <= registry[provider].avail.nMEM, "not enough mem");
        require(nDISK <= registry[provider].avail.nDISK, "not enough disk");

        // update available resource for provider
        registry[provider].avail.nCPU -= nCPU;
        registry[provider].avail.nGPU -= nGPU;
        registry[provider].avail.nMEM -= nMEM;
        registry[provider].avail.nDISK -= nDISK;
    }

    /**
     * @dev provider revise self cp info, only called by cp himself
    */
    function revise() public {

    }
}