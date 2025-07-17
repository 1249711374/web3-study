# web3 学习打卡任务


## 任务一
    打印结果:
        2025/07/17 13:03:35 block number: 0xc81fd177c188663bf18d03f894c8c4504219922bedaf40a7b1fd4bf4d3ccca59, tranaction count: 160, block.time: 1752721668 
        2025/07/17 13:03:37 txHash: 0x6ba87b08d12fb56dd796d5c6de057fcb8113ae2bdfa6f1a753159d2d4bffac77

## 任务二
    生成abi/bin 文件 `solc --abi --bin Visit.sol -o .`
    生成go文件 `abigen --bin=Visit.bin --abi=Visit.abi --pkg=visit --out=visit.go`

 ## 整体输出结果(截图见 go-ethereum-1.png)

    =========== 任务1 ===================>
    2025/07/17 15:24:21 block number: 0xc81fd177c188663bf18d03f894c8c4504219922bedaf40a7b1fd4bf4d3ccca59, tranaction count: 160, block.time: 1752721668
    2025/07/17 15:24:22 txHash: 0xd2a01c7255bac83bbd0425f0c7fcb040900d2ee8b3f9668f23953404a72d24fb
    =========== 任务2 ===================>
    2025/07/17 15:24:23 合约地址: 0xea6D7A662699Ee5FB4F8dD6E526a5622b925fA05, 交易hash: 0xdfcc896f9e376b1d6d1e30537a179149b06037b407c948900ef7c2346c537a6c  
    2025/07/17 15:24:25 部署交易确认✅: 0xdfcc896f9e376b1d6d1e30537a179149b06037b407c948900ef7c2346c537a6c, receipt.status: 1
    2025/07/17 15:24:28 交易hash: 0x584af4224896457a031de2fff2c58f2b58d8d85f548998abb14f89fc4269f375
    2025/07/17 15:26:36 交易确认✅: 0x584af4224896457a031de2fff2c58f2b58d8d85f548998abb14f89fc4269f375, receipt.status: 1
    2025/07/17 15:26:36 当前 count 值: 1
    2025/07/17 15:26:36 1 <nil>
    2025/07/17 15:26:39 交易hash: 0x6d401cbb65bbd7f3a6de42e868fcda75c6529da5ee2b0ca6d108ab57e057287c
    2025/07/17 15:26:48 交易确认✅: 0x6d401cbb65bbd7f3a6de42e868fcda75c6529da5ee2b0ca6d108ab57e057287c, receipt.status: 1
    2025/07/17 15:26:49 当前 count 值: 2
    2025/07/17 15:26:49 2 <nil>
    2025/07/17 15:26:50 交易hash: 0xe759654041668271353a4b35870eab622c4c103df89be05abbc303658d9a028d
    2025/07/17 15:27:12 交易确认✅: 0xe759654041668271353a4b35870eab622c4c103df89be05abbc303658d9a028d, receipt.status: 1
    2025/07/17 15:27:13 当前 count 值: 3
    2025/07/17 15:27:13 3 <nil>