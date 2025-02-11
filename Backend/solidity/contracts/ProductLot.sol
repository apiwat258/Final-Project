// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

contract ProductLot is Ownable {
    constructor(address initialOwner) Ownable(initialOwner) {}
    struct LotInfo {
        string productLotID;
        string productID;
        string qrCode;
        uint256 quantityBox;
        uint256 mfgDate;
        uint256 expDate;
        bool status;
        uint256 createdOn;
        string[] rawMilkIDs; // เก็บรายการของน้ำนมดิบที่ใช้ในล็อตนี้
    }

    mapping(string => LotInfo) public productLots;
    event ProductLotCreated(
        string productLotID,
        string productID,
        string qrCode,
        uint256 quantityBox,
        uint256 mfgDate,
        uint256 expDate,
        bool status,
        uint256 createdOn,
        string[] rawMilkIDs
    );

    function addProductLot(
        string memory _productLotID,
        string memory _productID,
        string memory _qrCode,
        uint256 _quantityBox,
        uint256 _mfgDate,
        uint256 _expDate,
        string[] memory _rawMilkIDs
    ) public onlyOwner {
        require(productLots[_productLotID].createdOn == 0, "Product Lot already exists");

        productLots[_productLotID] = LotInfo({
            productLotID: _productLotID,
            productID: _productID,
            qrCode: _qrCode,
            quantityBox: _quantityBox,
            mfgDate: _mfgDate,
            expDate: _expDate,
            status: true,
            createdOn: block.timestamp,
            rawMilkIDs: _rawMilkIDs
        });

        emit ProductLotCreated(_productLotID, _productID, _qrCode, _quantityBox, _mfgDate, _expDate, true, block.timestamp, _rawMilkIDs);
    }
}
