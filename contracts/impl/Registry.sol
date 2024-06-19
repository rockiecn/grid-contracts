// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Registry
 * @dev For cp registry
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
*/
contract Registry {

    struct Resources{
        uint64 CPU;
        uint64 GPU;
        uint64 MEM;
        uint64 DISK;    
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

        registry[msg.sender].total.CPU=cpu;
        registry[msg.sender].total.GPU=gpu;
        registry[msg.sender].total.MEM=mem;
        registry[msg.sender].total.DISK=disk;

        registry[msg.sender].avail.CPU=cpu;
        registry[msg.sender].avail.GPU=gpu;
        registry[msg.sender].avail.MEM=mem;
        registry[msg.sender].avail.DISK=disk;

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
    function update(uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) public {
        assert(cpu <= registry[msg.sender].avail.CPU);
        assert(gpu <= registry[msg.sender].avail.GPU);
        assert(mem <= registry[msg.sender].avail.MEM);
        assert(disk <= registry[msg.sender].avail.DISK);

        registry[msg.sender].avail.CPU -= cpu;
        registry[msg.sender].avail.GPU -= gpu;
        registry[msg.sender].avail.MEM -= mem;
        registry[msg.sender].avail.DISK -= disk;
    }

    /**
     * @dev provider revise self cp info, only called by cp himself
    */
    function revise() public {

    }

}