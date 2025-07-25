// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract Nft is ERC721URIStorage{

    uint256 public MaxNum = 100;

    constructor(string memory _name,string memory _symbol) ERC721( _name, _symbol) {

    }

    function mintNft(address recipient,uint256 tokenId) external  {
        require (tokenId < MaxNum && tokenId >= 0,"tokenId out of range");
        if (_ownerOf(tokenId) != address(0)){
            revert("tokenId is exist");
        }
        _mint(recipient, tokenId);
    }

    function setTokenURI(uint256 tokenId , string memory tokenURL) external  {
        _setTokenURI(tokenId,tokenURL);
    }

}



// 合约地址: 0x6b2E7ceE042aEaa166a07B80618c5EC142C65612
// tokenURI: https://ipfs.filebase.io/ipfs/QmR7FmyaYZXdppEobzLaeaD4zn2mkWbRTrLMbRhLWfqi5A