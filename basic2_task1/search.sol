// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 二分查找
contract Search {
    function binarySearch(uint[] calldata arr,uint target) public pure returns (uint index) {
        uint l = 0;
        uint r = arr.length - 1;


        while (l<=r) {
            uint mid = l + (r - l)/2;
            if (arr[mid] == target) {
                return mid;
            }

            if (arr[mid] < target) {
                l = mid+1;
            }else {
                r = mid-1;
            }
        }

        return 999999;
    }
}