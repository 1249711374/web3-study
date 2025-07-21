// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


// 合并有序数组
contract Array {
    function mergeSortedArray(uint[] memory arr1,uint[] memory arr2) public pure returns (uint[] memory) {
        uint len = arr1.length + arr2.length;
        uint[] memory result = new uint[](len);

        uint i = 0;
        uint j = 0;

        for (uint k = 0; k < arr1.length + arr2.length;k++) {
            while (i<arr1.length && j < arr2.length) {
                if (arr1[i] <= arr2[j]) {
                    result[k] = arr1[i];
                    k++;
                    i++;
                }else {
                    result[k] = arr2[j];
                    k++;
                    j++;
                }
            }

            if (i<arr1.length) {
                for (;i<arr1.length;i++) {
                    result[k] = arr1[i];
                    k++;
                }
            }
            if (j<arr2.length) {
                for (;j<arr2.length;j++) {
                    result[k] = arr2[j];
                    k++;
                }
            }
        }

        return result;
    }
}
