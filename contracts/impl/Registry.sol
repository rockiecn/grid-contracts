// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

/**
 * @title Registry
 * @dev For cp registry
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
*/
contract Registry {
    // record the price per hour for each resource type
    struct PricePerHour{
        uint64 pCPU;
        uint64 pGPU;
        uint64 pMEM;
        uint64 pDISK;
    }

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

        PricePerHour price; // price for each resource
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
    function register(
                string memory ip, 
                string memory domain, 
                string memory port, 
                uint64 cpu,  
                uint64 gpu,  
                uint64 mem,  
                uint64 disk, 
                uint64 pcpu,
                uint64 pgpu, 
                uint64 pmem, 
                uint64 pdisk) public {
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

        // resource price
        registry[msg.sender].price.pCPU=pcpu;
        registry[msg.sender].price.pGPU=pgpu;
        registry[msg.sender].price.pMEM=pmem;
        registry[msg.sender].price.pDISK=pdisk;

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
     * @dev provider revise self cp info
    */
    function revise(
        string memory ip, 
        string memory domain, 
        string memory port, 
        uint64 cpu,  
        uint64 gpu,  
        uint64 mem,  
        uint64 disk, 
        uint64 pcpu,
        uint64 pgpu, 
        uint64 pmem, 
        uint64 pdisk
    ) public {
        require(cpu >= 0, "cpu cannot less than 0");
        require(gpu >= 0, "gpu cannot less than 0");
        require(mem >= 0, "mem cannot less than 0");
        require(disk >= 0, "disk cannot less than 0");

        require(pcpu >= 0, "cpu price cannot less than 0");
        require(pgpu >= 0, "gpu price cannot less than 0");
        require(pmem >= 0, "mem price cannot less than 0");
        require(pdisk >= 0, "disk price cannot less than 0");

        // record the original info
        Info memory origin = registry[msg.sender];

        // revise base info
        registry[msg.sender].ip = ip;
        registry[msg.sender].domain = domain;
        registry[msg.sender].port = port;

        // revise number
        registry[msg.sender].total.nCPU = cpu;
        registry[msg.sender].total.nGPU = gpu;
        registry[msg.sender].total.nMEM = mem;
        registry[msg.sender].total.nDISK = disk;

        // revise price
        registry[msg.sender].price.pCPU = pcpu;
        registry[msg.sender].price.pGPU = pgpu;
        registry[msg.sender].price.pMEM = pmem;
        registry[msg.sender].price.pDISK = pdisk;

        // revise avail

        // cpu
        if (cpu >= origin.total.nCPU) {
            // avail increased
            uint64 incre = cpu - origin.total.nCPU;
            registry[msg.sender].avail.nCPU += incre;
        } else {
            // avail decreased
            uint64 decre = origin.total.nCPU - cpu;
            // avail decrease to 0
            if (decre >= registry[msg.sender].avail.nCPU) {
                registry[msg.sender].avail.nCPU = 0;
            } else {
                // decrease
                registry[msg.sender].avail.nCPU -= decre;
            }
        }

        // gpu
        if (gpu >= origin.total.nGPU) {
            // avail increased
            uint64 incre = gpu - origin.total.nGPU;
            registry[msg.sender].avail.nGPU += incre;
        } else {
            // avail decreased
            uint64 decre = origin.total.nGPU - gpu;
            // avail decrease to 0
            if (decre >= registry[msg.sender].avail.nGPU) {
                registry[msg.sender].avail.nGPU = 0;
            } else {
                // decrease
                registry[msg.sender].avail.nGPU -= decre;
            }
        }

        // mem
        if (mem >= origin.total.nMEM) {
            // avail increased
            uint64 incre = mem - origin.total.nMEM;
            registry[msg.sender].avail.nMEM += incre;
        } else {
            // avail decreased
            uint64 decre = origin.total.nMEM - mem;
            // avail decrease to 0
            if (decre >= registry[msg.sender].avail.nMEM) {
                registry[msg.sender].avail.nMEM = 0;
            } else {
                // decrease
                registry[msg.sender].avail.nMEM -= decre;
            }
        }

        // disk
        if (disk >= origin.total.nDISK) {
            // avail increased
            uint64 incre = disk - origin.total.nDISK;
            registry[msg.sender].avail.nDISK += incre;
        } else {
            // avail decreased
            uint64 decre = origin.total.nDISK - disk;
            // avail decrease to 0
            if (decre >= registry[msg.sender].avail.nDISK) {
                registry[msg.sender].avail.nDISK = 0;
            } else {
                // decrease
                registry[msg.sender].avail.nDISK -= decre;
            }
        }
    }


}