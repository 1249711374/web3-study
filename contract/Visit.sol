// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Visit {
    uint public count = 0;

    function visit() public returns (uint){
        count += 1;
        return count;
    }
}