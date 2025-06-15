/**
 *Submitted for verification at Etherscan.io on 2025-06-11
*/

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract StateRecords {
    address public owner;

    struct Record {
        bytes32 hash;
        address issuer;
        uint256 timestamp;
    }

    mapping(bytes32 => Record) private records;

    event RecordIssued(bytes32 indexed recordId, address indexed issuer, uint256 timestamp);

    modifier onlyOwner() {
        require(msg.sender == owner, "Not the contract owner.");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function issueRecord(bytes32 recordId, bytes32 hash) public onlyOwner {
        require(records[recordId].timestamp == 0, "Record already exists.");
        records[recordId] = Record(hash, msg.sender, block.timestamp);
        emit RecordIssued(recordId, msg.sender, block.timestamp);
    }

    function verifyRecord(bytes32 recordId, bytes32 hash) public view returns (bool) {
        Record memory rec = records[recordId];
        return (rec.hash == hash);
    }
}