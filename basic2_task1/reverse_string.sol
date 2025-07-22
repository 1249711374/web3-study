// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 反转字符串
contract RevertString {

    function revertString(string memory _input) public pure returns(string memory) {
        bytes memory inputBytes = bytes(_input);

        uint len = inputBytes.length;

        bytes memory revertBytes = new bytes(len);

        for (uint i = 0;i < len;i++){
            revertBytes[i] = inputBytes[len - 1 - i];
        }
        return string(revertBytes);
    }
}
