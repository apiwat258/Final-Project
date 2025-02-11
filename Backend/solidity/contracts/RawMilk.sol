// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

contract RawMilk is Ownable {
    constructor(address initialOwner) Ownable(initialOwner) {
        require(initialOwner != address(0), "Invalid owner address");
    }  // ✅ ป้องกัน Address เป็น 0x0

    struct MilkBatch {
        string rawMilkID;
        string farmerID;
        string batchNumber;
        uint256 quantity;
        uint256 temperature;
        uint256 timestamp;
        bool status;
    }

    struct QualityCheck {
        uint256 fatContent;
        uint256 proteinContent;
        uint256 bacteriaLevel;
    }

    mapping(string => MilkBatch) public milkBatches;
    mapping(string => QualityCheck) public milkQualityData;
    mapping(string => bytes32) public milkBatchHashes;

    event MilkBatchCreated(
        string rawMilkID,
        string farmerID,
        string batchNumber,
        uint256 quantity,
        uint256 temperature,
        uint256 timestamp
    );

    event MilkBatchStatusUpdated(string rawMilkID, bool status);
    event MilkBatchShipped(string rawMilkID, string transportedBy, uint256 shippedTimestamp);
    event MilkQualityChecked(string rawMilkID, uint256 fatContent, uint256 proteinContent, uint256 bacteriaLevel);

    function addMilkBatch(
        string memory _rawMilkID,
        string memory _farmerID,
        string memory _batchNumber,
        uint256 _quantity,
        uint256 _temperature
    ) public onlyOwner {
        require(milkBatches[_rawMilkID].timestamp == 0, "Milk batch already exists");

        milkBatches[_rawMilkID] = MilkBatch({
            rawMilkID: _rawMilkID,
            farmerID: _farmerID,
            batchNumber: _batchNumber,
            quantity: _quantity,
            temperature: _temperature,
            timestamp: block.timestamp,
            status: true
        });

        milkBatchHashes[_rawMilkID] = generateMilkBatchHash(_rawMilkID, _farmerID, _batchNumber, _quantity, _temperature);

        emit MilkBatchCreated(_rawMilkID, _farmerID, _batchNumber, _quantity, _temperature, block.timestamp);
    }

    function generateMilkBatchHash(
        string memory _rawMilkID,
        string memory _farmerID,
        string memory _batchNumber,
        uint256 _quantity,
        uint256 _temperature
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_rawMilkID, _farmerID, _batchNumber, _quantity, _temperature));
    }
}
