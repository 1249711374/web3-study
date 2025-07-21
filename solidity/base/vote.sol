// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    mapping(address=>uint) private voteMap;

    // 校验空地址
    modifier checkVoteAddr (address voteAddr) {
        require(voteAddr!= address(0));
        _;
    }

    // 为候选人投票
    function vote(address voteAddr,uint voteNumber) external checkVoteAddr(voteAddr) {
        voteMap[voteAddr] += voteNumber;
    }

    // 查询票数
    function getVotes(address voteAddr) public view  checkVoteAddr(voteAddr) returns (uint) {
        return voteMap[voteAddr];
    }

    // 重置候选人票数为0
    function resetVotes(address voteAddr) external checkVoteAddr(voteAddr) {
        voteMap[voteAddr] = 0;
    }

}