// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/access/Ownable.sol";


// 合约地址: 0x01e59edd055181c5b8638355414f4ffa395c1968

// 讨饭合约
contract BeggingContract is Ownable{
    constructor() Ownable(msg.sender) {
        // 部署者是合约的拥有者
    }

    // 2025-09-25 19:28:00
    uint deadline = 1758799680;
    // 存储前三名捐赠者的地址
    address[3] top3map = [address(0),address(0),address(0)];
    // 捐赠最高的前3名(依次递增)
    uint256[3] public top3arr=[0,0,0];

    // 收款
    receive() external payable {
        emit DonateLog(msg.sender, msg.value);
    }


    // 捐赠地址和金额
    mapping(address => uint256) public donaters;

    event DonateLog(address indexed from, uint256 amount);

    function donate() external payable  {
        require(block.timestamp < deadline ,"donate activity end");
        require(msg.value > 0, "number must great than 0");
        donaters[msg.sender] += msg.value;
        // 触发事件
        calcTop3(msg.sender, donaters[msg.sender]);
        emit DonateLog(msg.sender, 0);
    }

    function calcTop3(address donater, uint256 val) internal  {
        // 大于最大值
        if (val > top3arr[2]) {
            top3arr[0] = top3arr[1];
            top3arr[1] = top3arr[2];
            top3arr[2] = val;
            top3map[0] = top3map[1];
            top3map[1] = top3map[2];
            top3map[2] = donater;
        } else if (val > top3arr[1]) {
            top3arr[0] = top3arr[1];
            top3arr[1] = val;
            top3map[0] = top3map[1];
            top3map[1] = donater;
        } else if (val > top3arr[0]) {
            top3arr[0] = val;
            top3map[0] = donater;
        }
    }

    function donateTop3Addr() public view returns(address[3] memory) {
        return top3map;
    }

    function viewOwner() public view returns(address) {
        return owner();
    }

    // 查询地址 捐赠 金额
    function getDonation(address donater) external view returns (uint256) {
        return donaters[donater];
    }

    // 将合约的钱转给调用者(即部署者) 使用call提现
    function withdraw() external  {
        require(msg.sender == owner(),"not owner");
        (bool success,) = payable(msg.sender).call{value: address(this).balance}("");
        require(success,"withdraw failed");
    }

    // 使用transfer 提现
    function withdrawv2() external  {
        require(msg.sender == owner(),"not owner");
        payable(msg.sender).transfer(address(this).balance);
    }


    function arr() external pure returns(string[] memory ss) {
        return new string[](0);
    }
}