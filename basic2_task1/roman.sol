// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


contract Roman{
    // 罗马转数字
    mapping(bytes1 =>uint) public  romanMap ;
    struct ValueSymbol {
        uint value;
        bytes symbol;
    }

    ValueSymbol[] public valueSymbols;

    constructor() {
        romanMap["I"] = 1;
        romanMap["V"] = 5;
        romanMap["X"] = 10;
        romanMap["L"] = 50;
        romanMap["C"] = 100;
        romanMap["D"] = 500;
        romanMap["M"] = 1000;


        valueSymbols.push(ValueSymbol(1000, "M"));
        valueSymbols.push(ValueSymbol(900, "CM"));
        valueSymbols.push(ValueSymbol(500, "D"));
        valueSymbols.push(ValueSymbol(400, "CD"));
        valueSymbols.push(ValueSymbol(100, "C"));
        valueSymbols.push(ValueSymbol(90, "XC"));
        valueSymbols.push(ValueSymbol(50, "L"));
        valueSymbols.push(ValueSymbol(40, "XL"));
        valueSymbols.push(ValueSymbol(10, "X"));
        valueSymbols.push(ValueSymbol(9, "IX"));
        valueSymbols.push(ValueSymbol(5, "V"));
        valueSymbols.push(ValueSymbol(4, "IV"));
        valueSymbols.push(ValueSymbol(1, "I"));
    }

    // 罗马转数字
    function romanToInt(string memory _input) public view returns  (uint res) {
        bytes memory inputBytes = bytes(_input);
        for (uint i = 0; i < inputBytes.length; i++ ) {
            uint value = romanMap[inputBytes[i]];
            if (i < inputBytes.length - 1 && value < romanMap[inputBytes[i + 1]]){
                res -= romanMap[inputBytes[i]];
            }else {
                res += romanMap[inputBytes[i]];
            }

        }
        return res;
    }


    // 数字转罗马
    function intToRoman(uint input) public view returns (string memory res) {
        bytes memory result = new bytes(0);
        for (uint i = 0; i < valueSymbols.length; i++){
            uint val = valueSymbols[i].value;
            bytes memory symbol = valueSymbols[i].symbol;
            while  (input >= val) {
                input -= val;
                result  = abi.encodePacked(result,symbol);
            }

            if (input == 0){
                break;
            }

        }
        res = string(result);
        return res;
    }

    function getValueSymbols() public view returns(ValueSymbol[] memory) {
        return valueSymbols;
    }
}
