/**
 *Submitted for verification at Etherscan.io on 2025-05-27
*/

// Sources flattened with hardhat v2.17.0 https://hardhat.org

// File contracts/DiplomaManager.sol

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract DiplomaManager {
    // Event emitted when a diploma is stored
    event DiplomaIssued(bytes32 indexed diplomaHash, string issuerDID);

    // Mapping from diploma hash to issuer DID
    mapping(bytes32 => string) private diplomaToIssuer;

    // Store a new diploma hash with the issuer's DID
    function storeDiploma(bytes32 diplomaHash, string memory issuerDID) public {
        require(bytes(diplomaToIssuer[diplomaHash]).length == 0, "Diploma already exists");
        diplomaToIssuer[diplomaHash] = issuerDID;
        emit DiplomaIssued(diplomaHash, issuerDID);
    }

    // Retrieve the DID of the issuer by diploma hash
    function getIssuerDID(bytes32 diplomaHash) public view returns (string memory) {
        require(bytes(diplomaToIssuer[diplomaHash]).length > 0, "Diploma not found");
        return diplomaToIssuer[diplomaHash];
    }

    // Check if a diploma hash exists
    function isDiplomaStored(bytes32 diplomaHash) public view returns (bool) {
        return bytes(diplomaToIssuer[diplomaHash]).length > 0;
    }
}