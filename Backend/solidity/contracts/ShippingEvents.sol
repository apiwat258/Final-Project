// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

contract ShippingEvents is Ownable {
    constructor(address initialOwner) Ownable(initialOwner) {}

    struct ShippingEvent {
        string shippingID;
        string productLotID;
        string fromLocation;
        string toLocation;
        string transporter;
        uint256 timestamp;
        uint256 temperature;
        uint256 humidity;
        bool delivered;
        string qualityInspectionCID;
    }

    mapping(string => ShippingEvent) public shippingRecords;
    event ShippingEventCreated(
        string shippingID,
        string productLotID,
        string fromLocation,
        string toLocation,
        string transporter,
        uint256 timestamp,
        uint256 temperature,
        uint256 humidity,
        bool delivered,
        string qualityInspectionCID
    );

    function addShippingEvent(
        string memory _shippingID,
        string memory _productLotID,
        string memory _fromLocation,
        string memory _toLocation,
        string memory _transporter,
        uint256 _temperature,
        uint256 _humidity,
        string memory _qualityInspectionCID
    ) public onlyOwner {
        require(shippingRecords[_shippingID].timestamp == 0, "Shipping event already exists");

        shippingRecords[_shippingID] = ShippingEvent({
            shippingID: _shippingID,
            productLotID: _productLotID,
            fromLocation: _fromLocation,
            toLocation: _toLocation,
            transporter: _transporter,
            timestamp: block.timestamp,
            temperature: _temperature,
            humidity: _humidity,
            delivered: false,
            qualityInspectionCID: _qualityInspectionCID
        });

        emit ShippingEventCreated(_shippingID, _productLotID, _fromLocation, _toLocation, _transporter, block.timestamp, _temperature, _humidity, false, _qualityInspectionCID);
    }
}