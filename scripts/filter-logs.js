const { Web3 } = require('web3');

var web3 = new Web3("https://chain.metamemo.one:8501");

// 查询subOrder触发总次数

async function main() {
    var res = await web3.eth.getPastLogs({
        fromBlock: 13794630,
        toBlock: "latest",
        address: "0xbf29B46f21C6A5eE14C8E6aFC0Eaa1EA9F1F2c09", // 从指定合约地址中过滤
        topics: ["0x28c50dfc457ec963adc1b69cfb43284b512fe0b23719b3ab1cdf414d2ff3246c"]
        //"0x7f788365ed24e64ac0ae6aac49ec58094505982a137d3b3379afd2267fffdd75"
    })
    console.log(res);
    console.log("event nums:"+ res.length);
}

main()
.then(() => process.exit(0))
.catch((error) => {
    console.error(error);
    process.exit(1);
  });