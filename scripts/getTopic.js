// 根据输入的函数字符串，计算它的topic并输出
// 使用：node getTopic.js "ProWithdraw(address,uint64,uint8,uint256,uint256)"
//const hre = require("hardhat");
const web3 = require("web3")

async function main() {
    const args = process.argv.slice(2);
    //const str = args[0];

    const topic0 = web3.utils.sha3("Output(uint8)");
    const topic1 = web3.utils.sha3("Paytime(uint256)");


    console.log("event Output's topic ", topic0);
    console.log("event Paytime's topic ", topic1);


    //const topic = web3.utils.sha3(str);
    //console.log(str, "topic:", topic);
}

main()
.then(() => process.exit(0))
.catch((error) => {
    console.error(error);
    process.exit(1);
  });
